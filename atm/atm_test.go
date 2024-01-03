package main

import (
	"testing"
)

func TestWithdrawSingleBankNote(t *testing.T) {
	output := Withdraw(500)

	if output[0].value != 500 {
		t.Fatalf("Should be 500")
	}
}
func TestWithdrawMultipleBankNotes(t *testing.T) {
	output := Withdraw(1000)
	if output[0].value != 500 && output[1].value != 500 {
		t.Fatalf("Should be 500,500")
	}
}
