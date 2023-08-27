package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

// dengan goroutine kode program yang kita buat akan berjalan secara asynchronous
// namun berbeda dengan bahasa pemrograman lain yang harus menggunakan keyword async await
// di dalam bahasa golang kita cukup gunakan keyword go pada saat memanggil function
// yang ingin dijalankan menggunakan goroutine
// ketika menjalan goroutine, pastikan kode programnya tidak selesai terlebih dahulu
// namun penggunaan goroutine tidak cocok untuk function yang mengembalikan value

func TestCreateGoroutine(t *testing.T) {
	RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

// kita akan melakukan eksperimen yang membuktikan bahwa goroutine sangat ringan

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(2 * time.Second)
}
