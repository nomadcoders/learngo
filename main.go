package main

import (
	"fmt"

	"github.com/serranoarevalo/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
}
