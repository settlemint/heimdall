// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkpoint/v1beta/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	github_com_maticnetwork_heimdall_types "github.com/maticnetwork/heimdall/types"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	CheckpointBufferTime time.Duration `protobuf:"bytes,1,opt,name=checkpoint_buffer_time,json=checkpointBufferTime,proto3,stdduration" json:"checkpoint_buffer_time"`
	AvgCheckpointLength  uint64        `protobuf:"varint,2,opt,name=avg_checkpoint_length,json=avgCheckpointLength,proto3" json:"avg_checkpoint_length,omitempty"`
	MaxCheckpointLength  uint64        `protobuf:"varint,3,opt,name=max_checkpoint_length,json=maxCheckpointLength,proto3" json:"max_checkpoint_length,omitempty"`
	ChildBlockInterval   uint64        `protobuf:"varint,4,opt,name=child_block_interval,json=childBlockInterval,proto3" json:"child_block_interval,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a783d62b74a2d0, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params.Unmarshal(m, b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params.Marshal(b, m, deterministic)
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return xxx_messageInfo_Params.Size(m)
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// GenesisState defines the checkpoint module's genesis state.
type GenesisState struct {
	Params             Params                                               `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BufferedCheckpoint *github_com_maticnetwork_heimdall_types.Checkpoint   `protobuf:"bytes,2,opt,name=buffered_checkpoint,json=bufferedCheckpoint,proto3,casttype=github.com/maticnetwork/heimdall/types.Checkpoint" json:"buffered_checkpoint,omitempty"`
	LastNoACK          uint64                                               `protobuf:"varint,3,opt,name=last_no_ack,json=lastNoAck,proto3" json:"last_no_ack,omitempty"`
	AckCount           uint64                                               `protobuf:"varint,4,opt,name=ack_count,json=ackCount,proto3" json:"ack_count,omitempty"`
	Checkpoints        []*github_com_maticnetwork_heimdall_types.Checkpoint `protobuf:"bytes,5,rep,name=checkpoints,proto3,casttype=github.com/maticnetwork/heimdall/types.Checkpoint" json:"checkpoints,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a783d62b74a2d0, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisState.Unmarshal(m, b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return xxx_messageInfo_GenesisState.Size(m)
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "heimdall.checkpoint.x.checkpoint.types.Params")
	proto.RegisterType((*GenesisState)(nil), "heimdall.checkpoint.x.checkpoint.types.GenesisState")
}

func init() { proto.RegisterFile("checkpoint/v1beta/genesis.proto", fileDescriptor_99a783d62b74a2d0) }

var fileDescriptor_99a783d62b74a2d0 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xb1, 0x6f, 0xd3, 0x40,
	0x14, 0xc6, 0x9d, 0x26, 0x44, 0xc9, 0x05, 0x96, 0x6b, 0x40, 0x69, 0x91, 0xec, 0xa8, 0x03, 0xea,
	0xc2, 0x99, 0xa6, 0x62, 0xe9, 0x56, 0x07, 0x09, 0x21, 0x02, 0x42, 0x86, 0x05, 0x16, 0xeb, 0x7c,
	0xb9, 0xd8, 0x27, 0xdb, 0x77, 0x91, 0x7d, 0x0e, 0x89, 0xc4, 0x0e, 0x23, 0x63, 0x25, 0x96, 0xfe,
	0x39, 0x1d, 0x3b, 0x32, 0x05, 0x94, 0xfc, 0x17, 0x4c, 0xc8, 0x77, 0x0e, 0xb6, 0x28, 0x12, 0x0c,
	0x6c, 0xf7, 0xfc, 0xf9, 0xfb, 0xde, 0xbd, 0xdf, 0xd3, 0x01, 0x8b, 0x84, 0x94, 0x44, 0x73, 0xc1,
	0xb8, 0xb4, 0x17, 0x27, 0x3e, 0x95, 0xd8, 0x0e, 0x28, 0xa7, 0x19, 0xcb, 0xd0, 0x3c, 0x15, 0x52,
	0xc0, 0x07, 0x21, 0x65, 0xc9, 0x14, 0xc7, 0x31, 0xaa, 0xfe, 0x44, 0xcb, 0x7a, 0x21, 0x57, 0x73,
	0x9a, 0x1d, 0xf6, 0x03, 0x11, 0x08, 0x65, 0xb1, 0x8b, 0x93, 0x76, 0x1f, 0x1e, 0x04, 0x42, 0x04,
	0x31, 0xb5, 0x55, 0xe5, 0xe7, 0x33, 0x1b, 0xf3, 0x55, 0x29, 0x99, 0xbf, 0x4b, 0xd3, 0x3c, 0xc5,
	0x92, 0x09, 0xae, 0xf5, 0xa3, 0x8f, 0x7b, 0xa0, 0xfd, 0x0a, 0xa7, 0x38, 0xc9, 0xe0, 0x5b, 0x70,
	0xaf, 0xea, 0xe7, 0xf9, 0xf9, 0x6c, 0x46, 0x53, 0x4f, 0xb2, 0x84, 0x0e, 0x1a, 0xc3, 0xc6, 0x71,
	0x6f, 0x74, 0x80, 0x74, 0x16, 0xda, 0x65, 0xa1, 0x27, 0x65, 0x96, 0xd3, 0xb9, 0x5a, 0x5b, 0xc6,
	0xc5, 0x37, 0xab, 0xe1, 0xf6, 0xab, 0x08, 0x47, 0x25, 0xbc, 0x61, 0x09, 0x85, 0x23, 0x70, 0x17,
	0x2f, 0x02, 0xaf, 0x16, 0x1f, 0x53, 0x1e, 0xc8, 0x70, 0xb0, 0x37, 0x6c, 0x1c, 0xb7, 0xdc, 0x7d,
	0xbc, 0x08, 0xc6, 0xbf, 0xb4, 0x89, 0x92, 0x0a, 0x4f, 0x82, 0x97, 0x7f, 0xf0, 0x34, 0xb5, 0x27,
	0xc1, 0xcb, 0x1b, 0x9e, 0x47, 0xa0, 0x4f, 0x42, 0x16, 0x4f, 0x3d, 0x3f, 0x16, 0x24, 0xf2, 0x18,
	0x97, 0x34, 0x5d, 0xe0, 0x78, 0xd0, 0x52, 0x16, 0xa8, 0x34, 0xa7, 0x90, 0x9e, 0x95, 0xca, 0x59,
	0xe7, 0xd3, 0xa5, 0x65, 0x5c, 0x5c, 0x5a, 0xc6, 0xd1, 0x97, 0x26, 0xb8, 0xfd, 0x54, 0x2f, 0xe5,
	0xb5, 0xc4, 0x92, 0xc2, 0x09, 0x68, 0xcf, 0x15, 0x99, 0x72, 0x7e, 0x84, 0xfe, 0x6d, 0x49, 0x48,
	0xf3, 0x74, 0x5a, 0x05, 0x14, 0xb7, 0xcc, 0x80, 0x1f, 0xc0, 0xbe, 0x46, 0x4a, 0xa7, 0xb5, 0x99,
	0x14, 0x80, 0xde, 0xa8, 0x7f, 0x03, 0xed, 0x39, 0x5f, 0x39, 0x8f, 0x7f, 0xac, 0xad, 0x93, 0x80,
	0xc9, 0x30, 0xf7, 0x11, 0x11, 0x89, 0x9d, 0x60, 0xc9, 0x08, 0xa7, 0xf2, 0xbd, 0x48, 0x23, 0x7b,
	0x77, 0x17, 0x5b, 0xf7, 0xac, 0x50, 0xb8, 0x70, 0xd7, 0xa7, 0xfa, 0x06, 0x1f, 0x82, 0x5e, 0x8c,
	0x33, 0xe9, 0x71, 0xe1, 0x61, 0x12, 0x69, 0x84, 0xce, 0x9d, 0xcd, 0xda, 0xea, 0x4e, 0x70, 0x26,
	0x5f, 0x8a, 0xf3, 0xf1, 0x73, 0xb7, 0x1b, 0xeb, 0x23, 0x89, 0xe0, 0x7d, 0xd0, 0xc5, 0x24, 0xf2,
	0x88, 0xc8, 0xb9, 0x2c, 0xe1, 0x75, 0x30, 0x89, 0xc6, 0x45, 0x0d, 0x13, 0xd0, 0xab, 0x06, 0xc8,
	0x06, 0xb7, 0x86, 0xcd, 0xff, 0x3d, 0x41, 0x3d, 0xff, 0xac, 0x55, 0x6c, 0xc8, 0x79, 0x71, 0xb5,
	0x31, 0x8d, 0xeb, 0x8d, 0x69, 0x7c, 0xdf, 0x98, 0xc6, 0xe7, 0xad, 0x69, 0x5c, 0x6f, 0x4d, 0xe3,
	0xeb, 0xd6, 0x34, 0xde, 0x9d, 0xfe, 0x35, 0x7f, 0x69, 0xd7, 0x9e, 0x9f, 0x6a, 0xe6, 0xb7, 0xd5,
	0x35, 0x4f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x04, 0x7a, 0xdd, 0x99, 0x03, 0x00, 0x00,
}
