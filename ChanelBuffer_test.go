package Golang_Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestChanelBuffer(t *testing.T) {
	chanel := make(chan string, 2)
	defer close(chanel)

	chanel <- "Prayoga Boedihartoyo" // chanel is how the goroutine communication

	data := <-chanel

	fmt.Println(data)
	time.Sleep(2 * time.Second)
}
