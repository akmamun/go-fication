package helpers

import (
	"encoding/json"
	"io"
)

func DecodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	// decoder.DisallowUnknownFields()
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
func JsonString(data interface{}) string {
	err, _ := json.Marshal(data)
	if err != nil {
		return string(err)
	}
	return ""
}
