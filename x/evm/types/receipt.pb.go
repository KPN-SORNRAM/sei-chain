// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/receipt.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Log struct {
	Address string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Topics  []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	Data    []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_d864f6bdca684f52, []int{0}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Log.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return m.Size()
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Log) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *Log) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Receipt struct {
	TxType            uint32 `protobuf:"varint,1,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty" yaml:"tx_type"`
	CumulativeGasUsed uint64 `protobuf:"varint,2,opt,name=cumulative_gas_used,json=cumulativeGasUsed,proto3" json:"cumulative_gas_used,omitempty" yaml:"cumulative_gas_used"`
	ContractAddress   string `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`
	TxHashHex         string `protobuf:"bytes,4,opt,name=tx_hash_hex,json=txHashHex,proto3" json:"tx_hash_hex,omitempty" yaml:"tx_hash_hex"`
	GasUsed           uint64 `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty" yaml:"gas_used"`
	EffectiveGasPrice uint64 `protobuf:"varint,6,opt,name=effective_gas_price,json=effectiveGasPrice,proto3" json:"effective_gas_price,omitempty" yaml:"effective_gas_price"`
	BlockNumber       uint64 `protobuf:"varint,7,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty" yaml:"block_number"`
	TransactionIndex  uint32 `protobuf:"varint,8,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index,omitempty" yaml:"transaction_index"`
	Status            uint32 `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty" yaml:"status"`
	From              string `protobuf:"bytes,10,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	To                string `protobuf:"bytes,11,opt,name=to,proto3" json:"to,omitempty" yaml:"to"`
	VmError           string `protobuf:"bytes,12,opt,name=vm_error,json=vmError,proto3" json:"vm_error,omitempty" yaml:"vm_error"`
	Logs              []*Log `protobuf:"bytes,13,rep,name=logs,proto3" json:"logs,omitempty"`
	LogsBloom         []byte `protobuf:"bytes,14,opt,name=logsBloom,proto3" json:"logsBloom,omitempty"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_d864f6bdca684f52, []int{1}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return m.Size()
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetTxType() uint32 {
	if m != nil {
		return m.TxType
	}
	return 0
}

func (m *Receipt) GetCumulativeGasUsed() uint64 {
	if m != nil {
		return m.CumulativeGasUsed
	}
	return 0
}

func (m *Receipt) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Receipt) GetTxHashHex() string {
	if m != nil {
		return m.TxHashHex
	}
	return ""
}

func (m *Receipt) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *Receipt) GetEffectiveGasPrice() uint64 {
	if m != nil {
		return m.EffectiveGasPrice
	}
	return 0
}

func (m *Receipt) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *Receipt) GetTransactionIndex() uint32 {
	if m != nil {
		return m.TransactionIndex
	}
	return 0
}

func (m *Receipt) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Receipt) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Receipt) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Receipt) GetVmError() string {
	if m != nil {
		return m.VmError
	}
	return ""
}

func (m *Receipt) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Receipt) GetLogsBloom() []byte {
	if m != nil {
		return m.LogsBloom
	}
	return nil
}

func init() {
	proto.RegisterType((*Log)(nil), "seiprotocol.seichain.evm.Log")
	proto.RegisterType((*Receipt)(nil), "seiprotocol.seichain.evm.Receipt")
}

func init() { proto.RegisterFile("evm/receipt.proto", fileDescriptor_d864f6bdca684f52) }

var fileDescriptor_d864f6bdca684f52 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x41, 0x4f, 0xdb, 0x4c,
	0x10, 0xc5, 0x49, 0x48, 0xf0, 0x86, 0x00, 0xd9, 0x7c, 0x82, 0x15, 0x1f, 0xd8, 0xd1, 0xf6, 0x92,
	0xaa, 0xc2, 0x51, 0x5b, 0xa9, 0x07, 0x6e, 0x8d, 0xd4, 0x02, 0x2a, 0x42, 0xd5, 0xaa, 0xbd, 0xf4,
	0x62, 0x6d, 0x36, 0x8b, 0x63, 0xd5, 0xf6, 0x5a, 0xde, 0x75, 0x64, 0xfe, 0x45, 0x7f, 0x56, 0x8f,
	0x1c, 0x7b, 0xb2, 0x2a, 0xf8, 0x07, 0xbe, 0x57, 0xaa, 0xbc, 0xb6, 0x03, 0x42, 0xf4, 0x94, 0x99,
	0xf7, 0xde, 0xcc, 0xce, 0x9b, 0x78, 0xc0, 0x90, 0xaf, 0xc2, 0x69, 0xc2, 0x19, 0xf7, 0x63, 0xe5,
	0xc4, 0x89, 0x50, 0x02, 0x22, 0xc9, 0x7d, 0x1d, 0x31, 0x11, 0x38, 0x92, 0xfb, 0x6c, 0x49, 0xfd,
	0xc8, 0xe1, 0xab, 0xf0, 0xf0, 0x3f, 0x4f, 0x78, 0x42, 0x53, 0xd3, 0x32, 0xaa, 0xf4, 0xf8, 0x13,
	0x68, 0x5f, 0x0a, 0x0f, 0x22, 0xd0, 0xa3, 0x8b, 0x45, 0xc2, 0xa5, 0x44, 0xc6, 0xd8, 0x98, 0x98,
	0xa4, 0x49, 0xe1, 0x3e, 0xe8, 0x2a, 0x11, 0xfb, 0x4c, 0xa2, 0xd6, 0xb8, 0x3d, 0x31, 0x49, 0x9d,
	0x41, 0x08, 0x3a, 0x0b, 0xaa, 0x28, 0x6a, 0x8f, 0x8d, 0xc9, 0x36, 0xd1, 0x31, 0xfe, 0xb3, 0x09,
	0x7a, 0xa4, 0x1a, 0x07, 0xbe, 0x02, 0x3d, 0x95, 0xb9, 0xea, 0x26, 0xe6, 0xba, 0xe3, 0x60, 0x06,
	0x8b, 0xdc, 0xde, 0xb9, 0xa1, 0x61, 0x70, 0x8a, 0x6b, 0x02, 0x93, 0xae, 0xca, 0xbe, 0xdc, 0xc4,
	0x1c, 0x5e, 0x81, 0x11, 0x4b, 0xc3, 0x34, 0xa0, 0xca, 0x5f, 0x71, 0xd7, 0xa3, 0xd2, 0x4d, 0x25,
	0x5f, 0xa0, 0xd6, 0xd8, 0x98, 0x74, 0x66, 0x56, 0x91, 0xdb, 0x87, 0x55, 0xe1, 0x33, 0x22, 0x4c,
	0x86, 0x0f, 0xe8, 0x19, 0x95, 0x5f, 0x25, 0x5f, 0xc0, 0x8f, 0x60, 0x8f, 0x89, 0x48, 0x25, 0x94,
	0x29, 0xb7, 0xf1, 0x55, 0x0e, 0x6a, 0xce, 0xfe, 0x2f, 0x72, 0xfb, 0xa0, 0x6e, 0xf6, 0x44, 0x81,
	0xc9, 0x6e, 0x03, 0xbd, 0xaf, 0xcd, 0xbf, 0x03, 0x7d, 0x95, 0xb9, 0x4b, 0x2a, 0x97, 0xee, 0x92,
	0x67, 0xa8, 0xa3, 0x5b, 0xec, 0x17, 0xb9, 0x0d, 0xd7, 0x46, 0x1a, 0x12, 0x13, 0x53, 0x65, 0xe7,
	0x54, 0x2e, 0xcf, 0x79, 0x06, 0x1d, 0xb0, 0xb5, 0x36, 0xb1, 0xa9, 0x4d, 0x8c, 0x8a, 0xdc, 0xde,
	0xad, 0x8a, 0x1e, 0x26, 0xef, 0x79, 0xf5, 0xbc, 0x57, 0x60, 0xc4, 0xaf, 0xaf, 0x39, 0x5b, 0x3b,
	0x8b, 0x13, 0x9f, 0x71, 0xd4, 0x7d, 0xea, 0xff, 0x19, 0x11, 0x26, 0xc3, 0x35, 0x7a, 0x46, 0xe5,
	0xe7, 0x12, 0x83, 0xa7, 0x60, 0x7b, 0x1e, 0x08, 0xf6, 0xdd, 0x8d, 0xd2, 0x70, 0xce, 0x13, 0xd4,
	0xd3, 0x8d, 0x0e, 0x8a, 0xdc, 0x1e, 0x55, 0x8d, 0x1e, 0xb3, 0x98, 0xf4, 0x75, 0x7a, 0xa5, 0x33,
	0x78, 0x01, 0x86, 0x2a, 0xa1, 0x91, 0xa4, 0x4c, 0xf9, 0x22, 0x72, 0xfd, 0x68, 0xc1, 0x33, 0xb4,
	0xa5, 0xff, 0xc2, 0xa3, 0x22, 0xb7, 0x51, 0xed, 0xfc, 0xa9, 0x04, 0x93, 0xbd, 0x47, 0xd8, 0x45,
	0x09, 0xc1, 0x97, 0xa0, 0x2b, 0x15, 0x55, 0xa9, 0x44, 0xa6, 0xae, 0x1f, 0x16, 0xb9, 0x3d, 0xa8,
	0xea, 0x2b, 0x1c, 0x93, 0x5a, 0x00, 0x5f, 0x80, 0xce, 0x75, 0x22, 0x42, 0x04, 0xf4, 0x8a, 0x77,
	0x8b, 0xdc, 0xee, 0x57, 0xc2, 0x12, 0xc5, 0x44, 0x93, 0xf0, 0x18, 0xb4, 0x94, 0x40, 0x7d, 0x2d,
	0x19, 0x14, 0xb9, 0x6d, 0xd6, 0xb3, 0x08, 0x4c, 0x5a, 0x4a, 0x94, 0x5b, 0x5f, 0x85, 0x2e, 0x4f,
	0x12, 0x91, 0xa0, 0x6d, 0x2d, 0x7a, 0xb4, 0xf5, 0x86, 0xc1, 0xa4, 0xb7, 0x0a, 0x3f, 0x94, 0x11,
	0x7c, 0x0d, 0x3a, 0x81, 0xf0, 0x24, 0x1a, 0x8c, 0xdb, 0x93, 0xfe, 0x9b, 0x63, 0xe7, 0x5f, 0xa7,
	0xe3, 0x5c, 0x0a, 0x8f, 0x68, 0x29, 0x3c, 0x02, 0x66, 0xf9, 0x3b, 0x0b, 0x84, 0x08, 0xd1, 0x8e,
	0xfe, 0xf4, 0x1f, 0x80, 0xd9, 0xd9, 0xcf, 0x3b, 0xcb, 0xb8, 0xbd, 0xb3, 0x8c, 0xdf, 0x77, 0x96,
	0xf1, 0xe3, 0xde, 0xda, 0xb8, 0xbd, 0xb7, 0x36, 0x7e, 0xdd, 0x5b, 0x1b, 0xdf, 0x4e, 0x3c, 0x5f,
	0x2d, 0xd3, 0xb9, 0xc3, 0x44, 0x38, 0x95, 0xdc, 0x3f, 0x69, 0xde, 0xd1, 0x89, 0x7e, 0x68, 0x9a,
	0x4d, 0xcb, 0x6b, 0x2e, 0xef, 0x42, 0xce, 0xbb, 0x9a, 0x7f, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xfd, 0xe9, 0x49, 0x2e, 0xe1, 0x03, 0x00, 0x00,
}

func (m *Log) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Log) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Topics) > 0 {
		for iNdEx := len(m.Topics) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Topics[iNdEx])
			copy(dAtA[i:], m.Topics[iNdEx])
			i = encodeVarintReceipt(dAtA, i, uint64(len(m.Topics[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Receipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Receipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LogsBloom) > 0 {
		i -= len(m.LogsBloom)
		copy(dAtA[i:], m.LogsBloom)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.LogsBloom)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.Logs) > 0 {
		for iNdEx := len(m.Logs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Logs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReceipt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.VmError) > 0 {
		i -= len(m.VmError)
		copy(dAtA[i:], m.VmError)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.VmError)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x52
	}
	if m.Status != 0 {
		i = encodeVarintReceipt(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x48
	}
	if m.TransactionIndex != 0 {
		i = encodeVarintReceipt(dAtA, i, uint64(m.TransactionIndex))
		i--
		dAtA[i] = 0x40
	}
	if m.BlockNumber != 0 {
		i = encodeVarintReceipt(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x38
	}
	if m.EffectiveGasPrice != 0 {
		i = encodeVarintReceipt(dAtA, i, uint64(m.EffectiveGasPrice))
		i--
		dAtA[i] = 0x30
	}
	if m.GasUsed != 0 {
		i = encodeVarintReceipt(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x28
	}
	if len(m.TxHashHex) > 0 {
		i -= len(m.TxHashHex)
		copy(dAtA[i:], m.TxHashHex)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.TxHashHex)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.CumulativeGasUsed != 0 {
		i = encodeVarintReceipt(dAtA, i, uint64(m.CumulativeGasUsed))
		i--
		dAtA[i] = 0x10
	}
	if m.TxType != 0 {
		i = encodeVarintReceipt(dAtA, i, uint64(m.TxType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintReceipt(dAtA []byte, offset int, v uint64) int {
	offset -= sovReceipt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Log) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	if len(m.Topics) > 0 {
		for _, s := range m.Topics {
			l = len(s)
			n += 1 + l + sovReceipt(uint64(l))
		}
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	return n
}

func (m *Receipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxType != 0 {
		n += 1 + sovReceipt(uint64(m.TxType))
	}
	if m.CumulativeGasUsed != 0 {
		n += 1 + sovReceipt(uint64(m.CumulativeGasUsed))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	l = len(m.TxHashHex)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	if m.GasUsed != 0 {
		n += 1 + sovReceipt(uint64(m.GasUsed))
	}
	if m.EffectiveGasPrice != 0 {
		n += 1 + sovReceipt(uint64(m.EffectiveGasPrice))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovReceipt(uint64(m.BlockNumber))
	}
	if m.TransactionIndex != 0 {
		n += 1 + sovReceipt(uint64(m.TransactionIndex))
	}
	if m.Status != 0 {
		n += 1 + sovReceipt(uint64(m.Status))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	l = len(m.VmError)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovReceipt(uint64(l))
		}
	}
	l = len(m.LogsBloom)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	return n
}

func sovReceipt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReceipt(x uint64) (n int) {
	return sovReceipt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Log) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReceipt
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Log: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Log: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topics", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topics = append(m.Topics, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReceipt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReceipt
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Receipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReceipt
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Receipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			m.TxType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxType |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeGasUsed", wireType)
			}
			m.CumulativeGasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CumulativeGasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHashHex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHashHex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EffectiveGasPrice", wireType)
			}
			m.EffectiveGasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EffectiveGasPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionIndex", wireType)
			}
			m.TransactionIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VmError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VmError = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &Log{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogsBloom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogsBloom = append(m.LogsBloom[:0], dAtA[iNdEx:postIndex]...)
			if m.LogsBloom == nil {
				m.LogsBloom = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReceipt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReceipt
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipReceipt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReceipt
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthReceipt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReceipt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReceipt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReceipt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReceipt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReceipt = fmt.Errorf("proto: unexpected end of group")
)
