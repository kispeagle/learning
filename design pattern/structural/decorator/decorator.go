package decorator

import "fmt"

type Ipizza interface {
	GetPrice() int
	GetIntegration() string
}

type Pizza struct {
	Integrations []string
}

func (p Pizza) GetPrice() int {
	return 14
}

func (p Pizza) GetIntegration() string {
	result := "integration includes in: "
	for _, e := range p.Integrations {
		result += fmt.Sprintf(" + %s ", e)
	}
	return result
}

type VegPizza struct {
	Pizza Ipizza
}

func (v VegPizza) GetPrice() int {
	return 10 + v.Pizza.GetPrice()
}

func (v VegPizza) GetIntegration() string {
	return v.Pizza.GetIntegration()
}

// func main() {

// 	normalPizza := decorator.Pizza{[]string{"floor", "butter", "seafood"}}

// 	fmt.Println(normalPizza.GetIntegration())
// 	fmt.Println(normalPizza.GetPrice())

// 	vegPizza := decorator.VegPizza{decorator.Pizza{[]string{"floor", "butter", "bell peper", "onion"}}}
// 	fmt.Println(vegPizza.GetIntegration())
// 	fmt.Println(vegPizza.GetPrice())
// }
