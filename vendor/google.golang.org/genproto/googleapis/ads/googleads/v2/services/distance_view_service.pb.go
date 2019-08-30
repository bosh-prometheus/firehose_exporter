// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/distance_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [DistanceViewService.GetDistanceView][google.ads.googleads.v2.services.DistanceViewService.GetDistanceView].
type GetDistanceViewRequest struct {
	// The resource name of the distance view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDistanceViewRequest) Reset()         { *m = GetDistanceViewRequest{} }
func (m *GetDistanceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetDistanceViewRequest) ProtoMessage()    {}
func (*GetDistanceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4de60c8631d2d9d8, []int{0}
}

func (m *GetDistanceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDistanceViewRequest.Unmarshal(m, b)
}
func (m *GetDistanceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDistanceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetDistanceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDistanceViewRequest.Merge(m, src)
}
func (m *GetDistanceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetDistanceViewRequest.Size(m)
}
func (m *GetDistanceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDistanceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDistanceViewRequest proto.InternalMessageInfo

func (m *GetDistanceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDistanceViewRequest)(nil), "google.ads.googleads.v2.services.GetDistanceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/distance_view_service.proto", fileDescriptor_4de60c8631d2d9d8)
}

var fileDescriptor_4de60c8631d2d9d8 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x4a, 0xe3, 0x40,
	0x18, 0x27, 0x59, 0x58, 0xd8, 0xb0, 0xcb, 0x42, 0x16, 0x76, 0x4b, 0x76, 0x0f, 0xa5, 0xdb, 0xc3,
	0xd2, 0xc3, 0x0c, 0x1b, 0x11, 0x65, 0xb4, 0x87, 0x14, 0xa1, 0x9e, 0xa4, 0x54, 0xc8, 0x41, 0x02,
	0x65, 0x4c, 0x86, 0x30, 0xd0, 0xcc, 0xd4, 0x7c, 0xd3, 0xf4, 0x20, 0x5e, 0x7c, 0x05, 0xdf, 0xc0,
	0xa3, 0x77, 0x5f, 0xa2, 0x57, 0x5f, 0xc1, 0x53, 0x5f, 0x42, 0x49, 0x27, 0x93, 0x56, 0x6d, 0xe9,
	0xed, 0xc7, 0xf7, 0xfd, 0xfe, 0xcc, 0xf7, 0x4b, 0x9c, 0xe3, 0x54, 0xca, 0x74, 0xcc, 0x30, 0x4d,
	0x00, 0x6b, 0x58, 0xa2, 0xc2, 0xc7, 0xc0, 0xf2, 0x82, 0xc7, 0x0c, 0x70, 0xc2, 0x41, 0x51, 0x11,
	0xb3, 0x51, 0xc1, 0xd9, 0x6c, 0x54, 0x8d, 0xd1, 0x24, 0x97, 0x4a, 0xba, 0x4d, 0x2d, 0x41, 0x34,
	0x01, 0x54, 0xab, 0x51, 0xe1, 0x23, 0xa3, 0xf6, 0xf6, 0xb7, 0xf9, 0xe7, 0x0c, 0xe4, 0x34, 0xff,
	0x10, 0xa0, 0x8d, 0xbd, 0x3f, 0x46, 0x36, 0xe1, 0x98, 0x0a, 0x21, 0x15, 0x55, 0x5c, 0x0a, 0xa8,
	0xb6, 0xbf, 0xd6, 0xb6, 0xf1, 0x98, 0x33, 0xa1, 0xf4, 0xa2, 0xd5, 0x75, 0x7e, 0xf6, 0x99, 0x3a,
	0xa9, 0x0c, 0x43, 0xce, 0x66, 0x43, 0x76, 0x35, 0x65, 0xa0, 0xdc, 0xbf, 0xce, 0x37, 0x93, 0x38,
	0x12, 0x34, 0x63, 0x0d, 0xab, 0x69, 0xfd, 0xfb, 0x32, 0xfc, 0x6a, 0x86, 0x67, 0x34, 0x63, 0xfe,
	0xc2, 0x72, 0x7e, 0xac, 0x8b, 0xcf, 0xf5, 0x15, 0xee, 0xa3, 0xe5, 0x7c, 0x7f, 0xe7, 0xeb, 0x1e,
	0xa2, 0x5d, 0xb7, 0xa3, 0xcd, 0x4f, 0xf1, 0xf0, 0x56, 0x65, 0xdd, 0x09, 0x5a, 0xd7, 0xb5, 0x0e,
	0x6e, 0x9f, 0x9e, 0xef, 0xec, 0xff, 0x2e, 0x2e, 0x7b, 0xbb, 0x7e, 0x73, 0x46, 0x37, 0x9e, 0x82,
	0x92, 0x19, 0xcb, 0x01, 0x77, 0xea, 0x22, 0x4b, 0x11, 0xe0, 0xce, 0x8d, 0xf7, 0x7b, 0x1e, 0x34,
	0x56, 0x01, 0x15, 0x9a, 0x70, 0x40, 0xb1, 0xcc, 0x7a, 0x2f, 0x96, 0xd3, 0x8e, 0x65, 0xb6, 0xf3,
	0x8c, 0x5e, 0x63, 0x43, 0x25, 0x83, 0xb2, 0xee, 0x81, 0x75, 0x71, 0x5a, 0xa9, 0x53, 0x39, 0xa6,
	0x22, 0x45, 0x32, 0x4f, 0x71, 0xca, 0xc4, 0xf2, 0x63, 0xe0, 0x55, 0xde, 0xf6, 0xbf, 0xeb, 0xc8,
	0x80, 0x7b, 0xfb, 0x53, 0x3f, 0x08, 0x1e, 0xec, 0x66, 0x5f, 0x1b, 0x06, 0x09, 0x20, 0x0d, 0x4b,
	0x14, 0xfa, 0xa8, 0x0a, 0x86, 0xb9, 0xa1, 0x44, 0x41, 0x02, 0x51, 0x4d, 0x89, 0x42, 0x3f, 0x32,
	0x94, 0x85, 0xdd, 0xd6, 0x73, 0x42, 0x82, 0x04, 0x08, 0xa9, 0x49, 0x84, 0x84, 0x3e, 0x21, 0x86,
	0x76, 0xf9, 0x79, 0xf9, 0xce, 0xbd, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xd9, 0x38, 0x2a,
	0x04, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DistanceViewServiceClient is the client API for DistanceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DistanceViewServiceClient interface {
	// Returns the attributes of the requested distance view.
	GetDistanceView(ctx context.Context, in *GetDistanceViewRequest, opts ...grpc.CallOption) (*resources.DistanceView, error)
}

type distanceViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewDistanceViewServiceClient(cc *grpc.ClientConn) DistanceViewServiceClient {
	return &distanceViewServiceClient{cc}
}

func (c *distanceViewServiceClient) GetDistanceView(ctx context.Context, in *GetDistanceViewRequest, opts ...grpc.CallOption) (*resources.DistanceView, error) {
	out := new(resources.DistanceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.DistanceViewService/GetDistanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistanceViewServiceServer is the server API for DistanceViewService service.
type DistanceViewServiceServer interface {
	// Returns the attributes of the requested distance view.
	GetDistanceView(context.Context, *GetDistanceViewRequest) (*resources.DistanceView, error)
}

func RegisterDistanceViewServiceServer(s *grpc.Server, srv DistanceViewServiceServer) {
	s.RegisterService(&_DistanceViewService_serviceDesc, srv)
}

func _DistanceViewService_GetDistanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDistanceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistanceViewServiceServer).GetDistanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.DistanceViewService/GetDistanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistanceViewServiceServer).GetDistanceView(ctx, req.(*GetDistanceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistanceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.DistanceViewService",
	HandlerType: (*DistanceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDistanceView",
			Handler:    _DistanceViewService_GetDistanceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/distance_view_service.proto",
}