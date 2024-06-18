# json-utils

Package provides utilities for proceeding JSON-files.

## GetNumOfMessages

This function is designed to count the number of top-level JSON objects in a batch of byte slices. It ignores any nested JSON objects and only counts the top-level ones.

### Usage

```go
package main

import (
    "fmt"
    "github.com/edwardezs/json/utils"
)

func main() {
	// Example data
	data1 := []byte(`[{"name": "Alice"}]`)
	data2 := []byte(`[{"name": "Bob"}, {"name": "Charlie"}]`)

	// Call GetNumOfMessages function
	count, err := utils.GetNumOfMessages(data1, data2)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Number of top-level messages: %d", count) // 3
}
```

## GetNumOfMessagesV2

This function is modified version of the previous one. Does the same but works faster especially with large data. 
Based on [jsoniter](github.com/json-iterator/go).

### Usage

```go
package main

import (
    "fmt"
    "github.com/edwardezs/json/utils"
)

func main() {
	// Example data
	data1 := []byte(`[{"name": "Alice"}]`)
	data2 := []byte(`[{"name": "Bob"}, {"name": "Charlie"}]`)

	// Call GetNumOfMessagesV2 function
	count, err := utils.GetNumOfMessagesV2(data1, data2)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Number of top-level messages: %d", count) // 3
}
```
