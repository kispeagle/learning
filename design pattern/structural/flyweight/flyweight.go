package flyweight

import "fmt"

type ISoldier interface {
	Promote(string, int)
	SetName(string)
}

type Solidier struct {
	Name string
}

func (s *Solidier) SetName(name string) {
	s.Name = name
}

func (s *Solidier) Promote(name string, star int) {
	fmt.Printf("Solidier %s - start %d\n", name, star)
}

type Context struct {
	Name string
	Star int
}

type FactorySoldier struct {
	Solidiers map[string]ISoldier
}

func (f *FactorySoldier) CreateSoldier(c Context) ISoldier {
	if v, ok := f.Solidiers[c.Name]; ok {
		fmt.Println("Get Soldier: ", c.Name)
		fmt.Printf("%p\n", v)
		v.SetName(c.Name)
		return v
	} else {
		fmt.Println("Create New Soldier: ", c.Name)
		s := &Solidier{c.Name}
		f.Solidiers[c.Name] = s
		fmt.Printf("%p\n", s)
		return s
	}
}

// func main() {
// 	noNormal := 10
// 	noDocter := 4
// 	soldiers := []flyweight.ISoldier{}
// 	army := flyweight.FactorySoldier{make(map[string]flyweight.ISoldier)}
// 	for i := 0; i < noNormal; i++ {
// 		idxStr := strconv.Itoa(len(soldiers))
// 		context := flyweight.Context{"Normal ", 4}
// 		s := army.CreateSoldier(context)
// 		s.Promote(context.Name+idxStr, context.Star)
// 		soldiers = append(soldiers, s)
// 	}

// 	for i := 0; i < noDocter; i++ {
// 		idxStr := strconv.Itoa(len(soldiers))
// 		context := flyweight.Context{"Docter ", 2}
// 		s := army.CreateSoldier(context)
// 		s.Promote(context.Name+idxStr, context.Star)
// 		soldiers = append(soldiers, s)
// 	}
// }
