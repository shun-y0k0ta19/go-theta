// Copyright (c) 2017 "Shun Yokota" All rights reserved
//
// Part of the source code is adapted from https://github.com/google/go-github
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theta

import (
	"fmt"
	"testing"
)

func TestStringify(t *testing.T) {
	var nilPointer *string

	var tests = []struct {
		in  interface{}
		out string
	}{
		// basic types
		{"foo", `"foo"`},
		{123, `123`},
		{1.5, `1.5`},
		{false, `false`},
		{
			[]string{"a", "b"},
			`["a" "b"]`,
		},
		{
			struct {
				A []string
			}{nil},
			// nil slice is skipped
			`{}`,
		},
		{
			struct {
				A string
			}{"foo"},
			// structs not of a named type get no prefix
			`{A:"foo"}`,
		},

		// pointers
		{nilPointer, `<nil>`},
		{String("foo"), `"foo"`},
		{Int(123), `123`},
		{Bool(false), `false`},
		{
			[]*string{String("a"), String("b")},
			`["a" "b"]`,
		},
	}

	for i, tt := range tests {
		s := Stringify(tt.in)
		if s != tt.out {
			t.Errorf("%d. Stringify(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}

// Directly test the String() methods on various Theta API types. We don't do an
// exaustive test of all the various field types, since TestStringify() above
// takes care of that. Rather, we just make sure that Stringify() is being
// used to build the strings, which we do by verifying that pointers are
// stringified as their underlying value.
func TestString(t *testing.T) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		{CommandRequest{Name: String("name")}, `theta.CommandRequest{Name:"name"}`},
		{Parameters{Options: &Options{Aparture: Float64(1.9)}}, `theta.Parameters{Options:theta.Options{Aparture:1.9}}`},
		{CommandResponse{Name: String("name")}, `theta.CommandResponse{Name:"name"}`},
		{Results{Timeout: Int(5)}, `theta.Results{Timeout:5}`},
		{Entries{Name: String("name")}, `theta.Entries{Name:"name"}`},
		{EXIF{EXIFVersion: String("hoge")}, `theta.EXIF{EXIFVersion:"hoge"}`},
		{XMP{UsePanoramaViewer: Bool(true)}, `theta.XMP{UsePanoramaViewer:true}`},
		{Error{Code: "code", Message: "message"}, `theta.Error{Code:"code", Message:"message"}`},
		{Info{Endpoints: &Endpoints{HTTPPort: Int(80)}}, `theta.Info{Endpoints:theta.Endpoints{HTTPPort:80}}`},
		{Options{
			Aparture: Float64(1.0),
			AutoBracket: &Bracket{
				BracketNumber: Int(1),
				BracketParameters: &BracketParameters{
					ShutterSpeed: Float64(1.0),
					ISO:          Int(800),
				},
			},
		}, `theta.Options{Aparture:1, AutoBracket:theta.Bracket{BracketNumber:1, BracketParameters:theta.BracketParameters{ShutterSpeed:1, ISO:800}}}`},
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
