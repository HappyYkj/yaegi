// Code generated by 'github.com/containous/yaegi/extract crypto/md5'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"crypto/md5"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["crypto/md5"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BlockSize": reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"New":       reflect.ValueOf(md5.New),
		"Size":      reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"Sum":       reflect.ValueOf(md5.Sum),
	}
}
