package main

import "strings"

func cTypeToGoType(t string) string {
	if t == "" {
		return ""
	}
	replaces := []struct {
		left  string
		right string
	}{
		{"const ", ""},
		{" const", ""},
		{"volatile", ""},
		{"unsigned long long", "ulonglong"},
		{"unsigned long", "ulong"},
		{"unsigned int", "uint"},
		{"long double", "double"},
	}
	for _, e := range replaces {
		t = strings.Replace(t, e.left, e.right, -1)
		t = strings.TrimSpace(t)
	}
	pointerDepth := 0 // pointer depth
	for strings.HasSuffix(t, "*") {
		pointerDepth++
		t = t[:len(t)-1]
	}
	t = "C." + t                              // cgo type
	t = strings.Repeat("*", pointerDepth) + t // pointer
	return strings.TrimSpace(t)
}
