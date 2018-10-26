package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkStringConcat(b *testing.B) {
	v1, v2 := 1, 2
	for i := 0; i < b.N; i++ {
		var s string
		s = "(" + strconv.Itoa(v1) + "," + strconv.Itoa(v2) + ")"
		_ = s
	}
}

func BenchmarkStringSprintf(b *testing.B) {
	v1, v2 := 1, 2
	for i := 0; i < b.N; i++ {
		var s string
		s = fmt.Sprintf("(%d,%d)", v1, v2)
		_ = s
	}
}

func BenchmarkStringByteBuf(b *testing.B) {
	v1, v2 := 1, 2
	for i := 0; i < b.N; i++ {
		var s string
		var b bytes.Buffer
		b.WriteString("(")
		b.WriteString(strconv.Itoa(v1))
		b.WriteString(",")
		b.WriteString(strconv.Itoa(v2))
		b.WriteString(")")
		s = b.String()
		_ = s
	}
}

func BenchmarkStringStringBuilder(b *testing.B) {
	v1, v2 := 1, 2
	for i := 0; i < b.N; i++ {
		var s string
		var b strings.Builder
		b.WriteString("(")
		b.WriteString(strconv.Itoa(v1))
		b.WriteString(",")
		b.WriteString(strconv.Itoa(v2))
		b.WriteString(")")
		s = b.String()
		_ = s
	}
}
