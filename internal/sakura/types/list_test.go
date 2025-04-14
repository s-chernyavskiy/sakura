package types_test

import (
	"testing"

	"github.com/s-chernyavskiy/sakura/internal/sakura/types"
	"github.com/stretchr/testify/assert"
)

func TestList_Push(t *testing.T) {
	tests := []struct {
		name    string
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
			var list types.List
			gotErr := list.Push(tt.val...)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Push() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Push() succeeded unexpectedly")
			}
			assert.Equal(t, len(tt.val), list.Length(), "List lengths mismatch")
		})
	}
}
