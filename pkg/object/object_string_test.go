package object_test

import (
	"github.com/Capucinoxx/gox/pkg/object"
	"testing"
)

func TestAddString(t *testing.T) {
	cases := map[string]struct {
		a       object.String
		b       object.Object
		wantStr string
	}{
		"object.String.Add(object.Null): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.Null{},
			wantStr: object.ERROR_ADDITION_STRING + "null",
		},
		"object.String.Add(object.Array): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.Array{},
			wantStr: object.ERROR_ADDITION_STRING + "[]",
		},
		"object.String.Add(object.String): object.String": {
			a:       object.String{Value: "tt"},
			b:       object.String{Value: "tt"},
			wantStr: "tttt",
		},
		"object.String.Add(object.Number): object.String (1)": {
			a:       object.String{Value: "tt"},
			b:       object.Number{Value: 3},
			wantStr: "tt3",
		},
		"object.String.Add(object.Number): object.String (2)": {
			a:       object.String{Value: "tt"},
			b:       object.Number{Value: 3.14},
			wantStr: "tt3.14",
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

func TestSubString(t *testing.T) {
	cases := map[string]struct {
		a       object.String
		b       object.Object
		wantStr string
	}{
		"object.String.Sub(object.String): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.String{Value: "tt"},
			wantStr: object.ERROR_SUBTRACTION_STRING + "tt",
		},
		"object.String.Sub(object.Number): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.Number{Value: 3.14},
			wantStr: object.ERROR_SUBTRACTION_STRING + "3.14",
		},
		"object.String.Sub(object.Array): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.Array{},
			wantStr: object.ERROR_SUBTRACTION_STRING + "[]",
		},
		"object.String.Sub(object.Null): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.Null{},
			wantStr: object.ERROR_SUBTRACTION_STRING + "null",
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

func TestMulString(t *testing.T) {
	cases := map[string]struct {
		a       object.String
		b       object.Object
		wantStr string
	}{
		"object.String.Mul(object.String): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.String{Value: "tt"},
			wantStr: object.ERROR_MULTIPLICATION_STRING + "tt",
		},
		"object.String.Mul(object.Null): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.Null{},
			wantStr: object.ERROR_MULTIPLICATION_STRING + "null",
		},
		"object.String.Mul(object.Number): object.String": {
			a:       object.String{Value: "tt"},
			b:       object.Number{Value: 4},
			wantStr: "tttttttt",
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

func TestDivString(t *testing.T) {
	cases := map[string]struct {
		a       object.String
		b       object.Object
		wantStr string
	}{
		"object.String.Div(object.String): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.String{Value: "tt"},
			wantStr: object.ERROR_DIVISION_STRING + "tt",
		},
		"object.String.Div(object.Null): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.Null{},
			wantStr: object.ERROR_DIVISION_STRING + "null",
		},
		"object.String.Div(object.Number): object.Error": {
			a:       object.String{Value: "tt"},
			b:       object.Number{Value: 4},
			wantStr: object.ERROR_DIVISION_STRING + "4",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			res := tt.a.Div(tt.b)

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
