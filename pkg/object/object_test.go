package object_test

import (
	"github.com/Capucinoxx/gox/pkg/object"
	"testing"
)

func TestObject(t *testing.T) {
	cases := map[string]struct {
		obj     object.Object
		t       object.Type
		wantStr string
		value   interface{}
	}{
		"object number": {
			obj:     object.Number{Value: 64},
			t:       object.OBJECT_NUMBER,
			wantStr: "64",
			value:   64.0,
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

			if tt.obj.ToInterface() != tt.value {
				t.Errorf(
					"%q interface mismatch: exp=%v got=%q",
					tt.obj,
					tt.obj.ToInterface(),
					tt.value,
				)
			}
		})
	}
}
