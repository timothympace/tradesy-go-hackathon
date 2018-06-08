// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

package item

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

// The request message containing the user's name.
type Item struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Brand                string   `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_40ef58970b710acc, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (dst *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(dst, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Item) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type Id struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_40ef58970b710acc, []int{1}
}
func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (dst *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(dst, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_40ef58970b710acc, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Item)(nil), "item.Item")
	proto.RegisterType((*Id)(nil), "item.Id")
	proto.RegisterType((*Empty)(nil), "item.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ItemApiClient is the client API for ItemApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemApiClient interface {
	// Sends a greet`ing
	GetItem(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Item, error)
	GetItems(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ItemApi_GetItemsClient, error)
	AddItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Id, error)
	DeleteItem(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error)
	UpdateItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Id, error)
}

type itemApiClient struct {
	cc *grpc.ClientConn
}

func NewItemApiClient(cc *grpc.ClientConn) ItemApiClient {
	return &itemApiClient{cc}
}

func (c *itemApiClient) GetItem(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/item.ItemApi/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemApiClient) GetItems(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ItemApi_GetItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ItemApi_serviceDesc.Streams[0], "/item.ItemApi/GetItems", opts...)
	if err != nil {
		return nil, err
	}
	x := &itemApiGetItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemApi_GetItemsClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type itemApiGetItemsClient struct {
	grpc.ClientStream
}

func (x *itemApiGetItemsClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemApiClient) AddItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/item.ItemApi/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemApiClient) DeleteItem(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/item.ItemApi/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemApiClient) UpdateItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/item.ItemApi/UpdateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemApiServer is the server API for ItemApi service.
type ItemApiServer interface {
	// Sends a greet`ing
	GetItem(context.Context, *Id) (*Item, error)
	GetItems(*Empty, ItemApi_GetItemsServer) error
	AddItem(context.Context, *Item) (*Id, error)
	DeleteItem(context.Context, *Id) (*Id, error)
	UpdateItem(context.Context, *Item) (*Id, error)
}

func RegisterItemApiServer(s *grpc.Server, srv ItemApiServer) {
	s.RegisterService(&_ItemApi_serviceDesc, srv)
}

func _ItemApi_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemApiServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemApi/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemApiServer).GetItem(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemApi_GetItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemApiServer).GetItems(m, &itemApiGetItemsServer{stream})
}

type ItemApi_GetItemsServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type itemApiGetItemsServer struct {
	grpc.ServerStream
}

func (x *itemApiGetItemsServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

func _ItemApi_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemApiServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemApi/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemApiServer).AddItem(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemApi_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemApiServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemApi/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemApiServer).DeleteItem(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemApi_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemApiServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemApi/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemApiServer).UpdateItem(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

var _ItemApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "item.ItemApi",
	HandlerType: (*ItemApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItem",
			Handler:    _ItemApi_GetItem_Handler,
		},
		{
			MethodName: "AddItem",
			Handler:    _ItemApi_AddItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _ItemApi_DeleteItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _ItemApi_UpdateItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetItems",
			Handler:       _ItemApi_GetItems_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "item.proto",
}

func init() { proto.RegisterFile("item.proto", fileDescriptor_item_40ef58970b710acc) }

var fileDescriptor_item_40ef58970b710acc = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x9b, 0x98, 0xba, 0xf5, 0x09, 0x1e, 0x86, 0x1e, 0x42, 0x4f, 0x35, 0x08, 0xf6, 0x54,
	0x44, 0x7f, 0x41, 0x41, 0x91, 0xbd, 0x16, 0xf4, 0xbe, 0x75, 0xe6, 0x10, 0x30, 0x6d, 0x58, 0x73,
	0xf1, 0xef, 0xf9, 0xcb, 0x24, 0xb3, 0x8b, 0xd4, 0x42, 0x6f, 0x6f, 0xde, 0x3c, 0xbe, 0x79, 0x0c,
	0x10, 0x8b, 0xa4, 0x75, 0xee, 0x0f, 0xe5, 0x40, 0xae, 0xea, 0xf0, 0x0e, 0xd7, 0x16, 0x49, 0x74,
	0x03, 0x1b, 0xd9, 0x9b, 0xa5, 0x59, 0x5d, 0x6d, 0x6d, 0x64, 0x22, 0xb8, 0x7d, 0x97, 0xc4, 0x5b,
	0x75, 0x54, 0xd3, 0x1c, 0xd3, 0x5d, 0xdf, 0xed, 0xd9, 0x5f, 0xa8, 0x39, 0x0c, 0xd5, 0xcd, 0x7d,
	0xfc, 0x10, 0xef, 0x96, 0x66, 0x65, 0xb7, 0xc3, 0x10, 0xe6, 0xb0, 0x2d, 0x9f, 0x52, 0x43, 0x83,
	0xe9, 0x4b, 0xca, 0xe5, 0xfb, 0xf1, 0xc7, 0xa0, 0xa9, 0x77, 0x37, 0x39, 0xd2, 0x2d, 0x9a, 0x57,
	0x29, 0xda, 0x62, 0xb6, 0xd6, 0x82, 0x2d, 0x2f, 0x30, 0xaa, 0xda, 0x71, 0x42, 0xf7, 0x98, 0x8d,
	0x91, 0x2f, 0xba, 0x1e, 0x36, 0xca, 0xf9, 0x1f, 0x7b, 0x30, 0x95, 0xb5, 0x61, 0x56, 0xd6, 0xd1,
	0x6a, 0xf1, 0xc7, 0x0d, 0x13, 0x0a, 0xc0, 0xb3, 0x7c, 0x4a, 0x91, 0x93, 0x8b, 0xc7, 0x99, 0x3b,
	0xe0, 0x2d, 0x73, 0x37, 0x66, 0xce, 0x90, 0x76, 0x97, 0xfa, 0xc8, 0xa7, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x55, 0x08, 0xa9, 0x19, 0x56, 0x01, 0x00, 0x00,
}