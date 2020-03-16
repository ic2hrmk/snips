package rectangles

import (
	"reflect"
	"testing"
)

func TestFindRectangles(t *testing.T) {
	type args struct {
		points []*Point
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Simple rectangles",
			args: args{
				points: []*Point{
					{X: 1, Y: 2},
					{X: 1, Y: 6},
					{X: 7, Y: 2},
					{X: 9, Y: 4},
					{X: 5, Y: 4},
					{X: 7, Y: 6},
				}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindRectangles(tt.args.points)
			for i := range got {
				t.Log(got[i])
			}
		})
	}
}

func Test_getDistance(t *testing.T) {
	type args struct {
		a *Point
		b *Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Simple rectangles",
			args: args{
				a: &Point{X: 1, Y: 2},
				b: &Point{X: 1, Y: 6},
			},
			want: 4.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDistance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAngle(t *testing.T) {
	type args struct {
		corner *Point
		p1     *Point
		p2     *Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Simple rectangles",
			args: args{
				corner: &Point{X: 0, Y: 0},
				p1:     &Point{X: 1, Y: 3},
				p2:     &Point{X: 2, Y: 1},
			},
			want: 45.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			angle := getAngle(tt.args.corner, tt.args.p1, tt.args.p2)
			if got := roundFloat(angle, 2); got != tt.want {
				t.Errorf("getAngle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMiddlePoint(t *testing.T) {
	type args struct {
		p1 *Point
		p2 *Point
	}
	tests := []struct {
		name string
		args args
		want *Point
	}{
		{
			args: args{
				p1: &Point{X: 1, Y: 6},
				p2: &Point{X: 7, Y: 2},
			},
			want: &Point{X: 4, Y: 4},
		},
		{
			args: args{
				p1: &Point{X: -2, Y: 0},
				p2: &Point{X: 0, Y: 3},
			},
			want: &Point{X: -1, Y: 1.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMiddlePoint(tt.args.p1, tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMiddlePoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_IsEqual(t *testing.T) {
	type fields struct {
		Points [4]*Point
	}
	type args struct {
		r *Rectangle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				Points: [4]*Point{
					{X: 1, Y: 2},
					{X: 1, Y: 6},
					{X: 7, Y: 6},
					{X: 7, Y: 2},
				}},
			args: args{
				r: &Rectangle{Points: [4]*Point{
					{X: 1, Y: 6},
					{X: 7, Y: 6},
					{X: 7, Y: 2},
					{X: 1, Y: 2},
				}},
			},
			want: true,
		},
		{
			fields: fields{
				Points: [4]*Point{
					{X: 1, Y: 2},
					{X: 1, Y: 6},
					{X: 7, Y: 6},
					{X: 7, Y: 3},
				}},
			args: args{
				r: &Rectangle{Points: [4]*Point{
					{X: 1, Y: 6},
					{X: 7, Y: 6},
					{X: 7, Y: 2},
					{X: 1, Y: 2},
				}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rcv := &Rectangle{
				Points: tt.fields.Points,
			}
			if got := rcv.IsEqual(tt.args.r); got != tt.want {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
