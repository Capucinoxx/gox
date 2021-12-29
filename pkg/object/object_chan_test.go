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
				ch.Fn(tt.value)
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
			wantStr:  "-180",
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
			ch.GenFunc("-")

			for i := 0; i < tt.loopTime; i++ {
				ch.Fn(tt.value)
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

func TestChanMul(t *testing.T) {
	cases := map[string]struct {
		loopTime int
		value    object.Object
		wantStr  string
	}{
		"object.Chan.Mul(5) 5 times: object.Number": {
			loopTime: 5,
			value:    object.Number{Value: 5},
			wantStr:  "3125",
		},
		"object.Chan.Mul('t') 5 times: object.Error": {
			loopTime: 5,
			value:    object.String{Value: "t"},
			wantStr:  object.ERROR_MULTIPLICATION_STRING + "t",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			ch := object.Chan{}
			ch.Make()
			ch.GenFunc("*")

			for i := 0; i < tt.loopTime; i++ {
				ch.Fn(tt.value)
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

func TestChanDiv(t *testing.T) {
	cases := map[string]struct {
		loopTime int
		value    object.Object
		wantStr  string
	}{
		"object.Chan.Div(5) 5 times: object.Number": {
			loopTime: 5,
			value:    object.Number{Value: 5},
			wantStr:  "0.008",
		},
		"object.Chan.Div('t') 5 times: object.Error": {
			loopTime: 5,
			value:    object.String{Value: "t"},
			wantStr:  object.ERROR_DIVISION_STRING + "t",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			ch := object.Chan{}
			ch.Make()
			ch.GenFunc("/")

			for i := 0; i < tt.loopTime; i++ {
				ch.Fn(tt.value)
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
