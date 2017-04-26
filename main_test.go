package main

import (
	"testing"
)

var TESTSBNP = [1]BankNoteProblem{
	BankNoteProblem{
		[]int{
			10,
			20,
			30},
		[]BankNoteDeck{
			BankNoteDeck{1, 5},
			BankNoteDeck{5, 10},
			BankNoteDeck{10, 10},
			BankNoteDeck{20, 20},
			BankNoteDeck{50, 5},
			BankNoteDeck{100, 10}}}}

var TESTSBNS = [1]BankNoteSolution{
	BankNoteSolution{[]RobberAccount{
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 0},
				BankNoteDeck{5, 5},
				BankNoteDeck{10, 0},
				BankNoteDeck{20, 0},
				BankNoteDeck{50, 0},
				BankNoteDeck{100, 0}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 5},
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
				BankNoteDeck{100, 10}}}}}}

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
	err := bnp.validate()
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
	err := bnp.validate()
	if err != nil {
		t.Log("Good!")
	} else {
		t.Log("Expected to be error!")
		t.Fail()
	}
}

func Test_DefaultSolution_1(t *testing.T) {
	bnp := TESTSBNP[0]
	bns := bnp.getDefaultSolution()
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
	validateBnsEqual(t, &bnp, &bns, &expectedBns)

	err := bns.validate(&bnp)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func Test_GASolution_1(t *testing.T) {
	bnp := TESTSBNP[0]
	bns := bnp.getGeneticAlgorithmSolution()
	validateBnsEqual(t, &bnp, &bns, &TESTSBNS[0])
}