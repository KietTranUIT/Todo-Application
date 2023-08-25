package common

import (
	"fmt"
	"math/rand"
)

func GenerateRandomCode() string {
	numRands := make([]int, 6)

	for i := 0; i < 6; i++ {
		numRands[i] = rand.Intn(10)
	}

	code := fmt.Sprintf("%d%d%d%d%d%d", numRands[0], numRands[1], numRands[2], numRands[3], numRands[4], numRands[5])
	return code
}
