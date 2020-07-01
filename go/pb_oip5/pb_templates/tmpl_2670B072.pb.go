// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tmpl_2670B072.proto

package pb_templates

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/oipwg/proto/go/pb_oip"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Song
// 2670b072c07eda38a423bd1dbadc629df200d644ff98fc2768aaa00203c1efda
type Tmpl_2670B072 struct {
	Title    string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Genre    string `protobuf:"bytes,2,opt,name=genre" json:"genre,omitempty"`
	Duration uint32 `protobuf:"varint,3,opt,name=duration" json:"duration,omitempty"`
	Artist   string `protobuf:"bytes,4,opt,name=artist" json:"artist,omitempty"`
}

func (m *Tmpl_2670B072) Reset()                    { *m = Tmpl_2670B072{} }
func (m *Tmpl_2670B072) String() string            { return proto.CompactTextString(m) }
func (*Tmpl_2670B072) ProtoMessage()               {}
func (*Tmpl_2670B072) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Tmpl_2670B072) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Tmpl_2670B072) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func (m *Tmpl_2670B072) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Tmpl_2670B072) GetArtist() string {
	if m != nil {
		return m.Artist
	}
	return ""
}

func init() {
	proto.RegisterType((*Tmpl_2670B072)(nil), "oipProto.templates.tmpl_2670B072")
}

func init() { proto.RegisterFile("tmpl_2670B072.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc9, 0x2d, 0xc8,
	0x89, 0x37, 0x32, 0x33, 0x37, 0x70, 0x32, 0x30, 0x37, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xca, 0xcf, 0x2c, 0x08, 0x00, 0xb1, 0xf4, 0x4a, 0x52, 0x73, 0x0b, 0x72, 0x12, 0x4b, 0x52,
	0x8b, 0xa5, 0xb8, 0x4a, 0x2a, 0x32, 0x53, 0x20, 0xf2, 0x4a, 0xf9, 0x5c, 0xbc, 0x28, 0xda, 0x84,
	0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20,
	0x1c, 0x90, 0x68, 0x7a, 0x6a, 0x5e, 0x51, 0xaa, 0x04, 0x13, 0x44, 0x14, 0xcc, 0x11, 0x92, 0xe2,
	0xe2, 0x48, 0x29, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x0d,
	0x82, 0xf3, 0x85, 0xc4, 0xb8, 0xd8, 0x12, 0x8b, 0x4a, 0x32, 0x8b, 0x4b, 0x24, 0x58, 0xc0, 0x5a,
	0xa0, 0x3c, 0x27, 0xdb, 0x28, 0xeb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0xfd, 0xfc, 0xcc, 0x82, 0xf2, 0x74, 0x7d, 0xb0, 0x53, 0xf4, 0xd3, 0xf3, 0xf5, 0x0b, 0x92, 0xe2,
	0xf3, 0x33, 0x0b, 0x4c, 0x41, 0x34, 0xdc, 0xb9, 0xd6, 0xc8, 0x9c, 0x24, 0x36, 0xb0, 0x5a, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xb1, 0xc2, 0x74, 0xed, 0x00, 0x00, 0x00,
}
