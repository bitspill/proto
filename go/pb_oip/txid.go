package pb_oip

import (
	"encoding/hex"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
)

func TxidFromString(txid string) *Txid {
	if len(txid) != 64 {
		return nil
	}
	b, err := hex.DecodeString(txid)
	if err != nil {
		return nil
	}
	return &Txid{Raw: b}
}

func TxidToString(txid *Txid) string {
	if txid == nil || len(txid.Raw) != 32 {
		return ""
	}
	return hex.EncodeToString(txid.Raw)
}

func TxidPrefixToUint64(txid *Txid) uint64 {
	if len(txid.Raw) == 32 {
		return uint64(txid.Raw[0])<<56 |
			uint64(txid.Raw[1])<<48 |
			uint64(txid.Raw[2])<<40 |
			uint64(txid.Raw[3])<<32 |
			uint64(txid.Raw[4])<<24 |
			uint64(txid.Raw[5])<<16 |
			uint64(txid.Raw[6])<<8 |
			uint64(txid.Raw[7])<<0
	}
	return 0
}

func TxidPrefixToString(txid *Txid) string {
	if txid == nil || len(txid.Raw) != 32 {
		return ""
	}
	return hex.EncodeToString(txid.Raw[0:8])
}

func (m *Txid) UnmarshalJSONPB(u *jsonpb.Unmarshaler, b []byte) error {
	dst := make([]byte, len(b)/2)
	_, err := hex.Decode(dst, b)
	if err != nil {
		return err
	}
	return nil
}

func (m *Txid) MarshalJSONPB(marsh *jsonpb.Marshaler) ([]byte, error) {
	dst := make([]byte, len(m.Raw)*2)
	hex.Encode(dst, m.Raw)
	return json.Marshal(string(dst))
}
