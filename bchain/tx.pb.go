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
	Txid      []byte                       `protobuf:"bytes,1,opt,name=Txid,json=txid,proto3" json:"Txid,omitempty"`
	Hex       []byte                       `protobuf:"bytes,2,opt,name=Hex,json=hex,proto3" json:"Hex,omitempty"`
	Blocktime uint64                       `protobuf:"varint,3,opt,name=Blocktime,json=blocktime" json:"Blocktime,omitempty"`
	Locktime  uint32                       `protobuf:"varint,4,opt,name=Locktime,json=locktime" json:"Locktime,omitempty"`
	Height    uint32                       `protobuf:"varint,5,opt,name=Height,json=height" json:"Height,omitempty"`
	Vin       []*ProtoTransaction_VinType  `protobuf:"bytes,6,rep,name=Vin,json=vin" json:"Vin,omitempty"`
	Vout      []*ProtoTransaction_VoutType `protobuf:"bytes,7,rep,name=Vout,json=vout" json:"Vout,omitempty"`
	Version   int32                        `protobuf:"varint,8,opt,name=Version,json=version" json:"Version,omitempty"`
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
	Coinbase     string   `protobuf:"bytes,1,opt,name=Coinbase,json=coinbase" json:"Coinbase,omitempty"`
	Txid         []byte   `protobuf:"bytes,2,opt,name=Txid,json=txid,proto3" json:"Txid,omitempty"`
	Vout         uint32   `protobuf:"varint,3,opt,name=Vout,json=vout" json:"Vout,omitempty"`
	ScriptSigHex []byte   `protobuf:"bytes,4,opt,name=ScriptSigHex,json=scriptSigHex,proto3" json:"ScriptSigHex,omitempty"`
	Sequence     uint32   `protobuf:"varint,5,opt,name=Sequence,json=sequence" json:"Sequence,omitempty"`
	Addresses    []string `protobuf:"bytes,6,rep,name=Addresses,json=addresses" json:"Addresses,omitempty"`
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
	ValueSat        []byte   `protobuf:"bytes,1,opt,name=ValueSat,json=valueSat,proto3" json:"ValueSat,omitempty"`
	N               uint32   `protobuf:"varint,2,opt,name=N,json=n" json:"N,omitempty"`
	ScriptPubKeyHex []byte   `protobuf:"bytes,3,opt,name=ScriptPubKeyHex,json=scriptPubKeyHex,proto3" json:"ScriptPubKeyHex,omitempty"`
	Addresses       []string `protobuf:"bytes,4,rep,name=Addresses,json=addresses" json:"Addresses,omitempty"`
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

type ProtoTransaction_AssetAllocationType struct {
	AssetAllocationTuple         *ProtoTransaction_AssetAllocationType_AssetAllocationTupleType `protobuf:"bytes,1,opt,name=AssetAllocationTuple,json=assetAllocationTuple" json:"AssetAllocationTuple,omitempty"`
	ListSendingAllocationAmounts []*ProtoTransaction_AssetAllocationType_RangeAmountPairType    `protobuf:"bytes,2,rep,name=ListSendingAllocationAmounts,json=listSendingAllocationAmounts" json:"ListSendingAllocationAmounts,omitempty"`
}

func (m *ProtoTransaction_AssetAllocationType) Reset()         { *m = ProtoTransaction_AssetAllocationType{} }
func (m *ProtoTransaction_AssetAllocationType) String() string { return proto.CompactTextString(m) }
func (*ProtoTransaction_AssetAllocationType) ProtoMessage()    {}
func (*ProtoTransaction_AssetAllocationType) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2}
}

func (m *ProtoTransaction_AssetAllocationType) GetAssetAllocationTuple() *ProtoTransaction_AssetAllocationType_AssetAllocationTupleType {
	if m != nil {
		return m.AssetAllocationTuple
	}
	return nil
}

func (m *ProtoTransaction_AssetAllocationType) GetListSendingAllocationAmounts() []*ProtoTransaction_AssetAllocationType_RangeAmountPairType {
	if m != nil {
		return m.ListSendingAllocationAmounts
	}
	return nil
}

type ProtoTransaction_AssetAllocationType_WitnessAddressType struct {
	Version        uint32 `protobuf:"varint,1,opt,name=Version,json=version" json:"Version,omitempty"`
	WitnessProgram string `protobuf:"bytes,2,opt,name=WitnessProgram,json=witnessProgram" json:"WitnessProgram,omitempty"`
}

func (m *ProtoTransaction_AssetAllocationType_WitnessAddressType) Reset() {
	*m = ProtoTransaction_AssetAllocationType_WitnessAddressType{}
}
func (m *ProtoTransaction_AssetAllocationType_WitnessAddressType) String() string {
	return proto.CompactTextString(m)
}
func (*ProtoTransaction_AssetAllocationType_WitnessAddressType) ProtoMessage() {}
func (*ProtoTransaction_AssetAllocationType_WitnessAddressType) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2, 0}
}

func (m *ProtoTransaction_AssetAllocationType_WitnessAddressType) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ProtoTransaction_AssetAllocationType_WitnessAddressType) GetWitnessProgram() string {
	if m != nil {
		return m.WitnessProgram
	}
	return ""
}

type ProtoTransaction_AssetAllocationType_RangeAmountPairType struct {
	WitnessAddress *ProtoTransaction_AssetAllocationType_WitnessAddressType `protobuf:"bytes,1,opt,name=WitnessAddress,json=witnessAddress" json:"WitnessAddress,omitempty"`
	ValueSat       []byte                                                   `protobuf:"bytes,2,opt,name=ValueSat,json=valueSat,proto3" json:"ValueSat,omitempty"`
}

func (m *ProtoTransaction_AssetAllocationType_RangeAmountPairType) Reset() {
	*m = ProtoTransaction_AssetAllocationType_RangeAmountPairType{}
}
func (m *ProtoTransaction_AssetAllocationType_RangeAmountPairType) String() string {
	return proto.CompactTextString(m)
}
func (*ProtoTransaction_AssetAllocationType_RangeAmountPairType) ProtoMessage() {}
func (*ProtoTransaction_AssetAllocationType_RangeAmountPairType) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2, 1}
}

func (m *ProtoTransaction_AssetAllocationType_RangeAmountPairType) GetWitnessAddress() *ProtoTransaction_AssetAllocationType_WitnessAddressType {
	if m != nil {
		return m.WitnessAddress
	}
	return nil
}

func (m *ProtoTransaction_AssetAllocationType_RangeAmountPairType) GetValueSat() []byte {
	if m != nil {
		return m.ValueSat
	}
	return nil
}

type ProtoTransaction_AssetAllocationType_AssetAllocationTupleType struct {
	Asset          uint32                                                   `protobuf:"varint,1,opt,name=Asset,json=asset" json:"Asset,omitempty"`
	WitnessAddress *ProtoTransaction_AssetAllocationType_WitnessAddressType `protobuf:"bytes,2,opt,name=WitnessAddress,json=witnessAddress" json:"WitnessAddress,omitempty"`
}

func (m *ProtoTransaction_AssetAllocationType_AssetAllocationTupleType) Reset() {
	*m = ProtoTransaction_AssetAllocationType_AssetAllocationTupleType{}
}
func (m *ProtoTransaction_AssetAllocationType_AssetAllocationTupleType) String() string {
	return proto.CompactTextString(m)
}
func (*ProtoTransaction_AssetAllocationType_AssetAllocationTupleType) ProtoMessage() {}
func (*ProtoTransaction_AssetAllocationType_AssetAllocationTupleType) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2, 2}
}

func (m *ProtoTransaction_AssetAllocationType_AssetAllocationTupleType) GetAsset() uint32 {
	if m != nil {
		return m.Asset
	}
	return 0
}

func (m *ProtoTransaction_AssetAllocationType_AssetAllocationTupleType) GetWitnessAddress() *ProtoTransaction_AssetAllocationType_WitnessAddressType {
	if m != nil {
		return m.WitnessAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoTransaction)(nil), "bchain.ProtoTransaction")
	proto.RegisterType((*ProtoTransaction_VinType)(nil), "bchain.ProtoTransaction.VinType")
	proto.RegisterType((*ProtoTransaction_VoutType)(nil), "bchain.ProtoTransaction.VoutType")
	proto.RegisterType((*ProtoTransaction_AssetAllocationType)(nil), "bchain.ProtoTransaction.AssetAllocationType")
	proto.RegisterType((*ProtoTransaction_AssetAllocationType_WitnessAddressType)(nil), "bchain.ProtoTransaction.AssetAllocationType.WitnessAddressType")
	proto.RegisterType((*ProtoTransaction_AssetAllocationType_RangeAmountPairType)(nil), "bchain.ProtoTransaction.AssetAllocationType.RangeAmountPairType")
	proto.RegisterType((*ProtoTransaction_AssetAllocationType_AssetAllocationTupleType)(nil), "bchain.ProtoTransaction.AssetAllocationType.AssetAllocationTupleType")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0xaf, 0xd2, 0x40,
	0x14, 0xcd, 0xd0, 0x02, 0xed, 0x7d, 0xbc, 0x8f, 0xcc, 0x23, 0xa6, 0x69, 0x58, 0xd4, 0xb7, 0x30,
	0x5d, 0xb1, 0xc0, 0xb8, 0x56, 0x34, 0x26, 0x2f, 0xf1, 0xc5, 0x90, 0x81, 0xd4, 0xf5, 0x50, 0x26,
	0x65, 0x62, 0x99, 0xc1, 0xce, 0x14, 0x21, 0x71, 0xe9, 0x9f, 0xd0, 0x5f, 0x60, 0xe2, 0xff, 0xf2,
	0x77, 0x98, 0x4e, 0x07, 0x04, 0x04, 0x93, 0xb7, 0x70, 0x79, 0xcf, 0xed, 0x9d, 0x73, 0xee, 0x39,
	0x37, 0x05, 0x4f, 0xaf, 0xfb, 0xcb, 0x42, 0x6a, 0x89, 0x5b, 0xd3, 0x74, 0x4e, 0xb9, 0xb8, 0xfb,
	0xe5, 0xc3, 0xcd, 0xa8, 0x42, 0x26, 0x05, 0x15, 0x8a, 0xa6, 0x9a, 0x4b, 0x81, 0x31, 0xb8, 0x93,
	0x35, 0x9f, 0x05, 0x28, 0x42, 0x71, 0x87, 0xb8, 0x7a, 0xcd, 0x67, 0xf8, 0x06, 0x9c, 0x7b, 0xb6,
	0x0e, 0x1a, 0x06, 0x72, 0xe6, 0x6c, 0x8d, 0x7b, 0xe0, 0xbf, 0xce, 0x65, 0xfa, 0x51, 0xf3, 0x05,
	0x0b, 0x9c, 0x08, 0xc5, 0x2e, 0xf1, 0xa7, 0x5b, 0x00, 0x87, 0xe0, 0x3d, 0x6c, 0x9b, 0x6e, 0x84,
	0xe2, 0x4b, 0xe2, 0xed, 0x7a, 0x4f, 0xa0, 0x75, 0xcf, 0x78, 0x36, 0xd7, 0x41, 0xd3, 0x74, 0x5a,
	0x73, 0x53, 0xe1, 0x01, 0x38, 0x09, 0x17, 0x41, 0x2b, 0x72, 0xe2, 0x8b, 0x41, 0xd4, 0xaf, 0x25,
	0xf6, 0x8f, 0xe5, 0xf5, 0x13, 0x2e, 0x26, 0x9b, 0x25, 0x23, 0xce, 0x8a, 0x0b, 0xfc, 0x02, 0xdc,
	0x44, 0x96, 0x3a, 0x68, 0x9b, 0xa1, 0xa7, 0xe7, 0x87, 0x64, 0xa9, 0xcd, 0x94, 0xbb, 0x92, 0xa5,
	0xc6, 0x01, 0xb4, 0x13, 0x56, 0x28, 0x2e, 0x45, 0xe0, 0x45, 0x28, 0x6e, 0x92, 0xf6, 0xaa, 0x2e,
	0xc3, 0x9f, 0x08, 0xda, 0x96, 0xa1, 0x5a, 0xe2, 0x8d, 0xe4, 0x62, 0x4a, 0x15, 0x33, 0x66, 0xf8,
	0xc4, 0x4b, 0x6d, 0xbd, 0x33, 0xa9, 0xb1, 0x67, 0x12, 0xb6, 0x62, 0x1c, 0xb3, 0x56, 0xcd, 0x74,
	0x07, 0x9d, 0x71, 0x5a, 0xf0, 0xa5, 0x1e, 0xf3, 0xac, 0x72, 0xd0, 0x35, 0xdf, 0x77, 0xd4, 0x1e,
	0x56, 0xf1, 0x8c, 0xd9, 0xa7, 0x92, 0x89, 0x94, 0x59, 0x4b, 0x3c, 0x65, 0xeb, 0xca, 0xe6, 0xe1,
	0x6c, 0x56, 0x30, 0xa5, 0x98, 0x32, 0xd6, 0xf8, 0xc4, 0xa7, 0x5b, 0x20, 0xfc, 0x02, 0xde, 0x76,
	0xb3, 0xea, 0x95, 0x84, 0xe6, 0x25, 0x1b, 0x53, 0x6d, 0xa3, 0xf3, 0x56, 0xb6, 0xc6, 0x1d, 0x40,
	0xef, 0x8d, 0xd4, 0x4b, 0x82, 0x04, 0x8e, 0xe1, 0xba, 0xd6, 0x34, 0x2a, 0xa7, 0xef, 0xd8, 0xa6,
	0x92, 0xe5, 0x98, 0x81, 0x6b, 0x75, 0x08, 0x1f, 0xb2, 0xbb, 0xc7, 0xec, 0x3f, 0x9a, 0x70, 0x3b,
	0x54, 0x8a, 0xe9, 0x61, 0x9e, 0xcb, 0x94, 0x56, 0x46, 0x1b, 0x25, 0x1b, 0xe8, 0x1e, 0xc3, 0xe5,
	0x32, 0xaf, 0x3d, 0xbc, 0x18, 0xbc, 0x3d, 0x1b, 0xd2, 0x89, 0xb7, 0xfa, 0xa7, 0x1e, 0x32, 0x41,
	0x76, 0xe9, 0x89, 0x0e, 0xfe, 0x8a, 0xa0, 0xf7, 0xc0, 0x95, 0x1e, 0x33, 0x31, 0xe3, 0x22, 0xfb,
	0xd3, 0x1e, 0x2e, 0x64, 0x29, 0xb4, 0x0a, 0x1a, 0xe6, 0x50, 0x5e, 0x3d, 0x4a, 0x03, 0xa1, 0x22,
	0x63, 0xf5, 0x03, 0x23, 0xca, 0x0b, 0x43, 0xdf, 0xcb, 0xff, 0xc1, 0x12, 0x26, 0x80, 0x3f, 0x70,
	0x2d, 0x98, 0x52, 0xd6, 0x3e, 0xe3, 0xcb, 0xde, 0xd5, 0x21, 0x93, 0xc5, 0xf6, 0xea, 0xf0, 0x33,
	0xb8, 0xb2, 0xdf, 0x8f, 0x0a, 0x99, 0x15, 0x74, 0x61, 0xc2, 0xf2, 0xc9, 0xd5, 0xe7, 0x03, 0x34,
	0xfc, 0x8e, 0xe0, 0xf6, 0x84, 0x1a, 0x9c, 0xed, 0xe6, 0x2d, 0x9f, 0xf5, 0xfa, 0xe5, 0xa3, 0xf6,
	0xfc, 0x5b, 0xf2, 0x4e, 0x80, 0xc5, 0x0e, 0x8e, 0xac, 0x71, 0x78, 0x64, 0xe1, 0x37, 0x04, 0xc1,
	0xb9, 0xb8, 0x70, 0x17, 0x9a, 0xa6, 0x67, 0x37, 0x6f, 0x9a, 0xf4, 0x4e, 0xe8, 0x6e, 0xfc, 0x17,
	0xdd, 0xd3, 0x96, 0xf9, 0xef, 0x3d, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x25, 0xca, 0x6e, 0x65,
	0x03, 0x05, 0x00, 0x00,
}