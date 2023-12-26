package misezan

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Calc[T Number](a T, b T) float64 {
	var pattern = map[byte]byte{
		'6': '9',
		'9': '6',
		'2': '5',
		'5': '2',
	}

	if a == b {
		return 0
	}

	var min, max T

	if a > b {
		min, max = b, a
	} else {
		min, max = a, b
	}

	if int64(max) >= int64(min)*100 {
		return float64(max - min*17)
	}

	minString := fmt.Sprint(min)
	maxString := fmt.Sprint(max)

	if len(minString) != len(maxString) {
		return float64(max)
	}

	isReverse := true
	isSpecial := false
	for i := range minString {
		if maxStringAns, ok := pattern[minString[i]]; !ok || maxStringAns != maxString[len(maxString)-i-1] {
			isReverse = false
			break
		}
		if minString[i] == '2' || minString[i] == '5' {
			isSpecial = true
		}
	}
	if isReverse {
		if isSpecial {
			return 1.1
		} else {
			return 11
		}
	}

	return float64(max)
}
