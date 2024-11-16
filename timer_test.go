package golang_gorutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimee(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	fmt.Println(time.Now())

	time := <-timer.C

	fmt.Println(time)

}
func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)

	fmt.Println(time.Now())

	tick := <-channel

	fmt.Println(tick)

}

func TestAfterFunc(t *testing.T) {

	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Execute after 5 second")
		fmt.Println(time.Now())

		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()

}
