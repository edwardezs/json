package utils

import (
	"bytes"
	"encoding/json"
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// GetNumOfMessages returns number of consistent models in slice of json bytes using std json decoder
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

// GetNumOfMessagesV2 returns number of consistent models in slice of json bytes using jsoniter
func GetNumOfMessagesV2(data ...[]byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var msgCount int
	var models []map[string]interface{}

	for _, batch := range data {
		if err := json.Unmarshal(batch, &models); err != nil {
			return 0, errors.Wrap(err, "error while decoding json")
		}
		msgCount += len(models)
	}

	return msgCount, nil
}
