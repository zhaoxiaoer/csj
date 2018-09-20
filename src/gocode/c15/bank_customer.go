package c15

import (
	"fmt"
	"math"
)

var customers map[string]*BankCustomer

func init() {
	customers = make(map[string]*BankCustomer)
	customers["id001"] = &BankCustomer{"id001", "John", "Hacker", -3456.78}
	customers["id002"] = &BankCustomer{"id002", "Jane", "Hacker", 1234.56}
	customers["id003"] = &BankCustomer{"id003", "Juan", "Hacker", 987654.32}
}

func GetCustomer(id string) (*BankCustomer, error) {
	if _, ok := customers[id]; ok {
		return customers[id], nil
	} else {
		return nil, fmt.Errorf("%v", "Unknown customer id")
	}
}

type BankCustomer struct {
	Id        string
	FirstName string
	LastName  string
	Balance   float64
}

func (bc *BankCustomer) GetBalanceNoSign() float64 {
	return math.Abs(bc.Balance)
}
