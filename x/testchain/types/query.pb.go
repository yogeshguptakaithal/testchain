// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testchain/testchain/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f06cd066eb8a7bf0, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f06cd066eb8a7bf0, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetCommentRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetCommentRequest) Reset()         { *m = QueryGetCommentRequest{} }
func (m *QueryGetCommentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCommentRequest) ProtoMessage()    {}
func (*QueryGetCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f06cd066eb8a7bf0, []int{2}
}
func (m *QueryGetCommentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCommentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCommentRequest.Merge(m, src)
}
func (m *QueryGetCommentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCommentRequest proto.InternalMessageInfo

func (m *QueryGetCommentRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetCommentResponse struct {
	Comment Comment `protobuf:"bytes,1,opt,name=Comment,proto3" json:"Comment"`
}

func (m *QueryGetCommentResponse) Reset()         { *m = QueryGetCommentResponse{} }
func (m *QueryGetCommentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCommentResponse) ProtoMessage()    {}
func (*QueryGetCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f06cd066eb8a7bf0, []int{3}
}
func (m *QueryGetCommentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCommentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCommentResponse.Merge(m, src)
}
func (m *QueryGetCommentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCommentResponse proto.InternalMessageInfo

func (m *QueryGetCommentResponse) GetComment() Comment {
	if m != nil {
		return m.Comment
	}
	return Comment{}
}

type QueryAllCommentRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCommentRequest) Reset()         { *m = QueryAllCommentRequest{} }
func (m *QueryAllCommentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCommentRequest) ProtoMessage()    {}
func (*QueryAllCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f06cd066eb8a7bf0, []int{4}
}
func (m *QueryAllCommentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCommentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCommentRequest.Merge(m, src)
}
func (m *QueryAllCommentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCommentRequest proto.InternalMessageInfo

func (m *QueryAllCommentRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCommentResponse struct {
	Comment    []Comment           `protobuf:"bytes,1,rep,name=Comment,proto3" json:"Comment"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCommentResponse) Reset()         { *m = QueryAllCommentResponse{} }
func (m *QueryAllCommentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCommentResponse) ProtoMessage()    {}
func (*QueryAllCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f06cd066eb8a7bf0, []int{5}
}
func (m *QueryAllCommentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCommentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCommentResponse.Merge(m, src)
}
func (m *QueryAllCommentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCommentResponse proto.InternalMessageInfo

func (m *QueryAllCommentResponse) GetComment() []Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *QueryAllCommentResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "yogeshguptakaithal.testchain.testchain.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "yogeshguptakaithal.testchain.testchain.QueryParamsResponse")
	proto.RegisterType((*QueryGetCommentRequest)(nil), "yogeshguptakaithal.testchain.testchain.QueryGetCommentRequest")
	proto.RegisterType((*QueryGetCommentResponse)(nil), "yogeshguptakaithal.testchain.testchain.QueryGetCommentResponse")
	proto.RegisterType((*QueryAllCommentRequest)(nil), "yogeshguptakaithal.testchain.testchain.QueryAllCommentRequest")
	proto.RegisterType((*QueryAllCommentResponse)(nil), "yogeshguptakaithal.testchain.testchain.QueryAllCommentResponse")
}

func init() { proto.RegisterFile("testchain/testchain/query.proto", fileDescriptor_f06cd066eb8a7bf0) }

var fileDescriptor_f06cd066eb8a7bf0 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xe3, 0x10, 0x82, 0x64, 0x24, 0x0e, 0xa6, 0x02, 0x14, 0xa1, 0x6d, 0xf1, 0xa1, 0x54,
	0x1c, 0x6c, 0x52, 0x50, 0x11, 0x7f, 0x04, 0x6a, 0x91, 0xe8, 0x05, 0x89, 0x12, 0x6e, 0x9c, 0x70,
	0x36, 0x96, 0x63, 0xd8, 0x5d, 0x6f, 0x63, 0x07, 0x11, 0x21, 0x2e, 0x3c, 0x01, 0x12, 0x8f, 0xc2,
	0x89, 0x17, 0x40, 0x39, 0x56, 0xe2, 0x00, 0x27, 0x84, 0x12, 0x1e, 0x04, 0xc5, 0x9e, 0x36, 0x29,
	0x1b, 0x20, 0x9b, 0xde, 0xac, 0xcd, 0xcc, 0xf7, 0xfd, 0x3e, 0xcf, 0x38, 0x78, 0xd5, 0x49, 0xeb,
	0xe2, 0xae, 0xd0, 0x19, 0x9f, 0x9e, 0xf6, 0xfb, 0xb2, 0x37, 0x60, 0x79, 0xcf, 0x38, 0x43, 0xd6,
	0x07, 0x46, 0x49, 0xdb, 0x55, 0xfd, 0xdc, 0x89, 0x57, 0x42, 0xbb, 0xae, 0x48, 0xd8, 0x51, 0xe5,
	0xf4, 0xd4, 0x58, 0x51, 0x46, 0x19, 0xdf, 0xc2, 0x27, 0xa7, 0xd0, 0xdd, 0xb8, 0xac, 0x8c, 0x51,
	0x89, 0xe4, 0x22, 0xd7, 0x5c, 0x64, 0x99, 0x71, 0xc2, 0x69, 0x93, 0x59, 0xf8, 0xf5, 0x5a, 0x6c,
	0x6c, 0x6a, 0x2c, 0x6f, 0x0b, 0x2b, 0x83, 0x29, 0x7f, 0xdd, 0x6c, 0x4b, 0x27, 0x9a, 0x3c, 0x17,
	0x4a, 0x67, 0xbe, 0x18, 0x6a, 0xd7, 0xe6, 0x81, 0xe6, 0xa2, 0x27, 0xd2, 0x43, 0xb5, 0x2b, 0xf3,
	0x2a, 0x62, 0x93, 0xa6, 0x32, 0x73, 0xa1, 0x84, 0xae, 0x60, 0xf2, 0x74, 0x62, 0xb3, 0xe7, 0xfb,
	0x5a, 0x72, 0xbf, 0x2f, 0xad, 0xa3, 0x31, 0x3e, 0x7f, 0xec, 0xab, 0xcd, 0x4d, 0x66, 0x25, 0x79,
	0x8c, 0xeb, 0x41, 0xff, 0x12, 0x5a, 0x43, 0x1b, 0x67, 0x37, 0x19, 0x5b, 0xec, 0x2a, 0x58, 0xd0,
	0xd9, 0xa9, 0x0d, 0x7f, 0xac, 0x56, 0x5a, 0xa0, 0x41, 0x37, 0xf0, 0x05, 0x6f, 0xb2, 0x2b, 0xdd,
	0xc3, 0xc0, 0x04, 0xf6, 0xe4, 0x1c, 0xae, 0xea, 0x8e, 0xf7, 0xa8, 0xb5, 0xaa, 0xba, 0x43, 0x5f,
	0xe2, 0x8b, 0x85, 0x4a, 0x40, 0x7a, 0x82, 0xcf, 0xc0, 0x27, 0x60, 0xe2, 0x8b, 0x32, 0x41, 0x1b,
	0x40, 0x1d, 0xaa, 0xd0, 0x17, 0x40, 0xb5, 0x9d, 0x24, 0x7f, 0x50, 0x3d, 0xc2, 0x78, 0x3a, 0x03,
	0x70, 0x5b, 0x67, 0x61, 0x60, 0x6c, 0x32, 0x30, 0x16, 0xb6, 0x04, 0x06, 0xc6, 0xf6, 0x84, 0x92,
	0xd0, 0xdb, 0x9a, 0xe9, 0xa4, 0x9f, 0x10, 0xc4, 0x99, 0xb5, 0x98, 0x17, 0xe7, 0xd4, 0xc9, 0xe3,
	0x90, 0xdd, 0x63, 0xd0, 0x55, 0x0f, 0x7d, 0xf5, 0xbf, 0xd0, 0x81, 0x66, 0x96, 0x7a, 0xf3, 0x5b,
	0x0d, 0x9f, 0xf6, 0xd4, 0xe4, 0x33, 0xc2, 0xf5, 0x30, 0x50, 0x72, 0x67, 0x51, 0xba, 0xe2, 0x8e,
	0x35, 0xee, 0x2e, 0xd5, 0x1b, 0xc8, 0xe8, 0xd6, 0xfb, 0xaf, 0xbf, 0x3e, 0x56, 0xaf, 0x13, 0xc6,
	0x8b, 0x22, 0xfc, 0xef, 0xef, 0x82, 0x0c, 0xd1, 0xd1, 0x05, 0x93, 0xfb, 0xa5, 0x00, 0x0a, 0x5b,
	0xda, 0x78, 0xb0, 0x74, 0x3f, 0x84, 0xb8, 0xe7, 0x43, 0x6c, 0x91, 0x9b, 0x8b, 0x86, 0x80, 0xa7,
	0xcb, 0xdf, 0xea, 0xce, 0x3b, 0xf2, 0x05, 0x61, 0x0c, 0x8a, 0xdb, 0x49, 0x52, 0x32, 0x4d, 0x61,
	0xbb, 0x4b, 0xa6, 0x29, 0xae, 0x2e, 0xbd, 0xe5, 0xd3, 0x34, 0x09, 0x2f, 0x99, 0x66, 0xe7, 0xd9,
	0x70, 0x14, 0xa1, 0x83, 0x51, 0x84, 0x7e, 0x8e, 0x22, 0xf4, 0x61, 0x1c, 0x55, 0x0e, 0xc6, 0x51,
	0xe5, 0xfb, 0x38, 0xaa, 0x3c, 0xbf, 0xad, 0xb4, 0xeb, 0xf6, 0xdb, 0x2c, 0x36, 0xe9, 0xbf, 0x45,
	0xdf, 0xcc, 0x1a, 0x0c, 0x72, 0x69, 0xdb, 0x75, 0xff, 0xf7, 0x76, 0xe3, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x30, 0x69, 0xcf, 0xad, 0xce, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a Comment by id.
	Comment(ctx context.Context, in *QueryGetCommentRequest, opts ...grpc.CallOption) (*QueryGetCommentResponse, error)
	// Queries a list of Comment items.
	CommentAll(ctx context.Context, in *QueryAllCommentRequest, opts ...grpc.CallOption) (*QueryAllCommentResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/yogeshguptakaithal.testchain.testchain.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Comment(ctx context.Context, in *QueryGetCommentRequest, opts ...grpc.CallOption) (*QueryGetCommentResponse, error) {
	out := new(QueryGetCommentResponse)
	err := c.cc.Invoke(ctx, "/yogeshguptakaithal.testchain.testchain.Query/Comment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CommentAll(ctx context.Context, in *QueryAllCommentRequest, opts ...grpc.CallOption) (*QueryAllCommentResponse, error) {
	out := new(QueryAllCommentResponse)
	err := c.cc.Invoke(ctx, "/yogeshguptakaithal.testchain.testchain.Query/CommentAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a Comment by id.
	Comment(context.Context, *QueryGetCommentRequest) (*QueryGetCommentResponse, error)
	// Queries a list of Comment items.
	CommentAll(context.Context, *QueryAllCommentRequest) (*QueryAllCommentResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Comment(ctx context.Context, req *QueryGetCommentRequest) (*QueryGetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comment not implemented")
}
func (*UnimplementedQueryServer) CommentAll(ctx context.Context, req *QueryAllCommentRequest) (*QueryAllCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yogeshguptakaithal.testchain.testchain.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Comment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Comment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yogeshguptakaithal.testchain.testchain.Query/Comment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Comment(ctx, req.(*QueryGetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CommentAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommentAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yogeshguptakaithal.testchain.testchain.Query/CommentAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommentAll(ctx, req.(*QueryAllCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yogeshguptakaithal.testchain.testchain.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Comment",
			Handler:    _Query_Comment_Handler,
		},
		{
			MethodName: "CommentAll",
			Handler:    _Query_CommentAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testchain/testchain/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetCommentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCommentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCommentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCommentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCommentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCommentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Comment.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllCommentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCommentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCommentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCommentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCommentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCommentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Comment) > 0 {
		for iNdEx := len(m.Comment) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Comment[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetCommentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetCommentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Comment.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllCommentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCommentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Comment) > 0 {
		for _, e := range m.Comment {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCommentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCommentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCommentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCommentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCommentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCommentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Comment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCommentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCommentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCommentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCommentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCommentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCommentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comment = append(m.Comment, Comment{})
			if err := m.Comment[len(m.Comment)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
