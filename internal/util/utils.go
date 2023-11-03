package util

import "math/rand"

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func RollDice(number, mod int) int {
	total := 0

	for i := 0; i < number; i++ {
		total += rand.Intn(6) + 1
	}

	return total + mod
}
