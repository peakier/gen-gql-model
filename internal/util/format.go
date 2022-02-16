package util

import (
	"fmt"
	"go/format"
)

func FormatGoFile(content string) string {
	bFormattedGenerated, err := format.Source([]byte(content))
	if err != nil {
		fmt.Println("warning invalid syntax : can not format file")
		return content
	}
	return string(bFormattedGenerated)
}
