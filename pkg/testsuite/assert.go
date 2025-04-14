package testsuite

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, expected any, given any) {
	if expected != given {
		t.Errorf("assertion failed, want %T : given %T", expected, given)
	}
}

func AssertStringSliceEqual(t *testing.T, expected []string, given []string) {
	if !reflect.DeepEqual(expected, given) {
		t.Errorf("array assertion failed, want %v : given %v", expected, given)
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
