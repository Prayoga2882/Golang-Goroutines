package Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse1(chanel chan string) {
	time.Sleep(2 * time.Second)
	chanel <- "Prayoga Boedihartoyo"
}

func TestChanelSelectDefault(t *testing.T) {
	chanel1 := make(chan string)
	chanel2 := make(chan string)
	defer close(chanel1)
	defer close(chanel2)

	go GiveMeResponse1(chanel1)
	go GiveMeResponse1(chanel2)

	counter := 0
	for {
		select {
		case data := <-chanel1:
			fmt.Println("Data From Chanel 1", data)
			counter++
		case data := <-chanel2:
			fmt.Println("Data From Chanel 2", data)
			counter++
		default:
			fmt.Println("Waiting...")
		}
		if counter == 2 {
			break
		}
	}
}