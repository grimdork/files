package files

import (
	"encoding/json"
	"os"
)

// LoadJSON and unmarshal structure.
func LoadJSON(fn string, out interface{}) error {
	f, err := os.ReadFile(fn)
	if err != nil {
		return err
	}

	return json.Unmarshal(f, out)
}

// SaveJSON after marshalling neatly, logging any file errors to stderr.
func SaveJSON(path string, data interface{}) error {
	var b []byte
	var err error
	b, err = json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	WriteFile(path, b)
	return nil
}
