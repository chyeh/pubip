package pubip

import (
	"reflect"
	"testing"
)

func TestIsValidate(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{"aaa"}, ""},
		{[]string{"aaa", "aaa"}, ""},
		{[]string{"aaa", "aaa", "aaa"}, "aaa"},
		{[]string{"aaa", "bbb"}, ""},
		{[]string{"aaa", "aaa", "bbb"}, ""},
	}
	for i, v := range tests {
		actual, _ := validate(v.input)
		expected := v.expected
		t.Logf("Check case %d: %s(actual) == %s(expected)", i, actual, expected)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Error on case %d: %s(actual) != %s(expected)", i, actual, expected)
		}
	}
}
