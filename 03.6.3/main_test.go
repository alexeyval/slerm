package main

import "testing"

func Test_funcName(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{s: []string{"a", "b", "c"}},
			want: true,
		}, {
			name: "2",
			args: args{s: []string{"a", "b", "c", "c"}},
			want: true,
		}, {
			name: "3",
			args: args{s: []string{"a", "b", "c", "b"}},
			want: false,
		}, {
			name: "4",
			args: args{s: []string{"d", "b", "c"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcName(tt.args.s); got != tt.want {
				t.Errorf("funcName() = %v, want %v", got, tt.want)
			}
		})
	}
}
