package command

import "fmt"

type Bank struct {
	OpenAccount  Command
	CloseAccount Command
}

func (b Bank) ClickCloseAccount() {
	fmt.Println("User click close account")
	b.CloseAccount.Execute()
}

func (b Bank) ClickOpenAccount() {
	fmt.Println("User click open acocunt")
	b.OpenAccount.Execute()
}

type Account struct {
	Name string
}

func (a Account) Open() {
	fmt.Println(fmt.Sprintf("Account %s open", a.Name))
}

func (a Account) Close() {
	fmt.Println(fmt.Sprintf("Accounts %s close", a.Name))
}

type Command interface {
	Execute()
}

type OpenCommand struct {
	Acc Account
}

func (o OpenCommand) Execute() {
	o.Acc.Open()
}

type CloseCommand struct {
	Acc Account
}

func (c CloseCommand) Execute() {
	c.Acc.Close()
}

// func main() {

// 	account := command.Account{"Credit"}
// 	openCommand := command.OpenCommand{account}
// 	closeCommand := command.CloseCommand{account}
// 	VietComBank := command.Bank{OpenAccount: openCommand, CloseAccount: closeCommand}
// 	VietComBank.ClickOpenAccount()
// 	VietComBank.ClickCloseAccount()
// }
