package golang_gorutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// mengunakan mutex jika ada variable yg dishare ke beberapa goroutine
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {

	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)

				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("total balance", account.GetBalance())
}

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User1 ", user1.Name)
	user1.change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User2 ", user2.Name)
	user2.change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()

}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "bima",
		Balance: 100000,
	}

	user2 := UserBalance{
		Name:    "Komang",
		Balance: 30000,
	}
	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 10000)

	time.Sleep(1 * time.Second)

	fmt.Println("User1 ", user1.Name, "Balance", user1.Balance)
	fmt.Println("User2 ", user2.Name, "Balance", user2.Balance)
}