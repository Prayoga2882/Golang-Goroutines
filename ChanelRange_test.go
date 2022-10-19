package Golang_Goroutines

import (
	"fmt"
	"strconv"
	"testing"
)

func TestChanelRange(t *testing.T) {
	chanel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			chanel <- "Sending Data : " + strconv.Itoa(i)
		}
		close(chanel) // harus di close kalau gamau deadlock
	}()

	for data := range chanel {
		fmt.Println(data)
	}
}
