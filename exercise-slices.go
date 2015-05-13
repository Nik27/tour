package main

import "golang.org/x/tour/pic"
import "math"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		p[i][dx / 2] = 128		
	}
	for i := 0; i < dx; i++ {
		p[dy / 2][i] = 128	
	}
	r := float64(dx) / 2.0
	for i := 0; i < dx / 2; i++ {
		a := math.Pow(float64(r), 2.0 / 3.0)
		b := math.Pow(float64(i), 2.0 / 3.0)
		y := int(math.Pow(a - b, 3.0 / 2.0))
		v, w := dy / 2 + y, dx / 2 - i
		if v >= 0 && v < dy && w > 0 && w < dx {
			p[v][w] = 255
		}
		v, w = dy / 2 - y, dx / 2 - i
		if v >= 0 && v < dy && w > 0 && w < dx {
			p[v][w] = 255
		}
		v, w = dy / 2 + y, dx / 2 + i
		if v >= 0 && v < dy && w > 0 && w < dx {
			p[v][w] = 255
		}
		v, w = dy / 2 - y, dx / 2 + i
		if v >= 0 && v < dy && w > 0 && w < dx {
			p[v][w] = 255
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
