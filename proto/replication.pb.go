// Code generated by protoc-gen-go. DO NOT EDIT.
// source: replication.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type WriteAheadLog struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteAheadLog) Reset()         { *m = WriteAheadLog{} }
func (m *WriteAheadLog) String() string { return proto.CompactTextString(m) }
func (*WriteAheadLog) ProtoMessage()    {}
func (*WriteAheadLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed0454e9e09fb71a, []int{0}
}

func (m *WriteAheadLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteAheadLog.Unmarshal(m, b)
}
func (m *WriteAheadLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteAheadLog.Marshal(b, m, deterministic)
}
func (m *WriteAheadLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteAheadLog.Merge(m, src)
}
func (m *WriteAheadLog) XXX_Size() int {
	return xxx_messageInfo_WriteAheadLog.Size(m)
}
func (m *WriteAheadLog) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteAheadLog.DiscardUnknown(m)
}

var xxx_messageInfo_WriteAheadLog proto.InternalMessageInfo

type GetWALStreamRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWALStreamRequest) Reset()         { *m = GetWALStreamRequest{} }
func (m *GetWALStreamRequest) String() string { return proto.CompactTextString(m) }
func (*GetWALStreamRequest) ProtoMessage()    {}
func (*GetWALStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed0454e9e09fb71a, []int{1}
}

func (m *GetWALStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWALStreamRequest.Unmarshal(m, b)
}
func (m *GetWALStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWALStreamRequest.Marshal(b, m, deterministic)
}
func (m *GetWALStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWALStreamRequest.Merge(m, src)
}
func (m *GetWALStreamRequest) XXX_Size() int {
	return xxx_messageInfo_GetWALStreamRequest.Size(m)
}
func (m *GetWALStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWALStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWALStreamRequest proto.InternalMessageInfo

type GetWALStreamResponse struct {
	TransactionGroup     []byte   `protobuf:"bytes,1,opt,name=transaction_group,json=transactionGroup,proto3" json:"transaction_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWALStreamResponse) Reset()         { *m = GetWALStreamResponse{} }
func (m *GetWALStreamResponse) String() string { return proto.CompactTextString(m) }
func (*GetWALStreamResponse) ProtoMessage()    {}
func (*GetWALStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed0454e9e09fb71a, []int{2}
}

func (m *GetWALStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWALStreamResponse.Unmarshal(m, b)
}
func (m *GetWALStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWALStreamResponse.Marshal(b, m, deterministic)
}
func (m *GetWALStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWALStreamResponse.Merge(m, src)
}
func (m *GetWALStreamResponse) XXX_Size() int {
	return xxx_messageInfo_GetWALStreamResponse.Size(m)
}
func (m *GetWALStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWALStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWALStreamResponse proto.InternalMessageInfo

func (m *GetWALStreamResponse) GetTransactionGroup() []byte {
	if m != nil {
		return m.TransactionGroup
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteAheadLog)(nil), "proto.WriteAheadLog")
	proto.RegisterType((*GetWALStreamRequest)(nil), "proto.GetWALStreamRequest")
	proto.RegisterType((*GetWALStreamResponse)(nil), "proto.GetWALStreamResponse")
}

func init() {
	proto.RegisterFile("replication.proto", fileDescriptor_ed0454e9e09fb71a)
}

var fileDescriptor_ed0454e9e09fb71a = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4a, 0x2d, 0xc8,
	0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x53, 0x4a, 0xfc, 0x5c, 0xbc, 0xe1, 0x45, 0x99, 0x25, 0xa9, 0x8e, 0x19, 0xa9, 0x89, 0x29, 0x3e,
	0xf9, 0xe9, 0x4a, 0xa2, 0x5c, 0xc2, 0xee, 0xa9, 0x25, 0xe1, 0x8e, 0x3e, 0xc1, 0x25, 0x45, 0xa9,
	0x89, 0xb9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0xce, 0x5c, 0x22, 0xa8, 0xc2, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xda, 0x5c, 0x82, 0x25, 0x45, 0x89, 0x79, 0xc5, 0x89, 0xc9,
	0x20, 0xb3, 0xe3, 0xd3, 0x8b, 0xf2, 0x4b, 0x0b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x04,
	0x90, 0x24, 0xdc, 0x41, 0xe2, 0x46, 0x11, 0x5c, 0xdc, 0x41, 0x08, 0x87, 0x08, 0x79, 0x72, 0xf1,
	0x20, 0x9b, 0x29, 0x24, 0x05, 0x71, 0x9a, 0x1e, 0x16, 0xfb, 0xa5, 0xa4, 0xb1, 0xca, 0x41, 0x1c,
	0x61, 0xc0, 0x98, 0xc4, 0x06, 0x96, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x18, 0x7d,
	0x99, 0xe9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReplicationClient is the client API for Replication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReplicationClient interface {
	GetWALStream(ctx context.Context, in *GetWALStreamRequest, opts ...grpc.CallOption) (Replication_GetWALStreamClient, error)
}

type replicationClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationClient(cc grpc.ClientConnInterface) ReplicationClient {
	return &replicationClient{cc}
}

func (c *replicationClient) GetWALStream(ctx context.Context, in *GetWALStreamRequest, opts ...grpc.CallOption) (Replication_GetWALStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Replication_serviceDesc.Streams[0], "/proto.Replication/GetWALStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &replicationGetWALStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Replication_GetWALStreamClient interface {
	Recv() (*GetWALStreamResponse, error)
	grpc.ClientStream
}

type replicationGetWALStreamClient struct {
	grpc.ClientStream
}

func (x *replicationGetWALStreamClient) Recv() (*GetWALStreamResponse, error) {
	m := new(GetWALStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReplicationServer is the server API for Replication service.
type ReplicationServer interface {
	GetWALStream(*GetWALStreamRequest, Replication_GetWALStreamServer) error
}

// UnimplementedReplicationServer can be embedded to have forward compatible implementations.
type UnimplementedReplicationServer struct {
}

func (*UnimplementedReplicationServer) GetWALStream(req *GetWALStreamRequest, srv Replication_GetWALStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetWALStream not implemented")
}

func RegisterReplicationServer(s *grpc.Server, srv ReplicationServer) {
	s.RegisterService(&_Replication_serviceDesc, srv)
}

func _Replication_GetWALStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetWALStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReplicationServer).GetWALStream(m, &replicationGetWALStreamServer{stream})
}

type Replication_GetWALStreamServer interface {
	Send(*GetWALStreamResponse) error
	grpc.ServerStream
}

type replicationGetWALStreamServer struct {
	grpc.ServerStream
}

func (x *replicationGetWALStreamServer) Send(m *GetWALStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Replication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Replication",
	HandlerType: (*ReplicationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetWALStream",
			Handler:       _Replication_GetWALStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "replication.proto",
}
