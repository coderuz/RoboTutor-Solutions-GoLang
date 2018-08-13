package main

import (
	"golang.org/x/tour/pic"
	"math/rand"
)

func Pic(dx, dy int) [][]uint8 {
	rows := make([][]uint8, dy)
	for i := range rows{
		rows[i] =  make([]uint8, dx)
		for ith := range rows[i] {
			rows[i][ith] = uint8(rand.Intn(255))
		}
	}
	return rows
}

func main() {
	pic.Show(Pic)
}
