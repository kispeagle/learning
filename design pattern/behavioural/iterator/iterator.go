package iterator

type Item struct {
	Value int
}

type IIterator interface {
	HasNext() bool
	Next() Item
}

type ItemIterator struct {
	Items   []Item
	current int
}

func (i *ItemIterator) Next() Item {
	if i.current < len(i.Items) {
		value := i.Items[i.current]
		i.current++
		return value
	}
	return Item{}
}

func (i ItemIterator) HasNext() bool {
	return i.current < len(i.Items)
}

// func main() {
// 	i := iterator.ItemIterator{Items: []iterator.Item{iterator.Item{12}, iterator.Item{34}}}

// 	fmt.Println(len(i.Items))

// 	for i.HasNext() {
// 		fmt.Println(i.Next())
// 	}
// }
