package builder

import "fmt"

type Burger struct {
	Size string
}

type Coca struct {
	Size string
}

type FrenchFries struct {
	Size string
}

type Order struct {
	B Burger
	C Coca
	F FrenchFries
}

func (o Order) Invoice() string {
	invoice := ""
	if o.B != (Burger{}) {
		invoice += fmt.Sprintf("burger size: %s", o.B.Size)
	}
	if o.C != (Coca{}) {
		invoice += fmt.Sprintf(" - Cocal size: %s", o.C.Size)
	}

	if o.F != (FrenchFries{}) {
		invoice += fmt.Sprintf(" - FrenchFries size: %s", o.F.Size)
	}
	return invoice
}

func NewOrder(b Burger, c Coca, f FrenchFries) Order {
	return Order{b, c, f}
}

type IBuilder interface {
	OrderBurger(string) Burger
	OrderCoca(string) Coca
	OrderFriendFries(string) FrenchFries
	GetOrder() Order
}

type FastFoodOrderBuilder struct {
	B Burger
	C Coca
	F FrenchFries
}

func (f *FastFoodOrderBuilder) OrderBurger(size string) {
	f.B = Burger{size}
}

func (f *FastFoodOrderBuilder) OrderCoca(size string) {
	f.C = Coca{size}
}

func (f *FastFoodOrderBuilder) OrderFriendFries(size string) {
	f.F = FrenchFries{size}
}

func (f *FastFoodOrderBuilder) GetOrder() Order {
	return NewOrder(f.B, f.C, f.F)
}

// func main() {
// 	b := &builder.FastFoodOrderBuilder{}
// 	b.OrderBurger("small")
// 	b.OrderCoca("medium")
// 	b.OrderFriendFries("large")
// 	fmt.Println(b)
// 	order := b.GetOrder()
// 	fmt.Println(order.Invoice())

// }
