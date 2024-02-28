package util

import (
	"fmt"
	"math/rand"
	"strings"
)

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

func HexDigitToInt(digit string) (int, error) {
	// Make sure we're working with uppercase for simplicity
	digit = strings.ToUpper(digit)
	if len(digit) != 1 {
		return 0, fmt.Errorf("input must be a single character")
	}

	r := rune(digit[0])
	var value int
	if '0' <= r && r <= '9' {
		value = int(r - '0')
	} else if 'A' <= r && r <= 'Z' {
		// 'A' should correspond to 10, so we subtract 'A' and add 10
		value = int(r-'A') + 10
	} else {
		return 0, fmt.Errorf("invalid character: %s", digit)
	}

	return value, nil
}

func CleanUwp(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		n := rune(s[i])
		if (n >= '0' && n <= '9') || (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'Z') {
			result += string(n)
		}
	}

	return result
}
