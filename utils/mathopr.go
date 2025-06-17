package utils

import (
	"fmt"
)

func Multiple(a, b int) int {
	return a * b
}

func Factorial(a int) int {
	if a == 0 || a == 1 {
		return a
	}

	return a * Factorial(a-1)
}

func IterateStringArray(str []string) {
	for i := 0; i < len(str); i++ {
		fmt.Print(str[i])
	}
	for i, s := range str {
		fmt.Print(i, " Rangeable ", s)
	}

}
func IterateMap(str []string) {
	for i := 0; i < len(str); i++ {
		fmt.Print(str[i])
	}
	for i, s := range str {
		fmt.Print(i, " Rangeable ", s)
	}

}
