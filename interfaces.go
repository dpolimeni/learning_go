// code with the syntax error
// fix the syntax error by adding a semicolon at the end of the line
package main

import (
	"fmt"
	"time"
)

type email1 struct {
	toAddress   string
	fromAddress string
	message     string
	cost        float64
}

type sms struct {
	toNumber string
	message  string
	cost     float64
}

type expenseprinter interface {
	printExpense() (string, float64)
}

func getCost(e expenseprinter) float64 {
	em, isemail := e.(email1)
	if isemail {
		_, cost := em.printExpense()
		return cost + 0.1
	}
	sm, issms := e.(sms)
	if issms {
		_, cost := sm.printExpense()
		return cost + 0.2
	}
	return 0.0
}

func (e email1) printExpense() (string, float64) {
	return e.message, e.cost
}

func (s sms) printExpense() (string, float64) {
	return s.message, s.cost
}

func getcostSwitch(e expenseprinter) float64 {
	switch e.(type) {
	case email1:
		return 0.1
	case sms:
		return 0.2
	default:
		return 0.0
	}
}

func main() {
	// calculate time taken
	start := time.Now()

	var e1 email1 = email1{toAddress: "diego.com", fromAddress: "gianluca.com", message: "ciao", cost: 0.0}
	var s1 sms = sms{toNumber: "123456789", message: "ciao", cost: 0.0}

	e1.printExpense()
	s1.printExpense()
	end := time.Now()
	fmt.Println("Time taken: ", end.Sub(start), "seconds")

	fmt.Println("Cost of email: ", getCost(e1))
	fmt.Println("Cost of sms: ", getCost(s1))
	fmt.Println("Cost of email: ", getcostSwitch(e1))
}
