package util

import "math/rand"

func GenerateRandonNumber(min, max int) int {
	return rand.Intn(max - min) + min
}
