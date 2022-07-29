package prototype

type ICloneable interface {
	Clone() ICloneable
	SetName(string)
}

type Item struct {
	Name string
}

func (i *Item) SetName(name string) {
	i.Name = name
}

func (i *Item) Clone() ICloneable {
	return &Item{i.Name + "Clone"}
}

type Factory struct {
	ItemList []ICloneable
	Sample   Item
}

func (f *Factory) Produce(size int) {
	for i := 0; i < size; i++ {
		f.ItemList = append(f.ItemList, f.Sample.Clone())
	}
}

// func main() {
// 	sample := &prototype.Item{"Shoe"}
// 	f := prototype.Factory{Sample: *sample}
// 	for _, item := range f.ItemList {
// 		fmt.Println(item)
// 	}
// }
