package interview

import (
	"reflect"
	"testing"
)

func Test_findSubstringInWraproundString(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Changed requirements",
			args: args{p: "abc"},
			want: []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubstringInWraproundString(tt.args.p)

			t.Logf("want-> %v", tt.want)
			t.Logf("got -> %v", got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstringInWraproundString() = %v, want %v", got, tt.want)
			}
		})
	}
}
