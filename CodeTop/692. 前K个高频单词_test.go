package CodeTop

import (
	"reflect"
	"testing"
)

func Test_topKFrequent692(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2}, []string{"i", "love"}},
		{"2", args{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4}, []string{"the", "is", "sunny", "day"}},
		{"3", args{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 0}, []string{}},
		{"4", args{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 1}, []string{"the"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent692(tt.args.words, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent692() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_topKFrequent6922(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2}, []string{"i", "love"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent6922(tt.args.words, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent692() = %v, want %v", got, tt.want)
			}
		})
	}
}
