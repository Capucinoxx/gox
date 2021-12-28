package object_test

import (
	"github.com/Capucinoxx/gox/pkg/object"
	"testing"
)

func TestAddNumber(t *testing.T) {
	cases := map[string]struct {
		a       object.Number
		b       object.Object
		wantStr string
	}{
		"object.Number.Add(object.Null): object.Error": {
			a:       object.Number{Value: 3.14},
			b:       object.Null{},
			wantStr: object.ERROR_ADDITION_NUMBER + "null",
		},
		"object.Number.Add(object.Number): object.Number": {
			a:       object.Number{Value: 3.14},
			b:       object.Number{Value: 3.14},
			wantStr: "6.28",
		},
		"object.Number.Add(object.Array): object.Error": {
			a:       object.Number{Value: 3.14},
			b:       object.Array{},
			wantStr: object.ERROR_ADDITION_NUMBER + "[]",
		},
		"object.Number.Add(object.String): object.Error": {
			a:       object.Number{Value: 3.14},
			b:       object.String{},
			wantStr: object.ERROR_ADDITION_NUMBER + "",
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

func TestSubNumber(t *testing.T) {
	cases := map[string]struct {
		a       object.Number
		b       object.Object
		wantStr string
	}{
		"object.Number.Sub(object.Null): object.Error": {
			a:       object.Number{Value: 3.14},
			b:       object.Null{},
			wantStr: object.ERROR_SUBTRACTION_NUMBER + "null",
		},
		"object.Number.Sub(object.Number): object.Number": {
			a:       object.Number{Value: 3.14},
			b:       object.Number{Value: 3.14},
			wantStr: "0",
		},
		"object.Number.Sub(object.Array): object.Error": {
			a:       object.Number{Value: 3.14},
			b:       object.Array{},
			wantStr: object.ERROR_SUBTRACTION_NUMBER + "[]",
		},
		"object.Number.Sub(object.String): object.Error": {
			a:       object.Number{Value: 3.14},
			b:       object.String{},
			wantStr: object.ERROR_SUBTRACTION_NUMBER + "",
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

func TestMulNumber(t *testing.T) {
	cases := map[string]struct {
		a       object.Number
		b       object.Object
		wantStr string
	}{
		"object.Number.Mul(object.String): object.String": {
			a:       object.Number{Value: 3},
			b:       object.String{Value: "tow"},
			wantStr: "towtowtow",
		},
		"object.Number.Mul(object.Number): object.Number": {
			a:       object.Number{Value: 3.14},
			b:       object.Number{Value: 6},
			wantStr: "18.84",
		},
		"object.Number.Mul(object.Null): object.Error": {
			a:       object.Number{Value: 3.14},
			b:       object.Null{},
			wantStr: object.ERROR_MULTIPLICATION_NUMBER + "null",
		},
		"object.Number.Mul(object.Array): object.Error": {
			a:       object.Number{Value: 3.14},
			b:       object.Array{},
			wantStr: object.ERROR_MULTIPLICATION_NUMBER + "[]",
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

func TestDivNumber(t *testing.T) {
	cases := map[string]struct {
		a       object.Number
		b       object.Object
		wantStr string
	}{
		"object.Number.Div(object.String): object.Error": {
			a:       object.Number{Value: 3},
			b:       object.String{Value: "tow"},
			wantStr: object.ERROR_DIVISION_NUMBER + "tow",
		},
		"object.Number.Div(object.Number): object.Number": {
			a:       object.Number{Value: 3.15},
			b:       object.Number{Value: 6},
			wantStr: "0.525",
		},
		"object.Number.Div(object.Null): object.Error": {
			a:       object.Number{Value: 3.14},
			b:       object.Null{},
			wantStr: object.ERROR_DIVISION_NUMBER + "null",
		},
		"object.Number.Div(object.Array): object.Error": {
			a:       object.Number{Value: 3.14},
			b:       object.Array{},
			wantStr: object.ERROR_DIVISION_NUMBER + "[]",
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
