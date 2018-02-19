// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto

It has these top-level messages:
	IPTagging
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IPTagging_RequestType int32

const (
	IPTagging_BOTH     IPTagging_RequestType = 0
	IPTagging_INTERNAL IPTagging_RequestType = 1
	IPTagging_EXTERNAL IPTagging_RequestType = 2
)

var IPTagging_RequestType_name = map[int32]string{
	0: "BOTH",
	1: "INTERNAL",
	2: "EXTERNAL",
}
var IPTagging_RequestType_value = map[string]int32{
	"BOTH":     0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
}

func (x IPTagging_RequestType) String() string {
	return proto.EnumName(IPTagging_RequestType_name, int32(x))
}
func (IPTagging_RequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// This is an HTTP filter which enables Envoy to tag requests with extra
// information such as location, cloud source, and any extra data. This is
// useful to prevent against DDoS.
type IPTagging struct {
	// The type of requests the filter should apply to. The supported types
	// are internal, external or both. A request is considered internal if
	// x-envoy-internal is set to true. If x-envoy-internal is not set or
	// false, a request is considered external. The filter defaults to both,
	// and it will apply to all request types.
	RequestType IPTagging_RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,enum=envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType" json:"request_type,omitempty"`
	IpTags      []*IPTagging_IPTag    `protobuf:"bytes,2,rep,name=ip_tags,json=ipTags" json:"ip_tags,omitempty"`
}

func (m *IPTagging) Reset()                    { *m = IPTagging{} }
func (m *IPTagging) String() string            { return proto.CompactTextString(m) }
func (*IPTagging) ProtoMessage()               {}
func (*IPTagging) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IPTagging) GetRequestType() IPTagging_RequestType {
	if m != nil {
		return m.RequestType
	}
	return IPTagging_BOTH
}

func (m *IPTagging) GetIpTags() []*IPTagging_IPTag {
	if m != nil {
		return m.IpTags
	}
	return nil
}

type IPTagging_IPTag struct {
	// Specifies the ip tag name to apply.
	IpTagName string `protobuf:"bytes,1,opt,name=ip_tag_name,json=ipTagName" json:"ip_tag_name,omitempty"`
	// A list of IP address and subnet masks that will be tagged with the
	// ip_tag_name. Both IPv4 and IPv6 CIDR addresses are allowed here.
	IpList []*envoy_api_v2_core.CidrRange `protobuf:"bytes,2,rep,name=ip_list,json=ipList" json:"ip_list,omitempty"`
}

func (m *IPTagging_IPTag) Reset()                    { *m = IPTagging_IPTag{} }
func (m *IPTagging_IPTag) String() string            { return proto.CompactTextString(m) }
func (*IPTagging_IPTag) ProtoMessage()               {}
func (*IPTagging_IPTag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *IPTagging_IPTag) GetIpTagName() string {
	if m != nil {
		return m.IpTagName
	}
	return ""
}

func (m *IPTagging_IPTag) GetIpList() []*envoy_api_v2_core.CidrRange {
	if m != nil {
		return m.IpList
	}
	return nil
}

func init() {
	proto.RegisterType((*IPTagging)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging")
	proto.RegisterType((*IPTagging_IPTag)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging.IPTag")
	proto.RegisterEnum("envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType", IPTagging_RequestType_name, IPTagging_RequestType_value)
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0xdf, 0xf5, 0x9d, 0x73, 0x4d, 0x87, 0x8c, 0x9c, 0xc6, 0x10, 0x1d, 0x3b, 0xc8, 0x4e,
	0x09, 0x74, 0xc8, 0x4e, 0x1e, 0x9c, 0x0c, 0x1c, 0x8c, 0x3a, 0x42, 0x0f, 0xe2, 0xc1, 0x1a, 0xd7,
	0x2c, 0x06, 0xb6, 0x36, 0x26, 0xb1, 0xd0, 0xef, 0xe9, 0x07, 0x92, 0x24, 0x45, 0x7b, 0xd4, 0x5b,
	0xfe, 0xe1, 0xf9, 0xfd, 0x7f, 0xc9, 0x03, 0x16, 0xac, 0xa8, 0xca, 0x1a, 0xef, 0xca, 0x62, 0x2f,
	0x38, 0xde, 0x8b, 0x83, 0x61, 0x0a, 0xbf, 0x19, 0x23, 0xb1, 0x90, 0x99, 0xa1, 0x9c, 0x8b, 0x82,
	0xe3, 0x2a, 0x6e, 0x25, 0x24, 0x55, 0x69, 0x4a, 0x78, 0xe5, 0x40, 0xe4, 0x41, 0xe4, 0x41, 0x64,
	0x41, 0xd4, 0x1a, 0xad, 0xe2, 0xf1, 0xa5, 0x17, 0x50, 0x29, 0x6c, 0xcd, 0xae, 0x54, 0x0c, 0xd3,
	0x3c, 0x57, 0x4c, 0x6b, 0x5f, 0x34, 0xfd, 0x0c, 0x40, 0xb8, 0xde, 0xa6, 0x9e, 0x80, 0x2f, 0x60,
	0xa0, 0xd8, 0xfb, 0x07, 0xd3, 0x26, 0x33, 0xb5, 0x64, 0xa3, 0xce, 0xa4, 0x33, 0x3b, 0x8b, 0x6f,
	0xd0, 0xef, 0x6c, 0xe8, 0xbb, 0x08, 0x11, 0xdf, 0x92, 0xd6, 0x92, 0x91, 0x48, 0xfd, 0x04, 0xb8,
	0x05, 0xa7, 0x9e, 0xd1, 0xa3, 0x60, 0xf2, 0x7f, 0x16, 0xc5, 0x8b, 0xbf, 0x97, 0xbb, 0x13, 0xe9,
	0x09, 0x99, 0x52, 0xae, 0xc7, 0xcf, 0xe0, 0xc4, 0x5d, 0xc0, 0x0b, 0x10, 0x79, 0x22, 0x2b, 0xe8,
	0xd1, 0xbf, 0x3d, 0x24, 0xa1, 0x9b, 0x4a, 0xe8, 0x91, 0xc1, 0x6b, 0xa7, 0x3e, 0x08, 0x6d, 0x1a,
	0xf5, 0x79, 0xa3, 0xa6, 0x52, 0x58, 0x81, 0xdd, 0x0e, 0xba, 0x13, 0xb9, 0x22, 0xb4, 0xe0, 0xcc,
	0xf6, 0x6f, 0x84, 0x36, 0xd3, 0x39, 0x88, 0x5a, 0xbf, 0x81, 0x7d, 0xd0, 0x5d, 0x3e, 0xa4, 0xf7,
	0xc3, 0x7f, 0x70, 0x00, 0xfa, 0xeb, 0x24, 0x5d, 0x91, 0xe4, 0x76, 0x33, 0xec, 0xd8, 0xb4, 0x7a,
	0x6c, 0x52, 0xb0, 0xec, 0x3e, 0x05, 0x55, 0xfc, 0xda, 0x73, 0x3b, 0x9e, 0x7f, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x8b, 0x95, 0xc1, 0x5c, 0xe7, 0x01, 0x00, 0x00,
}