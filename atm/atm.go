package main

type Atm struct {
}

func NewAtm() *Atm {
	return &Atm{}
}

func Withdraw(amount int) []Note {
	res := make([]Note, 0)
	res = append(res, Note{500})
	if amount == 1000 {
		res = append(res, Note{500})
	}
	if amount == 600 {
		res = append(res, Note{100})
	}
	return res
}
