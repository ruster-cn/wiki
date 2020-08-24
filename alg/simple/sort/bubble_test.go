package sort

import (
	"testing"
)

func Test_bubble(t *testing.T) {
	type args struct {
		array   []interface{}
		compare Compare
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "bubble",
			args: args{
				array: []interface{}{4, 2, 6, 4, 7, 1, 8, 0},
				compare: func(a, b interface{}) bool {
					a1, _ := a.(int)
					b1, _ := b.(int)
					return a1 > b1
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubble(tt.args.array, tt.args.compare)
			t.Logf("%v", tt.args.array)
		})
	}
}

func Test_quick(t *testing.T) {
	type args struct {
		array   []interface{}
		compare Compare
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "quick",
			args: args{
				array: []interface{}{4, 2, 6, 4, 7, 1, 8, 0},
				compare: func(a, b interface{}) bool {
					a1, _ := a.(int)
					b1, _ := b.(int)
					return a1 > b1
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quick(tt.args.array, tt.args.compare)
			t.Logf("%v", tt.args.array)
		})
	}
}
