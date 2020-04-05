package anagrams

import "testing"

func Test_makeAnagram(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Example from site",
			args: args{
				a: "abcde",
				b: "cdefg",
			},
			want: 4,
		},
		{
			name: "Example from site 2",
			args: args{
				a: "abc",
				b: "abc",
			},
			want: 0,
		},
		{
			name: "Example from site 4",
			args: args{
				a: "abc",
				b: "def",
			},
			want: 6,
		},
		{
			name: "Example from site 3",
			args: args{
				a: "abcdhh",
				b: "cdh",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeAnagram(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("makeAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
