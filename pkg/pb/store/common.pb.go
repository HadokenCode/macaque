// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store/common.proto

/*
Package store is a generated protocol buffer package.

It is generated from these files:
	store/common.proto
	store/scheduler.proto
	store/service.proto

It has these top-level messages:
	EntityID
	SchedulerMeta
	SchedulerEntity
*/
package store

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EntityID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *EntityID) Reset()                    { *m = EntityID{} }
func (m *EntityID) String() string            { return proto.CompactTextString(m) }
func (*EntityID) ProtoMessage()               {}
func (*EntityID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EntityID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*EntityID)(nil), "store.EntityID")
}

func init() { proto.RegisterFile("store/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 81 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2e, 0xc9, 0x2f,
	0x4a, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x05, 0x8b, 0x29, 0x49, 0x71, 0x71, 0xb8, 0xe6, 0x95, 0x64, 0x96, 0x54, 0x7a, 0xba, 0x08, 0xf1,
	0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x24, 0xb1,
	0x81, 0x55, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x56, 0xba, 0x0c, 0x3f, 0x00, 0x00,
	0x00,
}
