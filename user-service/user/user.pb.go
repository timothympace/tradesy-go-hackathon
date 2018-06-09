// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

// Request message for creating a new user
type User struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string          `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string          `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Addresses            []*User_Address `protobuf:"bytes,5,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c2e712f5df49e26e, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetAddresses() []*User_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type User_Address struct {
	Street               string   `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City                 string   `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Zip                  string   `protobuf:"bytes,4,opt,name=zip,proto3" json:"zip,omitempty"`
	IsShippingAddress    bool     `protobuf:"varint,5,opt,name=isShippingAddress,proto3" json:"isShippingAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User_Address) Reset()         { *m = User_Address{} }
func (m *User_Address) String() string { return proto.CompactTextString(m) }
func (*User_Address) ProtoMessage()    {}
func (*User_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c2e712f5df49e26e, []int{0, 0}
}
func (m *User_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User_Address.Unmarshal(m, b)
}
func (m *User_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User_Address.Marshal(b, m, deterministic)
}
func (dst *User_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User_Address.Merge(dst, src)
}
func (m *User_Address) XXX_Size() int {
	return xxx_messageInfo_User_Address.Size(m)
}
func (m *User_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_User_Address.DiscardUnknown(m)
}

var xxx_messageInfo_User_Address proto.InternalMessageInfo

func (m *User_Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *User_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *User_Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *User_Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *User_Address) GetIsShippingAddress() bool {
	if m != nil {
		return m.IsShippingAddress
	}
	return false
}

type UserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c2e712f5df49e26e, []int{1}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type UserFilter struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserFilter) Reset()         { *m = UserFilter{} }
func (m *UserFilter) String() string { return proto.CompactTextString(m) }
func (*UserFilter) ProtoMessage()    {}
func (*UserFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c2e712f5df49e26e, []int{2}
}
func (m *UserFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFilter.Unmarshal(m, b)
}
func (m *UserFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFilter.Marshal(b, m, deterministic)
}
func (dst *UserFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFilter.Merge(dst, src)
}
func (m *UserFilter) XXX_Size() int {
	return xxx_messageInfo_UserFilter.Size(m)
}
func (m *UserFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFilter.DiscardUnknown(m)
}

var xxx_messageInfo_UserFilter proto.InternalMessageInfo

func (m *UserFilter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*User_Address)(nil), "user.User.Address")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
	proto.RegisterType((*UserFilter)(nil), "user.UserFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserApiClient is the client API for UserApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserApiClient interface {
	// Get User by ID - A server-to-client streaming RPC.
	GetUser(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (UserApi_GetUserClient, error)
	// Create a new User - A simple RPC
	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	// Update User
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	// Delete User
	DeleteUser(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (*UserResponse, error)
}

type userApiClient struct {
	cc *grpc.ClientConn
}

func NewUserApiClient(cc *grpc.ClientConn) UserApiClient {
	return &userApiClient{cc}
}

func (c *userApiClient) GetUser(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (UserApi_GetUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserApi_serviceDesc.Streams[0], "/user.UserApi/GetUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &userApiGetUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserApi_GetUserClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userApiGetUserClient struct {
	grpc.ClientStream
}

func (x *userApiGetUserClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userApiClient) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserApi/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserApi/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) DeleteUser(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserApi/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserApiServer is the server API for UserApi service.
type UserApiServer interface {
	// Get User by ID - A server-to-client streaming RPC.
	GetUser(*UserFilter, UserApi_GetUserServer) error
	// Create a new User - A simple RPC
	AddUser(context.Context, *User) (*UserResponse, error)
	// Update User
	UpdateUser(context.Context, *User) (*UserResponse, error)
	// Delete User
	DeleteUser(context.Context, *UserFilter) (*UserResponse, error)
}

func RegisterUserApiServer(s *grpc.Server, srv UserApiServer) {
	s.RegisterService(&_UserApi_serviceDesc, srv)
}

func _UserApi_GetUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserApiServer).GetUser(m, &userApiGetUserServer{stream})
}

type UserApi_GetUserServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userApiGetUserServer struct {
	grpc.ServerStream
}

func (x *userApiGetUserServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UserApi_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserApi/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).AddUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApi_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserApi/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApi_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserApi/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).DeleteUser(ctx, req.(*UserFilter))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserApi",
	HandlerType: (*UserApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserApi_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserApi_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserApi_DeleteUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUser",
			Handler:       _UserApi_GetUser_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_c2e712f5df49e26e) }

var fileDescriptor_user_c2e712f5df49e26e = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xed, 0xa6, 0x69, 0xf3, 0x75, 0x3e, 0x91, 0x3a, 0x88, 0x2c, 0xc5, 0x43, 0xc9, 0xa9, 0xa0,
	0x84, 0x52, 0x3d, 0x78, 0x2d, 0x88, 0xde, 0x23, 0xfd, 0x01, 0xb1, 0x3b, 0xd8, 0x85, 0x36, 0x59,
	0xb2, 0xdb, 0x83, 0xfe, 0x05, 0xef, 0xfe, 0x27, 0xff, 0x95, 0xcc, 0x66, 0x43, 0x0a, 0xf1, 0xe0,
	0x6d, 0xde, 0x9b, 0x37, 0xfb, 0x1e, 0x8f, 0x05, 0x38, 0x5a, 0xaa, 0x33, 0x53, 0x57, 0xae, 0xc2,
	0x98, 0xe7, 0xf4, 0x2b, 0x82, 0x78, 0x63, 0xa9, 0xc6, 0x73, 0x88, 0xb4, 0x92, 0x62, 0x2e, 0x16,
	0x93, 0x3c, 0xd2, 0x0a, 0x11, 0xe2, 0xb2, 0x38, 0x90, 0x8c, 0x3c, 0xe3, 0x67, 0xbc, 0x84, 0x11,
	0x1d, 0x0a, 0xbd, 0x97, 0x43, 0x4f, 0x36, 0x80, 0x59, 0xb3, 0xab, 0x4a, 0x92, 0x71, 0xc3, 0x7a,
	0x80, 0x4b, 0x98, 0x14, 0x4a, 0xd5, 0x64, 0x2d, 0x59, 0x39, 0x9a, 0x0f, 0x17, 0xff, 0x57, 0x98,
	0x79, 0x7b, 0xb6, 0xcb, 0xd6, 0xcd, 0x2e, 0xef, 0x44, 0xb3, 0x4f, 0x01, 0x49, 0xa0, 0xf1, 0x0a,
	0xc6, 0xd6, 0xd5, 0x44, 0x2e, 0x24, 0x0a, 0x88, 0x53, 0x6d, 0xb5, 0x7b, 0x6f, 0x53, 0xf1, 0xcc,
	0xfe, 0xd6, 0x15, 0x8e, 0xda, 0x54, 0x1e, 0xe0, 0x14, 0x86, 0x1f, 0xda, 0x84, 0x4c, 0x3c, 0xe2,
	0x2d, 0x5c, 0x68, 0xfb, 0xb2, 0xd3, 0xc6, 0xe8, 0xf2, 0x2d, 0x18, 0xc9, 0xd1, 0x5c, 0x2c, 0xfe,
	0xe5, 0xfd, 0x45, 0xfa, 0x00, 0x67, 0x1c, 0x34, 0x27, 0x6b, 0xaa, 0xd2, 0x52, 0xaf, 0x1f, 0x09,
	0x89, 0x3d, 0x6e, 0xb7, 0xfc, 0x46, 0xe4, 0xdf, 0x68, 0x61, 0x7a, 0x0d, 0xc0, 0x97, 0x4f, 0x7a,
	0xef, 0xfa, 0xbd, 0xae, 0xbe, 0x05, 0x24, 0xbc, 0x5e, 0x1b, 0x8d, 0x37, 0x90, 0x3c, 0x93, 0xf3,
	0xf5, 0x4f, 0xbb, 0x6e, 0x9a, 0xc3, 0x19, 0x74, 0x4c, 0x3a, 0x58, 0x0a, 0x16, 0xaf, 0x95, 0xf2,
	0xe2, 0x93, 0xd5, 0xec, 0xa4, 0xd4, 0x36, 0x6b, 0x3a, 0xc0, 0x0c, 0x60, 0x63, 0x54, 0xe1, 0xe8,
	0x8f, 0xfa, 0x7b, 0x80, 0x47, 0xda, 0x53, 0xd0, 0xf7, 0xc3, 0xfc, 0x7a, 0xf5, 0x3a, 0xf6, 0x3f,
	0xe9, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xcb, 0xf4, 0x9d, 0x57, 0x02, 0x00, 0x00,
}
