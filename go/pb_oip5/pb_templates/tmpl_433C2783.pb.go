// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tmpl_433C2783.proto

package pb_templates

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Registered Publisher
type Tmpl_433C2783 struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	FloBip44XPub string `protobuf:"bytes,2,opt,name=floBip44XPub" json:"floBip44XPub,omitempty"`
}

func (m *Tmpl_433C2783) Reset()                    { *m = Tmpl_433C2783{} }
func (m *Tmpl_433C2783) String() string            { return proto.CompactTextString(m) }
func (*Tmpl_433C2783) ProtoMessage()               {}
func (*Tmpl_433C2783) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Tmpl_433C2783) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tmpl_433C2783) GetFloBip44XPub() string {
	if m != nil {
		return m.FloBip44XPub
	}
	return ""
}

func init() {
	proto.RegisterType((*Tmpl_433C2783)(nil), "oipProto.templates.tmpl_433C2783")
}

func init() { proto.RegisterFile("tmpl_433C2783.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc9, 0x2d, 0xc8,
	0x89, 0x37, 0x31, 0x36, 0x76, 0x36, 0x32, 0xb7, 0x30, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xca, 0xcf, 0x2c, 0x08, 0x00, 0xb1, 0xf4, 0x4a, 0x52, 0x73, 0x0b, 0x72, 0x12, 0x4b, 0x52,
	0x8b, 0x95, 0xdc, 0xb9, 0x78, 0x51, 0x94, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x4a, 0x5c, 0x3c, 0x69, 0x39, 0xf9, 0x4e,
	0x99, 0x05, 0x26, 0x26, 0x11, 0x01, 0xa5, 0x49, 0x12, 0x4c, 0x60, 0x39, 0x14, 0x31, 0x27, 0xdb,
	0x28, 0xeb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xfc, 0xcc, 0x82,
	0xf2, 0x74, 0x7d, 0xb0, 0xb5, 0xfa, 0xe9, 0xf9, 0xfa, 0x05, 0x49, 0xf1, 0xf9, 0x99, 0x05, 0xa6,
	0x20, 0x1a, 0x6e, 0xb5, 0x35, 0x32, 0x27, 0x89, 0x0d, 0xac, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x56, 0x5a, 0xe4, 0xf1, 0xb9, 0x00, 0x00, 0x00,
}
