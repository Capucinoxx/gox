package object_test

import (
	"github.com/Capucinoxx/gox/pkg/object"
	"testing"
)

func TestChanAdd(t *testing.T) {
	cases := map[string]struct {
		loopTime int
		value    object.Object
		wantStr  string
	}{
		"object.Chan.Add(10) 20 times: object.Number": {
			loopTime: 20,
			value:    object.Number{Value: 10},
			wantStr:  "200",
		},
		"object.Chan.Add('t') 10 times: oject.String": {
			loopTime: 10,
			value:    object.String{Value: "t"},
			wantStr:  "tttttttttt",
		},
		"object.Chan.Add(object.Null) 10 times: oject.Null": {
			loopTime: 10,
			value:    object.Null{},
			wantStr:  "null",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			ch := object.Chan{}
			ch.Make()
			ch.GenFunc("+")

			for i := 0; i < tt.loopTime; i++ {
				ch.Add(tt.value)
			}

			res := ch.Eval()

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

func TestChanSub(t *testing.T) {
	cases := map[string]struct {
		loopTime int
		value    object.Object
		wantStr  string
	}{
		"object.Chan.Sub(10) 20 times: object.Number": {
			loopTime: 20,
			value:    object.Number{Value: 10},
			wantStr:  "-200",
		},
		"object.Chan.Sub('t') 10 times: object.Error": {
			loopTime: 10,
			value:    object.String{Value: "t"},
			wantStr:  object.ERROR_SUBTRACTION_STRING + "t",
		},
		"object.Chan.Sub(object.Null) 10 times: object.Null": {
			loopTime: 10,
			value:    object.Null{},
			wantStr:  "null",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			ch := object.Chan{}
			ch.Make()
			ch.GenFunc("+")

			for i := 0; i < tt.loopTime; i++ {
				ch.Sub(tt.value)
			}

			res := ch.Eval()

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
