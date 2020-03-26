package main

import (
	"golang.org/x/tour/pic"
	"math/rand"
)

func Pic(dx, dy int) [][]uint8 {
	var picMat [][]uint8
	picMat = make([][]uint8, dy, dy)
	for i := 0; i < dy; i++ {
		picMat[i] = make([]uint8, dx, dx)
		for j := 0; j < dx; j++ {
			picMat[i][j] = uint8(rand.Intn(int(255)))
		}
	}
	return picMat
}

func main() {
	pic.Show(Pic)
}


