package Golang_Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("User 1 Locking  : ", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("User 2 Locking :", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlockBalance(t *testing.T) {
	user1 := UserBalance{
		Name:    "Alex",
		Balance: 1000,
	}

	user2 := UserBalance{
		Name:    "Melvin",
		Balance: 1000,
	}

	go Transfer(&user1, &user2, 500)

	time.Sleep(2 * time.Second)

	fmt.Println("User 1 : ", user1.Name, "Balance : ", user1.Balance)
	fmt.Println("User 2 : ", user2.Name, "Balance : ", user2.Balance)
}
