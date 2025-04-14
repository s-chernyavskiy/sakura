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
	if !reflect.ValueOf(i).IsNil() {
		t.Errorf("wanted %v is not nil", i)
	}
}
