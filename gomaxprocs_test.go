package golang_gorutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGOMaxProcs(t *testing.T) {

	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	// group.Wait()
	totalCPU := runtime.NumCPU()

	fmt.Println("total CPU = ", totalCPU)

	totalTharead := runtime.GOMAXPROCS(-1)

	fmt.Println("Total Thread = ", totalTharead)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine = ", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {

	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	// group.Wait()
	totalCPU := runtime.NumCPU()

	fmt.Println("total CPU = ", totalCPU)
	runtime.GOMAXPROCS(10)
	totalTharead := runtime.GOMAXPROCS(-1)

	fmt.Println("Total Thread = ", totalTharead)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine = ", totalGoroutine)

	group.Wait()
}
