package pool

import (
	"fmt"
	"sync"
	"time"
)

type Taxi struct {
	Name string
}

type TaxiPool struct {
	Available []*Taxi
	InUse     []*Taxi
	Mutex     *sync.Mutex
}

func (t *TaxiPool) GetTaxi() {

	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	if len(t.Available) == 0 {
		fmt.Println("Client cancel")
		return
	}

	taxi := t.Available[0]
	t.InUse = append(t.InUse, taxi)
	t.Available = t.Available[1:]
	go t.Serving(taxi.Name)
}

func (t *TaxiPool) Serving(name string) {
	fmt.Printf("Taxi %s is serving\n", name)
	time.Sleep(1200 * time.Millisecond)
	t.Release()
}

func (t *TaxiPool) Release() {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	taxi := t.InUse[0]
	fmt.Printf("Taxi %s is  free \n", taxi.Name)
	t.InUse = t.InUse[1:]
	t.Available = append(t.Available, taxi)
}

type Client struct {
	Pool *TaxiPool
}

func (c Client) GrabTaxi() {
	c.Pool.GetTaxi()
}

// func main() {
// 	taxiPool := pool.TaxiPool{Available: make([]*pool.Taxi, 0), InUse: make([]*pool.Taxi, 0), Mutex: new(sync.Mutex)}

// 	numTaxi := 4

// 	for i := 0; i < numTaxi; i++ {
// 		taxi := pool.Taxi{fmt.Sprintf("taxi %d", i)}
// 		taxiPool.Available = append(taxiPool.Available, taxi)
// 	}

// 	noGuest := 10

// 	for i := 0; i < noGuest; i++ {
// 		client := pool.Client{Pool: &taxiPool}
// 		client.GrabTaxi()
// 		time.Sleep(200 * time.Millisecond)
// 	}

// }
