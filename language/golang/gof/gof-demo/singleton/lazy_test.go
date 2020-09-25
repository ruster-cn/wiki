package singleton

import (
	"testing"
)

func TestGetLazyCache(t *testing.T) {
	tests := []struct {
		name string
		want *lazyCache
	}{
		{
			name: "equal",
			want: GetLazyCache(),
		},
		{
			name: "not equal",
			want: &lazyCache{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLazyCache(); got != tt.want {
				t.Errorf("GetLazyCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
