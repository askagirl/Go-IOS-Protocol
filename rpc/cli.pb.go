package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Transaction struct {
	Tx                   []byte   `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}


func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_7f181ac8c7362e80, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_7f181ac8c7362e80, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}

func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type TransactionKey struct {
	Publisher            string   `protobuf:"bytes,1,opt,name=publisher" json:"publisher,omitempty"`
	Nonce                int64    `protobuf:"varint,2,opt,name=nonce" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionKey) Reset()         { *m = TransactionKey{} }
func (m *TransactionKey) String() string { return proto.CompactTextString(m) }

func (*TransactionKey) ProtoMessage()    {}
func (*TransactionKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_7f181ac8c7362e80, []int{2}
}
func (m *TransactionKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionKey.Unmarshal(m, b)
}
func (m *TransactionKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionKey.Marshal(b, m, deterministic)
}
func (dst *TransactionKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionKey.Merge(dst, src)
}
func (m *TransactionKey) XXX_Size() int {
	return xxx_messageInfo_TransactionKey.Size(m)
}
func (m *TransactionKey) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionKey.DiscardUnknown(m)
}
var xxx_messageInfo_TransactionKey proto.InternalMessageInfo

func (m *TransactionKey) GetPublisher() string {
if m != nil {
return m.Publisher
}
return ""
}

func (m *TransactionKey) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type Key struct {
	S                    string   `protobuf:"bytes,1,opt,name=s" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_7f181ac8c7362e80, []int{3}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (dst *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(dst, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CliClient is the client API for Cli service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CliClient interface {
	PublishTx(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Response, error)
	GetTransaction(ctx context.Context, in *TransactionKey, opts ...grpc.CallOption) (*Transaction, error)
	GetBalance(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	GetState(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	GetBlock(ctx context.Context, in *BlockKey, opts ...grpc.CallOption) (*BlockInfo, error)
}

// Server API for Cli service
type CliServer interface {
	PublishTx(context.Context, *Transaction) (*Response, error)
	GetTransaction(context.Context, *TransactionKey) (*Transaction, error)
	GetBalance(context.Context, *Key) (*Value, error)
	GetState(context.Context, *Key) (*Value, error)
	GetBlock(context.Context, *BlockKey) (*BlockInfo, error)
}


var _Cli_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Cli",
	HandlerType: (*CliServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishTx",
			Handler:    _Cli_PublishTx_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Cli_GetTransaction_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Cli_GetBalance_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _Cli_GetState_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _Cli_GetBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cli.proto",
}

func init() { proto.RegisterFile("cli.proto", fileDescriptor_cli_7f181ac8c7362e80) }

var fileDescriptor_cli_7f181ac8c7362e80 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0xf3, 0xf2, 0xcf, 0x9f, 0x64, 0xac, 0x41, 0xa6, 0x05, 0x4b, 0xf1, 0xa5, 0x2c, 0x1e,
	0x0a, 0x42, 0x0e, 0x7a, 0xf2, 0x5a, 0x85, 0x22, 0xbd, 0xc8, 0x5a, 0xbc, 0x6f, 0xd7, 0x15, 0xa3,
	0xcb, 0x6e, 0xc8, 0x6e, 0x4b, 0xfb, 0x89, 0xfd, 0x1a, 0x92, 0x49, 0x4b, 0x83, 0xe2, 0x6d, 0x66,
	0xf2, 0xfc, 0x26, 0xcf, 0x33, 0x0b, 0x99, 0xd4, 0x65, 0x51, 0xd5, 0xd6, 0x5b, 0x8c, 0xeb, 0x4a,
	0xb2, 0x73, 0x38, 0x5a, 0xd4, 0xc2, 0x38, 0x21, 0x7d, 0x69, 0x0d, 0xe6, 0x10, 0xf9, 0xcd, 0x30,
	0x1c, 0x87, 0x93, 0x1e, 0x8f, 0xfc, 0x86, 0x5d, 0x40, 0xca, 0x95, 0xab, 0xac, 0x71, 0x0a, 0x11,
	0xfe, 0x49, 0xfb, 0xaa, 0xe8, 0x6b, 0xc2, 0xa9, 0x66, 0x0f, 0x90, 0x77, 0xf0, 0xb9, 0xda, 0xe2,
	0x19, 0x64, 0xd5, 0x6a, 0xa9, 0x4b, 0xf7, 0xae, 0x6a, 0x92, 0x66, 0xfc, 0x30, 0xc0, 0x01, 0x24,
	0xc6, 0x1a, 0xa9, 0x86, 0xd1, 0x38, 0x9c, 0xc4, 0xbc, 0x6d, 0x58, 0x1f, 0xe2, 0x06, 0xed, 0x41,
	0xe8, 0x76, 0x48, 0xe8, 0xd8, 0x29, 0x24, 0x2f, 0x42, 0xaf, 0x54, 0xe3, 0xc9, 0xad, 0x09, 0xc8,
	0x78, 0xe4, 0xd6, 0x6c, 0x0c, 0xe9, 0x54, 0x5b, 0xf9, 0xd9, 0x20, 0x03, 0x48, 0xb4, 0xd8, 0xee,
	0xfe, 0x14, 0xf3, 0xb6, 0x61, 0x97, 0x90, 0x91, 0xe2, 0xd1, 0xbc, 0xd9, 0xc6, 0xf6, 0x87, 0xb3,
	0x66, 0xb7, 0x98, 0xea, 0x9b, 0xaf, 0x10, 0xe2, 0x7b, 0x5d, 0x62, 0x01, 0xd9, 0x53, 0xeb, 0x6d,
	0xb1, 0xc1, 0x93, 0xa2, 0xae, 0x64, 0xd1, 0x89, 0x33, 0x3a, 0xa6, 0xc9, 0xfe, 0x00, 0x2c, 0xc0,
	0x3b, 0xc8, 0x67, 0xca, 0x77, 0x0f, 0xd6, 0xff, 0x09, 0xcd, 0xd5, 0x76, 0xf4, 0x6b, 0x13, 0x0b,
	0xf0, 0x0a, 0x60, 0xa6, 0xfc, 0x54, 0x68, 0x61, 0xa4, 0xc2, 0x94, 0x14, 0x8d, 0x16, 0xa8, 0xa2,
	0xa4, 0x2c, 0x40, 0x06, 0xe9, 0x4c, 0xf9, 0x67, 0x2f, 0xfc, 0xdf, 0x9a, 0x6b, 0xd2, 0x50, 0x40,
	0x6c, 0x1d, 0xee, 0xcf, 0x31, 0xca, 0x0f, 0x6d, 0x93, 0x9d, 0x05, 0xcb, 0xff, 0xf4, 0xd6, 0xb7,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0xb9, 0xca, 0xed, 0xf8, 0x01, 0x00, 0x00,
}