package main

import (
	"fmt"
)

// fungsi untuk mendapatkan peringkat berdasarkan skor yang sudah terurut
func denseRanking(scores []int, gitsScores []int) []int {
	// menyimpan peringkat untuk setiap skor yang diperoleh GITS
	rankings := []int{}
	
	// menyaring skor duplikat untuk ranking yang tepat
	uniqueScores := []int{}
	for i := 0; i < len(scores); i++ {
		if len(uniqueScores) == 0 || scores[i] != uniqueScores[len(uniqueScores)-1] {
			uniqueScores = append(uniqueScores, scores[i])
		}
	}

	// iterasi untuk setiap skor GITS
	for _, gitsScore := range gitsScores {
		// temukan peringkat
		rank := 1
		for i := 0; i < len(uniqueScores); i++ {
			if gitsScore < uniqueScores[i] {
				rank++
			} else {
				break
			}
		}
		rankings = append(rankings, rank)
	}

	return rankings
}

func main() {
	// input
	fmt.Print("Masukkan input:\n")
	var n int
	fmt.Scan(&n)
	
	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}
	
	var m int
	fmt.Scan(&m)
	
	gitsScores := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&gitsScores[i])
	}
	
	// mendapatkan peringkat
	result := denseRanking(scores, gitsScores)
	
	// output
	fmt.Print("\nOutput: ")
	for i, rank := range result {
		if i == len(result)-1 {
			fmt.Println(rank)
		} else {
			fmt.Print(rank, " ")
		}
	}
}
