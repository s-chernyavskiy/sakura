package testsuite

import (
	"reflect"
	"testing"
	"slices"
)

func AssertEqual(t *testing.T, expected any, given any) {
	if expected != given {
		t.Errorf("assertion failed, want %v : given %v", expected, given)
	}
}

func AssertStringSliceEqual(t *testing.T, expected []string, given []string) {
	if !reflect.DeepEqual(expected, given) {
		t.Errorf("slice assertion failed, want %v : given %v", expected, given)
	}
}

func AssertNil(t *testing.T, i any) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice:
		if !val.IsNil() {
			t.Errorf("wanted %v is not nil", i)
		}
	default:
		if i != nil {
			t.Errorf("wanted %v for non nil type but is not nil", i)
		}
	}
}

func AssertErrorEqual(t *testing.T, expected error, given error) {
	if expected.Error() != given.Error() {
		t.Errorf("error assertion failed, want %v : given %v", expected.Error(), given.Error())
	}
}

func AssertContainsAllElements(t *testing.T, expected []string, given []string) {
	var match bool
	for _, i := range expected {
		match = slices.Contains(given, i)

		if !match {
			t.Errorf("element %v [%T] not found in array %v [%T]", i, i, given, given)
		}
	}
}
