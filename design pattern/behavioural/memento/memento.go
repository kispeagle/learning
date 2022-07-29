package memento

import "fmt"

type Originator struct {
	State string
}

func (o *Originator) SetState(state string) {
	o.State = state
}

func (o Originator) SaveMemento(c *CareTaker) {
	c.Add(Memento{o.State})
}

func (o *Originator) RestoreMemento(m Memento) {
	o.State = m.State
}

type Memento struct {
	State string
}

type CareTaker struct {
	Mementoes []Memento
}

func (c *CareTaker) Add(m Memento) {
	c.Mementoes = append(c.Mementoes, m)
	fmt.Println(c.Mementoes)
}

func (c CareTaker) Get(i int) Memento {
	return c.Mementoes[i]
}

// func main() {
// 	c := &memento.CareTaker{make([]memento.Memento, 0)}

// 	o := memento.Originator{}

// 	stateA := "A"
// 	o.SetState(stateA)
// 	fmt.Println("It's being state ", o.State)
// 	o.SaveMemento(c)
// 	stateB := "B"
// 	o.SetState(stateB)
// 	fmt.Println("It's being state ", o.State)
// 	o.SaveMemento(c)
// 	stateC := "C"
// 	o.SetState(stateC)
// 	o.SaveMemento(c)
// 	fmt.Println("It's being state ", o.State)
// 	o.RestoreMemento(c.Mementoes[1])
// 	fmt.Println("It's being state ", o.State)

// }
