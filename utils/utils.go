package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"slices"

	"github.com/pkg/errors"
)

// GetNumOfMessages returns number of consistent models in slice of json bytes
func GetNumOfMessages(data ...[]byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}

	var msgCount int
	var stack []json.Delim

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
					stack = append(stack, delim)
				case '}':
					stack = slices.Delete(stack, len(stack)-1, len(stack))
				default:
					continue
				}
			} else {
				continue
			}

			if len(stack) == 0 {
				msgCount++
			}
		}
	}
	return msgCount, nil
}
