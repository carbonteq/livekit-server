// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

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

type CreateRoomRequest struct {
	Room                 string   `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomRequest) Reset()         { *m = CreateRoomRequest{} }
func (m *CreateRoomRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoomRequest) ProtoMessage()    {}
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *CreateRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomRequest.Unmarshal(m, b)
}
func (m *CreateRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomRequest.Marshal(b, m, deterministic)
}
func (m *CreateRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomRequest.Merge(m, src)
}
func (m *CreateRoomRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoomRequest.Size(m)
}
func (m *CreateRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomRequest proto.InternalMessageInfo

func (m *CreateRoomRequest) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

type CreateRoomResponse struct {
	Room                 string   `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	NodeIp               string   `protobuf:"bytes,2,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip,omitempty"`
	NodeRtcPort          uint32   `protobuf:"varint,3,opt,name=node_rtc_port,json=nodeRtcPort,proto3" json:"node_rtc_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomResponse) Reset()         { *m = CreateRoomResponse{} }
func (m *CreateRoomResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRoomResponse) ProtoMessage()    {}
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *CreateRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomResponse.Unmarshal(m, b)
}
func (m *CreateRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomResponse.Marshal(b, m, deterministic)
}
func (m *CreateRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomResponse.Merge(m, src)
}
func (m *CreateRoomResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRoomResponse.Size(m)
}
func (m *CreateRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomResponse proto.InternalMessageInfo

func (m *CreateRoomResponse) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func (m *CreateRoomResponse) GetNodeIp() string {
	if m != nil {
		return m.NodeIp
	}
	return ""
}

func (m *CreateRoomResponse) GetNodeRtcPort() uint32 {
	if m != nil {
		return m.NodeRtcPort
	}
	return 0
}

type JoinRoomRequest struct {
	Room                 string   `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomRequest) Reset()         { *m = JoinRoomRequest{} }
func (m *JoinRoomRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRoomRequest) ProtoMessage()    {}
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *JoinRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomRequest.Unmarshal(m, b)
}
func (m *JoinRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomRequest.Marshal(b, m, deterministic)
}
func (m *JoinRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomRequest.Merge(m, src)
}
func (m *JoinRoomRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRoomRequest.Size(m)
}
func (m *JoinRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomRequest proto.InternalMessageInfo

func (m *JoinRoomRequest) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func (m *JoinRoomRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type JoinRoomResponse struct {
	NodeIp               string   `protobuf:"bytes,1,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip,omitempty"`
	NodeRtcPort          uint32   `protobuf:"varint,2,opt,name=node_rtc_port,json=nodeRtcPort,proto3" json:"node_rtc_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomResponse) Reset()         { *m = JoinRoomResponse{} }
func (m *JoinRoomResponse) String() string { return proto.CompactTextString(m) }
func (*JoinRoomResponse) ProtoMessage()    {}
func (*JoinRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *JoinRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomResponse.Unmarshal(m, b)
}
func (m *JoinRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomResponse.Marshal(b, m, deterministic)
}
func (m *JoinRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomResponse.Merge(m, src)
}
func (m *JoinRoomResponse) XXX_Size() int {
	return xxx_messageInfo_JoinRoomResponse.Size(m)
}
func (m *JoinRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomResponse proto.InternalMessageInfo

func (m *JoinRoomResponse) GetNodeIp() string {
	if m != nil {
		return m.NodeIp
	}
	return ""
}

func (m *JoinRoomResponse) GetNodeRtcPort() uint32 {
	if m != nil {
		return m.NodeRtcPort
	}
	return 0
}

type DeleteRoomRequest struct {
	Room                 string   `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoomRequest) Reset()         { *m = DeleteRoomRequest{} }
func (m *DeleteRoomRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRoomRequest) ProtoMessage()    {}
func (*DeleteRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *DeleteRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoomRequest.Unmarshal(m, b)
}
func (m *DeleteRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoomRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoomRequest.Merge(m, src)
}
func (m *DeleteRoomRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRoomRequest.Size(m)
}
func (m *DeleteRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoomRequest proto.InternalMessageInfo

func (m *DeleteRoomRequest) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

type DeleteRoomResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoomResponse) Reset()         { *m = DeleteRoomResponse{} }
func (m *DeleteRoomResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteRoomResponse) ProtoMessage()    {}
func (*DeleteRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *DeleteRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoomResponse.Unmarshal(m, b)
}
func (m *DeleteRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoomResponse.Marshal(b, m, deterministic)
}
func (m *DeleteRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoomResponse.Merge(m, src)
}
func (m *DeleteRoomResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteRoomResponse.Size(m)
}
func (m *DeleteRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoomResponse proto.InternalMessageInfo

type SignalRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalRequest) Reset()         { *m = SignalRequest{} }
func (m *SignalRequest) String() string { return proto.CompactTextString(m) }
func (*SignalRequest) ProtoMessage()    {}
func (*SignalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *SignalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalRequest.Unmarshal(m, b)
}
func (m *SignalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalRequest.Marshal(b, m, deterministic)
}
func (m *SignalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalRequest.Merge(m, src)
}
func (m *SignalRequest) XXX_Size() int {
	return xxx_messageInfo_SignalRequest.Size(m)
}
func (m *SignalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignalRequest proto.InternalMessageInfo

type SignalResponse struct {
	// Types that are valid to be assigned to Payload:
	//	*SignalResponse_Desc
	//	*SignalResponse_Trickle
	Payload              isSignalResponse_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SignalResponse) Reset()         { *m = SignalResponse{} }
func (m *SignalResponse) String() string { return proto.CompactTextString(m) }
func (*SignalResponse) ProtoMessage()    {}
func (*SignalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *SignalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalResponse.Unmarshal(m, b)
}
func (m *SignalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalResponse.Marshal(b, m, deterministic)
}
func (m *SignalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalResponse.Merge(m, src)
}
func (m *SignalResponse) XXX_Size() int {
	return xxx_messageInfo_SignalResponse.Size(m)
}
func (m *SignalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignalResponse proto.InternalMessageInfo

type isSignalResponse_Payload interface {
	isSignalResponse_Payload()
}

type SignalResponse_Desc struct {
	Desc *SessionDescription `protobuf:"bytes,1,opt,name=desc,proto3,oneof"`
}

type SignalResponse_Trickle struct {
	Trickle *TrickleRequest `protobuf:"bytes,2,opt,name=trickle,proto3,oneof"`
}

func (*SignalResponse_Desc) isSignalResponse_Payload() {}

func (*SignalResponse_Trickle) isSignalResponse_Payload() {}

func (m *SignalResponse) GetPayload() isSignalResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SignalResponse) GetDesc() *SessionDescription {
	if x, ok := m.GetPayload().(*SignalResponse_Desc); ok {
		return x.Desc
	}
	return nil
}

func (m *SignalResponse) GetTrickle() *TrickleRequest {
	if x, ok := m.GetPayload().(*SignalResponse_Trickle); ok {
		return x.Trickle
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignalResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignalResponse_Desc)(nil),
		(*SignalResponse_Trickle)(nil),
	}
}

type TrickleRequest struct {
	Candidate            string   `protobuf:"bytes,1,opt,name=candidate,proto3" json:"candidate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrickleRequest) Reset()         { *m = TrickleRequest{} }
func (m *TrickleRequest) String() string { return proto.CompactTextString(m) }
func (*TrickleRequest) ProtoMessage()    {}
func (*TrickleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *TrickleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrickleRequest.Unmarshal(m, b)
}
func (m *TrickleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrickleRequest.Marshal(b, m, deterministic)
}
func (m *TrickleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrickleRequest.Merge(m, src)
}
func (m *TrickleRequest) XXX_Size() int {
	return xxx_messageInfo_TrickleRequest.Size(m)
}
func (m *TrickleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TrickleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TrickleRequest proto.InternalMessageInfo

func (m *TrickleRequest) GetCandidate() string {
	if m != nil {
		return m.Candidate
	}
	return ""
}

type TrickleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrickleResponse) Reset()         { *m = TrickleResponse{} }
func (m *TrickleResponse) String() string { return proto.CompactTextString(m) }
func (*TrickleResponse) ProtoMessage()    {}
func (*TrickleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *TrickleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrickleResponse.Unmarshal(m, b)
}
func (m *TrickleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrickleResponse.Marshal(b, m, deterministic)
}
func (m *TrickleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrickleResponse.Merge(m, src)
}
func (m *TrickleResponse) XXX_Size() int {
	return xxx_messageInfo_TrickleResponse.Size(m)
}
func (m *TrickleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TrickleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TrickleResponse proto.InternalMessageInfo

type SessionDescription struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Sdp                  []byte   `protobuf:"bytes,2,opt,name=sdp,proto3" json:"sdp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionDescription) Reset()         { *m = SessionDescription{} }
func (m *SessionDescription) String() string { return proto.CompactTextString(m) }
func (*SessionDescription) ProtoMessage()    {}
func (*SessionDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{10}
}

func (m *SessionDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionDescription.Unmarshal(m, b)
}
func (m *SessionDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionDescription.Marshal(b, m, deterministic)
}
func (m *SessionDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionDescription.Merge(m, src)
}
func (m *SessionDescription) XXX_Size() int {
	return xxx_messageInfo_SessionDescription.Size(m)
}
func (m *SessionDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionDescription.DiscardUnknown(m)
}

var xxx_messageInfo_SessionDescription proto.InternalMessageInfo

func (m *SessionDescription) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SessionDescription) GetSdp() []byte {
	if m != nil {
		return m.Sdp
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRoomRequest)(nil), "livekit.CreateRoomRequest")
	proto.RegisterType((*CreateRoomResponse)(nil), "livekit.CreateRoomResponse")
	proto.RegisterType((*JoinRoomRequest)(nil), "livekit.JoinRoomRequest")
	proto.RegisterType((*JoinRoomResponse)(nil), "livekit.JoinRoomResponse")
	proto.RegisterType((*DeleteRoomRequest)(nil), "livekit.DeleteRoomRequest")
	proto.RegisterType((*DeleteRoomResponse)(nil), "livekit.DeleteRoomResponse")
	proto.RegisterType((*SignalRequest)(nil), "livekit.SignalRequest")
	proto.RegisterType((*SignalResponse)(nil), "livekit.SignalResponse")
	proto.RegisterType((*TrickleRequest)(nil), "livekit.TrickleRequest")
	proto.RegisterType((*TrickleResponse)(nil), "livekit.TrickleResponse")
	proto.RegisterType((*SessionDescription)(nil), "livekit.SessionDescription")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x4e, 0xdb, 0x40,
	0x10, 0xc5, 0x40, 0x31, 0x99, 0x34, 0x84, 0x8c, 0x2a, 0xe1, 0x3a, 0x3d, 0x20, 0x5f, 0x42, 0x0f,
	0x35, 0x6a, 0xb8, 0x55, 0x95, 0xaa, 0x06, 0x90, 0xa0, 0x17, 0xaa, 0x0d, 0xa7, 0x5e, 0x22, 0xb3,
	0x1e, 0xe8, 0x0a, 0xc7, 0xeb, 0xee, 0x2e, 0x48, 0x9c, 0x7b, 0xee, 0x1f, 0xf6, 0x63, 0x2a, 0xaf,
	0xd7, 0x89, 0x5b, 0xa7, 0xe1, 0xe4, 0xf5, 0xcc, 0xdb, 0x99, 0x37, 0xf3, 0x9e, 0x16, 0x7a, 0x9a,
	0xd4, 0xa3, 0xe0, 0x14, 0x17, 0x4a, 0x1a, 0x89, 0x7e, 0x26, 0x1e, 0xe9, 0x5e, 0x98, 0x68, 0x04,
	0x83, 0x53, 0x45, 0x89, 0x21, 0x26, 0xe5, 0x9c, 0xd1, 0x8f, 0x07, 0xd2, 0x06, 0x11, 0xb6, 0x95,
	0x94, 0xf3, 0xc0, 0x3b, 0xf4, 0x8e, 0x3a, 0xcc, 0x9e, 0x23, 0x02, 0x6c, 0x02, 0x75, 0x21, 0x73,
	0x4d, 0xab, 0x90, 0x78, 0x00, 0x7e, 0x2e, 0x53, 0x9a, 0x89, 0x22, 0xd8, 0xb4, 0xe1, 0x9d, 0xf2,
	0xf7, 0xb2, 0xc0, 0x08, 0x7a, 0x36, 0xa1, 0x0c, 0x9f, 0x15, 0x52, 0x99, 0x60, 0xeb, 0xd0, 0x3b,
	0xea, 0xb1, 0x6e, 0x19, 0x64, 0x86, 0x7f, 0x95, 0xca, 0x44, 0x13, 0xe8, 0x7f, 0x91, 0x22, 0x7f,
	0x86, 0x0d, 0x0e, 0xa1, 0xc3, 0x33, 0x41, 0xb9, 0x99, 0x89, 0xd4, 0x75, 0xd9, 0xad, 0x02, 0x97,
	0x69, 0x74, 0x05, 0xfb, 0xcb, 0x1a, 0x8e, 0x68, 0x83, 0x94, 0xb7, 0x9e, 0xd4, 0x66, 0x9b, 0xd4,
	0x08, 0x06, 0x67, 0x94, 0xd1, 0xf3, 0x4b, 0x7a, 0x05, 0xd8, 0x04, 0x56, 0xbd, 0xa3, 0x3e, 0xf4,
	0xa6, 0xe2, 0x2e, 0x4f, 0x32, 0x77, 0x35, 0xfa, 0xe9, 0xc1, 0x5e, 0x1d, 0x71, 0xfc, 0xde, 0xc3,
	0x76, 0x4a, 0x9a, 0xdb, 0x6a, 0xdd, 0xf1, 0x30, 0x76, 0xfa, 0xc4, 0x53, 0xd2, 0x5a, 0xc8, 0xfc,
	0x8c, 0x34, 0x57, 0xa2, 0x30, 0x42, 0xe6, 0x17, 0x1b, 0xcc, 0x42, 0xf1, 0x04, 0x7c, 0xa3, 0x04,
	0xbf, 0xcf, 0xc8, 0x72, 0xee, 0x8e, 0x0f, 0x16, 0xb7, 0xae, 0xab, 0xb8, 0xeb, 0x77, 0xb1, 0xc1,
	0x6a, 0xe4, 0xa4, 0x03, 0x7e, 0x91, 0x3c, 0x65, 0x32, 0x49, 0xa3, 0x18, 0xf6, 0xfe, 0xc6, 0xe1,
	0x1b, 0xe8, 0xf0, 0x24, 0x4f, 0x45, 0x9a, 0x18, 0x72, 0x73, 0x2d, 0x03, 0xd1, 0x00, 0xfa, 0x0b,
	0xbc, 0x9b, 0xec, 0x03, 0x60, 0x9b, 0x60, 0xb9, 0x19, 0xf3, 0x54, 0xd4, 0x15, 0xec, 0x19, 0xf7,
	0x61, 0x4b, 0xa7, 0x95, 0x21, 0x5e, 0xb2, 0xf2, 0x38, 0xfe, 0xed, 0x41, 0xb7, 0x5c, 0xd3, 0xb4,
	0x32, 0x26, 0x9e, 0x03, 0x2c, 0x0d, 0x86, 0xe1, 0x62, 0x96, 0x96, 0x3d, 0xc3, 0xe1, 0xca, 0x9c,
	0x5b, 0xe4, 0x27, 0xd8, 0xad, 0xc5, 0xc7, 0x60, 0x01, 0xfc, 0xc7, 0x53, 0xe1, 0xeb, 0x15, 0x19,
	0x57, 0xe0, 0x1c, 0x60, 0xa9, 0x61, 0x83, 0x47, 0xcb, 0x01, 0x0d, 0x1e, 0x6d, 0xd1, 0xc7, 0xbf,
	0x3c, 0x00, 0x76, 0x7d, 0x5a, 0x4f, 0xf7, 0x19, 0x5e, 0x5c, 0xdd, 0xde, 0x92, 0xc2, 0x75, 0xd2,
	0x86, 0xeb, 0x92, 0xf8, 0x11, 0x7c, 0xb7, 0x7f, 0xfc, 0x9f, 0xd2, 0x61, 0xd0, 0x4e, 0x54, 0x7c,
	0x26, 0x6f, 0xbf, 0x8d, 0xee, 0x84, 0xf9, 0xfe, 0x70, 0x13, 0x73, 0x39, 0x3f, 0x76, 0xa8, 0xfa,
	0xfb, 0xae, 0x7c, 0x1d, 0x48, 0x1d, 0xdb, 0xc7, 0xe1, 0x66, 0xc7, 0x7e, 0x4e, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x3b, 0x56, 0x68, 0x37, 0x34, 0x04, 0x00, 0x00,
}
