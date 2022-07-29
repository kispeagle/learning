package factory

type AdidasType string

type PumaType string

type Shoe interface {
	ProduceShoe() string
}

type Puma struct {
}

func (p Puma) ProduceShoe() string {
	return "Puma shoe"
}

type Adidas struct {
}

func (a Adidas) ProduceShoe() string {
	return "Adidas Shoe"
}

func ShoeFactory(shoe interface{}) Shoe {
	switch shoe.(type) {
	case AdidasType:
		return Adidas{}
	case PumaType:
		return Puma{}
	}
	return nil
}

// func main() {
// 	var shoe factory.AdidasType //= "sneaker"
// 	var shoe1 factory.PumaType
// 	brand := factory.ShoeFactory(shoe)
// 	fmt.Println(brand.ProduceShoe())

// 	brand2 := factory.ShoeFactory(shoe1)
// 	fmt.Println(brand2.ProduceShoe())
// }
