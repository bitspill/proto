// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tmpl_20AD45E7.proto

package pb_templates

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import oipProto "github.com/oipwg/proto/go/pb_oip"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Tmpl_20AD45E7_Language int32

const (
	Tmpl_20AD45E7_Language_UNDEFINED Tmpl_20AD45E7_Language = 0
	Tmpl_20AD45E7_Language_AF        Tmpl_20AD45E7_Language = 1
	Tmpl_20AD45E7_Language_AM        Tmpl_20AD45E7_Language = 2
	Tmpl_20AD45E7_Language_AR        Tmpl_20AD45E7_Language = 3
	Tmpl_20AD45E7_Language_ARN       Tmpl_20AD45E7_Language = 4
	Tmpl_20AD45E7_Language_AS        Tmpl_20AD45E7_Language = 5
	Tmpl_20AD45E7_Language_AZ        Tmpl_20AD45E7_Language = 6
	Tmpl_20AD45E7_Language_BA        Tmpl_20AD45E7_Language = 7
	Tmpl_20AD45E7_Language_BE        Tmpl_20AD45E7_Language = 8
	Tmpl_20AD45E7_Language_BG        Tmpl_20AD45E7_Language = 9
	Tmpl_20AD45E7_Language_BN        Tmpl_20AD45E7_Language = 10
	Tmpl_20AD45E7_Language_BO        Tmpl_20AD45E7_Language = 11
	Tmpl_20AD45E7_Language_BR        Tmpl_20AD45E7_Language = 12
	Tmpl_20AD45E7_Language_BS        Tmpl_20AD45E7_Language = 13
	Tmpl_20AD45E7_Language_CA        Tmpl_20AD45E7_Language = 14
	Tmpl_20AD45E7_Language_CO        Tmpl_20AD45E7_Language = 15
	Tmpl_20AD45E7_Language_CS        Tmpl_20AD45E7_Language = 16
	Tmpl_20AD45E7_Language_CY        Tmpl_20AD45E7_Language = 17
	Tmpl_20AD45E7_Language_DA        Tmpl_20AD45E7_Language = 18
	Tmpl_20AD45E7_Language_DE        Tmpl_20AD45E7_Language = 19
	Tmpl_20AD45E7_Language_DSB       Tmpl_20AD45E7_Language = 20
	Tmpl_20AD45E7_Language_DV        Tmpl_20AD45E7_Language = 21
	Tmpl_20AD45E7_Language_EL        Tmpl_20AD45E7_Language = 22
	Tmpl_20AD45E7_Language_EN        Tmpl_20AD45E7_Language = 23
	Tmpl_20AD45E7_Language_ES        Tmpl_20AD45E7_Language = 24
	Tmpl_20AD45E7_Language_ET        Tmpl_20AD45E7_Language = 25
	Tmpl_20AD45E7_Language_EU        Tmpl_20AD45E7_Language = 26
	Tmpl_20AD45E7_Language_FA        Tmpl_20AD45E7_Language = 27
	Tmpl_20AD45E7_Language_FI        Tmpl_20AD45E7_Language = 28
	Tmpl_20AD45E7_Language_FIL       Tmpl_20AD45E7_Language = 29
	Tmpl_20AD45E7_Language_FO        Tmpl_20AD45E7_Language = 30
	Tmpl_20AD45E7_Language_FR        Tmpl_20AD45E7_Language = 31
	Tmpl_20AD45E7_Language_FY        Tmpl_20AD45E7_Language = 32
	Tmpl_20AD45E7_Language_GA        Tmpl_20AD45E7_Language = 33
	Tmpl_20AD45E7_Language_GD        Tmpl_20AD45E7_Language = 34
	Tmpl_20AD45E7_Language_GL        Tmpl_20AD45E7_Language = 35
	Tmpl_20AD45E7_Language_GSW       Tmpl_20AD45E7_Language = 36
	Tmpl_20AD45E7_Language_GU        Tmpl_20AD45E7_Language = 37
	Tmpl_20AD45E7_Language_HA        Tmpl_20AD45E7_Language = 38
	Tmpl_20AD45E7_Language_HE        Tmpl_20AD45E7_Language = 39
	Tmpl_20AD45E7_Language_HI        Tmpl_20AD45E7_Language = 40
	Tmpl_20AD45E7_Language_HR        Tmpl_20AD45E7_Language = 41
	Tmpl_20AD45E7_Language_HSB       Tmpl_20AD45E7_Language = 42
	Tmpl_20AD45E7_Language_HU        Tmpl_20AD45E7_Language = 43
	Tmpl_20AD45E7_Language_HY        Tmpl_20AD45E7_Language = 44
	Tmpl_20AD45E7_Language_ID        Tmpl_20AD45E7_Language = 45
	Tmpl_20AD45E7_Language_IG        Tmpl_20AD45E7_Language = 46
	Tmpl_20AD45E7_Language_II        Tmpl_20AD45E7_Language = 47
	Tmpl_20AD45E7_Language_IS        Tmpl_20AD45E7_Language = 48
	Tmpl_20AD45E7_Language_IT        Tmpl_20AD45E7_Language = 49
	Tmpl_20AD45E7_Language_IU        Tmpl_20AD45E7_Language = 50
	Tmpl_20AD45E7_Language_JA        Tmpl_20AD45E7_Language = 51
	Tmpl_20AD45E7_Language_KA        Tmpl_20AD45E7_Language = 52
	Tmpl_20AD45E7_Language_KK        Tmpl_20AD45E7_Language = 53
	Tmpl_20AD45E7_Language_KL        Tmpl_20AD45E7_Language = 54
	Tmpl_20AD45E7_Language_KM        Tmpl_20AD45E7_Language = 55
	Tmpl_20AD45E7_Language_KN        Tmpl_20AD45E7_Language = 56
	Tmpl_20AD45E7_Language_KOK       Tmpl_20AD45E7_Language = 57
	Tmpl_20AD45E7_Language_KO        Tmpl_20AD45E7_Language = 58
	Tmpl_20AD45E7_Language_KY        Tmpl_20AD45E7_Language = 59
	Tmpl_20AD45E7_Language_LB        Tmpl_20AD45E7_Language = 60
	Tmpl_20AD45E7_Language_LO        Tmpl_20AD45E7_Language = 61
	Tmpl_20AD45E7_Language_LT        Tmpl_20AD45E7_Language = 62
	Tmpl_20AD45E7_Language_LV        Tmpl_20AD45E7_Language = 63
	Tmpl_20AD45E7_Language_MI        Tmpl_20AD45E7_Language = 64
	Tmpl_20AD45E7_Language_MK        Tmpl_20AD45E7_Language = 65
	Tmpl_20AD45E7_Language_ML        Tmpl_20AD45E7_Language = 66
	Tmpl_20AD45E7_Language_MN        Tmpl_20AD45E7_Language = 67
	Tmpl_20AD45E7_Language_MOH       Tmpl_20AD45E7_Language = 68
	Tmpl_20AD45E7_Language_MR        Tmpl_20AD45E7_Language = 69
	Tmpl_20AD45E7_Language_MS        Tmpl_20AD45E7_Language = 70
	Tmpl_20AD45E7_Language_MT        Tmpl_20AD45E7_Language = 71
	Tmpl_20AD45E7_Language_NB        Tmpl_20AD45E7_Language = 72
	Tmpl_20AD45E7_Language_NE        Tmpl_20AD45E7_Language = 73
	Tmpl_20AD45E7_Language_NL        Tmpl_20AD45E7_Language = 74
	Tmpl_20AD45E7_Language_NN        Tmpl_20AD45E7_Language = 75
	Tmpl_20AD45E7_Language_NSO       Tmpl_20AD45E7_Language = 76
	Tmpl_20AD45E7_Language_OC        Tmpl_20AD45E7_Language = 77
	Tmpl_20AD45E7_Language_OR        Tmpl_20AD45E7_Language = 78
	Tmpl_20AD45E7_Language_PA        Tmpl_20AD45E7_Language = 79
	Tmpl_20AD45E7_Language_PL        Tmpl_20AD45E7_Language = 80
	Tmpl_20AD45E7_Language_PRS       Tmpl_20AD45E7_Language = 81
	Tmpl_20AD45E7_Language_PT        Tmpl_20AD45E7_Language = 82
	Tmpl_20AD45E7_Language_QUT       Tmpl_20AD45E7_Language = 83
	Tmpl_20AD45E7_Language_QUZ       Tmpl_20AD45E7_Language = 84
	Tmpl_20AD45E7_Language_RM        Tmpl_20AD45E7_Language = 85
	Tmpl_20AD45E7_Language_RO        Tmpl_20AD45E7_Language = 86
	Tmpl_20AD45E7_Language_RU        Tmpl_20AD45E7_Language = 87
	Tmpl_20AD45E7_Language_RW        Tmpl_20AD45E7_Language = 88
	Tmpl_20AD45E7_Language_SAH       Tmpl_20AD45E7_Language = 89
	Tmpl_20AD45E7_Language_SA        Tmpl_20AD45E7_Language = 90
	Tmpl_20AD45E7_Language_SE        Tmpl_20AD45E7_Language = 91
	Tmpl_20AD45E7_Language_SI        Tmpl_20AD45E7_Language = 92
	Tmpl_20AD45E7_Language_SK        Tmpl_20AD45E7_Language = 93
	Tmpl_20AD45E7_Language_SL        Tmpl_20AD45E7_Language = 94
	Tmpl_20AD45E7_Language_SMA       Tmpl_20AD45E7_Language = 95
	Tmpl_20AD45E7_Language_SMJ       Tmpl_20AD45E7_Language = 96
	Tmpl_20AD45E7_Language_SMN       Tmpl_20AD45E7_Language = 97
	Tmpl_20AD45E7_Language_SQ        Tmpl_20AD45E7_Language = 98
	Tmpl_20AD45E7_Language_SR        Tmpl_20AD45E7_Language = 99
	Tmpl_20AD45E7_Language_SV        Tmpl_20AD45E7_Language = 100
	Tmpl_20AD45E7_Language_SW        Tmpl_20AD45E7_Language = 101
	Tmpl_20AD45E7_Language_SYR       Tmpl_20AD45E7_Language = 102
	Tmpl_20AD45E7_Language_TA        Tmpl_20AD45E7_Language = 103
	Tmpl_20AD45E7_Language_TE        Tmpl_20AD45E7_Language = 104
	Tmpl_20AD45E7_Language_TG        Tmpl_20AD45E7_Language = 105
	Tmpl_20AD45E7_Language_TH        Tmpl_20AD45E7_Language = 106
	Tmpl_20AD45E7_Language_TK        Tmpl_20AD45E7_Language = 107
	Tmpl_20AD45E7_Language_TN        Tmpl_20AD45E7_Language = 108
	Tmpl_20AD45E7_Language_TR        Tmpl_20AD45E7_Language = 109
	Tmpl_20AD45E7_Language_TT        Tmpl_20AD45E7_Language = 110
	Tmpl_20AD45E7_Language_TZM       Tmpl_20AD45E7_Language = 111
	Tmpl_20AD45E7_Language_UG        Tmpl_20AD45E7_Language = 112
	Tmpl_20AD45E7_Language_UK        Tmpl_20AD45E7_Language = 113
	Tmpl_20AD45E7_Language_UR        Tmpl_20AD45E7_Language = 114
	Tmpl_20AD45E7_Language_UZ        Tmpl_20AD45E7_Language = 115
	Tmpl_20AD45E7_Language_VI        Tmpl_20AD45E7_Language = 116
	Tmpl_20AD45E7_Language_WO        Tmpl_20AD45E7_Language = 117
	Tmpl_20AD45E7_Language_XH        Tmpl_20AD45E7_Language = 118
	Tmpl_20AD45E7_Language_YO        Tmpl_20AD45E7_Language = 119
	Tmpl_20AD45E7_Language_ZH        Tmpl_20AD45E7_Language = 120
	Tmpl_20AD45E7_Language_ZU        Tmpl_20AD45E7_Language = 121
)

var Tmpl_20AD45E7_Language_name = map[int32]string{
	0:   "Language_UNDEFINED",
	1:   "Language_AF",
	2:   "Language_AM",
	3:   "Language_AR",
	4:   "Language_ARN",
	5:   "Language_AS",
	6:   "Language_AZ",
	7:   "Language_BA",
	8:   "Language_BE",
	9:   "Language_BG",
	10:  "Language_BN",
	11:  "Language_BO",
	12:  "Language_BR",
	13:  "Language_BS",
	14:  "Language_CA",
	15:  "Language_CO",
	16:  "Language_CS",
	17:  "Language_CY",
	18:  "Language_DA",
	19:  "Language_DE",
	20:  "Language_DSB",
	21:  "Language_DV",
	22:  "Language_EL",
	23:  "Language_EN",
	24:  "Language_ES",
	25:  "Language_ET",
	26:  "Language_EU",
	27:  "Language_FA",
	28:  "Language_FI",
	29:  "Language_FIL",
	30:  "Language_FO",
	31:  "Language_FR",
	32:  "Language_FY",
	33:  "Language_GA",
	34:  "Language_GD",
	35:  "Language_GL",
	36:  "Language_GSW",
	37:  "Language_GU",
	38:  "Language_HA",
	39:  "Language_HE",
	40:  "Language_HI",
	41:  "Language_HR",
	42:  "Language_HSB",
	43:  "Language_HU",
	44:  "Language_HY",
	45:  "Language_ID",
	46:  "Language_IG",
	47:  "Language_II",
	48:  "Language_IS",
	49:  "Language_IT",
	50:  "Language_IU",
	51:  "Language_JA",
	52:  "Language_KA",
	53:  "Language_KK",
	54:  "Language_KL",
	55:  "Language_KM",
	56:  "Language_KN",
	57:  "Language_KOK",
	58:  "Language_KO",
	59:  "Language_KY",
	60:  "Language_LB",
	61:  "Language_LO",
	62:  "Language_LT",
	63:  "Language_LV",
	64:  "Language_MI",
	65:  "Language_MK",
	66:  "Language_ML",
	67:  "Language_MN",
	68:  "Language_MOH",
	69:  "Language_MR",
	70:  "Language_MS",
	71:  "Language_MT",
	72:  "Language_NB",
	73:  "Language_NE",
	74:  "Language_NL",
	75:  "Language_NN",
	76:  "Language_NSO",
	77:  "Language_OC",
	78:  "Language_OR",
	79:  "Language_PA",
	80:  "Language_PL",
	81:  "Language_PRS",
	82:  "Language_PT",
	83:  "Language_QUT",
	84:  "Language_QUZ",
	85:  "Language_RM",
	86:  "Language_RO",
	87:  "Language_RU",
	88:  "Language_RW",
	89:  "Language_SAH",
	90:  "Language_SA",
	91:  "Language_SE",
	92:  "Language_SI",
	93:  "Language_SK",
	94:  "Language_SL",
	95:  "Language_SMA",
	96:  "Language_SMJ",
	97:  "Language_SMN",
	98:  "Language_SQ",
	99:  "Language_SR",
	100: "Language_SV",
	101: "Language_SW",
	102: "Language_SYR",
	103: "Language_TA",
	104: "Language_TE",
	105: "Language_TG",
	106: "Language_TH",
	107: "Language_TK",
	108: "Language_TN",
	109: "Language_TR",
	110: "Language_TT",
	111: "Language_TZM",
	112: "Language_UG",
	113: "Language_UK",
	114: "Language_UR",
	115: "Language_UZ",
	116: "Language_VI",
	117: "Language_WO",
	118: "Language_XH",
	119: "Language_YO",
	120: "Language_ZH",
	121: "Language_ZU",
}
var Tmpl_20AD45E7_Language_value = map[string]int32{
	"Language_UNDEFINED": 0,
	"Language_AF":        1,
	"Language_AM":        2,
	"Language_AR":        3,
	"Language_ARN":       4,
	"Language_AS":        5,
	"Language_AZ":        6,
	"Language_BA":        7,
	"Language_BE":        8,
	"Language_BG":        9,
	"Language_BN":        10,
	"Language_BO":        11,
	"Language_BR":        12,
	"Language_BS":        13,
	"Language_CA":        14,
	"Language_CO":        15,
	"Language_CS":        16,
	"Language_CY":        17,
	"Language_DA":        18,
	"Language_DE":        19,
	"Language_DSB":       20,
	"Language_DV":        21,
	"Language_EL":        22,
	"Language_EN":        23,
	"Language_ES":        24,
	"Language_ET":        25,
	"Language_EU":        26,
	"Language_FA":        27,
	"Language_FI":        28,
	"Language_FIL":       29,
	"Language_FO":        30,
	"Language_FR":        31,
	"Language_FY":        32,
	"Language_GA":        33,
	"Language_GD":        34,
	"Language_GL":        35,
	"Language_GSW":       36,
	"Language_GU":        37,
	"Language_HA":        38,
	"Language_HE":        39,
	"Language_HI":        40,
	"Language_HR":        41,
	"Language_HSB":       42,
	"Language_HU":        43,
	"Language_HY":        44,
	"Language_ID":        45,
	"Language_IG":        46,
	"Language_II":        47,
	"Language_IS":        48,
	"Language_IT":        49,
	"Language_IU":        50,
	"Language_JA":        51,
	"Language_KA":        52,
	"Language_KK":        53,
	"Language_KL":        54,
	"Language_KM":        55,
	"Language_KN":        56,
	"Language_KOK":       57,
	"Language_KO":        58,
	"Language_KY":        59,
	"Language_LB":        60,
	"Language_LO":        61,
	"Language_LT":        62,
	"Language_LV":        63,
	"Language_MI":        64,
	"Language_MK":        65,
	"Language_ML":        66,
	"Language_MN":        67,
	"Language_MOH":       68,
	"Language_MR":        69,
	"Language_MS":        70,
	"Language_MT":        71,
	"Language_NB":        72,
	"Language_NE":        73,
	"Language_NL":        74,
	"Language_NN":        75,
	"Language_NSO":       76,
	"Language_OC":        77,
	"Language_OR":        78,
	"Language_PA":        79,
	"Language_PL":        80,
	"Language_PRS":       81,
	"Language_PT":        82,
	"Language_QUT":       83,
	"Language_QUZ":       84,
	"Language_RM":        85,
	"Language_RO":        86,
	"Language_RU":        87,
	"Language_RW":        88,
	"Language_SAH":       89,
	"Language_SA":        90,
	"Language_SE":        91,
	"Language_SI":        92,
	"Language_SK":        93,
	"Language_SL":        94,
	"Language_SMA":       95,
	"Language_SMJ":       96,
	"Language_SMN":       97,
	"Language_SQ":        98,
	"Language_SR":        99,
	"Language_SV":        100,
	"Language_SW":        101,
	"Language_SYR":       102,
	"Language_TA":        103,
	"Language_TE":        104,
	"Language_TG":        105,
	"Language_TH":        106,
	"Language_TK":        107,
	"Language_TN":        108,
	"Language_TR":        109,
	"Language_TT":        110,
	"Language_TZM":       111,
	"Language_UG":        112,
	"Language_UK":        113,
	"Language_UR":        114,
	"Language_UZ":        115,
	"Language_VI":        116,
	"Language_WO":        117,
	"Language_XH":        118,
	"Language_YO":        119,
	"Language_ZH":        120,
	"Language_ZU":        121,
}

func (x Tmpl_20AD45E7_Language) String() string {
	return proto.EnumName(Tmpl_20AD45E7_Language_name, int32(x))
}
func (Tmpl_20AD45E7_Language) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// Basic
// 20ad45e7bf208e03045a949b02860882a3dd93bf7f8efabca07da100eaa3eafc
type Tmpl_20AD45E7 struct {
	Name        string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Date        uint64                 `protobuf:"varint,3,opt,name=date" json:"date,omitempty"`
	Language    Tmpl_20AD45E7_Language `protobuf:"varint,4,opt,name=language,enum=oipProto.templates.Tmpl_20AD45E7_Language" json:"language,omitempty"`
	Avatar      *oipProto.Txid         `protobuf:"bytes,5,opt,name=avatar" json:"avatar,omitempty"`
	TagList     []string               `protobuf:"bytes,6,rep,name=tagList" json:"tagList,omitempty"`
	NoteList    []string               `protobuf:"bytes,7,rep,name=noteList" json:"noteList,omitempty"`
	UrlList     []*oipProto.Txid       `protobuf:"bytes,8,rep,name=urlList" json:"urlList,omitempty"`
}

func (m *Tmpl_20AD45E7) Reset()                    { *m = Tmpl_20AD45E7{} }
func (m *Tmpl_20AD45E7) String() string            { return proto.CompactTextString(m) }
func (*Tmpl_20AD45E7) ProtoMessage()               {}
func (*Tmpl_20AD45E7) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Tmpl_20AD45E7) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tmpl_20AD45E7) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Tmpl_20AD45E7) GetDate() uint64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Tmpl_20AD45E7) GetLanguage() Tmpl_20AD45E7_Language {
	if m != nil {
		return m.Language
	}
	return Tmpl_20AD45E7_Language_UNDEFINED
}

func (m *Tmpl_20AD45E7) GetAvatar() *oipProto.Txid {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *Tmpl_20AD45E7) GetTagList() []string {
	if m != nil {
		return m.TagList
	}
	return nil
}

func (m *Tmpl_20AD45E7) GetNoteList() []string {
	if m != nil {
		return m.NoteList
	}
	return nil
}

func (m *Tmpl_20AD45E7) GetUrlList() []*oipProto.Txid {
	if m != nil {
		return m.UrlList
	}
	return nil
}

func init() {
	proto.RegisterType((*Tmpl_20AD45E7)(nil), "oipProto.templates.tmpl_20AD45E7")
	proto.RegisterEnum("oipProto.templates.Tmpl_20AD45E7_Language", Tmpl_20AD45E7_Language_name, Tmpl_20AD45E7_Language_value)
}

func init() { proto.RegisterFile("tmpl_20AD45E7.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 792 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd6, 0x6d, 0x5f, 0xdb, 0x54,
	0x18, 0x06, 0x70, 0x3b, 0x18, 0x94, 0xb0, 0xb1, 0xcb, 0x33, 0x9d, 0x15, 0x9f, 0x2a, 0xea, 0xac,
	0x53, 0xdb, 0xc9, 0x86, 0x53, 0x71, 0xea, 0x29, 0x49, 0x9b, 0x90, 0x27, 0x38, 0x49, 0xda, 0xb5,
	0x3e, 0x60, 0xa0, 0xb1, 0x8b, 0xb6, 0x4d, 0x2c, 0x61, 0xc3, 0xef, 0xe6, 0x07, 0xf2, 0x63, 0xf8,
	0x23, 0x0e, 0xe4, 0x42, 0x5f, 0x71, 0xce, 0xff, 0xbe, 0xee, 0xfb, 0xdc, 0xf4, 0x55, 0xb4, 0xdb,
	0xc5, 0x34, 0x9f, 0x1c, 0x6c, 0xde, 0x97, 0xfa, 0xc3, 0x2d, 0xe3, 0x51, 0x33, 0x9f, 0x67, 0x45,
	0x26, 0x44, 0x96, 0xe6, 0x7b, 0x67, 0xa7, 0x66, 0x91, 0x4c, 0xf3, 0x49, 0x5c, 0x24, 0xc7, 0xeb,
	0x5a, 0x71, 0x9a, 0x8e, 0xfe, 0xa9, 0x6f, 0xfc, 0x25, 0xb4, 0x9b, 0xd4, 0x27, 0x84, 0xb6, 0x38,
	0x8b, 0xa7, 0x49, 0xad, 0x52, 0xaf, 0x34, 0x56, 0x54, 0x79, 0x16, 0x75, 0x6d, 0x75, 0x94, 0x1c,
	0x1f, 0xcd, 0xd3, 0xbc, 0x48, 0xb3, 0x59, 0xed, 0x5a, 0x59, 0xba, 0x4c, 0x67, 0x5d, 0xa3, 0xb8,
	0x48, 0x6a, 0x0b, 0xf5, 0x4a, 0x63, 0x51, 0x95, 0x67, 0xd1, 0xd1, 0xaa, 0x93, 0x78, 0x36, 0x3e,
	0x89, 0xc7, 0x49, 0x6d, 0xb1, 0x5e, 0x69, 0xac, 0x6d, 0xde, 0x6b, 0xfe, 0x77, 0x9d, 0x26, 0xaf,
	0xed, 0xbc, 0xe8, 0x50, 0x17, 0xbd, 0xe2, 0xae, 0xb6, 0x14, 0x3f, 0x8b, 0x8b, 0x78, 0x5e, 0xbb,
	0x5e, 0xaf, 0x34, 0x56, 0x37, 0xd7, 0xfe, 0x9d, 0x12, 0x9e, 0xa6, 0x23, 0xf5, 0xa2, 0x2a, 0x6a,
	0xda, 0x72, 0x11, 0x8f, 0x9d, 0xf4, 0xb8, 0xa8, 0x2d, 0xd5, 0x17, 0x1a, 0x2b, 0xea, 0xfc, 0x2a,
	0xd6, 0xb5, 0xea, 0x2c, 0x2b, 0x92, 0xb2, 0xb4, 0x5c, 0x96, 0x2e, 0xee, 0xa2, 0xa1, 0x2d, 0x9f,
	0xcc, 0x27, 0x65, 0xa9, 0x5a, 0x5f, 0xf8, 0x9f, 0xf1, 0xe7, 0xe5, 0x8d, 0x3f, 0xa1, 0x55, 0xcf,
	0xd7, 0x13, 0x77, 0x34, 0x71, 0x7e, 0x3e, 0x88, 0x3c, 0xdd, 0xe8, 0x58, 0x9e, 0xa1, 0xe3, 0x25,
	0x71, 0x4b, 0x5b, 0xbd, 0x70, 0xd9, 0x41, 0x85, 0xc1, 0xc5, 0x35, 0x06, 0x85, 0x05, 0x01, 0xed,
	0xc6, 0x25, 0xf0, 0xb0, 0xc8, 0x91, 0x00, 0xd7, 0x19, 0x86, 0x58, 0x22, 0x68, 0x4b, 0x2c, 0x33,
	0x18, 0xa8, 0x32, 0x74, 0xb1, 0xc2, 0xe0, 0x41, 0x63, 0xf0, 0xb1, 0xca, 0xa0, 0x70, 0x83, 0x21,
	0xc0, 0x4d, 0x82, 0x1d, 0x89, 0x35, 0x06, 0x1f, 0xb7, 0x18, 0x02, 0x80, 0x61, 0x80, 0x97, 0x09,
	0x74, 0x09, 0xc1, 0x60, 0xe0, 0x36, 0xfd, 0x20, 0x7a, 0xd0, 0xc6, 0x2b, 0x1c, 0xe9, 0xe1, 0x55,
	0x02, 0xc3, 0xc1, 0x1d, 0x06, 0x0f, 0xaf, 0x31, 0x04, 0xa8, 0x31, 0x84, 0x78, 0x9d, 0x21, 0xc2,
	0x3a, 0x41, 0x47, 0xe2, 0x0d, 0x06, 0x0b, 0x6f, 0xd2, 0x66, 0x1d, 0xcb, 0xc1, 0x5b, 0x1c, 0xf1,
	0xf1, 0x36, 0x83, 0xc2, 0x3b, 0x0c, 0x03, 0xd4, 0x09, 0xba, 0x12, 0xef, 0x32, 0xe8, 0xd8, 0x60,
	0x70, 0xf0, 0x1e, 0xbd, 0xdb, 0x0d, 0xfa, 0x78, 0x9f, 0x23, 0x11, 0x3e, 0x20, 0x30, 0x25, 0xee,
	0x32, 0x18, 0xf8, 0x90, 0xc1, 0x42, 0x83, 0x41, 0xe1, 0x23, 0x7a, 0xc6, 0x0c, 0xda, 0xb8, 0xc7,
	0x91, 0x08, 0x1f, 0x33, 0x0c, 0xf0, 0x09, 0x81, 0xa5, 0xe3, 0x53, 0x86, 0x2e, 0x9a, 0x0c, 0x16,
	0x5a, 0x0c, 0x01, 0xee, 0x33, 0x84, 0xf8, 0x8c, 0x21, 0xc2, 0x26, 0xc1, 0xae, 0xc4, 0x03, 0x02,
	0x5b, 0xe2, 0x21, 0x83, 0x8d, 0x2d, 0x06, 0x07, 0x9f, 0x33, 0xb8, 0x78, 0xc4, 0xe0, 0xe1, 0x0b,
	0xfa, 0xff, 0x6d, 0xdf, 0xc6, 0x97, 0x1c, 0xf1, 0xf1, 0x15, 0xc3, 0x00, 0xdb, 0x04, 0x4e, 0x1b,
	0x5f, 0x33, 0xf8, 0x78, 0xcc, 0x10, 0xe2, 0x1b, 0x86, 0x1e, 0xbe, 0x25, 0x70, 0x2d, 0x7c, 0xc7,
	0x60, 0x43, 0x32, 0x38, 0x68, 0x33, 0x78, 0xd8, 0xa1, 0xdd, 0x5d, 0xdf, 0x84, 0xce, 0x11, 0x05,
	0x83, 0x21, 0x40, 0x87, 0x21, 0x44, 0x97, 0xc0, 0x6b, 0xc3, 0x64, 0x30, 0x60, 0x31, 0x38, 0xd8,
	0x65, 0xf0, 0x60, 0xd3, 0x22, 0x5e, 0xe0, 0xc3, 0xa1, 0x88, 0xbf, 0x03, 0x97, 0x41, 0xc1, 0x23,
	0xd8, 0x93, 0xf0, 0x19, 0x1c, 0xec, 0xd1, 0xd4, 0x3d, 0x15, 0x60, 0x9f, 0x23, 0x21, 0x14, 0x45,
	0xf6, 0xa3, 0x10, 0xc1, 0x15, 0x19, 0x22, 0xa4, 0x26, 0xe5, 0x22, 0x62, 0xf0, 0xd1, 0x63, 0x88,
	0xd0, 0x67, 0xe8, 0xe3, 0x09, 0x4d, 0x0d, 0xa4, 0x89, 0x01, 0x45, 0x02, 0x89, 0x21, 0x83, 0x81,
	0xef, 0x19, 0x2c, 0xfc, 0xc0, 0x60, 0xe3, 0x47, 0x06, 0x07, 0x3f, 0xf1, 0x33, 0xae, 0xc4, 0xc1,
	0x15, 0xd9, 0xc5, 0xcf, 0x57, 0xc4, 0x43, 0xcc, 0x63, 0xf6, 0x71, 0xc8, 0xa0, 0x70, 0xc4, 0xd0,
	0xc3, 0x88, 0xa1, 0x8f, 0x84, 0xa7, 0x0e, 0x14, 0x7e, 0xa1, 0x48, 0x28, 0x31, 0x66, 0x30, 0xf0,
	0x94, 0xa1, 0x8b, 0x94, 0xc1, 0xc4, 0xaf, 0x0c, 0x36, 0x7e, 0x63, 0xf0, 0x30, 0x61, 0x50, 0x98,
	0x32, 0x84, 0x98, 0xd1, 0x66, 0xe1, 0xd0, 0x45, 0x46, 0x91, 0xa8, 0x8b, 0x9c, 0xc1, 0xc6, 0xef,
	0x0c, 0x0a, 0x73, 0x86, 0x21, 0x8e, 0x09, 0x7a, 0x16, 0x0a, 0x82, 0xbe, 0x8f, 0x13, 0x82, 0x27,
	0x26, 0x9e, 0x11, 0x0c, 0x7c, 0x3c, 0x27, 0x18, 0x9a, 0x38, 0x65, 0x88, 0xf0, 0x47, 0xfb, 0xf1,
	0x70, 0x7b, 0x9c, 0x16, 0x4f, 0x4f, 0x0e, 0x9b, 0x47, 0xd9, 0xb4, 0x95, 0xa5, 0xf9, 0xf3, 0x71,
	0xab, 0xfc, 0x08, 0x6b, 0x8d, 0xb3, 0x56, 0x7e, 0x78, 0x90, 0xa5, 0xf9, 0xd6, 0xd9, 0xdf, 0x8b,
	0x2f, 0xa3, 0xed, 0xcb, 0x97, 0xc3, 0xa5, 0x32, 0xfb, 0xe0, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd9, 0x44, 0xbe, 0xf5, 0xe7, 0x09, 0x00, 0x00,
}
