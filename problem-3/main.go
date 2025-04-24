package main

import (
	"fmt"
)

// fungsi untuk memeriksa apakah string memiliki bracket yang seimbang
func isBalanced(s string) string {
	stack := []rune{}

	// iterasi setiap karakter dalam string
	for _, char := range s {
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else if char == '}' || char == ']' || char == ')' {
			if len(stack) == 0 {
				return "NO"
			}

			// periksa apakah bracket penutup cocok dengan pembuka
			top := stack[len(stack)-1]
			if (char == '}' && top == '{') || (char == ']' && top == '[') || (char == ')' && top == '(') {
				stack = stack[:len(stack)-1] // hapus bracket pembuka yang sudah dicocokkan
			} else {
				return "NO"
			}
		}
	}

	if len(stack) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	fmt.Print("Masukkan input:\n")
	var input string
	fmt.Scan(&input)

	result := isBalanced(input)
	fmt.Print("\nOutput: ", result, "\n")
}
