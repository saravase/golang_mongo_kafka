package data

import (
	"encoding/json"
	"io"
)

// Deserialize the JSON format into raw data format
func FromJSON(i interface{}, reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(i)
}

// Serialize the raw data format into JSON format
func ToJSON(i interface{}, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(i)
}
