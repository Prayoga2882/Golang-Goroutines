package Golang_Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func ChanIn(chanel chan<- string) {
	time.Sleep(2 * time.Second)
	chanel <- "Prayoga Boedihartoyo"
}

func ChanOut(chanel <-chan string) {
	data := <-chanel
	fmt.Println(data)
	time.Sleep(2 * time.Second)
}

func TestChanelInOut(t *testing.T) {
	chanel := make(chan string)
	defer close(chanel)

	go ChanIn(chanel)
	go ChanOut(chanel)

	time.Sleep(2 * time.Second)
}
