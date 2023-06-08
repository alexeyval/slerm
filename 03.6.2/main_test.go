package main

import "testing"

func Test_funcName(t *testing.T) {
	type args struct {
		nums [9]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{nums: [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: false,
		},
		{
			name: "2",
			args: args{nums: [9]int{1, 2, 3, 4, 5, 6, 7, 9, 9}},
			want: true,
		},
		{
			name: "3",
			args: args{nums: [9]int{9, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: true,
		},
		{
			name: "4",
			args: args{nums: [9]int{1, 2, 3, 4, 5, 3, 7, 8, 9}},
			want: true,
		},
		{
			name: "5",
			args: args{nums: [9]int{1, 2, 3, 4, 5, 6, 7, 8, 1}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcName(tt.args.nums); got != tt.want {
				t.Errorf("funcName() = %v, want %v", got, tt.want)
			}
		})
	}
}
