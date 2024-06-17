package utils

import (
	"testing"
	"json-utils/utils"

	"github.com/stretchr/testify/require"
)

func Test_GetNumOfMessages(t *testing.T) {
	tests := []struct {
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
			name:     "Simple JSON array",
			data:     [][]byte{[]byte(`[{"msg": "hello"}, {"msg": "world"}]`)},
			expected: 2,
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := utils.GetNumOfMessages(test.data...)
			if test.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, test.expected, result)
		})
	}
}
