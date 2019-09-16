// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/done.proto

package done

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

type RegisterTokenRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterTokenRequest) Reset()         { *m = RegisterTokenRequest{} }
func (m *RegisterTokenRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterTokenRequest) ProtoMessage()    {}
func (*RegisterTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e503f38af1f412b3, []int{0}
}

func (m *RegisterTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterTokenRequest.Unmarshal(m, b)
}
func (m *RegisterTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterTokenRequest.Marshal(b, m, deterministic)
}
func (m *RegisterTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterTokenRequest.Merge(m, src)
}
func (m *RegisterTokenRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterTokenRequest.Size(m)
}
func (m *RegisterTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterTokenRequest proto.InternalMessageInfo

func (m *RegisterTokenRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type RegisterTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterTokenResponse) Reset()         { *m = RegisterTokenResponse{} }
func (m *RegisterTokenResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterTokenResponse) ProtoMessage()    {}
func (*RegisterTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e503f38af1f412b3, []int{1}
}

func (m *RegisterTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterTokenResponse.Unmarshal(m, b)
}
func (m *RegisterTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterTokenResponse.Marshal(b, m, deterministic)
}
func (m *RegisterTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterTokenResponse.Merge(m, src)
}
func (m *RegisterTokenResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterTokenResponse.Size(m)
}
func (m *RegisterTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterTokenResponse proto.InternalMessageInfo

func (m *RegisterTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegisterTokenResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type DeleteTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenRequest) Reset()         { *m = DeleteTokenRequest{} }
func (m *DeleteTokenRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenRequest) ProtoMessage()    {}
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e503f38af1f412b3, []int{2}
}

func (m *DeleteTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenRequest.Unmarshal(m, b)
}
func (m *DeleteTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenRequest.Merge(m, src)
}
func (m *DeleteTokenRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenRequest.Size(m)
}
func (m *DeleteTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenRequest proto.InternalMessageInfo

func (m *DeleteTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeleteTokenResponse struct {
	Deleted              bool     `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenResponse) Reset()         { *m = DeleteTokenResponse{} }
func (m *DeleteTokenResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenResponse) ProtoMessage()    {}
func (*DeleteTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e503f38af1f412b3, []int{3}
}

func (m *DeleteTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenResponse.Unmarshal(m, b)
}
func (m *DeleteTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenResponse.Merge(m, src)
}
func (m *DeleteTokenResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenResponse.Size(m)
}
func (m *DeleteTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenResponse proto.InternalMessageInfo

func (m *DeleteTokenResponse) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *DeleteTokenResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type ValidateTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenRequest) Reset()         { *m = ValidateTokenRequest{} }
func (m *ValidateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenRequest) ProtoMessage()    {}
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e503f38af1f412b3, []int{4}
}

func (m *ValidateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenRequest.Unmarshal(m, b)
}
func (m *ValidateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenRequest.Marshal(b, m, deterministic)
}
func (m *ValidateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenRequest.Merge(m, src)
}
func (m *ValidateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenRequest.Size(m)
}
func (m *ValidateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenRequest proto.InternalMessageInfo

func (m *ValidateTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateTokenResponse struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Error                *Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenResponse) Reset()         { *m = ValidateTokenResponse{} }
func (m *ValidateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenResponse) ProtoMessage()    {}
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e503f38af1f412b3, []int{5}
}

func (m *ValidateTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenResponse.Unmarshal(m, b)
}
func (m *ValidateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenResponse.Marshal(b, m, deterministic)
}
func (m *ValidateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenResponse.Merge(m, src)
}
func (m *ValidateTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenResponse.Size(m)
}
func (m *ValidateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenResponse proto.InternalMessageInfo

func (m *ValidateTokenResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *ValidateTokenResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ValidateTokenResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Json                 string   `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_e503f38af1f412b3, []int{6}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterTokenRequest)(nil), "done.RegisterTokenRequest")
	proto.RegisterType((*RegisterTokenResponse)(nil), "done.RegisterTokenResponse")
	proto.RegisterType((*DeleteTokenRequest)(nil), "done.DeleteTokenRequest")
	proto.RegisterType((*DeleteTokenResponse)(nil), "done.DeleteTokenResponse")
	proto.RegisterType((*ValidateTokenRequest)(nil), "done.ValidateTokenRequest")
	proto.RegisterType((*ValidateTokenResponse)(nil), "done.ValidateTokenResponse")
	proto.RegisterType((*Error)(nil), "done.Error")
}

func init() { proto.RegisterFile("proto/done.proto", fileDescriptor_e503f38af1f412b3) }

var fileDescriptor_e503f38af1f412b3 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0x32, 0x31,
	0x10, 0x86, 0x59, 0x3e, 0xe0, 0x93, 0x21, 0x24, 0xa6, 0x82, 0x59, 0xd7, 0x0b, 0xf6, 0x44, 0x8c,
	0x59, 0x13, 0xfc, 0x05, 0x26, 0x78, 0x80, 0x93, 0xa9, 0xc6, 0xfb, 0xba, 0x3b, 0x81, 0x55, 0x6c,
	0xb1, 0xed, 0xf2, 0xcf, 0xbd, 0x9b, 0x6d, 0xbb, 0xb2, 0x8b, 0x35, 0xe1, 0x36, 0xef, 0xcc, 0x64,
	0xe6, 0xe9, 0xbc, 0x85, 0xd3, 0xad, 0x14, 0x5a, 0xdc, 0x66, 0x82, 0x63, 0x6c, 0x42, 0xd2, 0x29,
	0x63, 0x1a, 0xc3, 0x88, 0xe1, 0x2a, 0x57, 0x1a, 0xe5, 0xb3, 0x78, 0x47, 0xce, 0xf0, 0xb3, 0x40,
	0xa5, 0xc9, 0x39, 0xf4, 0x0a, 0x85, 0x72, 0x91, 0x85, 0xc1, 0x24, 0x98, 0xf6, 0x99, 0x53, 0xf4,
	0x11, 0xc6, 0x07, 0xfd, 0x6a, 0x2b, 0xb8, 0x42, 0x32, 0x82, 0xae, 0x49, 0xb8, 0x7e, 0x2b, 0xc8,
	0x15, 0x74, 0x51, 0x4a, 0x21, 0xc3, 0xf6, 0x24, 0x98, 0x0e, 0x66, 0x83, 0xd8, 0x00, 0x3c, 0x94,
	0x29, 0x66, 0x2b, 0xf4, 0x1a, 0xc8, 0x1c, 0x37, 0xa8, 0xb1, 0xb1, 0xdf, 0x3b, 0x8e, 0x32, 0x38,
	0x6b, 0xf4, 0xba, 0xdd, 0x21, 0xfc, 0xcf, 0x4c, 0xda, 0xd2, 0x9e, 0xb0, 0x4a, 0x1e, 0xb3, 0xff,
	0x06, 0x46, 0x2f, 0xc9, 0x26, 0xcf, 0x92, 0xa3, 0x08, 0xd6, 0x30, 0x3e, 0xe8, 0xde, 0xbf, 0x7f,
	0x57, 0x16, 0x1c, 0x81, 0x15, 0xb5, 0x33, 0xb6, 0xeb, 0x67, 0xdc, 0x73, 0xfd, 0xfb, 0x93, 0x6b,
	0x01, 0x5d, 0xa3, 0x09, 0x81, 0x4e, 0x2a, 0x32, 0x74, 0x1c, 0x26, 0x2e, 0x73, 0x6f, 0x4a, 0x70,
	0x37, 0xd5, 0xc4, 0xe5, 0x15, 0x3e, 0x50, 0xa9, 0x64, 0x85, 0x66, 0x6a, 0x9f, 0x55, 0x72, 0xf6,
	0x15, 0xc0, 0xf8, 0xbe, 0xd0, 0x6b, 0xe4, 0x3a, 0x4f, 0x13, 0x9d, 0x0b, 0xfe, 0x84, 0x72, 0x97,
	0xa7, 0x48, 0x96, 0x30, 0x6c, 0xd8, 0x49, 0x22, 0x4b, 0xe2, 0xfb, 0x13, 0xd1, 0xa5, 0xb7, 0x66,
	0xdf, 0x4f, 0x5b, 0x64, 0x0e, 0x83, 0x9a, 0x39, 0x24, 0xb4, 0xdd, 0xbf, 0xbd, 0x8d, 0x2e, 0x3c,
	0x95, 0x9f, 0x29, 0x4b, 0x18, 0x36, 0x0e, 0x5c, 0x11, 0xf9, 0x3c, 0xaa, 0x88, 0xbc, 0x8e, 0xd0,
	0xd6, 0x6b, 0xcf, 0xfc, 0xf4, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x41, 0xa9, 0x2a,
	0xfd, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	RegisterToken(ctx context.Context, in *RegisterTokenRequest, opts ...grpc.CallOption) (*RegisterTokenResponse, error)
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
}

type authenticationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationServiceClient(cc *grpc.ClientConn) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) RegisterToken(ctx context.Context, in *RegisterTokenRequest, opts ...grpc.CallOption) (*RegisterTokenResponse, error) {
	out := new(RegisterTokenResponse)
	err := c.cc.Invoke(ctx, "/done.AuthenticationService/RegisterToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error) {
	out := new(DeleteTokenResponse)
	err := c.cc.Invoke(ctx, "/done.AuthenticationService/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, "/done.AuthenticationService/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	RegisterToken(context.Context, *RegisterTokenRequest) (*RegisterTokenResponse, error)
	DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenResponse, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
}

// UnimplementedAuthenticationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (*UnimplementedAuthenticationServiceServer) RegisterToken(ctx context.Context, req *RegisterTokenRequest) (*RegisterTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterToken not implemented")
}
func (*UnimplementedAuthenticationServiceServer) DeleteToken(ctx context.Context, req *DeleteTokenRequest) (*DeleteTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (*UnimplementedAuthenticationServiceServer) ValidateToken(ctx context.Context, req *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_RegisterToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).RegisterToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/done.AuthenticationService/RegisterToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).RegisterToken(ctx, req.(*RegisterTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/done.AuthenticationService/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/done.AuthenticationService/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "done.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterToken",
			Handler:    _AuthenticationService_RegisterToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _AuthenticationService_DeleteToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _AuthenticationService_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/done.proto",
}