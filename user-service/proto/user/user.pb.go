// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package laracom_service_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	StripeId             string   `protobuf:"bytes,6,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	CardBrand            string   `protobuf:"bytes,7,opt,name=card_brand,json=cardBrand,proto3" json:"card_brand,omitempty"`
	CardLastFour         string   `protobuf:"bytes,8,opt,name=card_last_four,json=cardLastFour,proto3" json:"card_last_four,omitempty"`
	TrialEndsAt          string   `protobuf:"bytes,9,opt,name=trial_ends_at,json=trialEndsAt,proto3" json:"trial_ends_at,omitempty"`
	RememberToken        string   `protobuf:"bytes,10,opt,name=remember_token,json=rememberToken,proto3" json:"remember_token,omitempty"`
	CreatedAt            string   `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
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

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *User) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *User) GetCardBrand() string {
	if m != nil {
		return m.CardBrand
	}
	return ""
}

func (m *User) GetCardLastFour() string {
	if m != nil {
		return m.CardLastFour
	}
	return ""
}

func (m *User) GetTrialEndsAt() string {
	if m != nil {
		return m.TrialEndsAt
	}
	return ""
}

func (m *User) GetRememberToken() string {
	if m != nil {
		return m.RememberToken
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
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

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "laracom.service.user.User")
	proto.RegisterType((*Request)(nil), "laracom.service.user.Request")
	proto.RegisterType((*Response)(nil), "laracom.service.user.Response")
	proto.RegisterType((*Error)(nil), "laracom.service.user.Error")
	proto.RegisterType((*Token)(nil), "laracom.service.user.Token")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0x1b, 0x27, 0x76, 0x93, 0x49, 0x93, 0xc3, 0x28, 0xef, 0xab, 0x55, 0xaa, 0xa2, 0xc8,
	0x02, 0x89, 0x93, 0x41, 0xed, 0x99, 0x83, 0xa9, 0xda, 0xaa, 0x82, 0x93, 0xf9, 0x73, 0xb5, 0x36,
	0xd9, 0x41, 0xb5, 0x70, 0xbc, 0x66, 0x77, 0x5d, 0x3e, 0x0f, 0x9f, 0x8c, 0x4f, 0xc1, 0x1d, 0xed,
	0xac, 0x83, 0x38, 0xa4, 0x29, 0x82, 0x4b, 0xb2, 0xf3, 0x7b, 0x9e, 0x19, 0x4d, 0x66, 0x46, 0x81,
	0xff, 0x5a, 0xa3, 0x9d, 0x7e, 0xd1, 0x59, 0x32, 0xfc, 0x91, 0x71, 0x8c, 0x8b, 0x5a, 0x1a, 0xb9,
	0xd1, 0xdb, 0xcc, 0x92, 0xb9, 0xaf, 0x36, 0x94, 0x79, 0x2d, 0xfd, 0x1e, 0xc1, 0xe8, 0x83, 0x25,
	0x83, 0x73, 0x88, 0x2a, 0x25, 0x06, 0xab, 0xc1, 0xf3, 0x49, 0x11, 0x55, 0x0a, 0x11, 0x46, 0x8d,
	0xdc, 0x92, 0x88, 0x98, 0xf0, 0x1b, 0x17, 0x10, 0xd3, 0x56, 0x56, 0xb5, 0x18, 0x32, 0x0c, 0x01,
	0x2e, 0x61, 0xdc, 0x4a, 0x6b, 0xbf, 0x6a, 0xa3, 0xc4, 0x88, 0x85, 0x5f, 0x31, 0xfe, 0x0f, 0x89,
	0x75, 0xd2, 0x75, 0x56, 0xc4, 0xac, 0xf4, 0x11, 0x9e, 0xc2, 0xc4, 0x3a, 0x53, 0xb5, 0x54, 0x56,
	0x4a, 0x24, 0x21, 0x29, 0x80, 0x5b, 0x85, 0x67, 0x00, 0x1b, 0x69, 0x54, 0xb9, 0x36, 0xb2, 0x51,
	0xe2, 0x98, 0xd5, 0x89, 0x27, 0xaf, 0x3d, 0xc0, 0xa7, 0x30, 0x67, 0xb9, 0x96, 0xd6, 0x95, 0x9f,
	0x74, 0x67, 0xc4, 0x98, 0x2d, 0x27, 0x9e, 0xbe, 0x95, 0xd6, 0x5d, 0xeb, 0xce, 0x60, 0x0a, 0x33,
	0x67, 0x2a, 0x59, 0x97, 0xd4, 0x28, 0x5b, 0x4a, 0x27, 0x26, 0x6c, 0x9a, 0x32, 0xbc, 0x6a, 0x94,
	0xcd, 0x1d, 0x3e, 0x83, 0xb9, 0xa1, 0x2d, 0x6d, 0xd7, 0x64, 0x4a, 0xa7, 0x3f, 0x53, 0x23, 0x80,
	0x4d, 0xb3, 0x1d, 0x7d, 0xef, 0x21, 0xf7, 0x63, 0x48, 0x3a, 0x52, 0xbe, 0xce, 0xb4, 0xef, 0x27,
	0x90, 0xdc, 0x79, 0xb9, 0x6b, 0xd5, 0x4e, 0x3e, 0x09, 0x72, 0x4f, 0x72, 0x97, 0x4e, 0xe0, 0xb8,
	0xa0, 0x2f, 0x1d, 0x59, 0x97, 0x7e, 0x1b, 0xc0, 0xb8, 0x20, 0xdb, 0xea, 0xc6, 0x12, 0x66, 0x30,
	0xf2, 0x1b, 0xe0, 0x91, 0x4f, 0xcf, 0x97, 0xd9, 0xbe, 0xf5, 0x64, 0x7e, 0x35, 0x05, 0xfb, 0xf0,
	0x25, 0xc4, 0xfe, 0xdb, 0x8a, 0x68, 0x35, 0x7c, 0x24, 0x21, 0x18, 0xf1, 0x02, 0x12, 0x32, 0x46,
	0x1b, 0x2b, 0x86, 0x9c, 0x72, 0xba, 0x3f, 0xe5, 0xca, 0x7b, 0x8a, 0xde, 0x9a, 0xbe, 0x82, 0x98,
	0x81, 0x3f, 0x80, 0x8d, 0x56, 0xc4, 0xfd, 0xc5, 0x05, 0xbf, 0x71, 0x05, 0x53, 0x45, 0x76, 0x63,
	0xaa, 0xd6, 0x55, 0xba, 0xe9, 0x6f, 0xe3, 0x77, 0x94, 0xde, 0x41, 0x1c, 0x86, 0xb6, 0x80, 0x38,
	0x8c, 0x34, 0x9c, 0x54, 0x08, 0x3c, 0xbd, 0x97, 0x75, 0xa5, 0x38, 0x75, 0x5c, 0x84, 0xe0, 0xaf,
	0x1a, 0x3d, 0xff, 0x11, 0xc1, 0xd4, 0xff, 0xda, 0x77, 0xc1, 0x82, 0xd7, 0x90, 0x5c, 0xf2, 0x4e,
	0xf0, 0xc0, 0x68, 0x96, 0x4f, 0xf6, 0x6b, 0xbb, 0xad, 0xa4, 0x47, 0x78, 0x09, 0xc3, 0x1b, 0x72,
	0xff, 0x58, 0xe4, 0x16, 0x92, 0x1b, 0x72, 0x79, 0x5d, 0xe3, 0xd9, 0x43, 0x5e, 0x3e, 0x89, 0x3f,
	0x28, 0x95, 0xc3, 0x28, 0xef, 0xdc, 0xdd, 0xc1, 0x86, 0x1e, 0x18, 0x18, 0x6f, 0x22, 0x3d, 0xc2,
	0x37, 0x30, 0xfb, 0xe8, 0x07, 0x2d, 0x1d, 0x85, 0xe5, 0x1c, 0xf2, 0x3f, 0x52, 0x6c, 0x9d, 0xf0,
	0xdf, 0xc9, 0xc5, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x31, 0x50, 0x68, 0x67, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "laracom.service.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}
