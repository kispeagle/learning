package state

type Context struct {
	state State
}

func (c *Context) SetState(s State) {
	c.state = s
}

func (c Context) ApplyState() string {
	return c.state.Handle()
}

type State interface {
	Handle() string
}

type CreateNewFile struct{}

func (c CreateNewFile) Handle() string {
	return "Create New File"
}

type SaveFile struct{}

func (s SaveFile) Handle() string {
	return "Save file"
}

// func main() {
// 	C := state.Context{}

// 	newFile := state.CreateNewFile{}
// 	save := state.SaveFile{}

// 	C.SetState(newFile)
// 	fmt.Println(C.ApplyState())

// 	C.SetState(save)
// 	fmt.Println(C.ApplyState())

// }
