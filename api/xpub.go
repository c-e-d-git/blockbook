package api

import (
	"blockbook/bchain"
	"blockbook/db"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"
	"strconv"

	"github.com/golang/glog"
	"github.com/juju/errors"
)

const defaultAddressesGap = 20
const maxAddressesGap = 10000

const txInput = 1
const txOutput = 2

const xpubCacheSize = 512
const xpubCacheExpirationSeconds = 7200

var cachedXpubs = make(map[string]xpubData)
var cachedXpubsMux sync.Mutex

type xpubTxid struct {
	txid        string
	height      uint32
	inputOutput byte
}

type xpubTxids []xpubTxid

func (a xpubTxids) Len() int      { return len(a) }
func (a xpubTxids) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a xpubTxids) Less(i, j int) bool {
	// if the heights are equal, make inputs less than outputs
	hi := a[i].height
	hj := a[j].height
	if hi == hj {
		return (a[i].inputOutput & txInput) >= (a[j].inputOutput & txInput)
	}
	return hi > hj
}

type xpubAddress struct {
	addrDesc  bchain.AddressDescriptor
	balance   *bchain.AddrBalance
	txs       uint32
	maxHeight uint32
	complete  bool
	txids     xpubTxids
}

type xpubData struct {
	gap             int
	accessed        int64
	basePath        string
	dataHeight      uint32
	dataHash        string
	txCountEstimate uint32
	sentSat         big.Int
	balanceSat      big.Int
	addresses       []xpubAddress
	changeAddresses []xpubAddress
}

func (w *Worker) xpubGetAddressTxids(addrDesc bchain.AddressDescriptor, mempool bool, fromHeight, toHeight uint32, maxResults int) ([]xpubTxid, bool, error) {
	var err error
	complete := true
	txs := make([]xpubTxid, 0, 4)
	var callback db.GetTransactionsCallback
	callback = func(txid string, height uint32, indexes []int32) error {
		// take all txs in the last found block even if it exceeds maxResults
		if len(txs) >= maxResults && txs[len(txs)-1].height != height {
			complete = false
			return &db.StopIteration{}
		}
		inputOutput := byte(0)
		for _, index := range indexes {
			if index < 0 {
				inputOutput |= txInput
			} else {
				inputOutput |= txOutput
			}
		}
		txs = append(txs, xpubTxid{txid, height, inputOutput})
		return nil
	}
	if mempool {
		uniqueTxs := make(map[string]int)
		o, err := w.mempool.GetAddrDescTransactions(addrDesc)
		if err != nil {
			return nil, false, err
		}
		for _, m := range o {
			if l, found := uniqueTxs[m.Txid]; !found {
				l = len(txs)
				callback(m.Txid, 0, []int32{m.Vout})
				if len(txs) > l {
					uniqueTxs[m.Txid] = l
				}
			} else {
				if m.Vout < 0 {
					txs[l].inputOutput |= txInput
				} else {
					txs[l].inputOutput |= txOutput
				}
			}
		}
	} else {
		err = w.db.GetAddrDescTransactions(addrDesc, fromHeight, toHeight, callback)
		if err != nil {
			return nil, false, err
		}
	}
	return txs, complete, nil
}

func (w *Worker) xpubCheckAndLoadTxids(ad *xpubAddress, filter *AddressFilter, maxHeight uint32, maxResults int) error {
	// skip if not used
	if ad.balance == nil {
		return nil
	}
	// if completely loaded, check if there are not some new txs and load if necessary
	if ad.complete {
		if ad.balance.Txs != ad.txs {
			newTxids, _, err := w.xpubGetAddressTxids(ad.addrDesc, false, ad.maxHeight+1, maxHeight, maxInt)
			if err == nil {
				ad.txids = append(newTxids, ad.txids...)
				ad.maxHeight = maxHeight
				ad.txs = uint32(len(ad.txids))
				if ad.txs != ad.balance.Txs {
					glog.Warning("xpubCheckAndLoadTxids inconsistency ", ad.addrDesc, ", ad.txs=", ad.txs, ", ad.balance.Txs=", ad.balance.Txs)
				}
			}
			return err
		}
		return nil
	}
	// load all txids to get paging correctly
	newTxids, complete, err := w.xpubGetAddressTxids(ad.addrDesc, false, 0, maxHeight, maxInt)
	if err != nil {
		return err
	}
	ad.txids = newTxids
	ad.complete = complete
	ad.maxHeight = maxHeight
	if complete {
		ad.txs = uint32(len(ad.txids))
		if ad.txs != ad.balance.Txs {
			glog.Warning("xpubCheckAndLoadTxids inconsistency ", ad.addrDesc, ", ad.txs=", ad.txs, ", ad.balance.Txs=", ad.balance.Txs)
		}
	}
	return nil
}

func (w *Worker) xpubDerivedAddressBalance(data *xpubData, ad *xpubAddress) (bool, error) {
	var err error
	if ad.balance, err = w.db.GetAddrDescBalance(ad.addrDesc, bchain.AddressBalanceDetailUTXO); err != nil {
		return false, err
	}
	if ad.balance != nil {
		data.txCountEstimate += ad.balance.Txs
		data.sentSat.Add(&data.sentSat, &ad.balance.SentSat)
		data.balanceSat.Add(&data.balanceSat, &ad.balance.BalanceSat)
		return true, nil
	}
	return false, nil
}

func (w *Worker) xpubScanAddresses(xpub string, data *xpubData, addresses []xpubAddress, gap int, change int, minDerivedIndex int, fork bool) (int, []xpubAddress, error) {
	// rescan known addresses
	lastUsed := 0
	for i := range addresses {
		ad := &addresses[i]
		if fork {
			// reset the cached data
			ad.txs = 0
			ad.maxHeight = 0
			ad.complete = false
			ad.txids = nil
		}
		used, err := w.xpubDerivedAddressBalance(data, ad)
		if err != nil {
			return 0, nil, err
		}
		if used {
			lastUsed = i
		}
	}
	// derive new addresses as necessary
	missing := len(addresses) - lastUsed
	for missing < gap {
		from := len(addresses)
		to := from + gap - missing
		if to < minDerivedIndex {
			to = minDerivedIndex
		}
		descriptors, err := w.chainParser.DeriveAddressDescriptorsFromTo(xpub, uint32(change), uint32(from), uint32(to))
		if err != nil {
			return 0, nil, err
		}
		for i, a := range descriptors {
			ad := xpubAddress{addrDesc: a}
			used, err := w.xpubDerivedAddressBalance(data, &ad)
			if err != nil {
				return 0, nil, err
			}
			if used {
				lastUsed = i + from
			}
			addresses = append(addresses, ad)
		}
		missing = len(addresses) - lastUsed
	}
	return lastUsed, addresses, nil
}

func (w *Worker) tokenFromXpubAddress(data *xpubData, ad *xpubAddress, changeIndex int, index int, option AccountDetails) []Token {
	a, _, _ := w.chainParser.GetAddressesFromAddrDesc(ad.addrDesc)
	var tokens []Token
	var address string
	if len(a) > 0 {
		address = a[0]
	}

	if ad.balance != nil {
		transfers := int(ad.balance.Txs)
		if option >= AccountDetailsTokenBalances {
			balance := &ad.balance.BalanceSat
			totalSent := &ad.balance.SentSat
			totalReceived := ad.balance.ReceivedSat()
			tokens = append(tokens, Token{
				Type:             XPUBAddressTokenType,
				Name:             address,
				Decimals:         w.chainParser.AmountDecimals(),
				BalanceSat:       (*Amount)(balance),
				TotalReceivedSat: (*Amount)(totalReceived),
				TotalSentSat:     (*Amount)(totalSent),
				Transfers:        transfers,
				Path:             fmt.Sprintf("%s/%d/%d", data.basePath, changeIndex, index),
			})
			for k, v := range ad.balance.AssetBalances {
				balanceAssetSat := &v.BalanceAssetSat
				sentAssetSat := &v.SentAssetSat
				totalReceived := balanceAssetSat.Add(balanceAssetSat, sentAssetSat)
				// add token as unallocated if address matches asset owner address other wise its allocated
				tokens = append(tokens, Token{
					Type:             SPTAllocatedTokenType,
					Name:             address,
					Decimals:         w.chainParser.AmountDecimals(),
					Symbol:			  "SPT",
					BalanceSat:       (*Amount)(balanceAssetSat),
					TotalReceivedSat: (*Amount)(totalReceived),
					TotalSentSat:     (*Amount)(sentAssetSat),
					Transfers:        transfers,
					Path:             fmt.Sprintf("%s/%d/%d", data.basePath, changeIndex, index),
					Contract:		  strconv.FormatUint(uint64(k), 10),
				})
			}
		}
	}
	return tokens
}

func evictXpubCacheItems() {
	var oldestKey string
	oldest := maxInt64
	now := time.Now().Unix()
	count := 0
	for k, v := range cachedXpubs {
		if v.accessed+xpubCacheExpirationSeconds < now {
			delete(cachedXpubs, k)
			count++
		}
		if v.accessed < oldest {
			oldestKey = k
			oldest = v.accessed
		}
	}
	if oldestKey != "" && oldest+xpubCacheExpirationSeconds >= now {
		delete(cachedXpubs, oldestKey)
		count++
	}
	glog.Info("Evicted ", count, " items from xpub cache, oldest item accessed at ", time.Unix(oldest, 0), ", cache size ", len(cachedXpubs))
}

func (w *Worker) getXpubData(xpub string, page int, txsOnPage int, option AccountDetails, filter *AddressFilter, gap int) (*xpubData, uint32, error) {
	if w.chainType != bchain.ChainBitcoinType {
		return nil, 0, ErrUnsupportedXpub
	}
	var (
		err        error
		bestheight uint32
		besthash   string
	)
	if gap <= 0 {
		gap = defaultAddressesGap
	} else if gap > maxAddressesGap {
		// limit the maximum gap to protect against unreasonably big values that could cause high load of the server
		gap = maxAddressesGap
	}
	// gap is increased one as there must be gap of empty addresses before the derivation is stopped
	gap++
	var processedHash string
	cachedXpubsMux.Lock()
	data, found := cachedXpubs[xpub]
	cachedXpubsMux.Unlock()
	// to load all data for xpub may take some time, do it in a loop to process a possible new block
	for {
		bestheight, besthash, err = w.db.GetBestBlock()
		if err != nil {
			return nil, 0, errors.Annotatef(err, "GetBestBlock")
		}
		if besthash == processedHash {
			break
		}
		fork := false
		if !found || data.gap != gap {
			data = xpubData{gap: gap}
			data.basePath, err = w.chainParser.DerivationBasePath(xpub)
			if err != nil {
				glog.Warning("DerivationBasePath error", err)
				data.basePath = "unknown"
			}
		} else {
			hash, err := w.db.GetBlockHash(data.dataHeight)
			if err != nil {
				return nil, 0, err
			}
			if hash != data.dataHash {
				// in case of for reset all cached data
				fork = true
			}
		}
		processedHash = besthash
		if data.dataHeight < bestheight || fork {
			data.dataHeight = bestheight
			data.dataHash = besthash
			data.balanceSat = *new(big.Int)
			data.sentSat = *new(big.Int)
			data.txCountEstimate = 0
			var lastUsedIndex int
			lastUsedIndex, data.addresses, err = w.xpubScanAddresses(xpub, &data, data.addresses, gap, 0, 0, fork)
			if err != nil {
				return nil, 0, err
			}
			_, data.changeAddresses, err = w.xpubScanAddresses(xpub, &data, data.changeAddresses, gap, 1, lastUsedIndex, fork)
			if err != nil {
				return nil, 0, err
			}
		}
		if option >= AccountDetailsTxidHistory {
			for _, da := range [][]xpubAddress{data.addresses, data.changeAddresses} {
				for i := range da {
					if err = w.xpubCheckAndLoadTxids(&da[i], filter, bestheight, (page+1)*txsOnPage); err != nil {
						return nil, 0, err
					}
				}
			}
		}
	}
	data.accessed = time.Now().Unix()
	cachedXpubsMux.Lock()
	if len(cachedXpubs) >= xpubCacheSize {
		evictXpubCacheItems()
	}
	cachedXpubs[xpub] = data
	cachedXpubsMux.Unlock()
	return &data, bestheight, nil
}

// GetXpubAddress computes address value and gets transactions for given address
func (w *Worker) GetXpubAddress(xpub string, page int, txsOnPage int, option AccountDetails, filter *AddressFilter, gap int) (*Address, error) {
	start := time.Now()
	page--
	if page < 0 {
		page = 0
	}
	type mempoolMap struct {
		tx          *Tx
		inputOutput byte
	}
	var (
		txc            xpubTxids
		txmMap         map[string]*Tx
		txCount        int
		txs            []*Tx
		txids          []string
		pg             Paging
		filtered       bool
		err            error
		uBalSat        big.Int
		unconfirmedTxs int
	)
	data, bestheight, err := w.getXpubData(xpub, page, txsOnPage, option, filter, gap)
	if err != nil {
		return nil, err
	}
	// setup filtering of txids
	var txidFilter func(txid *xpubTxid, ad *xpubAddress) bool
	if !(filter.FromHeight == 0 && filter.ToHeight == 0 && filter.Vout == AddressFilterVoutOff) {
		toHeight := maxUint32
		if filter.ToHeight != 0 {
			toHeight = filter.ToHeight
		}
		txidFilter = func(txid *xpubTxid, ad *xpubAddress) bool {
			if txid.height < filter.FromHeight || txid.height > toHeight {
				return false
			}
			if filter.Vout != AddressFilterVoutOff {
				if filter.Vout == AddressFilterVoutInputs && txid.inputOutput&txInput == 0 ||
					filter.Vout == AddressFilterVoutOutputs && txid.inputOutput&txOutput == 0 {
					return false
				}
			}
			return true
		}
		filtered = true
	}
	// process mempool, only if ToHeight is not specified
	if filter.ToHeight == 0 && !filter.OnlyConfirmed {
		txmMap = make(map[string]*Tx)
		mempoolEntries := make(bchain.MempoolTxidEntries, 0)
		for _, da := range [][]xpubAddress{data.addresses, data.changeAddresses} {
			for i := range da {
				ad := &da[i]
				newTxids, _, err := w.xpubGetAddressTxids(ad.addrDesc, true, 0, 0, maxInt)
				if err != nil {
					return nil, err
				}
				for _, txid := range newTxids {
					// the same tx can have multiple addresses from the same xpub, get it from backend it only once
					tx, foundTx := txmMap[txid.txid]
					if !foundTx {
						tx, err = w.GetTransaction(txid.txid, false, false)
						// mempool transaction may fail
						if err != nil || tx == nil {
							glog.Warning("GetTransaction in mempool: ", err)
							continue
						}
						txmMap[txid.txid] = tx
					}
					// skip already confirmed txs, mempool may be out of sync
					if tx.Confirmations == 0 {
						if !foundTx {
							unconfirmedTxs++
						}
						uBalSat.Add(&uBalSat, tx.getAddrVoutValue(ad.addrDesc))
						uBalSat.Sub(&uBalSat, tx.getAddrVinValue(ad.addrDesc))
						// mempool txs are returned only on the first page, uniquely and filtered
						if page == 0 && !foundTx && (txidFilter == nil || txidFilter(&txid, ad)) {
							mempoolEntries = append(mempoolEntries, bchain.MempoolTxidEntry{Txid: txid.txid, Time: uint32(tx.Blocktime)})
						}
					}
				}
			}
		}
		// sort the entries by time descending
		sort.Sort(mempoolEntries)
		for _, entry := range mempoolEntries {
			if option == AccountDetailsTxidHistory {
				txids = append(txids, entry.Txid)
			} else if option >= AccountDetailsTxHistoryLight {
				txs = append(txs, txmMap[entry.Txid])
			}
		}
	}
	if option >= AccountDetailsTxidHistory {
		txcMap := make(map[string]bool)
		txc = make(xpubTxids, 0, 32)
		for _, da := range [][]xpubAddress{data.addresses, data.changeAddresses} {
			for i := range da {
				ad := &da[i]
				for _, txid := range ad.txids {
					added, foundTx := txcMap[txid.txid]
					// count txs regardless of filter but only once
					if !foundTx {
						txCount++
					}
					// add tx only once
					if !added {
						add := txidFilter == nil || txidFilter(&txid, ad)
						txcMap[txid.txid] = add
						if add {
							txc = append(txc, txid)
						}
					}
				}
			}
		}
		sort.Stable(txc)
		txCount = len(txcMap)
		totalResults := txCount
		if filtered {
			totalResults = -1
		}
		var from, to int
		pg, from, to, page = computePaging(len(txc), page, txsOnPage)
		if len(txc) >= txsOnPage {
			if totalResults < 0 {
				pg.TotalPages = -1
			} else {
				pg, _, _, _ = computePaging(totalResults, page, txsOnPage)
			}
		}
		// get confirmed transactions
		for i := from; i < to; i++ {
			xpubTxid := &txc[i]
			if option == AccountDetailsTxidHistory {
				txids = append(txids, xpubTxid.txid)
			} else {
				tx, err := w.txFromTxid(xpubTxid.txid, bestheight, option, nil)
				if err != nil {
					return nil, err
				}
				txs = append(txs, tx)
			}
		}
	} else {
		txCount = int(data.txCountEstimate)
	}
	usedTokens := 0
	var tokens []Token
	var xpubAddresses map[string]struct{}
	if option > AccountDetailsBasic {
		tokens = make([]Token, 0, 4)
		xpubAddresses = make(map[string]struct{})
	}
	for ci, da := range [][]xpubAddress{data.addresses, data.changeAddresses} {
		for i := range da {
			ad := &da[i]
			if ad.balance != nil {
				usedTokens++
			}
			if option > AccountDetailsBasic {
				tokens := w.tokenFromXpubAddress(data, ad, ci, i, option)
				for _, token := range tokens {
					if filter.TokensToReturn == TokensToReturnDerived ||
						filter.TokensToReturn == TokensToReturnUsed && token.BalanceSat != nil ||
						filter.TokensToReturn == TokensToReturnNonzeroBalance && token.BalanceSat != nil && token.BalanceSat > 0 {
						tokens = append(tokens, token)
					}
					xpubAddresses[token.Name] = struct{}{}
				}
			}
		}
	}
	var totalReceived big.Int
	totalReceived.Add(&data.balanceSat, &data.sentSat)
	addr := Address{
		Paging:                pg,
		AddrStr:               xpub,
		BalanceSat:            (*Amount)(&data.balanceSat),
		TotalReceivedSat:      (*Amount)(&totalReceived),
		TotalSentSat:          (*Amount)(&data.sentSat),
		Txs:                   txCount,
		UnconfirmedBalanceSat: (*Amount)(&uBalSat),
		UnconfirmedTxs:        unconfirmedTxs,
		Transactions:          txs,
		Txids:                 txids,
		UsedTokens:            usedTokens,
		Tokens:                tokens,
		XPubAddresses:         xpubAddresses,
	}
	glog.Info("GetXpubAddress ", xpub[:16], ", ", len(data.addresses)+len(data.changeAddresses), " derived addresses, ", txCount, " confirmed txs, finished in ", time.Since(start))
	return &addr, nil
}

// GetXpubUtxo returns unspent outputs for given xpub
func (w *Worker) GetXpubUtxo(xpub string, onlyConfirmed bool, gap int) (Utxos, error) {
	start := time.Now()
	data, _, err := w.getXpubData(xpub, 0, 1, AccountDetailsBasic, &AddressFilter{
		Vout:          AddressFilterVoutOff,
		OnlyConfirmed: onlyConfirmed,
	}, gap)
	if err != nil {
		return nil, err
	}
	r := make(Utxos, 0, 8)
	for ci, da := range [][]xpubAddress{data.addresses, data.changeAddresses} {
		for i := range da {
			ad := &da[i]
			onlyMempool := false
			if ad.balance == nil {
				if onlyConfirmed {
					continue
				}
				onlyMempool = true
			}
			utxos, err := w.getAddrDescUtxo(ad.addrDesc, ad.balance, onlyConfirmed, onlyMempool)
			if err != nil {
				return nil, err
			}
			if len(utxos) > 0 {
				txs := w.tokenFromXpubAddress(data, ad, ci, i, AccountDetailsTokens)
				for _ , t := range txs {
					for j := range utxos {
						a := &utxos[j]
						a.Address = t.Name
						a.Path = t.Path
					}
				}
				r = append(r, utxos...)
			}
		}
	}
	sort.Stable(r)
	glog.Info("GetXpubUtxo ", xpub[:16], ", ", len(r), " utxos, finished in ", time.Since(start))
	return r, nil
}

// GetXpubBalanceHistory returns history of balance for given xpub
func (w *Worker) GetXpubBalanceHistory(xpub string, fromTime, toTime time.Time, fiat string, gap int, groupBy uint32) (BalanceHistories, error) {
	bhs := make(BalanceHistories, 0)
	start := time.Now()
	fromUnix, fromHeight, toUnix, toHeight := w.balanceHistoryHeightsFromTo(fromTime, toTime)
	if fromHeight >= toHeight {
		return bhs, nil
	}
	data, _, err := w.getXpubData(xpub, 0, 1, AccountDetailsTxidHistory, &AddressFilter{
		Vout:          AddressFilterVoutOff,
		OnlyConfirmed: true,
		FromHeight:    fromHeight,
		ToHeight:      toHeight,
	}, gap)
	if err != nil {
		return nil, err
	}
	for _, da := range [][]xpubAddress{data.addresses, data.changeAddresses} {
		for i := range da {
			ad := &da[i]
			txids := ad.txids
			for txi := len(txids) - 1; txi >= 0; txi-- {
				bh, err := w.balanceHistoryForTxid(ad.addrDesc, txids[txi].txid, fromUnix, toUnix)
				if err != nil {
					return nil, err
				}
				if bh != nil {
					bhs = append(bhs, *bh)
				}
			}
		}
	}
	bha := bhs.SortAndAggregate(groupBy)
	if fiat != "" {
		err = w.setFiatRateToBalanceHistories(bha, fiat)
		if err != nil {
			return nil, err
		}
	}
	glog.Info("GetUtxoBalanceHistory ", xpub[:16], ", blocks ", fromHeight, "-", toHeight, ", count ", len(bha), ", finished in ", time.Since(start))
	return bha, nil
}
