package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Race Condition
// memanipulasi variable yang sama dengan banyak goroutine
// kondisi di mana tiap goroutine berkompetisi untuk memanipulasi
// variable yang sama sehingga ada nilai-nilai yang hilang
// dan tidak tampak perubahan pada varibale yang dimanipulasi
// kondisi ini dapat titangani dengan sync.Mutex

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)
}

// RWMutex
// race condition yang terjadi tidak hanya pada read sebuah variable
// namun juga pada write sebuah variable

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	// account.RWMutex.RLock()
	balance := account.Balance
	// account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total balance", account.Balance)
}
