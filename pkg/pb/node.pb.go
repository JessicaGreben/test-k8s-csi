// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package pb

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

type NodePublishVolumeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodePublishVolumeRequest) Reset()         { *m = NodePublishVolumeRequest{} }
func (m *NodePublishVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*NodePublishVolumeRequest) ProtoMessage()    {}
func (*NodePublishVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

func (m *NodePublishVolumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePublishVolumeRequest.Unmarshal(m, b)
}
func (m *NodePublishVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePublishVolumeRequest.Marshal(b, m, deterministic)
}
func (m *NodePublishVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePublishVolumeRequest.Merge(m, src)
}
func (m *NodePublishVolumeRequest) XXX_Size() int {
	return xxx_messageInfo_NodePublishVolumeRequest.Size(m)
}
func (m *NodePublishVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePublishVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodePublishVolumeRequest proto.InternalMessageInfo

type NodePublishVolumeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodePublishVolumeResponse) Reset()         { *m = NodePublishVolumeResponse{} }
func (m *NodePublishVolumeResponse) String() string { return proto.CompactTextString(m) }
func (*NodePublishVolumeResponse) ProtoMessage()    {}
func (*NodePublishVolumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

func (m *NodePublishVolumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePublishVolumeResponse.Unmarshal(m, b)
}
func (m *NodePublishVolumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePublishVolumeResponse.Marshal(b, m, deterministic)
}
func (m *NodePublishVolumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePublishVolumeResponse.Merge(m, src)
}
func (m *NodePublishVolumeResponse) XXX_Size() int {
	return xxx_messageInfo_NodePublishVolumeResponse.Size(m)
}
func (m *NodePublishVolumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePublishVolumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodePublishVolumeResponse proto.InternalMessageInfo

type NodeUnpublishVolumeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeUnpublishVolumeRequest) Reset()         { *m = NodeUnpublishVolumeRequest{} }
func (m *NodeUnpublishVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*NodeUnpublishVolumeRequest) ProtoMessage()    {}
func (*NodeUnpublishVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{2}
}

func (m *NodeUnpublishVolumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeUnpublishVolumeRequest.Unmarshal(m, b)
}
func (m *NodeUnpublishVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeUnpublishVolumeRequest.Marshal(b, m, deterministic)
}
func (m *NodeUnpublishVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeUnpublishVolumeRequest.Merge(m, src)
}
func (m *NodeUnpublishVolumeRequest) XXX_Size() int {
	return xxx_messageInfo_NodeUnpublishVolumeRequest.Size(m)
}
func (m *NodeUnpublishVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeUnpublishVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeUnpublishVolumeRequest proto.InternalMessageInfo

type NodeUnpublishVolumeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeUnpublishVolumeResponse) Reset()         { *m = NodeUnpublishVolumeResponse{} }
func (m *NodeUnpublishVolumeResponse) String() string { return proto.CompactTextString(m) }
func (*NodeUnpublishVolumeResponse) ProtoMessage()    {}
func (*NodeUnpublishVolumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{3}
}

func (m *NodeUnpublishVolumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeUnpublishVolumeResponse.Unmarshal(m, b)
}
func (m *NodeUnpublishVolumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeUnpublishVolumeResponse.Marshal(b, m, deterministic)
}
func (m *NodeUnpublishVolumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeUnpublishVolumeResponse.Merge(m, src)
}
func (m *NodeUnpublishVolumeResponse) XXX_Size() int {
	return xxx_messageInfo_NodeUnpublishVolumeResponse.Size(m)
}
func (m *NodeUnpublishVolumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeUnpublishVolumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeUnpublishVolumeResponse proto.InternalMessageInfo

type NodeGetCapabilitiesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeGetCapabilitiesRequest) Reset()         { *m = NodeGetCapabilitiesRequest{} }
func (m *NodeGetCapabilitiesRequest) String() string { return proto.CompactTextString(m) }
func (*NodeGetCapabilitiesRequest) ProtoMessage()    {}
func (*NodeGetCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{4}
}

func (m *NodeGetCapabilitiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeGetCapabilitiesRequest.Unmarshal(m, b)
}
func (m *NodeGetCapabilitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeGetCapabilitiesRequest.Marshal(b, m, deterministic)
}
func (m *NodeGetCapabilitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeGetCapabilitiesRequest.Merge(m, src)
}
func (m *NodeGetCapabilitiesRequest) XXX_Size() int {
	return xxx_messageInfo_NodeGetCapabilitiesRequest.Size(m)
}
func (m *NodeGetCapabilitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeGetCapabilitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeGetCapabilitiesRequest proto.InternalMessageInfo

type NodeGetCapabilitiesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeGetCapabilitiesResponse) Reset()         { *m = NodeGetCapabilitiesResponse{} }
func (m *NodeGetCapabilitiesResponse) String() string { return proto.CompactTextString(m) }
func (*NodeGetCapabilitiesResponse) ProtoMessage()    {}
func (*NodeGetCapabilitiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{5}
}

func (m *NodeGetCapabilitiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeGetCapabilitiesResponse.Unmarshal(m, b)
}
func (m *NodeGetCapabilitiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeGetCapabilitiesResponse.Marshal(b, m, deterministic)
}
func (m *NodeGetCapabilitiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeGetCapabilitiesResponse.Merge(m, src)
}
func (m *NodeGetCapabilitiesResponse) XXX_Size() int {
	return xxx_messageInfo_NodeGetCapabilitiesResponse.Size(m)
}
func (m *NodeGetCapabilitiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeGetCapabilitiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeGetCapabilitiesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NodePublishVolumeRequest)(nil), "node.NodePublishVolumeRequest")
	proto.RegisterType((*NodePublishVolumeResponse)(nil), "node.NodePublishVolumeResponse")
	proto.RegisterType((*NodeUnpublishVolumeRequest)(nil), "node.NodeUnpublishVolumeRequest")
	proto.RegisterType((*NodeUnpublishVolumeResponse)(nil), "node.NodeUnpublishVolumeResponse")
	proto.RegisterType((*NodeGetCapabilitiesRequest)(nil), "node.NodeGetCapabilitiesRequest")
	proto.RegisterType((*NodeGetCapabilitiesResponse)(nil), "node.NodeGetCapabilitiesResponse")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcb, 0x4f, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xa4, 0xb8, 0x24, 0xfc, 0xf2,
	0x53, 0x52, 0x03, 0x4a, 0x93, 0x72, 0x32, 0x8b, 0x33, 0xc2, 0xf2, 0x73, 0x4a, 0x73, 0x53, 0x83,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xa4, 0xb9, 0x24, 0xb1, 0xc8, 0x15, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x2a, 0xc9, 0x70, 0x49, 0x81, 0x24, 0x43, 0xf3, 0x0a, 0xb0, 0x69, 0x95, 0xe5, 0x92,
	0xc6, 0x2a, 0x8b, 0xaa, 0xd9, 0x3d, 0xb5, 0xc4, 0x39, 0xb1, 0x20, 0x31, 0x29, 0x33, 0x27, 0xb3,
	0x24, 0x33, 0xb5, 0x18, 0x4d, 0x33, 0x86, 0x2c, 0x44, 0xb3, 0xd1, 0x2c, 0x26, 0x2e, 0x16, 0x90,
	0xbc, 0x50, 0x18, 0x97, 0x20, 0x86, 0xfb, 0x84, 0xe4, 0xf4, 0xc0, 0x7e, 0xc4, 0xe5, 0x29, 0x29,
	0x79, 0x9c, 0xf2, 0x50, 0xb7, 0x31, 0x08, 0xc5, 0x70, 0x09, 0x63, 0x71, 0xbc, 0x90, 0x02, 0x42,
	0x27, 0x76, 0x5f, 0x4b, 0x29, 0xe2, 0x51, 0x81, 0x6e, 0x3a, 0x9a, 0xef, 0x90, 0x4d, 0xc7, 0x1e,
	0x2c, 0xc8, 0xa6, 0xe3, 0x08, 0x1a, 0x25, 0x06, 0x27, 0x96, 0x28, 0xa6, 0x82, 0xa4, 0x24, 0x36,
	0x70, 0x14, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x48, 0xf6, 0x53, 0xf0, 0x01, 0x00,
	0x00,
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
	// necessary
	NodePublishVolume(ctx context.Context, in *NodePublishVolumeRequest, opts ...grpc.CallOption) (*NodePublishVolumeResponse, error)
	NodeUnpublishVolume(ctx context.Context, in *NodeUnpublishVolumeRequest, opts ...grpc.CallOption) (*NodeUnpublishVolumeResponse, error)
	NodeGetCapabilities(ctx context.Context, in *NodeGetCapabilitiesRequest, opts ...grpc.CallOption) (*NodeGetCapabilitiesResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) NodePublishVolume(ctx context.Context, in *NodePublishVolumeRequest, opts ...grpc.CallOption) (*NodePublishVolumeResponse, error) {
	out := new(NodePublishVolumeResponse)
	err := c.cc.Invoke(ctx, "/node.Node/NodePublishVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) NodeUnpublishVolume(ctx context.Context, in *NodeUnpublishVolumeRequest, opts ...grpc.CallOption) (*NodeUnpublishVolumeResponse, error) {
	out := new(NodeUnpublishVolumeResponse)
	err := c.cc.Invoke(ctx, "/node.Node/NodeUnpublishVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) NodeGetCapabilities(ctx context.Context, in *NodeGetCapabilitiesRequest, opts ...grpc.CallOption) (*NodeGetCapabilitiesResponse, error) {
	out := new(NodeGetCapabilitiesResponse)
	err := c.cc.Invoke(ctx, "/node.Node/NodeGetCapabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	// necessary
	NodePublishVolume(context.Context, *NodePublishVolumeRequest) (*NodePublishVolumeResponse, error)
	NodeUnpublishVolume(context.Context, *NodeUnpublishVolumeRequest) (*NodeUnpublishVolumeResponse, error)
	NodeGetCapabilities(context.Context, *NodeGetCapabilitiesRequest) (*NodeGetCapabilitiesResponse, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) NodePublishVolume(ctx context.Context, req *NodePublishVolumeRequest) (*NodePublishVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodePublishVolume not implemented")
}
func (*UnimplementedNodeServer) NodeUnpublishVolume(ctx context.Context, req *NodeUnpublishVolumeRequest) (*NodeUnpublishVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeUnpublishVolume not implemented")
}
func (*UnimplementedNodeServer) NodeGetCapabilities(ctx context.Context, req *NodeGetCapabilitiesRequest) (*NodeGetCapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGetCapabilities not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_NodePublishVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodePublishVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).NodePublishVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.Node/NodePublishVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).NodePublishVolume(ctx, req.(*NodePublishVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_NodeUnpublishVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeUnpublishVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).NodeUnpublishVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.Node/NodeUnpublishVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).NodeUnpublishVolume(ctx, req.(*NodeUnpublishVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_NodeGetCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGetCapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).NodeGetCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.Node/NodeGetCapabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).NodeGetCapabilities(ctx, req.(*NodeGetCapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "node.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NodePublishVolume",
			Handler:    _Node_NodePublishVolume_Handler,
		},
		{
			MethodName: "NodeUnpublishVolume",
			Handler:    _Node_NodeUnpublishVolume_Handler,
		},
		{
			MethodName: "NodeGetCapabilities",
			Handler:    _Node_NodeGetCapabilities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}