// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tx.proto

/*
Package bchain is a generated protocol buffer package.

It is generated from these files:
	tx.proto

It has these top-level messages:
	ProtoTransaction
*/
package bchain

import proto "github.com/golang/protobuf/proto"
import bech32 "github.com/martinboehm/btcutil/bech32"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtoTransaction struct {
	Txid      []byte                       `protobuf:"bytes,1,opt,name=Txid,proto3" json:"Txid,omitempty"`
	Hex       []byte                       `protobuf:"bytes,2,opt,name=Hex,proto3" json:"Hex,omitempty"`
	Blocktime uint64                       `protobuf:"varint,3,opt,name=Blocktime" json:"Blocktime,omitempty"`
	Locktime  uint32                       `protobuf:"varint,4,opt,name=Locktime" json:"Locktime,omitempty"`
	Height    uint32                       `protobuf:"varint,5,opt,name=Height" json:"Height,omitempty"`
	Vin       []*ProtoTransaction_VinType  `protobuf:"bytes,6,rep,name=Vin" json:"Vin,omitempty"`
	Vout      []*ProtoTransaction_VoutType `protobuf:"bytes,7,rep,name=Vout" json:"Vout,omitempty"`
	Version   int32                        `protobuf:"varint,8,opt,name=Version" json:"Version,omitempty"`
}

func (m *ProtoTransaction) Reset()                    { *m = ProtoTransaction{} }
func (m *ProtoTransaction) String() string            { return proto.CompactTextString(m) }
func (*ProtoTransaction) ProtoMessage()               {}
func (*ProtoTransaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoTransaction) GetTxid() []byte {
	if m != nil {
		return m.Txid
	}
	return nil
}

func (m *ProtoTransaction) GetHex() []byte {
	if m != nil {
		return m.Hex
	}
	return nil
}

func (m *ProtoTransaction) GetBlocktime() uint64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *ProtoTransaction) GetLocktime() uint32 {
	if m != nil {
		return m.Locktime
	}
	return 0
}

func (m *ProtoTransaction) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ProtoTransaction) GetVin() []*ProtoTransaction_VinType {
	if m != nil {
		return m.Vin
	}
	return nil
}

func (m *ProtoTransaction) GetVout() []*ProtoTransaction_VoutType {
	if m != nil {
		return m.Vout
	}
	return nil
}

func (m *ProtoTransaction) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type ProtoTransaction_VinType struct {
	Coinbase     string   `protobuf:"bytes,1,opt,name=Coinbase" json:"Coinbase,omitempty"`
	Txid         []byte   `protobuf:"bytes,2,opt,name=Txid,proto3" json:"Txid,omitempty"`
	Vout         uint32   `protobuf:"varint,3,opt,name=Vout" json:"Vout,omitempty"`
	ScriptSigHex []byte   `protobuf:"bytes,4,opt,name=ScriptSigHex,proto3" json:"ScriptSigHex,omitempty"`
	Sequence     uint32   `protobuf:"varint,5,opt,name=Sequence" json:"Sequence,omitempty"`
	Addresses    []string `protobuf:"bytes,6,rep,name=Addresses" json:"Addresses,omitempty"`
}

func (m *ProtoTransaction_VinType) Reset()                    { *m = ProtoTransaction_VinType{} }
func (m *ProtoTransaction_VinType) String() string            { return proto.CompactTextString(m) }
func (*ProtoTransaction_VinType) ProtoMessage()               {}
func (*ProtoTransaction_VinType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ProtoTransaction_VinType) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

func (m *ProtoTransaction_VinType) GetTxid() []byte {
	if m != nil {
		return m.Txid
	}
	return nil
}

func (m *ProtoTransaction_VinType) GetVout() uint32 {
	if m != nil {
		return m.Vout
	}
	return 0
}

func (m *ProtoTransaction_VinType) GetScriptSigHex() []byte {
	if m != nil {
		return m.ScriptSigHex
	}
	return nil
}

func (m *ProtoTransaction_VinType) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *ProtoTransaction_VinType) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type ProtoTransaction_VoutType struct {
	ValueSat        []byte   `protobuf:"bytes,1,opt,name=ValueSat,proto3" json:"ValueSat,omitempty"`
	N               uint32   `protobuf:"varint,2,opt,name=N" json:"N,omitempty"`
	ScriptPubKeyHex []byte   `protobuf:"bytes,3,opt,name=ScriptPubKeyHex,proto3" json:"ScriptPubKeyHex,omitempty"`
	Addresses       []string `protobuf:"bytes,4,rep,name=Addresses" json:"Addresses,omitempty"`
}

func (m *ProtoTransaction_VoutType) Reset()                    { *m = ProtoTransaction_VoutType{} }
func (m *ProtoTransaction_VoutType) String() string            { return proto.CompactTextString(m) }
func (*ProtoTransaction_VoutType) ProtoMessage()               {}
func (*ProtoTransaction_VoutType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *ProtoTransaction_VoutType) GetValueSat() []byte {
	if m != nil {
		return m.ValueSat
	}
	return nil
}

func (m *ProtoTransaction_VoutType) GetN() uint32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *ProtoTransaction_VoutType) GetScriptPubKeyHex() []byte {
	if m != nil {
		return m.ScriptPubKeyHex
	}
	return nil
}

func (m *ProtoTransaction_VoutType) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}
type ProtoTransaction_WitnessAddressType struct {
	Version			byte		`protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	WitnessProgram 	[]byte		`protobuf:"bytes,2,opt,name=WitnessProgram" json:"WitnessProgram,omitempty"`
}
type ProtoTransaction_RangeAmountPairType struct {
	witnessProgram 	*ProtoTransaction_WitnessAddressType	`protobuf:"bytes,1,opt,name=witnessProgram" json:"witnessProgram,omitempty"`
	Amount         	big.Int   	`protobuf:"bytes,2,opt,name=Amount" json:"Amount,omitempty"`
}
type ProtoTransaction_AssetAllocationTupleType struct {
	Asset			uint32	`protobuf:"varint,1,opt,name=Asset" json:"Asset,omitempty"`
	witnessProgram 	*ProtoTransaction_WitnessAddressType	`protobuf:"bytes,2,opt,name=witnessProgram" json:"witnessProgram,omitempty"`
}
type ProtoTransaction_AssetAllocationType struct {
	assetAllocationTuple     		*ProtoTransaction_AssetAllocationTupleType   `protobuf:"bytes,1,opt,name=assetAllocationTuple" json:"assetAllocationTuple,omitempty"`
	listSendingAllocationAmounts	[]*ProtoTransaction_RangeAmountPairType   	 `protobuf:"bytes,2,opt,name=listSendingAllocationAmounts" json:"listSendingAllocationAmounts,omitempty"`
}

func (m *ProtoTransaction_WitnessAddressType) ToString() string {
	if m != nil {
		if len(m.WitnessProgram) <= 4 && string(m.WitnessProgram) == "burn" {
			return "burn"
		}
		// Convert data to base32:
		conv, err := bech32.ConvertBits(m.WitnessProgram, 8, 5, true)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		encoded, err := bech32.Encode("sys", conv)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		return encoded
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoTransaction)(nil), "bchain.ProtoTransaction")
	proto.RegisterType((*ProtoTransaction_VinType)(nil), "bchain.ProtoTransaction.VinType")
	proto.RegisterType((*ProtoTransaction_VoutType)(nil), "bchain.ProtoTransaction.VoutType")
	proto.RegisterType((*ProtoTransaction_AssetAllocationType)(nil), "bchain.ProtoTransaction.AssetAllocationType")
	proto.RegisterType((*ProtoTransaction_AssetAllocationTupleType)(nil), "bchain.ProtoTransaction.AssetAllocationTupleType")
	proto.RegisterType((*ProtoTransaction_RangeAmountPairType)(nil), "bchain.ProtoTransaction.RangeAmountPairType")
	proto.RegisterType((*ProtoTransaction_WitnessAddressType)(nil), "bchain.ProtoTransaction.WitnessAddressType")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0x6e, 0xea, 0x30,
	0x14, 0x94, 0x89, 0x09, 0xe1, 0x7d, 0xd0, 0x47, 0x5e, 0x54, 0x16, 0xea, 0x22, 0x65, 0x95, 0x55,
	0x16, 0x54, 0x3d, 0x40, 0xdb, 0x0d, 0x52, 0x2b, 0x84, 0x1c, 0x94, 0x7d, 0x12, 0x2c, 0xb0, 0x4a,
	0x6d, 0x9a, 0x38, 0x12, 0x48, 0xbd, 0x51, 0x8f, 0xd0, 0xcb, 0x55, 0x7e, 0x84, 0x50, 0x90, 0xba,
	0xf3, 0x8c, 0xdf, 0x64, 0xe6, 0x4d, 0x0c, 0x81, 0xdd, 0xc7, 0xbb, 0xd2, 0x58, 0xc3, 0xfc, 0xbc,
	0xd8, 0x64, 0x4a, 0x4f, 0xbe, 0x29, 0x8c, 0x16, 0x8e, 0x59, 0x96, 0x99, 0xae, 0xb2, 0xc2, 0x2a,
	0xa3, 0x19, 0x03, 0xba, 0xdc, 0xab, 0x15, 0x27, 0x21, 0x89, 0x06, 0x02, 0xcf, 0x6c, 0x04, 0xde,
	0x4c, 0xee, 0x79, 0x07, 0x29, 0x77, 0x64, 0xb7, 0xd0, 0x7f, 0xda, 0x9a, 0xe2, 0xcd, 0xaa, 0x77,
	0xc9, 0xbd, 0x90, 0x44, 0x54, 0x9c, 0x09, 0x36, 0x86, 0xe0, 0xf5, 0x74, 0x49, 0x43, 0x12, 0x0d,
	0x45, 0x8b, 0xd9, 0x0d, 0xf8, 0x33, 0xa9, 0xd6, 0x1b, 0xcb, 0xbb, 0x78, 0xd3, 0x20, 0x36, 0x05,
	0x2f, 0x55, 0x9a, 0xfb, 0xa1, 0x17, 0xfd, 0x9b, 0x86, 0xf1, 0x31, 0x62, 0x7c, 0x1d, 0x2f, 0x4e,
	0x95, 0x5e, 0x1e, 0x76, 0x52, 0xb8, 0x61, 0xf6, 0x00, 0x34, 0x35, 0xb5, 0xe5, 0x3d, 0x14, 0xdd,
	0xfd, 0x2d, 0x32, 0xb5, 0x45, 0x15, 0x8e, 0x33, 0x0e, 0xbd, 0x54, 0x96, 0x95, 0x32, 0x9a, 0x07,
	0x21, 0x89, 0xba, 0xe2, 0x04, 0xc7, 0x5f, 0x04, 0x7a, 0x8d, 0x83, 0x5b, 0xe2, 0xd9, 0x28, 0x9d,
	0x67, 0x95, 0xc4, 0x32, 0xfa, 0xa2, 0xc5, 0x6d, 0x49, 0x9d, 0x5f, 0x25, 0xb1, 0x26, 0x8c, 0x87,
	0x6b, 0x1d, 0x9d, 0x26, 0x30, 0x48, 0x8a, 0x52, 0xed, 0x6c, 0xa2, 0xd6, 0xae, 0x41, 0x8a, 0xf3,
	0x17, 0x9c, 0xf3, 0x49, 0xe4, 0x47, 0x2d, 0x75, 0x21, 0x9b, 0x4a, 0x5a, 0xec, 0x6a, 0x7e, 0x5c,
	0xad, 0x4a, 0x59, 0x55, 0xb2, 0xc2, 0x6a, 0xfa, 0xe2, 0x4c, 0x8c, 0x3f, 0x21, 0x38, 0x6d, 0xe6,
	0xbe, 0x92, 0x66, 0xdb, 0x5a, 0x26, 0x99, 0x6d, 0x7e, 0x5d, 0x8b, 0xd9, 0x00, 0xc8, 0x1c, 0xa3,
	0x0e, 0x05, 0x99, 0xb3, 0x08, 0xfe, 0x1f, 0xfd, 0x17, 0x75, 0xfe, 0x22, 0x0f, 0x2e, 0x96, 0x87,
	0x82, 0x6b, 0xfa, 0xd2, 0x9d, 0x5e, 0xb9, 0xe7, 0x3e, 0x3e, 0xa6, 0xfb, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa1, 0x51, 0x2e, 0xba, 0x58, 0x02, 0x00, 0x00,
}
