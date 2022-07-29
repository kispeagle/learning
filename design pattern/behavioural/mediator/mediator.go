package mediator

import "fmt"

type User struct {
	H    Hub
	Name string
}

func (u User) Send(msg string) {
	fmt.Println(u.Name, " send message: ", msg)
	u.H.Send(msg, u)
}

func (u User) Recv(msg string) {
	fmt.Println(u.Name, " receive message: ", msg)
}

type Hub struct {
	Users map[string]User
}

func (h *Hub) Register(u User) {
	fmt.Println("Mediator add user: ", u.Name)
	h.Users[u.Name] = u
}

func (h *Hub) Deregister(u User) {
	fmt.Println("Mediator remove user: ", u.Name)
	if _, ok := h.Users[u.Name]; ok {
		delete(h.Users, u.Name)
	}
}

func (h Hub) Send(msg string, u User) {
	for _, v := range h.Users {
		if u.Name != v.Name {
			v.Recv(msg)
		}
	}
}

// func main() {
// 	h := mediator.Hub{
// 		Users: make(map[string]mediator.User),
// 	}

// 	user1 := mediator.User{h, "Trung"}
// 	user2 := mediator.User{h, "Anna"}

// 	h.Register(user1)
// 	h.Register(user2)

// 	user1.Send("Hi everyone")
// 	user2.Send("Nice to meet you")
// }
