package object_test

import (
	"github.com/Capucinoxx/gox/pkg/object"
	"reflect"
	"testing"
)

func TestObject(t *testing.T) {
	cases := map[string]struct {
		obj     object.Object
		t       object.Type
		wantStr string
	}{
		"object number": {
			obj:     object.Number{Value: 64},
			t:       object.OBJECT_NUMBER,
			wantStr: "64",
		},
		"object string": {
			obj:     object.String{Value: "test"},
			t:       object.OBJECT_STRING,
			wantStr: "test",
		},
		"object array": {
			obj: object.Array{
				Elements: []object.Object{
					object.Number{Value: 64},
					object.String{Value: "test"},
				},
			},
			t:       object.OBJECT_ARRAY,
			wantStr: "[64, test]",
		},
		"object null": {
			obj:     object.Null{},
			t:       object.OBJECT_NULL,
			wantStr: "null",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			if tt.obj.Type() != tt.t {
				t.Errorf(
					"%q type mismatch: exp=%d got=%d",
					tt.obj,
					tt.obj.Type(),
					tt.t,
				)
			}

			if tt.obj.ToString() != tt.wantStr {
				t.Errorf(
					"%q string mismatch: exp=%s got=%s",
					tt.obj,
					tt.obj.ToString(),
					tt.wantStr,
				)
			}

			if !reflect.DeepEqual((tt.obj.ToInterface()).(object.Object), tt.obj) {
				t.Errorf(
					"%q interface mismatch: exp=%v got=%q",
					tt.obj,
					tt.obj,
					(tt.obj.ToInterface()).(object.Object),
				)
			}
		})
	}
}

func TestChanObjectAdd(t *testing.T) {
	cases := map[string]struct {
		obj     object.Object
		wantStr string
	}{
		"number addition": {
			obj:     object.Number{Value: 13},
			wantStr: "52",
		},
		"string addition": {
			obj:     object.String{Value: "t"},
			wantStr: "tttt",
		},
		"array addition": {
			obj: object.Array{
				Elements: []object.Object{
					object.Number{Value: 64},
					object.String{Value: "test"},
				},
			},
			wantStr: "[64, test, 64, test, 64, test, 64, test]",
		},
		"null addition": {
			obj:     object.Null{},
			wantStr: "null",
		},
		"error addition": {
			obj:     object.Error{Error: "err"},
			wantStr: "err",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			ch := object.Chan{}
			ch.Make()

			ch.Add(tt.obj)
			ch.Add(tt.obj)
			ch.Add(tt.obj)
			ch.Add(tt.obj)

			res := ch.Eval().ToString()
			if tt.wantStr != res {
				t.Errorf(
					"%q addition mismatch: exp=%s got=%s",
					tt.obj,
					res,
					tt.wantStr,
				)
			}
		})
	}
}

func TestChanObjectSub(t *testing.T) {
	cases := map[string]struct {
		obj     object.Object
		wantStr string
	}{
		"number addition": {
			obj:     object.Number{Value: 13},
			wantStr: "0",
		},
		"string addition": {
			obj:     object.String{Value: "t"},
			wantStr: object.ERROR_SUBTRACTING_STRING + "t",
		},
		"array addition": {
			obj: object.Array{
				Elements: []object.Object{
					object.Number{Value: 64},
					object.String{Value: "test"},
				},
			},
			wantStr: "la soustraction n'est pas implémenté pour le type Array",
		},
		"null addition": {
			obj:     object.Null{},
			wantStr: "null",
		},
		"error addition": {
			obj:     object.Error{Error: "err"},
			wantStr: "err",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			ch := object.Chan{}
			ch.Make()

			ch.Add(tt.obj)
			ch.Sub(tt.obj)

			res := ch.Eval().ToString()
			if tt.wantStr != res {
				t.Errorf(
					"%q addition mismatch: exp=%s got=%s",
					tt.obj,
					res,
					tt.wantStr,
				)
			}
		})
	}
}

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
