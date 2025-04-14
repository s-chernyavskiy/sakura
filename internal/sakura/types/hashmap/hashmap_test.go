package hashmap_test

import (
	"testing"

	"github.com/s-chernyavskiy/sakura/internal/sakura/types/hashmap"
	"github.com/s-chernyavskiy/sakura/pkg/testsuite"
)

func TestHashMap_Set(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		key   string
		value string
	}{
		{
			name:  "add",
			key:   "foo",
			value: "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := hashmap.NewHashMap()
			m.Set(tt.key, tt.value)

			testsuite.AssertEqual(t, m.Exists(tt.key), true)
		})
	}
}

func TestHashMap_Get(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		key  string
		want string
	}{
		{
			name: "generic get",
			key:  "foo",
			want: "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := hashmap.NewHashMap()
			m.Set("foo", "bar")

			got := m.Get(tt.key)

			if tt.want != got {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Delete(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		keys []string
		want bool
	}{
		{
			name: "generic delete",
			keys: []string{"foo", "bar"},
			want: true,
		},
		{
			name: "generic single delete",
			keys: []string{"foo"},
			want: true,
		},
		{
			name: "failed to find single",
			keys: []string{"baz"},
			want: false,
		},
		{
			name: "failed to find multiple",
			keys: []string{"baz", "quo"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := hashmap.NewHashMap()
			m.Set("foo", "a")
			m.Set("bar", "a")

			got := m.Delete(tt.keys...)

			if tt.want != got {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Exists(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		key  string
		want bool
	}{
		{
			name: "generic exists",
			key:  "foo",
			want: true,
		},
		{
			name: "generic doesnt exist",
			key:  "bar",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := hashmap.NewHashMap()
			m.Set("foo", "a")

			got := m.Exists(tt.key)
			if tt.want != got {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
