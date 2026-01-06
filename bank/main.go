package main

import (
	"fmt"

	"github.com/rkweber-max/go-object-oriented/bank/accounts"
	"github.com/rkweber-max/go-object-oriented/bank/customers"
)

func main() {
	accountRodrigo := accounts.CurrentAccount{
		Owner:         customers.Customer{Name: "Rodrigo", CPF: "12345678900", Email: "rodrigo@example.com"},
		AgencyNumber:  123,
		AccountNumber: 123456,
		Balance:       1234.50,
	}

	accountAmanda := accounts.CurrentAccount{
		Owner:         customers.Customer{Name: "Amanda", CPF: "12345678900", Email: "amanda@example.com"},
		AgencyNumber:  123,
		AccountNumber: 123456,
		Balance:       1234.50,
	}

	fmt.Println("Balance Rodrigo:", accountRodrigo.Balance)
	fmt.Println(accountRodrigo.Transfer(100.00, &accountAmanda))
	fmt.Println("Balance Rodrigo after transfer:", accountRodrigo.Balance)
	fmt.Println("Balance Amanda after transfer:", accountAmanda.Balance)
}
