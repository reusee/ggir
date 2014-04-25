package main

import (
	"reflect"
	"strings"

	"./xml"
)

func dumpNotHandled(v reflect.Value) {
	if !v.IsValid() {
		return
	}
	t := v.Type()
	kind := t.Kind()
	if kind == reflect.Ptr {
		dumpNotHandled(v.Elem())
	} else if kind == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			dumpNotHandled(v.Index(i))
		}
	} else if kind == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			if field.Name == "NotHandled" {
				for _, n := range v.Field(i).Interface().([]*NotHandled) {
					n.dump()
				}
			} else if field.Name == "NotHandledAttr" {
				for _, attr := range v.Field(i).Interface().([]xml.Attr) {
					p("%s = %s %s\n", attr.Name.Local, attr.Value, strings.Repeat("=", 64))
				}
			} else {
				dumpNotHandled(v.Field(i))
			}
		}
	}
}

func (n *NotHandled) dump() {
	p("%s %s\n", n.XMLName.Local, strings.Repeat("=", 64))
	p("%s\n", n.Xml)
}
