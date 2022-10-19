package Golang_Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse(chanel chan string) {
	time.Sleep(2 * time.Second)
	chanel <- "Prayoga Boedihartoyo"
}

func TestChanelAsParameter(t *testing.T) {
	chanel := make(chan string)
	defer close(chanel)

	go GiveMeResponse(chanel)

	data := <-chanel
	fmt.Println(data)
	time.Sleep(2 * time.Second)
}
