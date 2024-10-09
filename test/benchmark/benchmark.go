package benchmark

import "strings"

func CreateStringWithoutStringBuilder() string {
	var s string = ""
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

func CreateStringWithStringBuilder() string {
	var sb strings.Builder
	for i := 0; i < 1000; i++ {
		sb.WriteString("a")
	}
	return sb.String()
}
