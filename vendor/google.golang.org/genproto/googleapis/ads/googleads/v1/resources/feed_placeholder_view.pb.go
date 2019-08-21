// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/feed_placeholder_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A feed placeholder view.
type FeedPlaceholderView struct {
	// The resource name of the feed placeholder view.
	// Feed placeholder view resource names have the form:
	//
	// `customers/{customer_id}/feedPlaceholderViews/{placeholder_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The placeholder type of the feed placeholder view.
	PlaceholderType      enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,2,opt,name=placeholder_type,json=placeholderType,proto3,enum=google.ads.googleads.v1.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *FeedPlaceholderView) Reset()         { *m = FeedPlaceholderView{} }
func (m *FeedPlaceholderView) String() string { return proto.CompactTextString(m) }
func (*FeedPlaceholderView) ProtoMessage()    {}
func (*FeedPlaceholderView) Descriptor() ([]byte, []int) {
	return fileDescriptor_84882bb262f263ab, []int{0}
}

func (m *FeedPlaceholderView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedPlaceholderView.Unmarshal(m, b)
}
func (m *FeedPlaceholderView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedPlaceholderView.Marshal(b, m, deterministic)
}
func (m *FeedPlaceholderView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedPlaceholderView.Merge(m, src)
}
func (m *FeedPlaceholderView) XXX_Size() int {
	return xxx_messageInfo_FeedPlaceholderView.Size(m)
}
func (m *FeedPlaceholderView) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedPlaceholderView.DiscardUnknown(m)
}

var xxx_messageInfo_FeedPlaceholderView proto.InternalMessageInfo

func (m *FeedPlaceholderView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *FeedPlaceholderView) GetPlaceholderType() enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderType
	}
	return enums.PlaceholderTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*FeedPlaceholderView)(nil), "google.ads.googleads.v1.resources.FeedPlaceholderView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/feed_placeholder_view.proto", fileDescriptor_84882bb262f263ab)
}

var fileDescriptor_84882bb262f263ab = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0x3b, 0x31,
	0x14, 0xc5, 0x99, 0xf9, 0xc3, 0x1f, 0x1c, 0xfc, 0xa2, 0x6e, 0x4a, 0x71, 0xd1, 0x2a, 0x85, 0xae,
	0x12, 0x46, 0x5d, 0x45, 0x5c, 0x4c, 0x41, 0x0b, 0x2e, 0xa4, 0x14, 0x99, 0x85, 0x0c, 0x94, 0xd8,
	0x5c, 0xe3, 0xc0, 0xe4, 0x83, 0xc9, 0xb4, 0xa5, 0x0f, 0xe0, 0x8b, 0xb8, 0xd3, 0x47, 0xf1, 0x51,
	0x7c, 0x0a, 0xe9, 0xc4, 0xc4, 0x5a, 0xac, 0xee, 0x0e, 0x37, 0xe7, 0x77, 0xcf, 0xbd, 0x37, 0xd1,
	0x05, 0x57, 0x8a, 0x17, 0x80, 0x29, 0x33, 0xd8, 0xca, 0xa5, 0x9a, 0xc5, 0xb8, 0x04, 0xa3, 0xa6,
	0xe5, 0x04, 0x0c, 0x7e, 0x00, 0x60, 0x63, 0x5d, 0xd0, 0x09, 0x3c, 0xaa, 0x82, 0x41, 0x39, 0x9e,
	0xe5, 0x30, 0x47, 0xba, 0x54, 0x95, 0x6a, 0x74, 0x2c, 0x83, 0x28, 0x33, 0xc8, 0xe3, 0x68, 0x16,
	0x23, 0x8f, 0xb7, 0xce, 0x36, 0x25, 0x80, 0x9c, 0x0a, 0x83, 0x57, 0x1b, 0x57, 0x0b, 0x0d, 0xb6,
	0x71, 0xeb, 0xd0, 0x51, 0x3a, 0xc7, 0x54, 0x4a, 0x55, 0xd1, 0x2a, 0x57, 0xd2, 0xd8, 0xd7, 0xa3,
	0x97, 0x20, 0x3a, 0xb8, 0x02, 0x60, 0xc3, 0x2f, 0x38, 0xcd, 0x61, 0xde, 0x38, 0x8e, 0x76, 0x5c,
	0xf0, 0x58, 0x52, 0x01, 0xcd, 0xa0, 0x1d, 0xf4, 0xb6, 0x46, 0xdb, 0xae, 0x78, 0x43, 0x05, 0x34,
	0x44, 0xb4, 0xbf, 0x1e, 0xda, 0x0c, 0xdb, 0x41, 0x6f, 0xf7, 0xa4, 0x8f, 0x36, 0xad, 0x53, 0xcf,
	0x8a, 0x56, 0xe2, 0x6e, 0x17, 0x1a, 0x2e, 0xe5, 0x54, 0xac, 0xd7, 0x46, 0x7b, 0xfa, 0x7b, 0xa1,
	0xff, 0x14, 0x46, 0xdd, 0x89, 0x12, 0xe8, 0xcf, 0x4b, 0xf5, 0x9b, 0x3f, 0xac, 0x34, 0x5c, 0xee,
	0x3b, 0x0c, 0xee, 0xae, 0x3f, 0x71, 0xae, 0x0a, 0x2a, 0x39, 0x52, 0x25, 0xc7, 0x1c, 0x64, 0x7d,
	0x0d, 0x77, 0x55, 0x9d, 0x9b, 0x5f, 0xbe, 0xf1, 0xdc, 0xab, 0xe7, 0xf0, 0xdf, 0x20, 0x49, 0x5e,
	0xc3, 0xce, 0xc0, 0xb6, 0x4c, 0x98, 0x41, 0x56, 0x2e, 0x55, 0x1a, 0xa3, 0x91, 0x73, 0xbe, 0x39,
	0x4f, 0x96, 0x30, 0x93, 0x79, 0x4f, 0x96, 0xc6, 0x99, 0xf7, 0xbc, 0x87, 0x5d, 0xfb, 0x40, 0x48,
	0xc2, 0x0c, 0x21, 0xde, 0x45, 0x48, 0x1a, 0x13, 0xe2, 0x7d, 0xf7, 0xff, 0xeb, 0x61, 0x4f, 0x3f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xf2, 0xef, 0x8d, 0x72, 0x02, 0x00, 0x00,
}
