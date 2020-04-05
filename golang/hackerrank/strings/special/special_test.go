package special

import "testing"

func Test_substrCount(t *testing.T) {
	type args struct {
		n int32
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example from site",
			args:args{
				n: 8,
				s: "mnonopoo",
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substrCount(tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("substrCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
