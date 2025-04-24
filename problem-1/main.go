package main

import (
	"fmt"
)

// fungsi untuk menghasilkan urutan A000124 sesuai dengan input
func generateA000124(n int) []int {
	sequence := make([]int, n)
	sequence[0] = 1 // nilai pertama adalah 1
	for i := 1; i < n; i++ {
		sequence[i] = sequence[i-1] + i
	}
	return sequence
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	// menghasilkan urutan A000124
	sequence := generateA000124(n)

	for i, v := range sequence {
		if i == len(sequence)-1 {
			fmt.Printf("%d\n", v)
		} else {
			fmt.Printf("%d-", v)
		}
	}
}
