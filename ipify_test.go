package pubip

import (
	"reflect"
	"testing"
)

//func TestGetIp(t *testing.T) {
//	originalApiUri := API_URI
//
//	_, err := GetIp()
//	if err != nil {
//		t.Error(err)
//	}
//
//	API_URI = "https://api.ipifyyyyyyyyyyyy.org"
//
//	_, err = GetIp()
//	if err == nil {
//		t.Error("Request to https://api.ipifyyyyyyyyyyyy.org should have failed, but succeeded.")
//	}
//
//	API_URI = originalApiUri
//}

func TestIsValidResultSet(t *testing.T) {
	tests := []struct {
		input    []string
		expected bool
	}{
		{[]string{"aaa", "aaa"}, true},
		{[]string{"aaa", "aaa", "aaa"}, true},
		{[]string{"aaa", "bbb"}, false},
		{[]string{"aaa", "aaa", "bbb"}, false},
		{[]string{"aaa"}, false},
		{[]string{}, false},
	}
	for i, v := range tests {
		actual := isValid(v.input)
		expected := v.expected
		t.Logf("Check case %d: %s(actual) == %s(expected)", i, actual, expected)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Error on case %d: %s(actual) != %s(expected)", i, actual, expected)
		}
	}
}
