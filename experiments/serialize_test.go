package experiments

import (
	"encoding/json"
	"testing"
)

func TestSerializeValues(t *testing.T) {
	type test struct {
		value interface{}
		want  string
	}

	tests := []test{
		// nums
		{1, "1"},
		{1.2567, "1.2567"},

		// bytes
		{[]byte{}, "\"\""},
		{[]byte{0}, "\"AA==\""},

		// strings
		{"a", "\"a\""},
		{"\"abba\"", "\"\\\"abba\\\"\""},

		// nil
		{nil, "null"},
		{(interface{})(nil), "null"},
		{(*interface{})(nil), "null"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			bts, err := json.Marshal(tt.value)
			if err != nil {
				t.Fatal(err)
			}
			if string(bts) != tt.want {
				t.Fatalf("get=%v, want='%v'", string(bts), tt.want)
			}
		})
	}
}

func TestSerializeMap(t *testing.T) {
	type A struct {
		A string
		B interface{}
	}

	type test struct {
		value interface{}
		want  string
	}

	tests := []test{
		{map[string]interface{}{"a": "b", "bob": nil}, "{\"a\":\"b\",\"bob\":null}"},
		{map[int]*A{1: nil, 0: {}}, "{\"0\":{\"A\":\"\",\"B\":null},\"1\":null}"},
		{map[int]interface{}{1: nil, 0: nil}, "{\"0\":null,\"1\":null}"},
		{map[string]interface{}{"": &A{}, "\"": &A{}}, "{\"\":{\"A\":\"\",\"B\":null},\"\\\"\":{\"A\":\"\",\"B\":null}}"},
		{&map[int]interface{}{1: nil, 0: A{}}, "{\"0\":{\"A\":\"\",\"B\":null},\"1\":null}"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			bts, err := json.Marshal(tt.value)
			if err != nil {
				t.Fatal(err)
			}
			if string(bts) != tt.want {
				t.Fatalf("get=%v, want='%v'", string(bts), tt.want)
			}
		})
	}
}

func TestSerializeStruct(t *testing.T) {
	type P struct {
		a string
		b interface{}
	}
	type A struct {
		A string
		B interface{}
	}

	type test struct {
		value interface{}
		want  string
	}

	tests := []test{
		// private
		{P{"bob", nil}, "{}"},

		{A{"bob", nil}, "{\"A\":\"bob\",\"B\":null}"},
		{&A{"bob", nil}, "{\"A\":\"bob\",\"B\":null}"},
		{&A{"bob", A{A: "mike"}}, "{\"A\":\"bob\",\"B\":{\"A\":\"mike\",\"B\":null}}"},
		{&A{"bob", &A{A: "mike"}}, "{\"A\":\"bob\",\"B\":{\"A\":\"mike\",\"B\":null}}"},
		{&A{"bob", A{A: "mike", B: nil}}, "{\"A\":\"bob\",\"B\":{\"A\":\"mike\",\"B\":null}}"},
		{&A{"bob", &A{A: "mike", B: nil}}, "{\"A\":\"bob\",\"B\":{\"A\":\"mike\",\"B\":null}}"},
		{&A{"bob", A{A: "mike", B: struct{}{}}}, "{\"A\":\"bob\",\"B\":{\"A\":\"mike\",\"B\":{}}}"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			bts, err := json.Marshal(tt.value)
			if err != nil {
				t.Fatal(err)
			}
			if string(bts) != tt.want {
				t.Fatalf("get=%v, want='%v'", string(bts), tt.want)
			}
		})
	}
}
