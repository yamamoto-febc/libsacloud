// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sacloud.proto

package sacloud

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ApplianceConnectedSwitch struct {
	Id                   int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Scope                types.Scope `protobuf:"varint,2,opt,name=scope,proto3,enum=types.Scope" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ApplianceConnectedSwitch) Reset()         { *m = ApplianceConnectedSwitch{} }
func (m *ApplianceConnectedSwitch) String() string { return proto.CompactTextString(m) }
func (*ApplianceConnectedSwitch) ProtoMessage()    {}
func (*ApplianceConnectedSwitch) Descriptor() ([]byte, []int) {
	return fileDescriptor_sacloud_5822f2e6502f99fa, []int{0}
}
func (m *ApplianceConnectedSwitch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplianceConnectedSwitch.Unmarshal(m, b)
}
func (m *ApplianceConnectedSwitch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplianceConnectedSwitch.Marshal(b, m, deterministic)
}
func (dst *ApplianceConnectedSwitch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplianceConnectedSwitch.Merge(dst, src)
}
func (m *ApplianceConnectedSwitch) XXX_Size() int {
	return xxx_messageInfo_ApplianceConnectedSwitch.Size(m)
}
func (m *ApplianceConnectedSwitch) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplianceConnectedSwitch.DiscardUnknown(m)
}

var xxx_messageInfo_ApplianceConnectedSwitch proto.InternalMessageInfo

func (m *ApplianceConnectedSwitch) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ApplianceConnectedSwitch) GetScope() types.Scope {
	if m != nil {
		return m.Scope
	}
	return types.Scope_SCOPE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*ApplianceConnectedSwitch)(nil), "sacloud.ApplianceConnectedSwitch")
}

func init() { proto.RegisterFile("sacloud.proto", fileDescriptor_sacloud_5822f2e6502f99fa) }

var fileDescriptor_sacloud_5822f2e6502f99fa = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x4c, 0xce,
	0xc9, 0x2f, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x04, 0x4b,
	0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0xc1, 0x24, 0x44, 0x4e, 0xc9, 0x8f, 0x4b, 0xc2, 0xb1, 0xa0, 0x20,
	0x27, 0x33, 0x31, 0x2f, 0x39, 0xd5, 0x39, 0x3f, 0x2f, 0x2f, 0x35, 0xb9, 0x24, 0x35, 0x25, 0xb8,
	0x3c, 0xb3, 0x24, 0x39, 0x43, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x39, 0x88, 0x29, 0x33, 0x45, 0x48, 0x89, 0x8b, 0xb5, 0x38, 0x39, 0xbf, 0x20, 0x55, 0x82, 0x49,
	0x81, 0x51, 0x83, 0xcf, 0x88, 0x47, 0x0f, 0x62, 0x50, 0x30, 0x48, 0x2c, 0x08, 0x22, 0x95, 0xc4,
	0x06, 0x36, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x4b, 0x65, 0x40, 0x83, 0x00, 0x00,
	0x00,
}
