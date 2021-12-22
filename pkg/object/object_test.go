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
