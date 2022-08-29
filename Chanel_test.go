package Goroutines

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
