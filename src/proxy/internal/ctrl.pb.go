// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/ctrl.proto

package internal

import (
	context "context"
	fmt "fmt"
	wire "github.com/Fantom-foundation/go-lachesis/src/inter/wire"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ID struct {
	Hex                  string   `protobuf:"bytes,1,opt,name=hex,proto3" json:"hex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4c68a24d38d4c7, []int{0}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetHex() string {
	if m != nil {
		return m.Hex
	}
	return ""
}

type Balance struct {
	Amount               uint64   `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4c68a24d38d4c7, []int{1}
}

func (m *Balance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Balance.Unmarshal(m, b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return xxx_messageInfo_Balance.Size(m)
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type TransferRequest struct {
	Nonce                uint64   `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Receiver             *ID      `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount               uint64   `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Until                uint64   `protobuf:"varint,4,opt,name=until,proto3" json:"until,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4c68a24d38d4c7, []int{2}
}

func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (m *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(m, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TransferRequest) GetReceiver() *ID {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *TransferRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransferRequest) GetUntil() uint64 {
	if m != nil {
		return m.Until
	}
	return 0
}

type TransferResponse struct {
	Hex                  string   `protobuf:"bytes,1,opt,name=hex,proto3" json:"hex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferResponse) Reset()         { *m = TransferResponse{} }
func (m *TransferResponse) String() string { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()    {}
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4c68a24d38d4c7, []int{3}
}

func (m *TransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferResponse.Unmarshal(m, b)
}
func (m *TransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferResponse.Marshal(b, m, deterministic)
}
func (m *TransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferResponse.Merge(m, src)
}
func (m *TransferResponse) XXX_Size() int {
	return xxx_messageInfo_TransferResponse.Size(m)
}
func (m *TransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferResponse proto.InternalMessageInfo

func (m *TransferResponse) GetHex() string {
	if m != nil {
		return m.Hex
	}
	return ""
}

type TransactionRequest struct {
	Hex                  string   `protobuf:"bytes,1,opt,name=hex,proto3" json:"hex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4c68a24d38d4c7, []int{4}
}

func (m *TransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionRequest.Unmarshal(m, b)
}
func (m *TransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
}
func (m *TransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *TransactionRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionRequest.Size(m)
}
func (m *TransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func (m *TransactionRequest) GetHex() string {
	if m != nil {
		return m.Hex
	}
	return ""
}

type TransactionResponse struct {
	Txn                  *wire.InternalTransaction `protobuf:"bytes,1,opt,name=txn,proto3" json:"txn,omitempty"`
	Event                *wire.Event               `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Block                *wire.Block               `protobuf:"bytes,3,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4c68a24d38d4c7, []int{5}
}

func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetTxn() *wire.InternalTransaction {
	if m != nil {
		return m.Txn
	}
	return nil
}

func (m *TransactionResponse) GetEvent() *wire.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *TransactionResponse) GetBlock() *wire.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type LogLevel struct {
	Level                string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevel) Reset()         { *m = LogLevel{} }
func (m *LogLevel) String() string { return proto.CompactTextString(m) }
func (*LogLevel) ProtoMessage()    {}
func (*LogLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4c68a24d38d4c7, []int{6}
}

func (m *LogLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevel.Unmarshal(m, b)
}
func (m *LogLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevel.Marshal(b, m, deterministic)
}
func (m *LogLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevel.Merge(m, src)
}
func (m *LogLevel) XXX_Size() int {
	return xxx_messageInfo_LogLevel.Size(m)
}
func (m *LogLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevel.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevel proto.InternalMessageInfo

func (m *LogLevel) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func init() {
	proto.RegisterType((*ID)(nil), "internal.ID")
	proto.RegisterType((*Balance)(nil), "internal.Balance")
	proto.RegisterType((*TransferRequest)(nil), "internal.TransferRequest")
	proto.RegisterType((*TransferResponse)(nil), "internal.TransferResponse")
	proto.RegisterType((*TransactionRequest)(nil), "internal.TransactionRequest")
	proto.RegisterType((*TransactionResponse)(nil), "internal.TransactionResponse")
	proto.RegisterType((*LogLevel)(nil), "internal.LogLevel")
}

func init() { proto.RegisterFile("internal/ctrl.proto", fileDescriptor_af4c68a24d38d4c7) }

var fileDescriptor_af4c68a24d38d4c7 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xdd, 0x6e, 0xd3, 0x40,
	0x10, 0x85, 0xf3, 0xd7, 0x34, 0x4c, 0x90, 0x28, 0x53, 0x88, 0xd2, 0x05, 0xa4, 0x74, 0x85, 0x50,
	0x24, 0x90, 0x83, 0xc2, 0x25, 0x57, 0x54, 0xa9, 0x90, 0x45, 0x05, 0x92, 0x93, 0x17, 0x70, 0xdc,
	0x71, 0xb0, 0xba, 0xdd, 0x0d, 0x9b, 0x4d, 0x08, 0x57, 0xdc, 0xf2, 0x3e, 0xbc, 0x20, 0xda, 0x5d,
	0xff, 0xc4, 0xa5, 0xbd, 0xb1, 0x3c, 0xe7, 0x7c, 0x9a, 0x9d, 0x99, 0x03, 0xa7, 0x99, 0x34, 0xa4,
	0x65, 0x2c, 0x26, 0x89, 0xd1, 0x22, 0x58, 0x6b, 0x65, 0x14, 0xf6, 0x0a, 0x91, 0xbd, 0x58, 0x29,
	0xb5, 0x12, 0x34, 0x71, 0xfa, 0x72, 0x9b, 0x4e, 0xe8, 0x76, 0x6d, 0x7e, 0x79, 0x8c, 0x3d, 0x77,
	0xd8, 0xe4, 0x67, 0xa6, 0xc9, 0x7d, 0xbc, 0xcc, 0x07, 0xd0, 0x0a, 0x67, 0x78, 0x02, 0xed, 0xef,
	0xb4, 0x1f, 0x36, 0x47, 0xcd, 0xf1, 0xa3, 0xc8, 0xfe, 0xf2, 0x73, 0x38, 0xbe, 0x88, 0x45, 0x2c,
	0x13, 0xc2, 0x01, 0x74, 0xe3, 0x5b, 0xb5, 0x95, 0xc6, 0xf9, 0x9d, 0x28, 0xaf, 0xf8, 0x6f, 0x78,
	0xb2, 0xd0, 0xb1, 0xdc, 0xa4, 0xa4, 0x23, 0xfa, 0xb1, 0xa5, 0x8d, 0xc1, 0x67, 0x70, 0x24, 0x95,
	0x4c, 0x28, 0x27, 0x7d, 0x81, 0x63, 0xe8, 0x69, 0x4a, 0x28, 0xdb, 0x91, 0x1e, 0xb6, 0x46, 0xcd,
	0x71, 0x7f, 0xfa, 0x38, 0x28, 0x86, 0x0e, 0xc2, 0x59, 0x54, 0xba, 0x07, 0x4f, 0xb5, 0x0f, 0x9f,
	0xb2, 0x7d, 0xb7, 0xd2, 0x64, 0x62, 0xd8, 0xf1, 0x7d, 0x5d, 0xc1, 0x5f, 0xc3, 0x49, 0x35, 0xc0,
	0x66, 0xad, 0xe4, 0x86, 0xee, 0xd9, 0xe4, 0x0d, 0xa0, 0xa3, 0xe2, 0xc4, 0x64, 0x4a, 0x16, 0x93,
	0xfe, 0xcf, 0xfd, 0x69, 0xc2, 0x69, 0x0d, 0xcc, 0x3b, 0xbe, 0x85, 0xb6, 0xd9, 0x4b, 0x47, 0xf6,
	0xa7, 0x67, 0x81, 0xbb, 0x5d, 0x98, 0x4f, 0x7f, 0xc8, 0x5b, 0x0a, 0xcf, 0xe1, 0x88, 0x76, 0x24,
	0x4d, 0xbe, 0x67, 0xdf, 0xe3, 0x97, 0x56, 0x8a, 0xbc, 0x63, 0x91, 0xa5, 0x50, 0xc9, 0x8d, 0x5b,
	0xb1, 0x44, 0x2e, 0xac, 0x14, 0x79, 0x87, 0x8f, 0xa0, 0x77, 0xa5, 0x56, 0x57, 0xb4, 0x23, 0x61,
	0x57, 0x17, 0xf6, 0x27, 0x1f, 0xd5, 0x17, 0xd3, 0xbf, 0x2d, 0xe8, 0x7c, 0x55, 0xd7, 0x84, 0xef,
	0xa1, 0x3b, 0x27, 0x91, 0x86, 0x33, 0x1c, 0x04, 0x3e, 0xfe, 0xa0, 0x88, 0x3f, 0xb8, 0xb4, 0xf1,
	0xb3, 0xda, 0xad, 0x79, 0x03, 0xdf, 0xc1, 0xf1, 0xdc, 0xc4, 0x37, 0xf4, 0x2d, 0xc5, 0x9a, 0xc5,
	0x9e, 0x56, 0x55, 0x1e, 0x3d, 0x6f, 0xe0, 0x27, 0xdb, 0x5f, 0x5e, 0x2f, 0x14, 0x9e, 0x55, 0xf6,
	0x9d, 0xd8, 0x19, 0xbb, 0xcf, 0xf2, 0xe7, 0xe3, 0x0d, 0xfc, 0x02, 0xf0, 0x99, 0xcc, 0x62, 0x2f,
	0x43, 0x99, 0x2a, 0x7c, 0x79, 0x87, 0xad, 0xc5, 0xc2, 0x5e, 0x3d, 0xe0, 0x96, 0xcd, 0x3e, 0x42,
	0x7f, 0x4e, 0xa6, 0xbc, 0x0e, 0x56, 0x7c, 0xa1, 0xb1, 0x07, 0x0e, 0xc1, 0x1b, 0xcb, 0xae, 0x53,
	0x3e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x8a, 0x96, 0xf2, 0x48, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	SelfID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ID, error)
	StakeOf(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Balance, error)
	SendTo(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	GetTxnInfo(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	SetLogLevel(ctx context.Context, in *LogLevel, opts ...grpc.CallOption) (*empty.Empty, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) SelfID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/internal.Node/SelfID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) StakeOf(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Balance, error) {
	out := new(Balance)
	err := c.cc.Invoke(ctx, "/internal.Node/StakeOf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) SendTo(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, "/internal.Node/SendTo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetTxnInfo(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/internal.Node/GetTxnInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) SetLogLevel(ctx context.Context, in *LogLevel, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/internal.Node/SetLogLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	SelfID(context.Context, *empty.Empty) (*ID, error)
	StakeOf(context.Context, *ID) (*Balance, error)
	SendTo(context.Context, *TransferRequest) (*TransferResponse, error)
	GetTxnInfo(context.Context, *TransactionRequest) (*TransactionResponse, error)
	SetLogLevel(context.Context, *LogLevel) (*empty.Empty, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) SelfID(ctx context.Context, req *empty.Empty) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelfID not implemented")
}
func (*UnimplementedNodeServer) StakeOf(ctx context.Context, req *ID) (*Balance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakeOf not implemented")
}
func (*UnimplementedNodeServer) SendTo(ctx context.Context, req *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTo not implemented")
}
func (*UnimplementedNodeServer) GetTxnInfo(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxnInfo not implemented")
}
func (*UnimplementedNodeServer) SetLogLevel(ctx context.Context, req *LogLevel) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLogLevel not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_SelfID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).SelfID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Node/SelfID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).SelfID(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_StakeOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).StakeOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Node/StakeOf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).StakeOf(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_SendTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).SendTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Node/SendTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).SendTo(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetTxnInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetTxnInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Node/GetTxnInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetTxnInfo(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_SetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLevel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).SetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Node/SetLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).SetLogLevel(ctx, req.(*LogLevel))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SelfID",
			Handler:    _Node_SelfID_Handler,
		},
		{
			MethodName: "StakeOf",
			Handler:    _Node_StakeOf_Handler,
		},
		{
			MethodName: "SendTo",
			Handler:    _Node_SendTo_Handler,
		},
		{
			MethodName: "GetTxnInfo",
			Handler:    _Node_GetTxnInfo_Handler,
		},
		{
			MethodName: "SetLogLevel",
			Handler:    _Node_SetLogLevel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/ctrl.proto",
}
