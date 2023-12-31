// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discovery.proto

package discovery

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Discovery_Type int32

const (
	Discovery_CONNECT Discovery_Type = 0
	Discovery_SUCCESS Discovery_Type = 1
	Discovery_FAILED  Discovery_Type = 2
)

var Discovery_Type_name = map[int32]string{
	0: "CONNECT",
	1: "SUCCESS",
	2: "FAILED",
}

var Discovery_Type_value = map[string]int32{
	"CONNECT": 0,
	"SUCCESS": 1,
	"FAILED":  2,
}

func (x Discovery_Type) Enum() *Discovery_Type {
	p := new(Discovery_Type)
	*p = x
	return p
}

func (x Discovery_Type) String() string {
	return proto.EnumName(Discovery_Type_name, int32(x))
}

func (x *Discovery_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Discovery_Type_value, data, "Discovery_Type")
	if err != nil {
		return err
	}
	*x = Discovery_Type(value)
	return nil
}

func (Discovery_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{0, 0}
}

type Discovery struct {
	Type                 *Discovery_Type `protobuf:"varint,1,opt,name=type,enum=discovery.Discovery_Type" json:"type,omitempty"`
	Protocol             *string         `protobuf:"bytes,2,opt,name=pb" json:"pb,omitempty"`
	NodeName             *string         `protobuf:"bytes,3,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Discovery) Reset()         { *m = Discovery{} }
func (m *Discovery) String() string { return proto.CompactTextString(m) }
func (*Discovery) ProtoMessage()    {}
func (*Discovery) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{0}
}

func (m *Discovery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Discovery.Unmarshal(m, b)
}
func (m *Discovery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Discovery.Marshal(b, m, deterministic)
}
func (m *Discovery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Discovery.Merge(m, src)
}
func (m *Discovery) XXX_Size() int {
	return xxx_messageInfo_Discovery.Size(m)
}
func (m *Discovery) XXX_DiscardUnknown() {
	xxx_messageInfo_Discovery.DiscardUnknown(m)
}

var xxx_messageInfo_Discovery proto.InternalMessageInfo

func (m *Discovery) GetType() Discovery_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Discovery_CONNECT
}

func (m *Discovery) GetProtocol() string {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return ""
}

func (m *Discovery) GetNodeName() string {
	if m != nil && m.NodeName != nil {
		return *m.NodeName
	}
	return ""
}

func init() {
	proto.RegisterEnum("discovery.Discovery_Type", Discovery_Type_name, Discovery_Type_value)
	proto.RegisterType((*Discovery)(nil), "discovery.Discovery")
}

func init() { proto.RegisterFile("discovery.proto", fileDescriptor_1e7ff60feb39c8d0) }

var fileDescriptor_1e7ff60feb39c8d0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xc9, 0x2c, 0x4e,
	0xce, 0x2f, 0x4b, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x2d, 0x64, 0xe4, 0xe2, 0x74, 0x81, 0xf1, 0x84, 0x74, 0xb9, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25,
	0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x24, 0xf5, 0x10, 0x1a, 0xe1, 0x6a, 0xf4, 0x42, 0x2a, 0x0b,
	0x52, 0x83, 0xc0, 0xca, 0x84, 0xa4, 0xb8, 0x38, 0xc0, 0x06, 0x26, 0xe7, 0xe7, 0x48, 0x30, 0x29,
	0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x42, 0xd2, 0x5c, 0x9c, 0x79, 0xf9, 0x29, 0xa9, 0xf1, 0x79,
	0x89, 0xb9, 0xa9, 0x12, 0xcc, 0x10, 0x49, 0x90, 0x80, 0x5f, 0x62, 0x6e, 0xaa, 0x92, 0x0e, 0x17,
	0x0b, 0xc8, 0x18, 0x21, 0x6e, 0x2e, 0x76, 0x67, 0x7f, 0x3f, 0x3f, 0x57, 0xe7, 0x10, 0x01, 0x06,
	0x10, 0x27, 0x38, 0xd4, 0xd9, 0xd9, 0x35, 0x38, 0x58, 0x80, 0x51, 0x88, 0x8b, 0x8b, 0xcd, 0xcd,
	0xd1, 0xd3, 0xc7, 0xd5, 0x45, 0x80, 0x09, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x89, 0xaf, 0x65,
	0xc0, 0x00, 0x00, 0x00,
}
