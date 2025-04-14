package testsuite_test

import (
	"errors"
	"testing"

	"github.com/s-chernyavskiy/sakura/pkg/testsuite"
)

func TestAssertEqual(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		expected any
		given    any
		wantErr  bool
	}{
		{
			name:     "string equal",
			expected: "foo",
			given:    "foo",
			wantErr:  false,
		},
		{
			name:     "number equal",
			expected: 1,
			given:    1,
			wantErr:  false,
		},
		{
			name:     "nil equal",
			expected: nil,
			given:    nil,
			wantErr:  false,
		},
		{
			name:     "string not equal",
			expected: "foo",
			given:    "bar",
			wantErr:  true,
		},
		{
			name:     "number not equal",
			expected: 1,
			given:    2,
			wantErr:  true,
		},
		{
			name:     "nil not equal",
			expected: nil,
			given:    "not nil",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			testsuite.AssertEqual(test, tt.expected, tt.given)

			if !tt.wantErr && test.Failed() {
				t.Errorf("AssertEqual() failed when it should have passed for %v == %v", tt.expected, tt.given)
			} else if tt.wantErr && !test.Failed() {
				t.Errorf("AssertEqual() passed when it should have failed for %v != %v", tt.expected, tt.given)
			}
		})
	}
}

func TestAssertStringSliceEqual(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		expected []string
		given    []string
		wantErr  bool
	}{
		{
			name:     "equal",
			expected: []string{"1", "2", "3"},
			given:    []string{"1", "2", "3"},
			wantErr:  false,
		},
		{
			name:     "not equal",
			expected: []string{"1", "2", "3"},
			given:    []string{"1", "2", "4"},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			testsuite.AssertStringSliceEqual(test, tt.expected, tt.given)

			if !tt.wantErr && test.Failed() {
				t.Errorf("AssertStringSliceEqual() failed when it should have passed for %v == %v", tt.expected, tt.given)
			} else if tt.wantErr && !test.Failed() {
				t.Errorf("AssertStringSliceEqual() passed when it should have failed for %v != %v", tt.expected, tt.given)
			}
		})
	}
}

func TestAssertNil(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		i       any
		wantErr bool
	}{
		{
			name:    "nil value",
			i:       nil,
			wantErr: false,
		},
		{
			name:    "nil pointer",
			i:       (*int)(nil),
			wantErr: false,
		},
		{
			name:    "nil slice",
			i:       []int(nil),
			wantErr: false,
		},
		{
			name:    "nil map",
			i:       map[string]int(nil),
			wantErr: false,
		},
		{
			name:    "nil channel",
			i:       (chan int)(nil),
			wantErr: false,
		},
		{
			name:    "nil function",
			i:       (func())(nil),
			wantErr: false,
		},
		{
			name:    "non-nil pointer",
			i:       new(int),
			wantErr: true,
		},
		{
			name:    "non-nil slice",
			i:       []int{},
			wantErr: true,
		},
		{
			name:    "non-nil map",
			i:       make(map[string]int),
			wantErr: true,
		},
		{
			name:    "non-nil channel",
			i:       make(chan int),
			wantErr: true,
		},
		{
			name:    "non-nil function",
			i:       func() {},
			wantErr: true,
		},
		{
			name:    "non-pointer value",
			i:       42,
			wantErr: true,
		},
		{
			name:    "empty string",
			i:       "",
			wantErr: true,
		},
		{
			name:    "zero struct",
			i:       struct{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			testsuite.AssertNil(test, tt.i)

			if !tt.wantErr && test.Failed() {
				t.Errorf("AssertNil() failed when it should have passed for %v", tt.i)
			} else if tt.wantErr && !test.Failed() {
				t.Errorf("AssertNil() passed when it should have failed for %v", tt.i)
			}
		})
	}
}

func TestAssertErrorEqual(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		expected error
		given    error
		wantErr  bool
	}{
		{
			name:     "generic error",
			expected: errors.New("foo"),
			given:    errors.New("foo"),
			wantErr:  false,
		},
		{
			name:     "generic error fail",
			expected: errors.New("foo"),
			given:    errors.New("bar"),
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			testsuite.AssertErrorEqual(test, tt.expected, tt.given)

			if !tt.wantErr && test.Failed() {
				t.Errorf("AssertErrorEqual() failed when it should have passed for %v == %v", tt.expected, tt.given)
			} else if tt.wantErr && !test.Failed() {
				t.Errorf("AssertErrorEqual() passed when it should have failed for %v != %v", tt.expected, tt.given)
			}
		})
	}
}

func TestAssertContainsAllElements(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		expected []string
		given    []string
		wantErr  bool
	}{
		{
			name:     "generic slice",
			expected: []string{"1", "2", "3"},
			given:    []string{"1", "2", "3"},
			wantErr:  false,
		},
		{
			name:     "generic slice failure",
			expected: []string{"1", "2", "3"},
			given:    []string{"1", "2", "4"},
			wantErr:  true,
		},
		{
			name:     "empty lists",
			expected: []string{},
			given:    []string{},
			wantErr:  false,
		},
		{
			name:     "empty and non empty lists",
			expected: []string{},
			given:    []string{"1"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			testsuite.AssertContainsAllElements(test, tt.expected, tt.given)

			if !tt.wantErr && test.Failed() {
				t.Errorf("AssertContainsAllElements() failed when it should have passed for %v == %v", tt.expected, tt.given)
			} else if tt.wantErr && !test.Failed() {
				t.Errorf("AssertContainsAllElements() passed when it should have failed for %v != %v", tt.expected, tt.given)
			}
		})
	}
}
