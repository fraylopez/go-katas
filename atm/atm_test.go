package main

import (
	"testing"
)

func TestWithdrawSingleBankNote(t *testing.T) {
	output, _ := LoadedAtm().Withdraw(500)

	if output[0].value != 500 {
		t.Fatalf("Should be 500")
	}
}
func TestWithdrawMultipleBankNotes(t *testing.T) {
	output, _ := LoadedAtm().Withdraw(1000)
	if output[0].value != 500 && output[1].value != 500 {
		t.Fatalf("Should be 500,500")
	}
}
func TestWithdrawDifferentBankNotes(t *testing.T) {
	output, _ := LoadedAtm().Withdraw(600)
	if output[0].value != 500 || output[1].value != 100 {
		t.Fatalf("Should be 500,100")
	}
}

func TestWithdraw650(t *testing.T) {
	output, _ := LoadedAtm().Withdraw(650)
	if output[0].value != 500 || output[1].value != 100 || output[2].value != 50 {
		t.Fatalf("Should be 500,100,50")
	}
}

func TestRunOutOfNotes(t *testing.T) {
	atm := NewAtm()
	atm.Insert(Note{500})

	_, err := atm.Withdraw(1000)

	if err.Error() != "Not enough banknotes" {
		t.Fatalf("Should be No denomination found")
	}
}
