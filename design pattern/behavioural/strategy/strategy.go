package strategy

import "fmt"

type SortAlgorithmStrategy interface {
	Sort()
}

type QuickSort struct {
}

func (q QuickSort) Sort() {
	fmt.Println("Quick Sort algorithm")
}

type MergeSort struct {
}

func (m MergeSort) Sort() {
	fmt.Println("Merge Sort algorithm")
}

type ShellSort struct{}

func (s ShellSort) Sort() {
	fmt.Println("Shell sort algorithm")
}

type ListItem struct {
	Sort  SortAlgorithmStrategy
	Items []string
}

func (l *ListItem) SetSort(s SortAlgorithmStrategy) {
	l.Sort = s
}

func (l ListItem) SortItem() {
	l.Sort.Sort()
}

// func main() {
// 	fmt.Println("Create list item")
// 	listItem := strategy.ListItem{Items: []string{"Laptop", "Iphone", "Clothes"}, Sort: nil}
// 	listItem.SetSort(strategy.MergeSort{})
// 	listItem.SortItem()

// 	listItem.SetSort(strategy.QuickSort{})
// 	listItem.SortItem()
// }
