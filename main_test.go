package main

import (
	"testing"
)

func Test_validateBankNoteProblem_success(t *testing.T) {
	var bnp BankNoteProblem
	bnp.robberShare = []int{
		10,
		20,
		30}
	bnp.bankNoteDecks = []BankNoteDeck{
		BankNoteDeck{1, 5},
		BankNoteDeck{5, 10},
		BankNoteDeck{10, 10},
		BankNoteDeck{20, 20},
		BankNoteDeck{50, 5},
		BankNoteDeck{100, 10}}
	err := validateBankNoteProblem(&bnp)
	if err != nil {
		t.Log("error but not expected!")
		t.Fail()
	} else {
		t.Log("Good!")
	}
}

func Test_validateBankNoteProblem_failed(t *testing.T) {
	var bnp BankNoteProblem
	robberShareArray := [...]int{
		10,
		20,
		31}
	bnp.robberShare = robberShareArray[:]
	bankNoteDecksArray := [...]BankNoteDeck{
		BankNoteDeck{1, 5},
		BankNoteDeck{5, 10},
		BankNoteDeck{10, 10},
		BankNoteDeck{20, 20},
		BankNoteDeck{50, 5},
		BankNoteDeck{100, 10}}
	bnp.bankNoteDecks = bankNoteDecksArray[:]
	err := validateBankNoteProblem(&bnp)
	if err != nil {
		t.Log("Good!")
	} else {
		t.Log("Expected to be error!")
		t.Fail()
	}
}

func Test_DefaultSolution_1(t *testing.T) {
	var bnp BankNoteProblem
	bnp.robberShare = []int{
		10,
		20,
		30}
	bnp.bankNoteDecks = []BankNoteDeck{
		BankNoteDeck{1, 5},
		BankNoteDeck{5, 10},
		BankNoteDeck{10, 10},
		BankNoteDeck{20, 20},
		BankNoteDeck{50, 5},
		BankNoteDeck{100, 10}}
	bns := getDefaultSolution(&bnp)
	if len(bns.robberAccounts) != len(bnp.robberShare) {
		t.Log("RobberAccounts No. != RobberShare No.")
		t.Fail()
	}
	expectedBns := BankNoteSolution{[]RobberAccount{
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 5},
				BankNoteDeck{5, 5},
				BankNoteDeck{10, 0},
				BankNoteDeck{20, 0},
				BankNoteDeck{50, 0},
				BankNoteDeck{100, 0}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 0},
				BankNoteDeck{5, 5},
				BankNoteDeck{10, 10},
				BankNoteDeck{20, 5},
				BankNoteDeck{50, 0},
				BankNoteDeck{100, 0}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 0},
				BankNoteDeck{5, 0},
				BankNoteDeck{10, 0},
				BankNoteDeck{20, 15},
				BankNoteDeck{50, 5},
				BankNoteDeck{100, 10}}}}}
	for i := 0; i < len(bns.robberAccounts); i++ {
		robberAccount1 := bns.robberAccounts[i]
		robberAccount2 := expectedBns.robberAccounts[i]
		for j := 0; j < len(bnp.bankNoteDecks); j++ {
			bankNoteDeck1 := robberAccount1.bankNoteDecks[j]
			bankNoteDeck2 := robberAccount2.bankNoteDecks[j]
			if (bankNoteDeck1.faceValue != bankNoteDeck2.faceValue) {
				t.Log("face value mismatched")
				t.Fail()
			}

			if (bankNoteDeck1.quantity != bankNoteDeck2.quantity) {
				t.Log("quantity mismatched")
				t.Fail()
			}
		}
	}
}