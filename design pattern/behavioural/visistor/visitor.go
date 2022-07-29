package visitor

import "fmt"

type IVistor interface {
	VisitForSquare(Square) string
	VisitForCircle(Circle) string
	VisitForTriangle(Triangle) string
}

type Area struct{}

func (a Area) VisitForSquare(s Square) string {
	return "Calculate area of square"
}

func (a Area) VisitForCircle(c Circle) string {
	return "Calculate area of circle"
}

func (a Area) VisitForTriangle(t Triangle) string {
	return "Calculate area of triangle"
}

type Shape interface {
	Accept(IVistor)
}

type Square struct {
	Width int
}

func (s Square) Accept(v IVistor) {
	fmt.Println(v.VisitForSquare(s))
}

type Triangle struct {
	Width int
}

func (t Triangle) Accept(v IVistor) {
	fmt.Println(v.VisitForTriangle(t))
}

type Circle struct {
	R int
}

func (c Circle) Accept(v IVistor) {
	fmt.Println(v.VisitForCircle(c))
}

// func main() {
// 	a := visitor.Area{}

// 	s := visitor.Square{2}
// 	c := visitor.Circle{2}
// 	t := visitor.Triangle{2}

// 	s.Accept(a)
// 	c.Accept(a)
// 	t.Accept(a)

// }
