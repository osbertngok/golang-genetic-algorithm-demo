package main

import (
	"fmt"
	"testing"
)

var TESTSBNP = [3]BankNoteProblem{
	{
		[]int{
			10,
			20,
			30},
		[]BankNoteDeck{
			{1, 5},
			{5, 10},
			{10, 10},
			{20, 20},
			{50, 5},
			{100, 10}}},
	{
		[]int{
			10,
			20,
			30,
			40},
		[]BankNoteDeck{
			{1, 20},
			{72, 50},
			{100, 30}}},
	{
		[]int{
			13,
			15,
			18,
			20},
		[]BankNoteDeck{
			{2.56, 20},
			{3.72, 20},
			{4.55, 26}}}}

var TESTSBNS = [3]BankNoteSolution{

	{[]RobberAccount{
		{
			[]BankNoteDeck{
				{1, 1},
				{5, 0},
				{10, 5},
				{20, 0},
				{50, 3},
				{100, 1}}},
		{
			[]BankNoteDeck{
				{1, 2},
				{5, 8},
				{10, 4},
				{20, 1},
				{50, 0},
				{100, 5}}},
		{
			[]BankNoteDeck{
				{1, 2},
				{5, 2},
				{10, 1},
				{20, 19},
				{50, 2},
				{100, 4}}}}},
	{[]RobberAccount{
		{
			[]BankNoteDeck{
				{1, 2},
				{72, 5},
				{100, 3}}},
		{
			[]BankNoteDeck{
				{1, 4},
				{72, 10},
				{100, 6}}},
		{
			[]BankNoteDeck{
				{1, 6},
				{72, 15},
				{100, 9}}},
		{
			[]BankNoteDeck{
				{1, 8},
				{72, 20},
				{100, 12}}}}},
	{[]RobberAccount{
		{
			[]BankNoteDeck{
				{2.56, 1},
				{3.72, 11},
				{4.55, 1}}},
		{
			[]BankNoteDeck{
				{2.56, 6},
				{3.72, 1},
				{4.55, 8}}},
		{
			[]BankNoteDeck{
				{2.56, 6},
				{3.72, 4},
				{4.55, 8}}},
		{
			[]BankNoteDeck{
				{2.56, 7},
				{3.72, 4},
				{4.55, 9}}}}}}

func Test_validateBankNoteProblem_success(t *testing.T) {
	var bnp BankNoteProblem
	bnp.robberShare = []int{
		10,
		20,
		30}
	bnp.bankNoteDecks = []BankNoteDeck{
		{1, 5},
		{5, 10},
		{10, 10},
		{20, 20},
		{50, 5},
		{100, 10}}
	if err := bnp.validate(); err != nil {
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
		{1, 5},
		{5, 10},
		{10, 10},
		{20, 20},
		{50, 5},
		{100, 10}}
	bnp.bankNoteDecks = bankNoteDecksArray[:]
	if err := bnp.validate(); err != nil {
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
		{
			[]BankNoteDeck{
				{1, 5},
				{5, 5},
				{10, 0},
				{20, 0},
				{50, 0},
				{100, 0}}},
		{
			[]BankNoteDeck{
				{1, 0},
				{5, 5},
				{10, 10},
				{20, 5},
				{50, 0},
				{100, 0}}},
		{
			[]BankNoteDeck{
				{1, 0},
				{5, 0},
				{10, 0},
				{20, 15},
				{50, 5},
				{100, 10}}}}}
	if fmt.Sprint(bns) != fmt.Sprint(expectedBns) {
		t.Fail()
	}

	if err := bns.validate(&bnp); err != nil {
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

	if err = bns.validate(&bnp); err != nil {
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
	if err := newBns.validate(&bnp); err != nil {
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
