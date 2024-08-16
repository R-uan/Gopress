package gopress

import (
	"fmt"
	"strconv"
)

func ParseContentLength(value string) int64 {
	if value == "" {
		return 0
	}
	length, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("Error parsing Content-Length:", err)
		return 0
	}
	return length
}