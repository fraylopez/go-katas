package main

type Atm struct {
}

func NewAtm() *Atm {
	return &Atm{}
}

func Withdraw(amount int) []Note {
	res := make([]Note, 0)
	remainder := amount
	for remainder > 0 {
		if remainder >= 500 {
			res = append(res, Note{500})
			remainder -= 500
		} else if remainder >= 100 {
			res = append(res, Note{100})
			remainder -= 100
		}
	}
	return res
}

func AccumulateBill(bunch *[]Note) {

}
