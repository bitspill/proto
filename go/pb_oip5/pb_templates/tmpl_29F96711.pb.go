// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tmpl_29F96711.proto

package pb_templates

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Coin
// 29f9671160b8f6f0ab11cebd8911725262bc7180302f4ce50556c091f7e26696
type Tmpl_29F96711 struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Ticker   string `protobuf:"bytes,2,opt,name=ticker" json:"ticker,omitempty"`
	Source   string `protobuf:"bytes,3,opt,name=source" json:"source,omitempty"`
	Homepage string `protobuf:"bytes,4,opt,name=homepage" json:"homepage,omitempty"`
}

func (m *Tmpl_29F96711) Reset()                    { *m = Tmpl_29F96711{} }
func (m *Tmpl_29F96711) String() string            { return proto.CompactTextString(m) }
func (*Tmpl_29F96711) ProtoMessage()               {}
func (*Tmpl_29F96711) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Tmpl_29F96711) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tmpl_29F96711) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *Tmpl_29F96711) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Tmpl_29F96711) GetHomepage() string {
	if m != nil {
		return m.Homepage
	}
	return ""
}

func init() {
	proto.RegisterType((*Tmpl_29F96711)(nil), "oipProto.templates.tmpl_29F96711")
}

func init() { proto.RegisterFile("tmpl_29F96711.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc9, 0x2d, 0xc8,
	0x89, 0x37, 0xb2, 0x74, 0xb3, 0x34, 0x33, 0x37, 0x34, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xca, 0xcf, 0x2c, 0x08, 0x00, 0xb1, 0xf4, 0x4a, 0x52, 0x73, 0x0b, 0x72, 0x12, 0x4b, 0x52,
	0x8b, 0x95, 0xf2, 0xb9, 0x78, 0x51, 0x94, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x62, 0x5c, 0x6c, 0x25, 0x99, 0xc9, 0xd9,
	0xa9, 0x45, 0x12, 0x4c, 0x60, 0x51, 0x28, 0x0f, 0x24, 0x5e, 0x9c, 0x5f, 0x5a, 0x94, 0x9c, 0x2a,
	0xc1, 0x0c, 0x11, 0x87, 0xf0, 0x84, 0xa4, 0xb8, 0x38, 0x32, 0xf2, 0x73, 0x53, 0x0b, 0x12, 0xd3,
	0x53, 0x25, 0x58, 0xc0, 0x32, 0x70, 0xbe, 0x93, 0x6d, 0x94, 0x75, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7e, 0x66, 0x41, 0x79, 0xba, 0x3e, 0xd8, 0x79, 0xfa, 0xe9,
	0xf9, 0xfa, 0x05, 0x49, 0xf1, 0xf9, 0x99, 0x05, 0xa6, 0x20, 0x1a, 0xee, 0x44, 0x6b, 0x64, 0x4e,
	0x12, 0x1b, 0x58, 0xad, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xd0, 0x57, 0x79, 0xe1, 0x00,
	0x00, 0x00,
}
