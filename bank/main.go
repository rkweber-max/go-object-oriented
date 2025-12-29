package main

import "fmt"

type CurrentAccount struct {
	owner         string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	ownerRodrigo := CurrentAccount{
		owner:         "Rodrigo",
		agencyNumber:  123,
		accountNumber: 123456,
		balance:       1234.50,
	}

	fmt.Println(ownerRodrigo)

	var accountRodrigo *CurrentAccount
	accountRodrigo = new(CurrentAccount)
	accountRodrigo.owner = "Rodrigo"

	fmt.Println(*accountRodrigo)
}
