package types_test

import (
	"testing"

	"github.com/s-chernyavskiy/sakura/internal/sakura/types"
	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, len(tt.val), l.Length(), "List lengths mismatch")
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

			assert.Equal(t, len(tt.val), l.Length(), "List lengths mismatch")
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
