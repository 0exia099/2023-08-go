package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

type Account struct {
	Balance int
}

func main() {
	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go DepositAndWithdraw(account)
	}
	wg.Wait()
	fmt.Println(account.Balance)
}
func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
	wg.Done()
}
