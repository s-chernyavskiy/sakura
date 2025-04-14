package types_test

import (
	"testing"

	"github.com/s-chernyavskiy/sakura/internal/sakura/types"
	"github.com/s-chernyavskiy/sakura/pkg/testsuite"
)

func TestList_PushFront(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		val     []string
		wantErr bool
	}{
		{
			name:    "generic add",
			val:     []string{"foo", "bar", "baz"},
			wantErr: false,
		},
		{
			name:    "no values",
			val:     []string{},
			wantErr: true,
		},
		{
			name:    "single value",
			val:     []string{"single"},
			wantErr: false,
		},
		{
			name:    "special characters",
			val:     []string{"!@#$%^&*()", "hello", "world"},
			wantErr: false,
		},
		{
			name:    "long strings",
			val:     []string{string(make([]byte, 1000))},
			wantErr: false,
		},
		{
			name:    "duplicate values",
			val:     []string{"duplicate", "duplicate"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l types.List
			gotErr := l.PushFront(tt.val...)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("PushFront() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("PushFront() succeeded unexpectedly")
			}
			testsuite.AssertEqual(t, len(tt.val), l.Length())
		})
	}
}

func TestList_PushBack(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		val     []string
		wantErr bool
	}{
		{
			name:    "generic add",
			val:     []string{"foo", "bar", "baz"},
			wantErr: false,
		},
		{
			name:    "no values",
			val:     []string{},
			wantErr: true,
		},
		{
			name:    "single value",
			val:     []string{"single"},
			wantErr: false,
		},
		{
			name:    "special characters",
			val:     []string{"!@#$%^&*()", "hello", "world"},
			wantErr: false,
		},
		{
			name:    "long strings",
			val:     []string{string(make([]byte, 1000))},
			wantErr: false,
		},
		{
			name:    "duplicate values",
			val:     []string{"duplicate", "duplicate"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l types.List
			gotErr := l.PushBack(tt.val...)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("PushBack() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("PushBack() succeeded unexpectedly")
			}

			testsuite.AssertEqual(t, len(tt.val), l.Length())
		})
	}
}

func TestList_Head(t *testing.T) {
	tests := []struct {
		name        string // description of this test case
		want        string
		toAdd       bool
		valuesToAdd []string
	}{
		{
			name:        "generic list",
			want:        "1",
			toAdd:       true,
			valuesToAdd: []string{"1", "2", "3"},
		},
		{
			name:        "nil head",
			want:        "",
			toAdd:       false,
			valuesToAdd: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l types.List
			if tt.toAdd {
				l.PushBack(tt.valuesToAdd...)
			}

			got := l.Head()
			if got == nil && tt.toAdd {
				t.Errorf("Head() = <nil>, want %v", tt.want)
			}

			if got == nil {
				return
			}

			if got.Value != tt.want {
				t.Errorf("Head().Value = %v, want %v", got.Value, tt.want)
			}
		})
	}
}

func TestList_Tail(t *testing.T) {
	tests := []struct {
		name        string // description of this test case
		want        string
		toAdd       bool
		valuesToAdd []string
	}{
		{
			name:        "generic list",
			want:        "3",
			toAdd:       true,
			valuesToAdd: []string{"1", "2", "3"},
		},
		{
			name:        "nil tail",
			want:        "",
			toAdd:       false,
			valuesToAdd: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l types.List
			if tt.toAdd {
				l.PushBack(tt.valuesToAdd...)
			}

			got := l.Tail()
			if got == nil && tt.toAdd {
				t.Errorf("Tail() = <nil>, want %v", tt.want)
			}

			if got == nil {
				return
			}

			if got.Value != tt.want {
				t.Errorf("Tail().Value = %v, want %v", got.Value, tt.want)
			}
		})
	}
}

func TestList_PopBack(t *testing.T) {
	tests := []struct {
		name   string
		want   string
		values []string
	}{
		{
			name:   "generic pop back",
			want:   "3",
			values: []string{"1", "2", "3"},
		},
		{
			name:   "null value pop back",
			want:   "",
			values: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l types.List
			l.PushBack(tt.values...)

			got := l.PopBack()
			if got != tt.want {
				t.Errorf("PopBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_PopFront(t *testing.T) {
	tests := []struct {
		name   string // description of this test case
		want   any
		values []string
	}{
		{
			name:   "generic pop front",
			want:   "1",
			values: []string{"1", "2", "3"},
		},
		{
			name:   "null value pop front",
			want:   "",
			values: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l types.List
			l.PushBack(tt.values...)

			got := l.PopFront()
			if got != tt.want {
				t.Errorf("PopFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Range(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		start  int
		stop   int
		want   []string
		values []string
	}{
		{
			name:   "generic range",
			start:  2,
			stop:   3,
			want:   []string{"3", "4"},
			values: []string{"1", "2", "3", "4", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l types.List
			l.PushBack(tt.values...)

			got := l.Range(tt.start, tt.stop)
			testsuite.AssertStringSliceEqual(t, tt.want, got)
		})
	}
}
