// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

package types

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

type Agent struct {
	// Output only. SPIFFE ID of the agent.
	Id *SPIFFEID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The method by which the agent attested.
	AttestationType string `protobuf:"bytes,2,opt,name=attestation_type,json=attestationType,proto3" json:"attestation_type,omitempty"`
	// Output only. The X509-SVID serial number.
	X509SvidSerialNumber string `protobuf:"bytes,3,opt,name=x509svid_serial_number,json=x509svidSerialNumber,proto3" json:"x509svid_serial_number,omitempty"`
	// Output only. The X509-SVID expiration (seconds since Unix epoch).
	X509SvidExpiresAt int64 `protobuf:"varint,4,opt,name=x509svid_expires_at,json=x509svidExpiresAt,proto3" json:"x509svid_expires_at,omitempty"`
	// Output only. The selectors attributed to the agent during attestation.
	Selectors []*Selector `protobuf:"bytes,5,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// Output only. Whether or not the agent is banned.
	Banned               bool     `protobuf:"varint,6,opt,name=banned,proto3" json:"banned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Agent) Reset()         { *m = Agent{} }
func (m *Agent) String() string { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()    {}
func (*Agent) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}

func (m *Agent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Agent.Unmarshal(m, b)
}
func (m *Agent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Agent.Marshal(b, m, deterministic)
}
func (m *Agent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Agent.Merge(m, src)
}
func (m *Agent) XXX_Size() int {
	return xxx_messageInfo_Agent.Size(m)
}
func (m *Agent) XXX_DiscardUnknown() {
	xxx_messageInfo_Agent.DiscardUnknown(m)
}

var xxx_messageInfo_Agent proto.InternalMessageInfo

func (m *Agent) GetId() *SPIFFEID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Agent) GetAttestationType() string {
	if m != nil {
		return m.AttestationType
	}
	return ""
}

func (m *Agent) GetX509SvidSerialNumber() string {
	if m != nil {
		return m.X509SvidSerialNumber
	}
	return ""
}

func (m *Agent) GetX509SvidExpiresAt() int64 {
	if m != nil {
		return m.X509SvidExpiresAt
	}
	return 0
}

func (m *Agent) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *Agent) GetBanned() bool {
	if m != nil {
		return m.Banned
	}
	return false
}

type AgentMask struct {
	// attestation_type field mask
	AttestationType bool `protobuf:"varint,2,opt,name=attestation_type,json=attestationType,proto3" json:"attestation_type,omitempty"`
	// x509svid_serial_number field mask
	X509SvidSerialNumber bool `protobuf:"varint,3,opt,name=x509svid_serial_number,json=x509svidSerialNumber,proto3" json:"x509svid_serial_number,omitempty"`
	// x509svid_expires_at field mask
	X509SvidExpiresAt bool `protobuf:"varint,4,opt,name=x509svid_expires_at,json=x509svidExpiresAt,proto3" json:"x509svid_expires_at,omitempty"`
	// selectors field mask
	Selectors bool `protobuf:"varint,5,opt,name=selectors,proto3" json:"selectors,omitempty"`
	// banned field mask
	Banned               bool     `protobuf:"varint,6,opt,name=banned,proto3" json:"banned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentMask) Reset()         { *m = AgentMask{} }
func (m *AgentMask) String() string { return proto.CompactTextString(m) }
func (*AgentMask) ProtoMessage()    {}
func (*AgentMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{1}
}

func (m *AgentMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentMask.Unmarshal(m, b)
}
func (m *AgentMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentMask.Marshal(b, m, deterministic)
}
func (m *AgentMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentMask.Merge(m, src)
}
func (m *AgentMask) XXX_Size() int {
	return xxx_messageInfo_AgentMask.Size(m)
}
func (m *AgentMask) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentMask.DiscardUnknown(m)
}

var xxx_messageInfo_AgentMask proto.InternalMessageInfo

func (m *AgentMask) GetAttestationType() bool {
	if m != nil {
		return m.AttestationType
	}
	return false
}

func (m *AgentMask) GetX509SvidSerialNumber() bool {
	if m != nil {
		return m.X509SvidSerialNumber
	}
	return false
}

func (m *AgentMask) GetX509SvidExpiresAt() bool {
	if m != nil {
		return m.X509SvidExpiresAt
	}
	return false
}

func (m *AgentMask) GetSelectors() bool {
	if m != nil {
		return m.Selectors
	}
	return false
}

func (m *AgentMask) GetBanned() bool {
	if m != nil {
		return m.Banned
	}
	return false
}

func init() {
	proto.RegisterType((*Agent)(nil), "spire.types.Agent")
	proto.RegisterType((*AgentMask)(nil), "spire.types.AgentMask")
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77) }

var fileDescriptor_56ede974c0020f77 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0xb3, 0xe5, 0x0f, 0x29, 0xcb, 0xe1, 0xaf, 0xab, 0x92, 0x8d, 0x31, 0xb1, 0x21, 0x31,
	0xa9, 0x07, 0xb7, 0x44, 0xf4, 0xe0, 0x11, 0x23, 0x24, 0x1c, 0x34, 0xa6, 0x78, 0xf2, 0xd2, 0x6c,
	0xe9, 0x80, 0x1b, 0x61, 0xdb, 0x74, 0x07, 0x03, 0xef, 0xe0, 0xd3, 0xf9, 0x44, 0xa6, 0x0b, 0x28,
	0x2a, 0xc6, 0x70, 0x9c, 0xf9, 0xbe, 0xd9, 0x9d, 0xdf, 0xe4, 0xa3, 0x35, 0x39, 0x02, 0x8d, 0x22,
	0xcb, 0x53, 0x4c, 0x59, 0xcd, 0x64, 0x2a, 0x07, 0x81, 0xf3, 0x0c, 0xcc, 0xe1, 0xb1, 0x2d, 0xce,
	0x34, 0xcc, 0x30, 0xb0, 0x9d, 0xc0, 0xc0, 0x18, 0x06, 0x98, 0xe6, 0x0b, 0xf7, 0x26, 0x43, 0xa6,
	0x86, 0x43, 0x50, 0xc9, 0xc2, 0xd0, 0x78, 0x75, 0x68, 0xb9, 0x5d, 0x3c, 0xcf, 0x4e, 0xa8, 0xa3,
	0x12, 0x4e, 0x3c, 0xe2, 0xd7, 0xce, 0x0f, 0xc4, 0xda, 0x2f, 0xa2, 0x7f, 0xdf, 0xeb, 0x76, 0x3b,
	0xbd, 0x9b, 0xd0, 0x51, 0x09, 0x3b, 0xa5, 0x3b, 0x12, 0x11, 0x0c, 0x4a, 0x54, 0xa9, 0x8e, 0x0a,
	0x07, 0x77, 0x3c, 0xe2, 0x57, 0xc3, 0xff, 0x6b, 0xfd, 0x87, 0x79, 0x06, 0xec, 0x82, 0xd6, 0x67,
	0x97, 0xcd, 0x2b, 0xf3, 0xa2, 0x92, 0xc8, 0x40, 0xae, 0xe4, 0x38, 0xd2, 0xd3, 0x49, 0x0c, 0x39,
	0x2f, 0xd9, 0x81, 0xfd, 0x95, 0xda, 0xb7, 0xe2, 0x9d, 0xd5, 0x98, 0xa0, 0x7b, 0x1f, 0x53, 0x30,
	0x2b, 0xd6, 0x30, 0x91, 0x44, 0xfe, 0xcf, 0x23, 0x7e, 0x29, 0xdc, 0x5d, 0x49, 0x9d, 0x85, 0xd2,
	0x46, 0xd6, 0xa2, 0xd5, 0x15, 0xb4, 0xe1, 0x65, 0xaf, 0xf4, 0x73, 0xfd, 0xa5, 0x1a, 0x7e, 0xfa,
	0x58, 0x9d, 0x56, 0x62, 0xa9, 0x35, 0x24, 0xbc, 0xe2, 0x11, 0xdf, 0x0d, 0x97, 0x55, 0xe3, 0x8d,
	0xd0, 0xaa, 0x3d, 0xc7, 0xad, 0x34, 0xcf, 0xbf, 0xb2, 0xba, 0xdb, 0xb2, 0xba, 0xdb, 0xb3, 0xba,
	0x9b, 0x58, 0x8f, 0xbe, 0xb2, 0x16, 0xae, 0xbf, 0xa1, 0xae, 0x9b, 0x8f, 0x62, 0xa4, 0xf0, 0x69,
	0x1a, 0x8b, 0x41, 0x3a, 0x59, 0x06, 0x20, 0xb0, 0x17, 0x0a, 0x6c, 0x08, 0x82, 0xef, 0x21, 0x89,
	0x2b, 0xb6, 0xdf, 0x7a, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x46, 0x3a, 0x1a, 0x7a, 0x02, 0x00,
	0x00,
}