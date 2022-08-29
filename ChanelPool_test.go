package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	var group = sync.WaitGroup{}

	pool.Put("Prayoga")
	pool.Put("Boedihartoyo")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(5 * time.Second)
			pool.Put(data)
		}()
	}
	group.Wait()
	fmt.Println("Done")
}
