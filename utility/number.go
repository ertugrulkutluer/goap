package utility

import "math/rand"

func RandomInt(low, hi int) int {
	return low + rand.Intn(hi-low)
}
