package main

import "errors"

type Atm struct {
}

func NewAtm() *Atm {
	return &Atm{}
}

func Withdraw(amount int) ([]Note, error) {
	res := make([]Note, 0)
	remainder := amount
	for remainder > 0 {
		_res, err := GetBiggerDenomination(remainder)
		if err != nil {
			return nil, err
		}
		res = append(res, _res)
		remainder -= res[len(res)-1].value
	}
	return res, nil
}

func GetBiggerDenomination(amount int) (Note, error) {
	denominations := []int{500, 100, 50}
	for _, denomination := range denominations {
		if amount >= denomination {
			return Note{denomination}, nil
		}
	}
	return InvalidNote(), errors.New("No denomination found")
}
