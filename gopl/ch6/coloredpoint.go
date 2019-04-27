package main

import (
	"fmt"
	"image/color"
	"math"
)

//嵌入式结构体的使用

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	dx := q.X - p.X
	dy := q.Y - p.Y
	return math.Sqrt(dx*dx + dy*dy)
}

//如果需要对结构体的数据做更改，需要设置为指针接收者，否则修改无法全局生效,作用域为函数内部
//func (p Point) ScaleBy(factor float64){
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
	fmt.Println("func internal x:", p.X)
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) //5

	p.ScaleBy(2) // x 2,  2
	//fmt.Println(p.X,p.Y)
	q.ScaleBy(2) // 10,8
	//fmt.Println(q.X,q.Y)
	fmt.Println(p.Distance(q.Point)) // 10

}

func init() {
	//初始化函数，在main函数之前运行或者导包自动运行
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)

	fmt.Println("*********************")
}

func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point))
	q.Point = p.Point //q p 现在共享相同的指针
	p.ScaleBy(2)
	fmt.Println(*p.Point, q.Point)
	fmt.Println("------------------")
}
