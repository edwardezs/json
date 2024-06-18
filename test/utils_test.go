package utils

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/edwardezs/json/utils"
	"github.com/stretchr/testify/require"
)

var tests = []struct {
	name     string
	data     [][]byte
	expected int
	hasError bool
}{
	{
		name:     "Empty slice",
		data:     [][]byte{},
		expected: 0,
		hasError: false,
	},
	{
		name:     "Simple JSON array with three messages",
		data:     [][]byte{[]byte(`[{"msg": "hello"}, {"msg": "world"}, {"msg": "!"}]`)},
		expected: 3,
		hasError: false,
	},
	{
		name:     "Invalid JSON",
		data:     [][]byte{[]byte(`invalid json`)},
		expected: 0,
		hasError: true,
	},
	{
		name:     "Nested JSON objects",
		data:     [][]byte{[]byte(`[{"msg": {"text": "hello"}}, {"msg": {"text": "world"}}]`)},
		expected: 2,
		hasError: false,
	},
	{
		name:     "JSON with nested arrays",
		data:     [][]byte{[]byte(`[{"msg": ["hello", "world"]}, {"msg": ["foo", "bar"]}]`)},
		expected: 2,
		hasError: false,
	},
	{
		name: "Mixed nested JSON objects and arrays",
		data: [][]byte{[]byte(`[
			{"msg": {"text": "hello", "details": ["a", "b"]}},
			{"msg": {"text": "world", "details": ["c", "d"]}},
			{"msg": ["nested", {"object": "value"}]}
		]`)},
		expected: 3,
		hasError: false,
	},
	{
		name: "Complex nested structure",
		data: [][]byte{[]byte(`[
			{
				"msg": {
					"text": "hello",
					"details": {
						"part1": {"subpart": ["a", "b"]},
						"part2": ["c", "d"]
					}
				}
			},
			{"msg": {"text": "world"}}
		]`)},
		expected: 2,
		hasError: false,
	},
	{
		name: "Multiple JSON batches",
		data: [][]byte{
			[]byte(`[{"msg": "batch1-message1"}, {"msg": "batch1-message2"}]`),
			[]byte(`[{"msg": "batch2-message1"}, {"msg": "batch2-message2"}, {"msg": "batch2-message3"}]`),
			[]byte(`[{"msg": "batch3-message1"}]`),
		},
		expected: 6,
		hasError: false,
	},
	{
		name:     "Large JSON",
		data:     [][]byte{generateJson(20000000)},
		expected: 20000000,
		hasError: false,
	},
}

func Test_GetNumOfMessages(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := time.Now()
			result, err := utils.GetNumOfMessages(test.data...)
			duration := time.Since(start)
			if test.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, test.expected, result)
			t.Logf("Execution time: %s", duration)
		})
	}
}

func Test_GetNumOfMessagesV2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := time.Now()
			result, err := utils.GetNumOfMessagesV2(test.data...)
			duration := time.Since(start)
			if test.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, test.expected, result)
			t.Logf("Execution time: %s", duration)
		})
	}
}

func generateJson(size int) []byte {
	var buf bytes.Buffer
	buf.WriteString(`[`)
	for i := 0; i < size; i++ {
		buf.WriteString(fmt.Sprintf(`{"msg": "%d"},`, i))
	}
	buf.Truncate(buf.Len() - 1)
	buf.WriteString(`]`)
	return buf.Bytes()
}
