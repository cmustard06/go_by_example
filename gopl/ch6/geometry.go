package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

//传统函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{2, 4}
	q := Point{3, 5}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	path := Path{Point{2, 4}, Point{3, 5}}
	fmt.Println(path.Distance())

}
