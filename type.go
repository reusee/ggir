package main

import "strings"

func cTypeToGoType(t string) string {
	t = strings.Replace(t, "const", "", -1) // const and non-const pointer is compatible in Go 1.2
	t = strings.TrimSpace(t)
	t = strings.Replace(t, "volatile", "", -1) // no volatile
	t = strings.TrimSpace(t)
	t = strings.Replace(t, "long double", "double", -1) // no long double, FIXME
	t = strings.TrimSpace(t)
	pointerDepth := 0 // pointer depth
	for strings.HasSuffix(t, "*") {
		pointerDepth++
		t = t[:len(t)-1]
	}
	t = "C." + t                              // cgo type
	t = strings.Repeat("*", pointerDepth) + t // pointer
	return strings.TrimSpace(t)
}

func (self *Param) MapType() {
	//FIXME
}
