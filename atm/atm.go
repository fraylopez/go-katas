package main

import "errors"

type Atm struct {
	Vault map[Note]int
}

func NewAtm() *Atm {
	var atm Atm
	atm.Vault = map[Note]int{Note{500}: 0}
	return &atm
}
func LoadedAtm() *Atm {
	var atm Atm
	atm.Vault = map[Note]int{Note{500}: 10, Note{200}: 10, Note{100}: 10, Note{50}: 10, Note{20}: 10, Note{10}: 10, Note{5}: 10, Note{2}: 10, Note{1}: 10}
	return &atm
}

func (atm Atm) Withdraw(amount int) ([]Note, error) {
	res := make([]Note, 0)
	remainder := amount
	for remainder > 0 {
		_res, err := GetBiggerDenomination(remainder)
		if err != nil {
			return nil, err
		}
		if atm.Vault[_res] == 0 {
			return nil, errors.New("Not enough banknotes")
		}
		res = append(res, _res)
		remainder -= res[len(res)-1].value
		atm.Vault[_res]--
	}
	return res, nil
}

func (atm Atm) Insert(note Note) {
	atm.Vault[note]++
}

func GetBiggerDenomination(amount int) (Note, error) {
	denominations := []int{500, 200, 100, 50, 20, 10, 5, 2, 1}
	for _, denomination := range denominations {
		if amount >= denomination {
			return Note{denomination}, nil
		}
	}
	return InvalidNote(), errors.New("No denomination found")
}
