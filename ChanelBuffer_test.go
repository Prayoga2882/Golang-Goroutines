package Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestChanelBuffer(t *testing.T) {
	chanel := make(chan string, 2)
	defer close(chanel)

	chanel <- "Prayoga Boeduhartoyo" // mengirim chanel harus menggunakan goroutine biar gak deadlock

	data := <-chanel

	fmt.Println(data)
	time.Sleep(2 * time.Second)
}
