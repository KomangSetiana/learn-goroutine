package golang_gorutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {

	tiker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		tiker.Stop()
	}()
	for time := range tiker.C {
		fmt.Println(time)
	}

}

func TestTick(t *testing.T) {

	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}

}
