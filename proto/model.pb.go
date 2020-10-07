// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package proto

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

type Node struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip                   string     `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	RtcPort              uint32     `protobuf:"varint,3,opt,name=rtc_port,json=rtcPort,proto3" json:"rtc_port,omitempty"`
	Stats                *NodeStats `protobuf:"bytes,4,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Node) GetRtcPort() uint32 {
	if m != nil {
		return m.RtcPort
	}
	return 0
}

func (m *Node) GetStats() *NodeStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type NodeStats struct {
	NumRooms             int32    `protobuf:"varint,1,opt,name=num_rooms,json=numRooms,proto3" json:"num_rooms,omitempty"`
	NumClients           int32    `protobuf:"varint,2,opt,name=num_clients,json=numClients,proto3" json:"num_clients,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStats) Reset()         { *m = NodeStats{} }
func (m *NodeStats) String() string { return proto.CompactTextString(m) }
func (*NodeStats) ProtoMessage()    {}
func (*NodeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

func (m *NodeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStats.Unmarshal(m, b)
}
func (m *NodeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStats.Marshal(b, m, deterministic)
}
func (m *NodeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStats.Merge(m, src)
}
func (m *NodeStats) XXX_Size() int {
	return xxx_messageInfo_NodeStats.Size(m)
}
func (m *NodeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStats.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStats proto.InternalMessageInfo

func (m *NodeStats) GetNumRooms() int32 {
	if m != nil {
		return m.NumRooms
	}
	return 0
}

func (m *NodeStats) GetNumClients() int32 {
	if m != nil {
		return m.NumClients
	}
	return 0
}

func init() {
	proto.RegisterType((*Node)(nil), "livekit.Node")
	proto.RegisterType((*NodeStats)(nil), "livekit.NodeStats")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xd0, 0xd0, 0xe6, 0x22, 0x18, 0x3c, 0x05, 0x31, 0x10, 0x75, 0x21, 0x0c, 0xa4,
	0x12, 0xfc, 0x03, 0x98, 0x58, 0x10, 0x32, 0x1b, 0x4b, 0x45, 0x63, 0x0b, 0xac, 0xc6, 0x3e, 0xeb,
	0x7c, 0xee, 0xef, 0x47, 0x76, 0x42, 0xa7, 0xa7, 0xef, 0x3d, 0xeb, 0x3d, 0x1f, 0x34, 0x16, 0x95,
	0x9e, 0x06, 0x4f, 0xc8, 0x28, 0xd6, 0x93, 0x39, 0xe9, 0xa3, 0xe1, 0xed, 0x11, 0x56, 0xef, 0xa8,
	0xb4, 0xb8, 0x86, 0xd2, 0xa8, 0xb6, 0xe8, 0x8a, 0xbe, 0x96, 0xa5, 0x51, 0x99, 0x7d, 0x5b, 0x2e,
	0xec, 0xc5, 0x0d, 0x6c, 0x88, 0xc7, 0xbd, 0x47, 0xe2, 0xf6, 0xa2, 0x2b, 0xfa, 0x2b, 0xb9, 0x26,
	0x1e, 0x3f, 0x90, 0x58, 0xf4, 0x50, 0x05, 0xfe, 0xe6, 0xd0, 0xae, 0xba, 0xa2, 0x6f, 0x9e, 0xc4,
	0xb0, 0x74, 0x0f, 0xa9, 0xf8, 0x33, 0x25, 0x72, 0x7e, 0xb0, 0x7d, 0x83, 0xfa, 0xec, 0x89, 0x5b,
	0xa8, 0x5d, 0xb4, 0x7b, 0x42, 0xb4, 0x21, 0x0f, 0x57, 0x72, 0xe3, 0xa2, 0x95, 0x89, 0xc5, 0x1d,
	0x34, 0x29, 0x1c, 0x27, 0xa3, 0x1d, 0x87, 0xfc, 0x8f, 0x4a, 0x82, 0x8b, 0xf6, 0x75, 0x76, 0x5e,
	0x1e, 0xbe, 0xee, 0x7f, 0x0c, 0xff, 0xc6, 0xc3, 0x30, 0xa2, 0xdd, 0x2d, 0x8b, 0xff, 0xfa, 0x18,
	0x34, 0x9d, 0x34, 0xed, 0xf2, 0xad, 0x87, 0xcb, 0x2c, 0xcf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x52, 0x03, 0x38, 0xe7, 0x01, 0x01, 0x00, 0x00,
}
