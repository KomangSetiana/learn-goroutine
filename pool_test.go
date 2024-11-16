package golang_gorutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Komang")
	pool.Put("Bima")
	pool.Put("Sena")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
			group.Done()
		}()
	}
	group.Wait()
}
