package json

import (
	"testing"
)

func TestSort(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: `{"c":3,"b":2,"a":1}`, want: `{"a":1,"b":2,"c":3}`},
		{input: `{"c":3,"b":{"z":2,"f":1},"a":1}`, want: `{"a":1,"b":{"f":1,"z":2},"c":3}`},
		{input: `{"a":3,"2":2,"1":1}`, want: `{"1":1,"2":2,"a":3}`},
	}

	for _, tc := range tests {
		got, err := Sort(tc.input)
		if err != nil {
			t.Fatalf("error occured: %v", err)
		}
		if tc.want != string(got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
