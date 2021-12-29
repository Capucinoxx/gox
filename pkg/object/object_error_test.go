package object_test

import (
	"github.com/Capucinoxx/gox/pkg/object"
	"testing"
)

func TestAddError(t *testing.T) {
	cases := map[string]struct {
		a       object.Error
		b       object.Object
		wantStr string
	}{
		"object.Error.Add(object.String): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.String{},
			wantStr: "err random",
		},
		"object.Error.Add(object.Number): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.Number{},
			wantStr: "err random",
		},
		"object.Error.Add(object.Null): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.Null{},
			wantStr: "err random",
		},
		"object.Error.Add(object.Array): object.Array": {
			a: object.Error{Error: "err random 2"},
			b: object.Array{
				Elements: []object.Object{
					object.Error{Error: "err random 1"},
				},
			},
			wantStr: "[err random 2, err random 1]",
		},
		"object.Error.Add(object.Error): object.Array": {
			a:       object.Error{Error: "err random 1"},
			b:       object.Error{Error: "err random 2"},
			wantStr: "[err random 1, err random 2]",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			res := tt.a.Add(tt.b)

			if tt.wantStr != res.ToString() {
				t.Errorf(
					"%q addition mismatch: exp=%s got=%s",
					tt.b,
					res,
					tt.wantStr,
				)
			}
		})
	}
}

func TestSubError(t *testing.T) {
	cases := map[string]struct {
		a       object.Error
		b       object.Object
		wantStr string
	}{
		"object.Error.Sub(object.String): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.String{},
			wantStr: "err random",
		},
		"object.Error.Sub(object.Number): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.Number{},
			wantStr: "err random",
		},
		"object.Error.Sub(object.Null): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.Null{},
			wantStr: "err random",
		},
		"object.Error.Sub(object.Array): object.Error": {
			a: object.Error{Error: "err random 2"},
			b: object.Array{
				Elements: []object.Object{
					object.Error{Error: "err random 1"},
				},
			},
			wantStr: "err random 2",
		},
		"object.Error.Sub(object.Error): object.Error": {
			a:       object.Error{Error: "err random 1"},
			b:       object.Error{Error: "err random 2"},
			wantStr: "err random 1",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			res := tt.a.Sub(tt.b)

			if tt.wantStr != res.ToString() {
				t.Errorf(
					"%q subtraction mismatch: exp=%s got=%s",
					tt.b,
					res,
					tt.wantStr,
				)
			}
		})
	}
}

func TestMulError(t *testing.T) {
	cases := map[string]struct {
		a       object.Error
		b       object.Object
		wantStr string
	}{
		"object.Error.Mul(object.String): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.String{},
			wantStr: "err random",
		},
		"object.Error.Mul(object.Number): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.Number{},
			wantStr: "err random",
		},
		"object.Error.Mul(object.Null): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.Null{},
			wantStr: "err random",
		},
		"object.Error.Mul(object.Array): object.Error": {
			a: object.Error{Error: "err random 2"},
			b: object.Array{
				Elements: []object.Object{
					object.Error{Error: "err random 1"},
				},
			},
			wantStr: "err random 2",
		},
		"object.Error.Mul(object.Error): object.Error": {
			a:       object.Error{Error: "err random 1"},
			b:       object.Error{Error: "err random 2"},
			wantStr: "err random 1",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			res := tt.a.Mul(tt.b)

			if tt.wantStr != res.ToString() {
				t.Errorf(
					"%q multiplication mismatch: exp=%s got=%s",
					tt.b,
					res,
					tt.wantStr,
				)
			}
		})
	}
}

func TestDivError(t *testing.T) {
	cases := map[string]struct {
		a       object.Error
		b       object.Object
		wantStr string
	}{
		"object.Error.Div(object.String): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.String{},
			wantStr: "err random",
		},
		"object.Error.Div(object.Number): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.Number{},
			wantStr: "err random",
		},
		"object.Error.Div(object.Null): object.Error": {
			a:       object.Error{Error: "err random"},
			b:       object.Null{},
			wantStr: "err random",
		},
		"object.Error.Div(object.Array): object.Error": {
			a: object.Error{Error: "err random 2"},
			b: object.Array{
				Elements: []object.Object{
					object.Error{Error: "err random 1"},
				},
			},
			wantStr: "err random 2",
		},
		"object.Error.Div(object.Error): object.Error": {
			a:       object.Error{Error: "err random 1"},
			b:       object.Error{Error: "err random 2"},
			wantStr: "err random 1",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			res := tt.a.Div(tt.b)

			if tt.wantStr != res.ToString() {
				t.Errorf(
					"%q division mismatch: exp=%s got=%s",
					tt.b,
					res,
					tt.wantStr,
				)
			}
		})
	}
}
