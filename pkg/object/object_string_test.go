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
		"failed aditionying String with Null": {
			a:       object.String{Value: "tt"},
			b:       object.Null{},
			wantStr: object.ERROR_ADDITINYING_STRING + "null",
		},
		"failed aditionying String with Array": {
			a:       object.String{Value: "tt"},
			b:       object.Array{},
			wantStr: object.ERROR_ADDITINYING_STRING + "[]",
		},
		"success aditionying String with String": {
			a:       object.String{Value: "tt"},
			b:       object.String{Value: "tt"},
			wantStr: "tttt",
		},
		"success aditionying String with Number 1": {
			a:       object.String{Value: "tt"},
			b:       object.Number{Value: 3},
			wantStr: "tt3",
		},
		"success aditionying String with Number 2": {
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
		"failed subtraction String with String": {
			a:       object.String{Value: "tt"},
			b:       object.String{Value: "tt"},
			wantStr: object.ERROR_SUBTRACTING_STRING + "tt",
		},
		"failed subtraction String with Number": {
			a:       object.String{Value: "tt"},
			b:       object.Number{Value: 3.14},
			wantStr: object.ERROR_SUBTRACTING_STRING + "3.14",
		},
		"failed subtraction String with Array": {
			a:       object.String{Value: "tt"},
			b:       object.Array{},
			wantStr: object.ERROR_SUBTRACTING_STRING + "[]",
		},
		"failed subtraction String with Null": {
			a:       object.String{Value: "tt"},
			b:       object.Null{},
			wantStr: object.ERROR_SUBTRACTING_STRING + "null",
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
		"failed multiply String with String": {
			a:       object.String{Value: "tt"},
			b:       object.String{Value: "tt"},
			wantStr: object.ERROR_MULTIPLYING_STRING + "tt",
		},
		"failed mul String with Null": {
			a:       object.String{Value: "tt"},
			b:       object.Null{},
			wantStr: object.ERROR_MULTIPLYING_STRING + "null",
		},
		"success multiply String with Number": {
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
		"failed division String with String": {
			a:       object.String{Value: "tt"},
			b:       object.String{Value: "tt"},
			wantStr: object.ERROR_DIVISION_STRING + "tt",
		},
		"failed division String with Null": {
			a:       object.String{Value: "tt"},
			b:       object.Null{},
			wantStr: object.ERROR_DIVISION_STRING + "null",
		},
		"failed division String with Number": {
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
