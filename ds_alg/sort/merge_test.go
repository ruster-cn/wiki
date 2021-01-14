package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test one",
			args: args{
				array: []int{3, 6, 4, 7, 9, 1, 0},
			},
			want: []int{0, 1, 3, 4, 6, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
