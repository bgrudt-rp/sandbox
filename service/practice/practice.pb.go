// Code generated by protoc-gen-go. DO NOT EDIT.
// source: practice/practice.proto

package practice

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

type Practice struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Practice) Reset()         { *m = Practice{} }
func (m *Practice) String() string { return proto.CompactTextString(m) }
func (*Practice) ProtoMessage()    {}
func (*Practice) Descriptor() ([]byte, []int) {
	return fileDescriptor_887b89dde555b786, []int{0}
}

func (m *Practice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Practice.Unmarshal(m, b)
}
func (m *Practice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Practice.Marshal(b, m, deterministic)
}
func (m *Practice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Practice.Merge(m, src)
}
func (m *Practice) XXX_Size() int {
	return xxx_messageInfo_Practice.Size(m)
}
func (m *Practice) XXX_DiscardUnknown() {
	xxx_messageInfo_Practice.DiscardUnknown(m)
}

var xxx_messageInfo_Practice proto.InternalMessageInfo

func (m *Practice) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Practice) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Practice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Subpractice struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	PracticeCode         string   `protobuf:"bytes,4,opt,name=practice_code,json=practiceCode,proto3" json:"practice_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subpractice) Reset()         { *m = Subpractice{} }
func (m *Subpractice) String() string { return proto.CompactTextString(m) }
func (*Subpractice) ProtoMessage()    {}
func (*Subpractice) Descriptor() ([]byte, []int) {
	return fileDescriptor_887b89dde555b786, []int{1}
}

func (m *Subpractice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subpractice.Unmarshal(m, b)
}
func (m *Subpractice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subpractice.Marshal(b, m, deterministic)
}
func (m *Subpractice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subpractice.Merge(m, src)
}
func (m *Subpractice) XXX_Size() int {
	return xxx_messageInfo_Subpractice.Size(m)
}
func (m *Subpractice) XXX_DiscardUnknown() {
	xxx_messageInfo_Subpractice.DiscardUnknown(m)
}

var xxx_messageInfo_Subpractice proto.InternalMessageInfo

func (m *Subpractice) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Subpractice) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Subpractice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subpractice) GetPracticeCode() string {
	if m != nil {
		return m.PracticeCode
	}
	return ""
}

type Response struct {
	Created              bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Practice             *Practice `protobuf:"bytes,2,opt,name=practice,proto3" json:"practice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_887b89dde555b786, []int{2}
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

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetPractice() *Practice {
	if m != nil {
		return m.Practice
	}
	return nil
}

func init() {
	proto.RegisterType((*Practice)(nil), "practice.Practice")
	proto.RegisterType((*Subpractice)(nil), "practice.Subpractice")
	proto.RegisterType((*Response)(nil), "practice.Response")
}

func init() {
	proto.RegisterFile("practice/practice.proto", fileDescriptor_887b89dde555b786)
}

var fileDescriptor_887b89dde555b786 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0x4a, 0x4c,
	0x2e, 0xc9, 0x4c, 0x4e, 0xd5, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0x25, 0x27, 0x2e, 0x8e, 0x00, 0x28, 0x5b, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b, 0x25, 0x39, 0x3f, 0x25, 0x55,
	0x82, 0x09, 0x2c, 0x02, 0x66, 0x83, 0xc4, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x21, 0x62, 0x20,
	0xb6, 0x52, 0x16, 0x17, 0x77, 0x70, 0x69, 0x52, 0x01, 0x85, 0xc6, 0x08, 0x29, 0x73, 0xf1, 0xc2,
	0xcc, 0x88, 0x07, 0x6b, 0x60, 0x01, 0x4b, 0xf2, 0xc0, 0x04, 0x9d, 0xf3, 0x53, 0x52, 0x95, 0x42,
	0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x8b,
	0x52, 0x13, 0x4b, 0x52, 0x21, 0xb6, 0x71, 0x04, 0xc1, 0xb8, 0x42, 0x7a, 0x5c, 0x70, 0x1f, 0x82,
	0xad, 0xe5, 0x36, 0x12, 0xd2, 0x83, 0x07, 0x01, 0xcc, 0xbf, 0x41, 0x70, 0x35, 0x46, 0x1e, 0x5c,
	0x5c, 0x41, 0x01, 0xf0, 0x70, 0xb0, 0xe2, 0xe2, 0x73, 0x06, 0x1b, 0x04, 0x17, 0xc1, 0xa2, 0x5b,
	0x0a, 0x49, 0x0c, 0xe6, 0x22, 0x25, 0x86, 0x24, 0x36, 0x70, 0x00, 0x1b, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x6c, 0xa6, 0xa8, 0x87, 0x7b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RPPracticeClient is the client API for RPPractice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPPracticeClient interface {
	CreatePractice(ctx context.Context, in *Practice, opts ...grpc.CallOption) (*Response, error)
}

type rPPracticeClient struct {
	cc grpc.ClientConnInterface
}

func NewRPPracticeClient(cc grpc.ClientConnInterface) RPPracticeClient {
	return &rPPracticeClient{cc}
}

func (c *rPPracticeClient) CreatePractice(ctx context.Context, in *Practice, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/practice.RPPractice/CreatePractice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPPracticeServer is the server API for RPPractice service.
type RPPracticeServer interface {
	CreatePractice(context.Context, *Practice) (*Response, error)
}

// UnimplementedRPPracticeServer can be embedded to have forward compatible implementations.
type UnimplementedRPPracticeServer struct {
}

func (*UnimplementedRPPracticeServer) CreatePractice(ctx context.Context, req *Practice) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePractice not implemented")
}

func RegisterRPPracticeServer(s *grpc.Server, srv RPPracticeServer) {
	s.RegisterService(&_RPPractice_serviceDesc, srv)
}

func _RPPractice_CreatePractice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Practice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPPracticeServer).CreatePractice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/practice.RPPractice/CreatePractice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPPracticeServer).CreatePractice(ctx, req.(*Practice))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPPractice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "practice.RPPractice",
	HandlerType: (*RPPracticeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePractice",
			Handler:    _RPPractice_CreatePractice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "practice/practice.proto",
}