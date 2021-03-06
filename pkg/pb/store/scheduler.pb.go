// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store/scheduler.proto

package store

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SchedulerType int32

const (
	SchedulerType_PUBLIC  SchedulerType = 0
	SchedulerType_PRIVATE SchedulerType = 1
)

var SchedulerType_name = map[int32]string{
	0: "PUBLIC",
	1: "PRIVATE",
}
var SchedulerType_value = map[string]int32{
	"PUBLIC":  0,
	"PRIVATE": 1,
}

func (x SchedulerType) String() string {
	return proto.EnumName(SchedulerType_name, int32(x))
}
func (SchedulerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type SchedulerStatus int32

const (
	SchedulerStatus_CREATED   SchedulerStatus = 0
	SchedulerStatus_SUSPENDED SchedulerStatus = 1
	SchedulerStatus_DELETED   SchedulerStatus = 3
)

var SchedulerStatus_name = map[int32]string{
	0: "CREATED",
	1: "SUSPENDED",
	3: "DELETED",
}
var SchedulerStatus_value = map[string]int32{
	"CREATED":   0,
	"SUSPENDED": 1,
	"DELETED":   3,
}

func (x SchedulerStatus) String() string {
	return proto.EnumName(SchedulerStatus_name, int32(x))
}
func (SchedulerStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type SchedulerMeta struct {
	CreatedBy          string                     `protobuf:"bytes,1,opt,name=createdBy" json:"createdBy,omitempty"`
	CreatedOn          *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=createdOn" json:"createdOn,omitempty"`
	LastModificationBy string                     `protobuf:"bytes,3,opt,name=lastModificationBy" json:"lastModificationBy,omitempty"`
	LastModificationOn *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=lastModificationOn" json:"lastModificationOn,omitempty"`
}

func (m *SchedulerMeta) Reset()                    { *m = SchedulerMeta{} }
func (m *SchedulerMeta) String() string            { return proto.CompactTextString(m) }
func (*SchedulerMeta) ProtoMessage()               {}
func (*SchedulerMeta) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SchedulerMeta) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *SchedulerMeta) GetCreatedOn() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedOn
	}
	return nil
}

func (m *SchedulerMeta) GetLastModificationBy() string {
	if m != nil {
		return m.LastModificationBy
	}
	return ""
}

func (m *SchedulerMeta) GetLastModificationOn() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastModificationOn
	}
	return nil
}

type SchedulerEntity struct {
	Id          *EntityID                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string                     `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Expression  string                     `protobuf:"bytes,3,opt,name=expression" json:"expression,omitempty"`
	Labels      []string                   `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty"`
	Type        SchedulerType              `protobuf:"varint,5,opt,name=type,enum=store.SchedulerType" json:"type,omitempty"`
	Status      SchedulerStatus            `protobuf:"varint,6,opt,name=status,enum=store.SchedulerStatus" json:"status,omitempty"`
	Meta        *SchedulerMeta             `protobuf:"bytes,7,opt,name=meta" json:"meta,omitempty"`
	StartDate   *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=startDate" json:"startDate,omitempty"`
	EndDate     *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=endDate" json:"endDate,omitempty"`
}

func (m *SchedulerEntity) Reset()                    { *m = SchedulerEntity{} }
func (m *SchedulerEntity) String() string            { return proto.CompactTextString(m) }
func (*SchedulerEntity) ProtoMessage()               {}
func (*SchedulerEntity) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SchedulerEntity) GetId() *EntityID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SchedulerEntity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SchedulerEntity) GetExpression() string {
	if m != nil {
		return m.Expression
	}
	return ""
}

func (m *SchedulerEntity) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *SchedulerEntity) GetType() SchedulerType {
	if m != nil {
		return m.Type
	}
	return SchedulerType_PUBLIC
}

func (m *SchedulerEntity) GetStatus() SchedulerStatus {
	if m != nil {
		return m.Status
	}
	return SchedulerStatus_CREATED
}

func (m *SchedulerEntity) GetMeta() *SchedulerMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *SchedulerEntity) GetStartDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *SchedulerEntity) GetEndDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func init() {
	proto.RegisterType((*SchedulerMeta)(nil), "store.SchedulerMeta")
	proto.RegisterType((*SchedulerEntity)(nil), "store.SchedulerEntity")
	proto.RegisterEnum("store.SchedulerType", SchedulerType_name, SchedulerType_value)
	proto.RegisterEnum("store.SchedulerStatus", SchedulerStatus_name, SchedulerStatus_value)
}

func init() { proto.RegisterFile("store/scheduler.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6b, 0xdb, 0x40,
	0x10, 0x85, 0x23, 0xdb, 0x91, 0xd1, 0x98, 0x34, 0x66, 0x69, 0xc3, 0x62, 0x4a, 0x23, 0x72, 0x12,
	0x39, 0xc8, 0xe0, 0xf6, 0x50, 0x7a, 0x8b, 0x23, 0x1d, 0x5c, 0x92, 0xc6, 0xac, 0x9d, 0xde, 0xd7,
	0xd2, 0x24, 0x15, 0x48, 0x5a, 0xa1, 0x1d, 0x43, 0x75, 0xeb, 0x7f, 0xed, 0x1f, 0x29, 0xbb, 0xb2,
	0x1d, 0xd1, 0x04, 0x7c, 0xdc, 0x79, 0xdf, 0xcc, 0x9b, 0xa7, 0x11, 0x7c, 0xd0, 0xa4, 0x6a, 0x9c,
	0xea, 0xe4, 0x17, 0xa6, 0xdb, 0x1c, 0xeb, 0xb0, 0xaa, 0x15, 0x29, 0x76, 0x6a, 0xcb, 0x13, 0xd6,
	0xaa, 0x89, 0x2a, 0x0a, 0x55, 0xb6, 0xd2, 0xe4, 0xf2, 0x59, 0xa9, 0xe7, 0x1c, 0xa7, 0xf6, 0xb5,
	0xd9, 0x3e, 0x4d, 0x29, 0x2b, 0x50, 0x93, 0x2c, 0xaa, 0x16, 0xb8, 0xfa, 0xeb, 0xc0, 0xd9, 0x6a,
	0x3f, 0xef, 0x1e, 0x49, 0xb2, 0x8f, 0xe0, 0x25, 0x35, 0x4a, 0xc2, 0x74, 0xde, 0x70, 0xc7, 0x77,
	0x02, 0x4f, 0xbc, 0x14, 0xd8, 0xd7, 0x83, 0xfa, 0x50, 0xf2, 0x9e, 0xef, 0x04, 0xa3, 0xd9, 0x24,
	0x6c, 0x4d, 0xc2, 0xbd, 0x49, 0xb8, 0xde, 0x9b, 0x88, 0x17, 0x98, 0x85, 0xc0, 0x72, 0xa9, 0xe9,
	0x5e, 0xa5, 0xd9, 0x53, 0x96, 0x48, 0xca, 0x54, 0x39, 0x6f, 0x78, 0xdf, 0x1a, 0xbc, 0xa1, 0xb0,
	0xef, 0xaf, 0xf9, 0x87, 0x92, 0x0f, 0x8e, 0x5a, 0xbe, 0xd1, 0x75, 0xf5, 0xa7, 0x0f, 0xe7, 0x87,
	0x94, 0x71, 0x49, 0x19, 0x35, 0xec, 0x12, 0x7a, 0x59, 0x6a, 0x03, 0x8e, 0x66, 0xe7, 0xa1, 0xfd,
	0x76, 0x61, 0x2b, 0x2d, 0x22, 0xd1, 0xcb, 0x52, 0xe6, 0xc3, 0x28, 0x45, 0x9d, 0xd4, 0x59, 0x65,
	0xa6, 0xd8, 0xb0, 0x9e, 0xe8, 0x96, 0xd8, 0x27, 0x00, 0xfc, 0x5d, 0xd5, 0xa8, 0xb5, 0x01, 0xda,
	0x28, 0x9d, 0x0a, 0xbb, 0x00, 0x37, 0x97, 0x1b, 0xcc, 0x35, 0x1f, 0xf8, 0xfd, 0xc0, 0x13, 0xbb,
	0x17, 0x0b, 0x60, 0x40, 0x4d, 0x85, 0xfc, 0xd4, 0x77, 0x82, 0x77, 0xb3, 0xf7, 0x3b, 0xf3, 0xc3,
	0x82, 0xeb, 0xa6, 0x42, 0x61, 0x09, 0x16, 0x82, 0xab, 0x49, 0xd2, 0x56, 0x73, 0xd7, 0xb2, 0x17,
	0xff, 0xb3, 0x2b, 0xab, 0x8a, 0x1d, 0x65, 0x26, 0x17, 0x48, 0x92, 0x0f, 0x6d, 0xac, 0x57, 0x93,
	0xcd, 0x81, 0x85, 0x25, 0xcc, 0x21, 0x35, 0xc9, 0x9a, 0x22, 0x49, 0xc8, 0xbd, 0xe3, 0x87, 0x3c,
	0xc0, 0xec, 0x0b, 0x0c, 0xb1, 0x4c, 0x6d, 0x1f, 0x1c, 0xed, 0xdb, 0xa3, 0xd7, 0x41, 0xe7, 0x3f,
	0x33, 0x01, 0x19, 0x80, 0xbb, 0x7c, 0x9c, 0xdf, 0x2d, 0x6e, 0xc7, 0x27, 0x6c, 0x04, 0xc3, 0xa5,
	0x58, 0xfc, 0xbc, 0x59, 0xc7, 0x63, 0xe7, 0xfa, 0x5b, 0xe7, 0x56, 0x6d, 0x3c, 0xa3, 0xdf, 0x8a,
	0xf8, 0x66, 0x1d, 0x47, 0xe3, 0x13, 0x76, 0x06, 0xde, 0xea, 0x71, 0xb5, 0x8c, 0x7f, 0x44, 0x71,
	0x34, 0x76, 0x8c, 0x16, 0xc5, 0x77, 0xb1, 0xd1, 0xfa, 0x1b, 0xd7, 0xae, 0xf0, 0xf9, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x48, 0xaa, 0x14, 0xc4, 0x2a, 0x03, 0x00, 0x00,
}
