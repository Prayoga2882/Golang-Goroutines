package Golang_Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWmutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWmutex.Lock()
	account.Balance = account.Balance + amount
	account.RWmutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWmutex.RLock()
	balance := account.Balance
	account.RWmutex.RUnlock()
	return balance
}

func TestRWmutex(t *testing.T) {
	account := BankAccount{}
	x := 0
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				account.AddBalance(1)
				x += 1
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Final Balance", account.GetBalance())
	fmt.Println("X : ", x)
}
