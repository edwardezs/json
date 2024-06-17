package utils

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// GetNumOfMessages returns number of consistent models in slice of json bytes
func GetNumOfMessages(data ...[]byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}

	var msgCount int
	var stack int

	for _, batch := range data {
		decoder := json.NewDecoder(bytes.NewReader(batch))

		for {
			t, err := decoder.Token()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return 0, errors.Wrap(err, "error while decoding json")
			}

			if delim, ok := t.(json.Delim); ok {
				switch delim {
				case '{':
					stack++
				case '}':
					stack--
				default:
					continue
				}
			} else {
				continue
			}

			if stack == 0 {
				msgCount++
			}
		}
	}
	return msgCount, nil
}
