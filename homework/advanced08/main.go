package main

import "fmt"

type Shape interface {
}

func Area(s *Shape) {

}

func Perimeter(s *Shape) {

}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	r := Rectangle{width: 10, height: 5}
	c := Circle{radius: 5}

	fmt.Printf("长方体的周长：%.2f\n", r.Perimeter())

	fmt.Printf("园的周长：%.2f\n", c.Perimeter())
	fmt.Printf("长方体的面积：%.2f\n", r.Area())
	fmt.Printf("园的面积：%.2f\n", c.Area())
}
