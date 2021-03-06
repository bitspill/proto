// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oip5.proto

package pb_oip5

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import oipProto_templates "github.com/oipwg/proto/go/pb_oip5/pb_templates"
import oipProto "github.com/oipwg/proto/go/pb_oip"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OipFive struct {
	RecordTemplate *oipProto_templates.RecordTemplateProto `protobuf:"bytes,1,opt,name=recordTemplate" json:"recordTemplate,omitempty"`
	Record         *RecordProto                            `protobuf:"bytes,2,opt,name=record" json:"record,omitempty"`
	Normalize      *NormalizeRecordProto                   `protobuf:"bytes,3,opt,name=normalize" json:"normalize,omitempty"`
	Transfer       *Transfer                               `protobuf:"bytes,7,opt,name=transfer" json:"transfer,omitempty"`
	Deactivate     *Deactivate                             `protobuf:"bytes,8,opt,name=deactivate" json:"deactivate,omitempty"`
	Edit           *EditProto                              `protobuf:"bytes,9,opt,name=edit" json:"edit,omitempty"`
}

func (m *OipFive) Reset()                    { *m = OipFive{} }
func (m *OipFive) String() string            { return proto.CompactTextString(m) }
func (*OipFive) ProtoMessage()               {}
func (*OipFive) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *OipFive) GetRecordTemplate() *oipProto_templates.RecordTemplateProto {
	if m != nil {
		return m.RecordTemplate
	}
	return nil
}

func (m *OipFive) GetRecord() *RecordProto {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *OipFive) GetNormalize() *NormalizeRecordProto {
	if m != nil {
		return m.Normalize
	}
	return nil
}

func (m *OipFive) GetTransfer() *Transfer {
	if m != nil {
		return m.Transfer
	}
	return nil
}

func (m *OipFive) GetDeactivate() *Deactivate {
	if m != nil {
		return m.Deactivate
	}
	return nil
}

func (m *OipFive) GetEdit() *EditProto {
	if m != nil {
		return m.Edit
	}
	return nil
}

type Transfer struct {
}

func (m *Transfer) Reset()                    { *m = Transfer{} }
func (m *Transfer) String() string            { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()               {}
func (*Transfer) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type Deactivate struct {
	Txid *oipProto.Txid `protobuf:"bytes,1,opt,name=txid" json:"txid,omitempty"`
}

func (m *Deactivate) Reset()                    { *m = Deactivate{} }
func (m *Deactivate) String() string            { return proto.CompactTextString(m) }
func (*Deactivate) ProtoMessage()               {}
func (*Deactivate) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Deactivate) GetTxid() *oipProto.Txid {
	if m != nil {
		return m.Txid
	}
	return nil
}

func init() {
	proto.RegisterType((*OipFive)(nil), "oipProto.oip5.OipFive")
	proto.RegisterType((*Transfer)(nil), "oipProto.oip5.Transfer")
	proto.RegisterType((*Deactivate)(nil), "oipProto.oip5.Deactivate")
}

func init() { proto.RegisterFile("oip5.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xd5, 0xfe, 0x55, 0x2f, 0xf3, 0x43, 0x17, 0x96, 0x10, 0xa6, 0x2b, 0x94, 0x2e, 0x00,
	0x81, 0x1c, 0xd4, 0x8a, 0x05, 0x62, 0x05, 0x02, 0x96, 0x14, 0x45, 0x5d, 0xb1, 0x41, 0x69, 0x63,
	0xc2, 0x48, 0x4d, 0x6d, 0xb9, 0xa6, 0x54, 0xbc, 0x34, 0xaf, 0x80, 0x7c, 0x49, 0x42, 0x2c, 0x56,
	0x71, 0x66, 0xce, 0x77, 0x3c, 0xc7, 0x03, 0x20, 0x50, 0x5e, 0x31, 0xa9, 0x84, 0x16, 0x64, 0x5f,
	0xa0, 0x7c, 0x36, 0x27, 0x66, 0x8a, 0xa3, 0xbd, 0x84, 0x2f, 0x85, 0xca, 0x5c, 0x73, 0x34, 0xd6,
	0xbc, 0x90, 0xab, 0x54, 0xf3, 0x4d, 0xec, 0xea, 0x73, 0xff, 0xef, 0x08, 0x27, 0x3a, 0x78, 0x12,
	0xaa, 0x48, 0x57, 0xf8, 0xc5, 0x1b, 0x2c, 0xf0, 0x0c, 0x75, 0x79, 0xd6, 0x3b, 0xf4, 0xf5, 0xe8,
	0xbb, 0x0d, 0xbd, 0x19, 0xca, 0x47, 0xdc, 0x72, 0x32, 0x83, 0xa1, 0x6a, 0xf8, 0xd2, 0xd6, 0x71,
	0xeb, 0xf4, 0xff, 0xe4, 0x84, 0x55, 0x53, 0x55, 0x13, 0xb0, 0x3f, 0x26, 0x48, 0x02, 0x9c, 0x4c,
	0xa0, 0xeb, 0x2a, 0xb4, 0x6d, 0x8d, 0x46, 0xac, 0x11, 0xcf, 0x7b, 0x38, 0xd6, 0x2b, 0xc9, 0x2d,
	0x0c, 0xd6, 0x65, 0x02, 0xfa, 0xcf, 0x62, 0xe3, 0x00, 0x0b, 0x12, 0x3a, 0xbe, 0xa6, 0xc8, 0x14,
	0xfa, 0x5a, 0xa5, 0xeb, 0xcd, 0x1b, 0x57, 0xb4, 0x67, 0x1d, 0x0e, 0x03, 0x87, 0xb9, 0x6f, 0x27,
	0x95, 0x90, 0x5c, 0x03, 0x64, 0x3c, 0x5d, 0x6a, 0xdc, 0x9a, 0xe0, 0x7d, 0x8b, 0x1d, 0x05, 0xd8,
	0x7d, 0x25, 0x48, 0x7e, 0x89, 0xc9, 0x05, 0x74, 0xcc, 0xeb, 0xd2, 0x81, 0x85, 0x68, 0x00, 0x3d,
	0x64, 0xa8, 0xdd, 0x88, 0x56, 0x15, 0x01, 0xf4, 0xcb, 0xeb, 0xa3, 0x4b, 0x80, 0xda, 0x93, 0x44,
	0xd0, 0x31, 0x9b, 0xf1, 0xaf, 0x3e, 0xac, 0x7d, 0xe6, 0x3b, 0xcc, 0x12, 0xdb, 0xbb, 0x3b, 0x7f,
	0x39, 0xcb, 0x51, 0xbf, 0x7f, 0x2c, 0xd8, 0x52, 0x14, 0xb1, 0x40, 0xf9, 0x99, 0xc7, 0x76, 0x93,
	0x71, 0x2e, 0x62, 0xb9, 0x78, 0x35, 0x57, 0xde, 0xf8, 0xef, 0xa2, 0x6b, 0x3b, 0xd3, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x75, 0xf3, 0xfa, 0x5a, 0x62, 0x02, 0x00, 0x00,
}
