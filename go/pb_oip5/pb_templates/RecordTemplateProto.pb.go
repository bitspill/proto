// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RecordTemplateProto.proto

/*
Package pb_templates is a generated protocol buffer package.

It is generated from these files:
	RecordTemplateProto.proto
	tmpl_20AD45E7.proto
	tmpl_2670B072.proto
	tmpl_29F96711.proto
	tmpl_433C2783.proto
	tmpl_6E6D471D.proto
	tmpl_D8D0F22C.proto
	tmpl_DE84D583.proto
	tmpl_payments.proto

It has these top-level messages:
	RecordTemplateProto
	Tmpl_20AD45E7
	Tmpl_2670B072
	Tmpl_29F96711
	Tmpl_433C2783
	Tmpl_6E6D471D
	Tmpl_D8D0F22C
	Tmpl_DE84D583
	P
*/
package pb_templates

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

type RecordTemplateProto struct {
	// Human readable name to quickly identify type (non-unique)
	FriendlyName string `protobuf:"bytes,1,opt,name=friendlyName" json:"friendlyName,omitempty"`
	// Description of the purpose behind this new type
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Compiled descriptor including dependencies; Defines fields
	DescriptorSetProto []byte `protobuf:"bytes,4,opt,name=DescriptorSetProto,proto3" json:"DescriptorSetProto,omitempty"`
	// List of unique template identifiers required for use with this template
	Extends []uint32 `protobuf:"fixed32,13,rep,packed,name=extends" json:"extends,omitempty"`
	// Populated by oipd with the unique identifier for this type
	Identifier uint32 `protobuf:"fixed32,10,opt,name=identifier" json:"identifier,omitempty"`
}

func (m *RecordTemplateProto) Reset()                    { *m = RecordTemplateProto{} }
func (m *RecordTemplateProto) String() string            { return proto.CompactTextString(m) }
func (*RecordTemplateProto) ProtoMessage()               {}
func (*RecordTemplateProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RecordTemplateProto) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *RecordTemplateProto) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RecordTemplateProto) GetDescriptorSetProto() []byte {
	if m != nil {
		return m.DescriptorSetProto
	}
	return nil
}

func (m *RecordTemplateProto) GetExtends() []uint32 {
	if m != nil {
		return m.Extends
	}
	return nil
}

func (m *RecordTemplateProto) GetIdentifier() uint32 {
	if m != nil {
		return m.Identifier
	}
	return 0
}

func init() {
	proto.RegisterType((*RecordTemplateProto)(nil), "oipProto.templates.RecordTemplateProto")
}

func init() { proto.RegisterFile("RecordTemplateProto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x59, 0x3d, 0x8c, 0x8e, 0x39, 0x90, 0xb5, 0x59, 0x1b, 0x59, 0xae, 0x4a, 0x95, 0x14,
	0x62, 0x75, 0xd8, 0x88, 0x95, 0x85, 0xc8, 0x6a, 0x65, 0x23, 0x97, 0xec, 0x5c, 0x1c, 0xb8, 0x64,
	0x96, 0xcd, 0x88, 0xfa, 0x1b, 0xfd, 0x53, 0x72, 0xeb, 0x29, 0x39, 0x48, 0x33, 0xc3, 0xf7, 0xde,
	0x9b, 0x81, 0x07, 0x17, 0x0e, 0x1b, 0x8e, 0xfe, 0x19, 0xbb, 0xb0, 0x59, 0x09, 0x3e, 0x46, 0x16,
	0x2e, 0xc3, 0x76, 0x6a, 0xcd, 0x14, 0x7e, 0x59, 0x76, 0xee, 0xb0, 0xf8, 0x56, 0x70, 0x3e, 0x71,
	0xa1, 0x17, 0x90, 0xaf, 0x23, 0x61, 0xef, 0x37, 0x5f, 0x0f, 0xab, 0x0e, 0x8d, 0xb2, 0xaa, 0x38,
	0x71, 0x7b, 0x9a, 0xb6, 0x70, 0xea, 0x71, 0x68, 0x22, 0x05, 0x21, 0xee, 0xcd, 0x41, 0x8a, 0x8c,
	0x25, 0x5d, 0x82, 0xbe, 0xdb, 0x21, 0xc7, 0x27, 0x94, 0xf4, 0xdb, 0xcc, 0xac, 0x2a, 0x72, 0x37,
	0xe1, 0x68, 0x03, 0x19, 0x7e, 0x0a, 0xf6, 0x7e, 0x30, 0x73, 0x7b, 0x58, 0x64, 0xee, 0x0f, 0xf5,
	0x25, 0x00, 0x79, 0xec, 0x85, 0xd6, 0x84, 0xd1, 0x80, 0x55, 0x45, 0xe6, 0x46, 0xca, 0xfd, 0xec,
	0x38, 0x3f, 0x9b, 0xdf, 0xde, 0xbc, 0x2c, 0x5b, 0x92, 0xb7, 0xf7, 0xba, 0x6c, 0xb8, 0xab, 0x98,
	0xc2, 0x47, 0x5b, 0xa5, 0xee, 0x55, 0xcb, 0x55, 0xa8, 0x5f, 0x99, 0xc2, 0xf5, 0x76, 0xff, 0xf7,
	0x5f, 0x8e, 0xa1, 0x3e, 0x4a, 0xd9, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x3e, 0x77,
	0x66, 0x44, 0x01, 0x00, 0x00,
}
