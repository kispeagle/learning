package interpreter

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// interpreter translate statement into state
// A instance operate relying on state

type Interpreter struct {
	add AddOperation
	sub SubtractOperation
}

func New(add AddOperation, sub SubtractOperation) *Interpreter {
	return &Interpreter{add, sub}
}

func (i Interpreter) Translate(statement string) []int {
	pattern := regexp.MustCompile("[^0-9]+")
	statement = pattern.ReplaceAllString(statement, " ")
	operands := strings.Split(statement, " ")

	numbers := []int{}
	for _, e := range operands {
		num, _ := strconv.Atoi(e)
		numbers = append(numbers, num)
	}

	return numbers
}

func (i Interpreter) Intepret(statement string) int {
	if strings.Index(statement, "add") > -1 {
		operands := i.Translate(statement)
		return i.add.Add(operands[0], operands[1])
	} else if strings.Index(statement, "sub") > -1 {
		operands := i.Translate(statement)
		return i.sub.Subtract(operands[0], operands[1])
	}

	fmt.Println("Doesn't support")
	return int(math.NaN())
}

type AddOperation struct{}

func (a AddOperation) Add(x, y int) int {
	return x + y
}

type SubtractOperation struct{}

func (a SubtractOperation) Subtract(x, y int) int {
	return x - y
}

// func main() {
// 	add := interpreter.AddOperation{}
// 	sub := interpreter.SubtractOperation{}
// 	interpreter := interpreter.New(add, sub)

// 	statement := "10 add 10"
// 	fmt.Println(statement, " = ", interpreter.Intepret(statement))
// 	statement = "10 sub 10"
// 	fmt.Println(statement, " = ", interpreter.Intepret(statement))

// }
