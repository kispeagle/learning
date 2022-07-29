package adapter

import "fmt"

type Projector struct {
}

func (p Projector) Connect(c Device) {
	fmt.Println("Projector plugged into device through vga port")
	c.Receive()
}

type Device interface {
	Receive()
}

type Computer struct{}

func (c Computer) Receive() {
	fmt.Println("To accept device to connect through vga port")
}

type Phone struct{}

func (p Phone) charge() {
	fmt.Println("To accept device to connect through type-c port")
}

type Adapter struct {
	P *Phone
}

func (a Adapter) Receive() {
	fmt.Println("Convert vga jack to type-c jack")
	a.P.charge()
}

// func main() {
// 	p := adapter.Projector{}

// 	c := adapter.VgaComputer{}
// 	p.PluginToVGA(c)

// 	h := new(adapter.HdmiComputer)
// 	a := adapter.Adapter{h}

// 	p.PluginToVGA(a)
// }
