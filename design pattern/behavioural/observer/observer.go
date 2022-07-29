package observer

import "fmt"

type Login struct {
	State     bool
	Observers []Observer
}

func (l *Login) Login() {
	if !l.State {
		fmt.Println("User login Successfully")
	} else {
		fmt.Println("User logout")
	}
	l.ChangeState()
}

func (l *Login) ChangeState() {
	l.State = !l.State

	for _, e := range l.Observers {
		e.Update(l.State)
	}
}

type Observer interface {
	Update(bool)
}

type DB struct {
}

func (db DB) Update(state bool) {
	fmt.Println("Database: user online is ", state)
}

type Console struct {
}

func (c Console) Update(state bool) {
	fmt.Println("Console: user online is ", state)
}

type Email struct {
}

func (e Email) Update(state bool) {
	fmt.Println("Email: user online is ", state)
}

// func main() {

// 	email := observer.Email{}
// 	console := observer.Console{}
// 	db := observer.DB{}

// 	user := observer.Login{false, []observer.Observer{email, console, db}}

// 	user.Login()
// 	user.Login()
// }
