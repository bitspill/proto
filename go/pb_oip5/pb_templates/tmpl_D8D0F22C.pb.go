// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tmpl_D8D0F22C.proto

package pb_templates

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import oipProto "github.com/oipwg/proto/go/pb_oip"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Tmpl_D8D0F22C_Network int32

const (
	Tmpl_D8D0F22C_NETWORK_UNDEFINED  Tmpl_D8D0F22C_Network = 0
	Tmpl_D8D0F22C_NETWORK_IPFS       Tmpl_D8D0F22C_Network = 1
	Tmpl_D8D0F22C_NETWORK_BITTORRENT Tmpl_D8D0F22C_Network = 2
	Tmpl_D8D0F22C_NETWORK_SIA        Tmpl_D8D0F22C_Network = 3
	Tmpl_D8D0F22C_NETWORK_STORJ      Tmpl_D8D0F22C_Network = 4
)

var Tmpl_D8D0F22C_Network_name = map[int32]string{
	0: "NETWORK_UNDEFINED",
	1: "NETWORK_IPFS",
	2: "NETWORK_BITTORRENT",
	3: "NETWORK_SIA",
	4: "NETWORK_STORJ",
}
var Tmpl_D8D0F22C_Network_value = map[string]int32{
	"NETWORK_UNDEFINED":  0,
	"NETWORK_IPFS":       1,
	"NETWORK_BITTORRENT": 2,
	"NETWORK_SIA":        3,
	"NETWORK_STORJ":      4,
}

func (x Tmpl_D8D0F22C_Network) String() string {
	return proto.EnumName(Tmpl_D8D0F22C_Network_name, int32(x))
}
func (Tmpl_D8D0F22C_Network) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 0} }

// CommercialContent
// d8d0f22c79f3d5670874a4f0e439b7cdc130dd213759e56eb8d2f3081e0e7dda
type Tmpl_D8D0F22C struct {
	Location      string                `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Network       Tmpl_D8D0F22C_Network `protobuf:"varint,2,opt,name=network,enum=oipProto.templates.Tmpl_D8D0F22C_Network" json:"network,omitempty"`
	Terms         []*oipProto.Txid      `protobuf:"bytes,3,rep,name=terms" json:"terms,omitempty"`
	EmbeddedTerms []uint32              `protobuf:"fixed32,4,rep,packed,name=embedded_terms,json=embeddedTerms" json:"embedded_terms,omitempty"`
}

func (m *Tmpl_D8D0F22C) Reset()                    { *m = Tmpl_D8D0F22C{} }
func (m *Tmpl_D8D0F22C) String() string            { return proto.CompactTextString(m) }
func (*Tmpl_D8D0F22C) ProtoMessage()               {}
func (*Tmpl_D8D0F22C) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Tmpl_D8D0F22C) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Tmpl_D8D0F22C) GetNetwork() Tmpl_D8D0F22C_Network {
	if m != nil {
		return m.Network
	}
	return Tmpl_D8D0F22C_NETWORK_UNDEFINED
}

func (m *Tmpl_D8D0F22C) GetTerms() []*oipProto.Txid {
	if m != nil {
		return m.Terms
	}
	return nil
}

func (m *Tmpl_D8D0F22C) GetEmbeddedTerms() []uint32 {
	if m != nil {
		return m.EmbeddedTerms
	}
	return nil
}

func init() {
	proto.RegisterType((*Tmpl_D8D0F22C)(nil), "oipProto.templates.tmpl_D8D0F22C")
	proto.RegisterEnum("oipProto.templates.Tmpl_D8D0F22C_Network", Tmpl_D8D0F22C_Network_name, Tmpl_D8D0F22C_Network_value)
}

func init() { proto.RegisterFile("tmpl_D8D0F22C.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x51, 0x4b, 0x3a, 0x41,
	0x14, 0xc5, 0xff, 0xeb, 0xfa, 0xcf, 0xba, 0xa6, 0x8d, 0x37, 0x0a, 0xf1, 0x69, 0x91, 0x82, 0xed,
	0x65, 0x37, 0x36, 0x82, 0x40, 0x7a, 0x48, 0x77, 0x85, 0x2d, 0x58, 0x65, 0x9c, 0x08, 0x7a, 0x59,
	0x5c, 0x77, 0xb0, 0x21, 0xd7, 0x19, 0x74, 0x42, 0xbf, 0x54, 0xdf, 0x31, 0x74, 0x5b, 0x4b, 0x7a,
	0x9a, 0xb9, 0xbf, 0x7b, 0xce, 0x99, 0xe1, 0xc0, 0xa9, 0xce, 0xd4, 0x2c, 0xf6, 0xef, 0xfc, 0xeb,
	0xbe, 0xe7, 0xf5, 0x1c, 0xb5, 0x90, 0x5a, 0x22, 0x4a, 0xa1, 0x86, 0x9b, 0x9b, 0xa3, 0x79, 0xa6,
	0x66, 0x63, 0xcd, 0x97, 0x2d, 0xd0, 0x6b, 0x91, 0xe6, 0xfb, 0xf6, 0x67, 0x09, 0x6a, 0x7b, 0x3e,
	0x6c, 0xc1, 0xe1, 0x4c, 0x4e, 0xc6, 0x5a, 0xc8, 0x79, 0xd3, 0xb0, 0x0c, 0xfb, 0x88, 0xee, 0x66,
	0xec, 0x41, 0x65, 0xce, 0xf5, 0x4a, 0x2e, 0xde, 0x9b, 0x25, 0xcb, 0xb0, 0xeb, 0xde, 0x95, 0xf3,
	0x37, 0xdf, 0xd9, 0xff, 0x47, 0x94, 0x1b, 0x68, 0xe1, 0xc4, 0x0b, 0xf8, 0xaf, 0xf9, 0x22, 0x5b,
	0x36, 0x4d, 0xcb, 0xb4, 0xab, 0x5e, 0xfd, 0x27, 0x82, 0xad, 0x45, 0x4a, 0xf3, 0x25, 0x5e, 0x42,
	0x9d, 0x67, 0x09, 0x4f, 0x53, 0x9e, 0xc6, 0xb9, 0xbc, 0x6c, 0x99, 0x76, 0x85, 0xd6, 0x0a, 0xca,
	0x36, 0xb0, 0x3d, 0x87, 0xca, 0xf7, 0x03, 0x78, 0x06, 0x8d, 0x28, 0x60, 0x2f, 0x03, 0xfa, 0x14,
	0x3f, 0x47, 0x7e, 0xd0, 0x0f, 0xa3, 0xc0, 0x27, 0xff, 0x90, 0xc0, 0x71, 0x81, 0xc3, 0x61, 0x7f,
	0x44, 0x0c, 0x3c, 0x07, 0x2c, 0x48, 0x37, 0x64, 0x6c, 0x40, 0x69, 0x10, 0x31, 0x52, 0xc2, 0x13,
	0xa8, 0x16, 0x7c, 0x14, 0x3e, 0x10, 0x13, 0x1b, 0x50, 0xdb, 0x01, 0x36, 0xa0, 0x8f, 0xa4, 0xdc,
	0xbd, 0x7f, 0xed, 0x4c, 0x85, 0x7e, 0xfb, 0x48, 0x9c, 0x89, 0xcc, 0x5c, 0x29, 0xd4, 0x6a, 0xea,
	0x6e, 0x9b, 0x74, 0xa7, 0xd2, 0x55, 0x49, 0x2c, 0x85, 0xba, 0xdd, 0x9c, 0xbb, 0x36, 0x3a, 0xbf,
	0x87, 0xe4, 0x60, 0xab, 0xbd, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x6a, 0x9a, 0x56, 0xac,
	0x01, 0x00, 0x00,
}
