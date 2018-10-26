package main

import (
	"strconv"
	"testing"
)

// Index is a type for representing a generic index.
type Index interface{ String() string }

// intIdx is an int that implements Index
type intIdx int

func (i intIdx) String() string { return strconv.Itoa(int(i)) }

type structIdx struct{ X, Y int }

func (s structIdx) String() string {
	return "(" + strconv.Itoa(s.X) + "," + strconv.Itoa(s.Y) + ")"
}

func BenchmarkIfaceOverheadNativePrimitive(b *testing.B) {
	m := make(map[int]struct{})
	for i := 0; i < b.N; i++ {
		m[i] = struct{}{}
	}
	for i := 0; i < b.N; i++ {
		// force a lookup
		if _, ok := m[i]; ok {
		}
	}
}

func BenchmarkIfaceOverheadNativeStruct(b *testing.B) {
	m := make(map[struct{ X, Y int }]struct{})
	for i := 0; i < b.N; i++ {
		m[struct{ X, Y int }{i, i}] = struct{}{}
	}
	for i := 0; i < b.N; i++ {
		// force a lookup
		if _, ok := m[struct{ X, Y int }{i, i}]; ok {
		}
	}
}

func BenchmarkIfaceOverheadIfacePrimitive(b *testing.B) {
	m := make(map[Index]struct{})
	for i := 0; i < b.N; i++ {
		m[intIdx(i)] = struct{}{}
	}
	for i := 0; i < b.N; i++ {
		// force a lookup
		if _, ok := m[intIdx(i)]; ok {
		}
	}
}

func BenchmarkIfaceOverheadIfaceStruct(b *testing.B) {
	m := make(map[Index]struct{})
	for i := 0; i < b.N; i++ {
		m[structIdx{X: i, Y: i}] = struct{}{}
	}
	for i := 0; i < b.N; i++ {
		// force a lookup
		if _, ok := m[structIdx{X: i, Y: i}]; ok {
		}
	}
}
