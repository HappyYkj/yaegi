// Code generated by 'goexports encoding/gob'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"encoding/gob"
	"reflect"
)

func init() {
	Symbols["encoding/gob"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewDecoder":   reflect.ValueOf(gob.NewDecoder),
		"NewEncoder":   reflect.ValueOf(gob.NewEncoder),
		"Register":     reflect.ValueOf(gob.Register),
		"RegisterName": reflect.ValueOf(gob.RegisterName),

		// type definitions
		"CommonType": reflect.ValueOf((*gob.CommonType)(nil)),
		"Decoder":    reflect.ValueOf((*gob.Decoder)(nil)),
		"Encoder":    reflect.ValueOf((*gob.Encoder)(nil)),
		"GobDecoder": reflect.ValueOf((*gob.GobDecoder)(nil)),
		"GobEncoder": reflect.ValueOf((*gob.GobEncoder)(nil)),

		// interface wrapper definitions
		"_GobDecoder": reflect.ValueOf((*_encoding_gob_GobDecoder)(nil)),
		"_GobEncoder": reflect.ValueOf((*_encoding_gob_GobEncoder)(nil)),
	}
}

// _encoding_gob_GobDecoder is an interface wrapper for GobDecoder type
type _encoding_gob_GobDecoder struct {
	WGobDecode func(a0 []byte) error
}

func (W _encoding_gob_GobDecoder) GobDecode(a0 []byte) error { return W.WGobDecode(a0) }

// _encoding_gob_GobEncoder is an interface wrapper for GobEncoder type
type _encoding_gob_GobEncoder struct {
	WGobEncode func() ([]byte, error)
}

func (W _encoding_gob_GobEncoder) GobEncode() ([]byte, error) { return W.WGobEncode() }