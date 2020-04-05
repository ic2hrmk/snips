package staircase

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	const (
		stair = `#`
		space = ` `
	)

	for i := int32(0); i < n; i++ {
		fmt.Println(strings.Repeat(space, int(n-i-1)) + strings.Repeat(stair, int(i+1)))
	}
}
