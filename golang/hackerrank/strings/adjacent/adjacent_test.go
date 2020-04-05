package adjacent

import "testing"

func Test_alternatingCharacters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Example from site",
			args: args{
				s: "AABBABABABAB",
			},
			want: 2,
		},
		{
			name: "Example from site 1",
			args: args{
				s: "AAAA",
			},
			want: 3,
		},
		{
			name: "Example from site 2",
			args: args{
				s: "BBBBB",
			},
			want: 4,
		},
		{
			name: "Example from site 3",
			args: args{
				s: "ABABABAB",
			},
			want: 0,
		},
		{
			name: "Example from site 4",
			args: args{
				s: "BABABA",
			},
			want: 0,
		},
		{
			name: "Example from site 5",
			args: args{
				s: "AAABBB",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternatingCharacters(tt.args.s); got != tt.want {
				t.Errorf("alternatingCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
