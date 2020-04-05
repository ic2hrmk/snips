package staircase

import "testing"

func Test_staircase(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Of 4",
			args: args{n: 4},
		}, {
			name: "Of 6",
			args: args{n: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			staircase(tt.args.n)
		})
	}
}
