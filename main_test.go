package main

import (
	"testing"
	"fmt"
)

var TESTSBNP = [3]BankNoteProblem{
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
			BankNoteDeck{100, 10}}},
	BankNoteProblem{
		[]int{
			10,
			20,
			30,
			40},
		[]BankNoteDeck{
			BankNoteDeck{1, 20},
			BankNoteDeck{72, 50},
			BankNoteDeck{100, 30}}},
	BankNoteProblem{
		[]int{
			13,
			15,
			18,
			20},
		[]BankNoteDeck{
			BankNoteDeck{2.56, 20},
			BankNoteDeck{3.72, 20},
			BankNoteDeck{4.55, 26}}}}

var TESTSBNS = [3]BankNoteSolution{
	
	BankNoteSolution{[]RobberAccount{
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 1},
				BankNoteDeck{5, 0},
				BankNoteDeck{10, 5},
				BankNoteDeck{20, 0},
				BankNoteDeck{50, 3},
				BankNoteDeck{100, 1}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 2},
				BankNoteDeck{5, 8},
				BankNoteDeck{10, 4},
				BankNoteDeck{20, 1},
				BankNoteDeck{50, 0},
				BankNoteDeck{100, 5}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 2},
				BankNoteDeck{5, 2},
				BankNoteDeck{10, 1},
				BankNoteDeck{20, 19},
				BankNoteDeck{50, 2},
				BankNoteDeck{100, 4}}}}},
	BankNoteSolution{[]RobberAccount{
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 2},
				BankNoteDeck{72, 5},
				BankNoteDeck{100, 3}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 4},
				BankNoteDeck{72, 10},
				BankNoteDeck{100, 6}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 6},
				BankNoteDeck{72, 15},
				BankNoteDeck{100, 9}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{1, 8},
				BankNoteDeck{72, 20},
				BankNoteDeck{100, 12}}}}},
		BankNoteSolution{[]RobberAccount{
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{2.56, 1},
				BankNoteDeck{3.72, 11},
				BankNoteDeck{4.55, 1}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{2.56, 6},
				BankNoteDeck{3.72, 1},
				BankNoteDeck{4.55, 8}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{2.56, 6},
				BankNoteDeck{3.72, 4},
				BankNoteDeck{4.55, 8}}},
		RobberAccount{
			[]BankNoteDeck{
				BankNoteDeck{2.56, 7},
				BankNoteDeck{3.72, 4},
				BankNoteDeck{4.55, 9}}}}}}

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
	if fmt.Sprint(bns) != fmt.Sprint(expectedBns) {
		t.Fail()
	}

	err := bns.validate(&bnp)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}


func Test_GASolution_1(t *testing.T) {
	bnp := TESTSBNP[2]
	bns, err := bnp.getGeneticAlgorithmSolution()
	if err != nil {
		t.Log(err)
		t.Fatal()
	}

	err = bns.validate(&bnp)
	if err != nil {
		t.Log(err)
		t.Fatal()
	}
	fmt.Println(bns)
	if fmt.Sprint(bns) != fmt.Sprint(TESTSBNS[2]) {
		t.Fail()
	}
}


func Test_Mutate(t *testing.T) {
	bnp := TESTSBNP[0]
	bns := bnp.getDefaultSolution()
	newBns := bns.clone()
	maxMutateCount := 10
	maxAttemptCount := 100
	mutateFuncGenerator(maxMutateCount, maxAttemptCount)(&newBns, 1.0)
	err := newBns.validate(&bnp)
	if err != nil {
		t.Log(err)
		t.Fatal()
	}

	if fmt.Sprint(bns) == fmt.Sprint(newBns) {
		t.Log("not expected to be equal")
		t.Fail()
	}
}

func Test_Clone(t *testing.T) {
	bnp := TESTSBNP[0]
	bns := bnp.getDefaultSolution()
	newBns := bns.clone()
	if fmt.Sprint(bns) != fmt.Sprint(newBns) {
		t.Fail()
	}
}