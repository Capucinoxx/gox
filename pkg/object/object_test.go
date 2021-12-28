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
