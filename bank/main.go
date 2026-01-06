package main

import (
	"fmt"

	"github.com/rkweber-max/go-object-oriented/bank/accounts"
)

func main() {

	accountExample := accounts.CurrentAccount{}
	accountExample.Deposit(100.00)
	fmt.Println(accountExample.GetBalance())
}
