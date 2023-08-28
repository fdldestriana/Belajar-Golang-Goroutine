package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// channel adalah cara komunikasi antara dua atau lebih goroutine
// komunikasi secara synchronous, dalam artian setiap goroutine yang mengirimkan data
// ke dalam sebuah channel, maka ia akan terblock sampai ada goroutine lain
// yang mengambil datanya, dan kita hanya bisa mengirim satu data ke dalam channel
// lalu jika sebuah goroutine akan mengambil data dari sebuah channel,
// maka ia akan menuggu sampai data tersedia di dalam channel tersebut
// channel ini mirip dengan async await pada bahasa pemrograman lain
// channel hanya bisa menerima satu jenis tipe data
// channel bisa diambil dari lebih dari satu goroutine, tapi tiap goroutine harus menunggu
// channel harus diclose secara rekomendasi agar dihapus oleh garbage collector golang

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // untuk memastikan channel diclose bisa gunakan defer

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Eko Kurniawan Khannedy"
		fmt.Println("Selesai mengirim data ke channel")
	}() // gunakan () pada anon function untuk langsung mengeksekusi function tersebut
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// channel sebagai parameter
// untuk menggunakan channel sebagai parameter tidak perlu menggunakan asterisk
// karena tanpa asterisk pun channel sudah merupakan sebuah referensi ke dalam memory address

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khannedy"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println("Ini adalah nilai dari variable data", data)
	time.Sleep(2 * time.Second)
}

// saat menggunakan channel sebagai parameter, di dalam body function
// channel tersebut bisa digunakan untuk menerim datau mengirim data
// namun kita dapat mempertegas apakah parameter tersebut untuk mengirim
// atau menerima data pada braces function tersebut

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khannedy"
	channel <- "Fadly Destriana Rusmana"
	// channel <- "Destriana Fadly Rusmana"
	// channel <- "Rusmana Fadly Destriana "
	// channel <- "Rusmana Destriana Fadly"
	// channel <- "Fadly Rusmana Destriana"
}
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("Ini adalah data pertama dari function OnlyOut", data)
	data2 := <-channel
	fmt.Println("Ini adalah data kedua dari function OnlyOut", data2)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
}

// range channel
// mengirim dan mengambil data ke dan dari channel menggunakan for loop
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke" + strconv.Itoa(i)
		}
		close(channel) // jika tidak close maka akan terkena deadlock
	}()
	for data := range channel {
		fmt.Println("Menerima data", data)
	}
}

// select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel pertama", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel kedua", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}
