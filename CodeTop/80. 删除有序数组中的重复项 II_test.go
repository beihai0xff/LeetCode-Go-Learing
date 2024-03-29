package CodeTop

import "testing"

func Test_removeDuplicates80(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 1, 2, 2, 3}}, 5},
		{"2", args{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates80(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates80() = %v, want %v", got, tt.want)
			}
		})
	}
}
