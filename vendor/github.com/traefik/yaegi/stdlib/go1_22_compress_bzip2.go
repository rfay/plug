// Code generated by 'yaegi extract compress/bzip2'. DO NOT EDIT.

//go:build go1.22
// +build go1.22

package stdlib

import (
	"compress/bzip2"
	"reflect"
)

func init() {
	Symbols["compress/bzip2/bzip2"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewReader": reflect.ValueOf(bzip2.NewReader),

		// type definitions
		"StructuralError": reflect.ValueOf((*bzip2.StructuralError)(nil)),
	}
}
