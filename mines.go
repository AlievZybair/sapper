package main

import "math/rand"

const MINE = -1

func calculate(a, b int) int {
	kol := (a * b - (a * b) % 7) / 7
	return kol
}

func arrange(a, b, kol int, field *[][]int) {
	for i := 0; i < kol; i++ {
		(*field)[rand.Intn(a)][rand.Intn(b)] = MINE
	}
}
