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
