// Code generated by protoc-gen-go-json. DO NOT EDIT.
// source: e2e.proto

package e2e

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
)

// MarshalJSON implements json.Marshaler
func (msg *Basic) MarshalJSON() ([]byte, error) {
	var m jsonpb.Marshaler
	var buf bytes.Buffer
	err := m.Marshal(&buf, msg)
	return buf.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *Basic) UnmarshalJSON(b []byte) error {
	return jsonpb.Unmarshal(bytes.NewReader(b), msg)
}
