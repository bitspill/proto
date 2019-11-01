// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Record.proto

package pb_oip5

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RecordProto struct {
	// Searchable user-tags
	Tags []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	// Payment terms
	Payment *Payment `protobuf:"bytes,6,opt,name=payment" json:"payment,omitempty"`
	// Record specific additional information
	Details     *OipDetails  `protobuf:"bytes,7,opt,name=details" json:"details,omitempty"`
	Permissions *Permissions `protobuf:"bytes,8,opt,name=permissions" json:"permissions,omitempty"`
}

func (m *RecordProto) Reset()                    { *m = RecordProto{} }
func (m *RecordProto) String() string            { return proto.CompactTextString(m) }
func (*RecordProto) ProtoMessage()               {}
func (*RecordProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *RecordProto) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *RecordProto) GetPayment() *Payment {
	if m != nil {
		return m.Payment
	}
	return nil
}

func (m *RecordProto) GetDetails() *OipDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *RecordProto) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Permissions struct {
}

func (m *Permissions) Reset()                    { *m = Permissions{} }
func (m *Permissions) String() string            { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()               {}
func (*Permissions) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type Payment struct {
}

func (m *Payment) Reset()                    { *m = Payment{} }
func (m *Payment) String() string            { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()               {}
func (*Payment) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type OipDetails struct {
	Details []*google_protobuf.Any `protobuf:"bytes,1,rep,name=details" json:"details,omitempty"`
}

func (m *OipDetails) Reset()                    { *m = OipDetails{} }
func (m *OipDetails) String() string            { return proto.CompactTextString(m) }
func (*OipDetails) ProtoMessage()               {}
func (*OipDetails) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *OipDetails) GetDetails() []*google_protobuf.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*RecordProto)(nil), "oipProto.oip5.RecordProto")
	proto.RegisterType((*Permissions)(nil), "oipProto.oip5.Permissions")
	proto.RegisterType((*Payment)(nil), "oipProto.oip5.Payment")
	proto.RegisterType((*OipDetails)(nil), "oipProto.oip5.OipDetails")
}

func init() { proto.RegisterFile("Record.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x29, 0x1b, 0xab, 0xfb, 0xea, 0x2e, 0x41, 0xa4, 0xdb, 0xa9, 0xf4, 0x54, 0x11, 0xbe,
	0xc8, 0x86, 0x27, 0x77, 0x51, 0xbc, 0x3b, 0x72, 0xf4, 0x22, 0xed, 0x16, 0x63, 0x60, 0xed, 0x17,
	0x9a, 0x0c, 0xe9, 0x3f, 0xe8, 0xdf, 0x25, 0x4b, 0x56, 0x3b, 0x77, 0xca, 0x07, 0xef, 0xf7, 0xf2,
	0x1e, 0x0f, 0xae, 0x85, 0xdc, 0x52, 0xbb, 0x43, 0xd3, 0x92, 0x23, 0x36, 0x23, 0x6d, 0x36, 0xc7,
	0x0b, 0x49, 0x9b, 0xc7, 0xc5, 0x5c, 0x11, 0xa9, 0xbd, 0xe4, 0x5e, 0xac, 0x0e, 0x9f, 0xbc, 0x6c,
	0xba, 0x40, 0xe6, 0x3f, 0x11, 0x24, 0xc1, 0xea, 0x79, 0xc6, 0x60, 0xec, 0x4a, 0x65, 0xd3, 0x71,
	0x36, 0x2a, 0xa6, 0xc2, 0xdf, 0xec, 0x01, 0x62, 0x53, 0x76, 0xb5, 0x6c, 0x5c, 0x3a, 0xc9, 0xa2,
	0x22, 0x59, 0xde, 0xe2, 0xbf, 0xff, 0x71, 0x13, 0x54, 0xd1, 0x63, 0x6c, 0x05, 0xf1, 0x4e, 0xba,
	0x52, 0xef, 0x6d, 0x1a, 0x7b, 0xc7, 0xfc, 0xc2, 0xf1, 0xa6, 0xcd, 0x6b, 0x00, 0x44, 0x4f, 0xb2,
	0x35, 0x24, 0x46, 0xb6, 0xb5, 0xb6, 0x56, 0x53, 0x63, 0xd3, 0x2b, 0x6f, 0x5c, 0x5c, 0x46, 0x0d,
	0x84, 0x38, 0xc7, 0xf3, 0x19, 0x24, 0x67, 0x5a, 0x3e, 0x85, 0xf8, 0xd4, 0x2a, 0x5f, 0x03, 0x0c,
	0x71, 0x0c, 0x87, 0x6a, 0x51, 0x36, 0x2a, 0x92, 0xe5, 0x0d, 0x86, 0x75, 0xb0, 0x5f, 0x07, 0x9f,
	0x9b, 0xee, 0xaf, 0xd5, 0xcb, 0xfd, 0xfb, 0x9d, 0xd2, 0xee, 0xeb, 0x50, 0xe1, 0x96, 0x6a, 0x4e,
	0xda, 0x7c, 0xab, 0xb0, 0x23, 0x57, 0xc4, 0x4d, 0xf5, 0x71, 0x6c, 0xf5, 0x74, 0x7a, 0xab, 0x89,
	0x57, 0x56, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x1b, 0x17, 0x77, 0x8e, 0x01, 0x00, 0x00,
}
