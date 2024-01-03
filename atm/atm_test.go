package main

import (
	"testing"
)

func TestATM(t *testing.T) {
	var output = Withdraw(500)
	if output != "500" {
		t.Fatalf("Should be 500")
	}
}
