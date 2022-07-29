package singleton

import "fmt"

type Database struct {
	name string
}

var instance *Database = newDatabase("Mongo")

func newDatabase(name string) *Database {
	fmt.Println("Create new database")
	return &Database{name}
}

func GetInstance() *Database {
	fmt.Println("Get singleton database instance")
	return instance
}

func (d Database) Connect() {
	fmt.Println("Connect to ", d.name)
}

// func main() {
// 	db := singleton.GetInstance()
// 	db.Connect()

// 	fmt.Printf("%p\n", db)

// 	db1 := singleton.GetInstance()
// 	fmt.Printf("%p\n", db1)

// 	db1 := singleton.Database{"mysql"}
// 	a := singleton.newDatabase("postgres")
// }
