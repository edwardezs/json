# json-utils

## GetNumOfMessages

This function is designed to count the number of top-level JSON objects in a batch of byte slices. It ignores any nested JSON objects and only counts the top-level ones.

### Usage

```go
package main

import (
    "fmt"
    "github.com/edwardezs/json-utils/utils"
)

func main() {
	// Example data
	data1 := []byte(`{"name": "Alice"}`)
	data2 := []byte(`{"name": "Bob"}{"name": "Charlie"}`)

	// Call GetNumOfMessages function
	count, err := utils.GetNumOfMessages(data1, data2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Number of top-level messages:", count) // 3
}
```