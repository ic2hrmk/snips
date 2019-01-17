package main

import (
	"fmt"
	"strings"
)

const (
	water = cell(0)
	land  = cell(1)
)

type cell int

func (c cell) isLand() bool {
	return c == land
}

func (c cell) isWater() bool {
	return c == water
}

type cellCoordinate struct {
	x int
	y int
}

type island struct {
	coordinates []*cellCoordinate
}

var testField1 = [][]cell{
	{1, 1, 1, 1, 0},
	{1, 1, 0, 1, 0},
	{1, 1, 0, 0, 0},
	{0, 0, 0, 0, 0},
}

var testField2 = [][]cell{
	{1, 1, 1, 1, 0},
	{1, 1, 0, 1, 0},
	{1, 1, 0, 0, 0},
	{0, 0, 1, 1, 1},
}

func main() {
	islands := detectIslands(testField2)
	fmt.Printf("There are [%d] islands\n\n", len(islands))

	for i := range islands {
		printIsland(len(testField1), len(testField1[0]), islands[i])
	}

}

func detectIslands(field [][]cell) []*island {
	var islands []*island

	if field == nil || len(field) == 0 {
		return islands
	}

	//
	// Create a copy of input
	//
	copiedField := make([][]cell, len(field))
	for i := range field {
		copiedField[i] = make([]cell, len(field[i]))
		copy(copiedField[i], field[i])
	}

	//
	// Execute look for each cell
	//
	for i := range copiedField {
		for j := range copiedField[i] {
			// Skip water cells
			if copiedField[i][j].isWater() {
				continue
			}

			currentIsland := &island{}
			walkIslandFromPoint(i, j, copiedField, currentIsland)

			// Try to discover an island
			islands = append(islands, currentIsland)
		}
	}

	return islands
}

func walkIslandFromPoint(i, j int, field [][]cell, isl *island) {
	// Mark current cell as visited
	field[i][j] = water
	isl.coordinates = append(isl.coordinates, &cellCoordinate{
		x: i,
		y: j,
	})

	//
	// Try left
	//
	if j-1 >= 0 && field[i][j-1].isLand() {
		walkIslandFromPoint(i, j-1, field, isl)
	}

	//
	// Try right
	//
	if j+1 < len(field[i]) && field[i][j+1].isLand()  {
		walkIslandFromPoint(i, j+1, field, isl)
	}

	//
	// Try up
	//
	if i-1 >= 0 && field[i-1][j].isLand() {
		walkIslandFromPoint(i-1, j, field, isl)
	}

	//
	// Try bottom
	//
	if i+1 < len(field) && field[i+1][j].isLand() {
		walkIslandFromPoint(i+1, j, field, isl)
	}
}

func printField(field [][]cell) {
	const (
		regularCellPattern   = "  %s  "

		landCellPattern   = "*"
		watterCellPattern = "~"
	)

	cellPattern := ""

	for i := range field {
		for j := range field[i] {
			if field[i][j].isWater() {
				cellPattern = watterCellPattern
			} else {
				cellPattern = landCellPattern
			}

			fmt.Printf(regularCellPattern, cellPattern)
		}
		fmt.Println()
	}
}

func printIsland(xFieldSize, yFieldSize int, isl *island) {
	var line = strings.Repeat("=", xFieldSize+1)

	fmt.Println(line)
	defer fmt.Println(line)

	field := make([][]cell, xFieldSize)
	for i := range field {
		field[i] = make([]cell, yFieldSize)
		for j := range field[i] {
			field[i][j] = water
		}
	}

	for _ , c := range isl.coordinates {
		field[c.x][c.y] = land
	}

	printField(field)
}
