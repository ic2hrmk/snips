/*

	The package provides method to find all rectangles in a given points set

*/

package rectangles

import (
	"fmt"
	"math"
	"strconv"
)

type Point struct {
	X, Y float64
}

func (rcv *Point) IsEqual(p *Point) bool {
	if rcv == nil || p == nil {
		return false
	}

	return rcv.X == p.X && rcv.Y == p.Y
}

func (rcv *Point) String() string {
	if rcv == nil {
		return "[X=NaN;Y=NaN]"
	}

	return fmt.Sprintf("[X=%.2f;Y=%.2f]", rcv.X, rcv.Y)
}

type Rectangle struct {
	Points [4]*Point
}

func (rcv *Rectangle) IsEqual(r *Rectangle) bool {
	if rcv == nil || r == nil {
		return false
	}

	const edgesInRectangle = 4
	var sameEdgesCount = 0

	for i := range rcv.Points {
		for j := range r.Points {
			if rcv.Points[i].IsEqual(r.Points[j]) {
				sameEdgesCount++
				break
			}
		}
	}

	return sameEdgesCount == edgesInRectangle
}

func (rcv *Rectangle) String() string {
	return fmt.Sprintf("A=%s;B=%s;C=%s;D=%s",
		rcv.Points[0], rcv.Points[1],
		rcv.Points[2], rcv.Points[3],
	)
}

func FindRectangles(points []*Point) []*Rectangle {
	var rectangles []*Rectangle

	isKnownPoint := func(p *Point, l []*Point) bool {
		for i := range l {
			if p.IsEqual(l[i]) {
				return true
			}
		}
		return false
	}

	for i := range points {
		for j := range points {
			for k := range points {

				if i == j || i == k || j == k {
					continue
				}

				if !is90degreeAngle(points[i], points[j], points[k]) {
					continue
				}

				var (
					middlePoint = getMiddlePoint(points[j], points[k])

					dx = middlePoint.X - points[i].X
					dy = middlePoint.Y - points[i].Y

					calculatedPoint = &Point{
						X: middlePoint.X + dx,
						Y: middlePoint.Y + dy,
					}
				)

				if !isKnownPoint(calculatedPoint, points) {
					continue // The point doesn't exists
				}

				rectangles = append(rectangles,
					&Rectangle{Points: [4]*Point{
						points[i],
						points[j],
						calculatedPoint,
						points[k],
					}})
			}
		}
	}

	return getDistinctRectangles(rectangles)
}

// Returns a distance between two points in 2D
func getDistance(a, b *Point) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

// Returns an angle in degrees
func getAngle(corner, p1, p2 *Point) float64 {
	var (
		vectorA = &Point{X: corner.X - p1.X, Y: corner.Y - p1.Y}
		vectorB = &Point{X: corner.X - p2.X, Y: corner.Y - p2.Y}
	)

	angleRad := math.Acos(float64(vectorA.X*vectorB.X+vectorA.Y*vectorB.Y) /
		(getDistance(&Point{X: 0, Y: 0}, vectorA) * getDistance(&Point{X: 0, Y: 0}, vectorB)))

	return angleRad * (180.0 / math.Pi)
}

func getMiddlePoint(p1, p2 *Point) *Point {
	return &Point{
		X: (p1.X + p2.X) / 2,
		Y: (p1.Y + p2.Y) / 2,
	}
}

func getDistinctRectangles(in []*Rectangle) []*Rectangle {
	var distinct []*Rectangle

	isInList := func(r *Rectangle, l []*Rectangle) bool {
		for i := range l {
			if r.IsEqual(l[i]) {
				return true
			}
		}
		return false
	}

	for i := range in {
		if !isInList(in[i], distinct) {
			distinct = append(distinct, in[i])
		}
	}

	return distinct
}

func is90degreeAngle(corner, p1, p2 *Point) bool {
	return roundFloat(getAngle(corner, p1, p2), 0) == 90.0
}

func roundFloat(x float64, prec int) float64 {
	f, _ := strconv.ParseFloat(strconv.FormatFloat(x, 'g', prec, 64), 64)
	return f
}
