package Golang_Goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func Asyncronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("hello")
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go Asyncronous(group)
	}
	group.Wait()
	fmt.Println("All Done !")

}
