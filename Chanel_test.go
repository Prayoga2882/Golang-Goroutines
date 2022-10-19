package Golang_Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	chanelPengirim := make(chan string)
	defer close(chanelPengirim)

	go func() {
		time.Sleep(5 * time.Second)
		chanelPengirim <- "Melvin" // akan terjadi deadlock
		fmt.Println("Everything is done")
	}()

	data := <-chanelPengirim
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestGo(t *testing.T) {
	chanel := make(chan string)
	//defer close(chanel)

	go sender(chanel, "melvin sending data")
	go reciever(chanel)
	fmt.Println("all done")
	time.Sleep(3 * time.Second)
}

func sender(chanel chan string, value string) {
	chanel <- value
}

func reciever(chanel chan string) {
	data := <-chanel
	fmt.Println("success receive data from chanel\n", data)
}
