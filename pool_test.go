package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// pool adalah implementasi dari design pattern bernama object pool pattern
// pool di golang digunakan untuk concurrency dan parallel programming
// implementasinya adalah untuk menyimpan data, kemudian ketika ingin
// menggunakan sebuah data, kita mengambilnya dari pool dan kemudian
// ketika telah selesai menggunakannya kita simpan data tersebut ke dalam pool lagi
// implementasi pool di golang sudah aman dari problem race condition
// dan biasanya digunakan untuk membuat koneksi ke database
// misal saat awal aplikasi dibuka, kita membuat koneksi ke database
// lalu koneksi tersebut kita simpan pada pool dan kemudian ketika akan
// menggunakan koneksi tersebut kita ambil dari pool
// kode di bawah berbeda dengan video udemy dari pak Eko
// kode di bawah menggunakan pointer

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface {
		} {
			return "New"
		},
	}
	group := sync.WaitGroup{}

	eko := "Eko"
	kurniawan := "Kurniawan"
	khannedy := "Khannedy"

	pool.Put(&eko)
	pool.Put(&kurniawan)
	pool.Put(&khannedy)

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			data := pool.Get()
			if data != "New" {
				value := data.(*string)
				fmt.Println(*value)
			} else {
				fmt.Println(data)
			}

			time.Sleep(1 * time.Second)
			pool.Put(data)
			defer group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Selesai")
}
