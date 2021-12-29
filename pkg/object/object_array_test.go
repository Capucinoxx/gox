package object_test

import (
	"github.com/Capucinoxx/gox/pkg/object"
	"testing"
)

func TestArrayAdd(t *testing.T) {
	cases := map[string]struct {
		a       object.Array
		b       object.Object
		wantStr string
	}{
		"object.Array.Add(object.String): object.Array": {
			a:       object.Array{Elements: []object.Object{object.String{Value: "xxx"}}},
			b:       object.String{Value: "ttt"},
			wantStr: "[xxx, ttt]",
		},
		"object.Array.Add(object.Number): object.Array": {
			a:       object.Array{Elements: []object.Object{object.Number{Value: 3.14}}},
			b:       object.Number{Value: 3.14},
			wantStr: "[3.14, 3.14]",
		},
		"object.Array.Add(object.Array): object.Array": {
			a:       object.Array{Elements: []object.Object{object.Number{Value: 3.14}}},
			b:       object.Array{Elements: []object.Object{object.Number{Value: 3.14}}},
			wantStr: "[3.14, 3.14]",
		},
		"object.Array.Add(object.Error): object.Array": {
			a:       object.Array{Elements: []object.Object{object.Number{Value: 3.14}}},
			b:       object.Error{Error: "err"},
			wantStr: "[3.14, err]",
		},
		"object.Array.Add(object.Null): object.Array": {
			a:       object.Array{Elements: []object.Object{object.Number{Value: 3.14}}},
			b:       object.Null{},
			wantStr: "[3.14]",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			res := tt.a.Add(tt.b)

			if res.ToString() != tt.wantStr {
				t.Errorf(
					"%q string mismatch: exp=%s got=%s",
					res,
					tt.wantStr,
					res.ToString(),
				)
			}
		})
	}
}
