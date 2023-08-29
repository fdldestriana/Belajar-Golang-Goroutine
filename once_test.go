package main

import (
	"fmt"
	"sync"
	"testing"
)

// once adalah fitur yang dimiliki golang untuk memastikan bahwa sebuah function hanya dapat dipanggil satu kali
// apabila banyak goroutine memanggil function tersebut, maka hanya goroutine pertama lah yang akan dapat memanggilnya

var counter = 0

func OnlyOnce() {

	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		group.Add(1)
		go func() {
			once.Do(OnlyOnce) // akan menghasilkan 1
			// OnlyOnce() // akan menghasilkan random karena race condition
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter", counter)
}
