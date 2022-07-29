package bridge

import "fmt"

type Shape interface {
	ShowShape() string
	ShowInfo()
}

type Square struct {
	C Color
}

func (s Square) ShowShape() string {
	return "Square"
}

func (s Square) ShowInfo() {
	fmt.Println("This is " + s.ShowShape() + " with " + s.C.ShowColor() + " color")
}

type Rectangular struct {
	C Color
}

func (r Rectangular) ShowShape() string {
	return "Rectangular"
}

func (r Rectangular) ShowInfo() {
	fmt.Println("This is " + r.ShowShape() + " with " + r.C.ShowColor() + " color")
}

type Color interface {
	ShowColor() string
}

type Red struct{}

func (r Red) ShowColor() string {
	return "Red"
}

type Yellow struct{}

func (y Yellow) ShowColor() string {
	return "Yellow"
}

// func main() {

// 	red := bridge.Red{}
// 	yellow := bridge.Yellow{}
// 	r := bridge.Rectangular{red}
// 	s := bridge.Square{yellow}

// 	r.ShowInfo()
// 	s.ShowInfo()
// }
