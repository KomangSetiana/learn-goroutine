package golang_gorutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateGoRoutineChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Bima"
		fmt.Println("selesai mengirim data ke channel")
	}()

	data := <-channel

	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)

	channel <- "Bimasena"

}
func TestChannelParameter(t *testing.T) {

	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)

	data := <-channel

	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

func OnlySendChannel(channel chan<- string) {
	time.Sleep(2 * time.Second)

	channel <- "Bimasena"
}

func OnlyReciveChannel(channel <-chan string) {
	time.Sleep(2 * time.Second)

	data := <-channel

	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlySendChannel(channel)
	go OnlyReciveChannel(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)
	go func() {
		channel <- "Bima"
		channel <- "Sena"
		channel <- "Ganteng"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
}

// menerima jumlah data tidak pasti pada channel
func TestRangeChannel(t *testing.T) {

	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke : " + strconv.Itoa(i)
		}
		//wajib close
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Done")
}

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
			fmt.Println("menerima data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("menerima data dari channel 2", data)
			counter++
			// melakukan sesuatu kalau data belum ada di channel
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}

	}
}
