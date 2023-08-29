package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// menggunakan wait group untuk menunggu goroutine selesai dieksekusi adalah metode yang terbaik
// menggunakan time sleep hanya mengira - ngira bahwa dalam kurun waktu sleep semua goroutine
// akan selesai dieksekusi,

// gunakan Add function untuk mengindikasikan bahwa kita memiliki goroutine yang akan dieksekusi dan melakukan await
// gunakan Done function yang akan mengindikasikan bahwa goroutine telah selesai dieksekusi sehingga program akan selesai
// melakukan await dan menyelesaikan block code function goroutine yang sedang dieksekusi

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			RunAsynchronous(&group)
		}()
	}
	group.Wait()
	fmt.Println("Selesai")
}
