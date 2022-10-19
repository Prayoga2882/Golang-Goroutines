package Golang_Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello guys")
}

func TestGoroutines(t *testing.T) {
	go HelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumberGoroutine(number int) {
	fmt.Println("Display : ", number)
}

func TestDisplayGoroutines(t *testing.T) {
	for i := 0; i < 100; i++ {
		go DisplayNumberGoroutine(i)
	}

	time.Sleep(10 * time.Second)
}
