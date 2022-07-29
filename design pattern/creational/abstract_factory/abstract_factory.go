package abfactory

import "fmt"

func AbstractFactory(factory string) SportFactory {
	switch factory {
	case "AdidasFactory":
		return AdidasFactory{}

	case "PumaFactory":
		return PumaFactory{}
	default:
		fmt.Println("Don't support this type of factory")
		return nil
	}
}

type SportFactory interface {
	CreateShoe() string
	CreateSock() string
}

type Shoe interface {
	CreateShoe() string
	CreateSock() string
}

type PumaShoe struct{}

func (p PumaShoe) CreateShoe() string {
	return "Create Puma shoe"
}

type AdidasShoe struct{}

func (a AdidasShoe) CreatShoe() string {
	return "Create Adidas shoe"
}

type Sock interface {
	CreateSock() string
}

type PumaSock struct{}

func (p PumaSock) CreateSock() string {
	return "Create Puma sock"
}

type AdidasSock struct{}

func (a AdidasSock) CreateSock() string {
	return "Create Adidas sock"
}

type AdidasFactory struct {
	Sock AdidasSock
	Shoe AdidasShoe
}

func (a AdidasFactory) CreateShoe() string {
	return a.Shoe.CreatShoe()
}

func (a AdidasFactory) CreateSock() string {
	return a.Sock.CreateSock()
}

type PumaFactory struct {
	Sock PumaSock
	Shoe PumaShoe
}

func (p PumaFactory) CreateShoe() string {
	return p.Shoe.CreateShoe()
}

func (p PumaFactory) CreateSock() string {
	return p.Sock.CreateSock()
}

// func main() {
// 	adidasF := abfactory.AbstractFactory("AdidasFactory")
// 	fmt.Println(adidasF.CreateShoe())
// 	fmt.Println(adidasF.CreateSock())

// 	pumaF := abfactory.AbstractFactory("PumaFactory")
// 	fmt.Println(pumaF.CreateShoe())
// 	fmt.Println(pumaF.CreateSock())
// }
