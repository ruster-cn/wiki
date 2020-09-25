package singleton

import (
	"testing"
)

func TestGetCache(t *testing.T) {
	tests := []struct {
		name string
		want *cache
	}{
		{
			name: "equal",
			want: GetCache(),
		},
		{
			name: "not equal",
			want: &cache{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCache(); got != tt.want {
				t.Errorf("GetCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
