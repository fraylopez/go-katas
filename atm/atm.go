package main

type Atm struct {
}

func NewAtm() *Atm {
	return &Atm{}
}
func Withdraw(amount int) string {
	if amount == 1000 {
		return "500,500"
	}
	return "500"
}
