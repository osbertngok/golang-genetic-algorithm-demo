package main

import (
	"fmt"
	"testing"
)

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
	bnp := testsbnp[0]
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
	bnp := testsbnp[2]
	config := GeneticAlgorithmConfig{
		maxGenerationCount:         20,
		maxCandidateCount:          100,
		noOfMutantForEachCandidate: 100,
		maxMutateCount:             10,
		maxAttemptCount:            100,
		intensityFunc: func(generationCount, maxGenerationCount int) float64 {
			return 1.0
			/*
				intensity := 1.0 - float64(generationCount) / float64(maxGenerationCount)
				if intensity < 0.1 {
					intensity = 0.1
				}
				return intensity
			*/
		}}
	bns, err := bnp.getGeneticAlgorithmSolution(config)
	if err != nil {
		t.Log(err)
		t.Fatal()
	}

	if err = bns.validate(&bnp); err != nil {
		t.Log(err)
		t.Fatal()
	}
	fmt.Println(bns)
	if fmt.Sprint(bns) != fmt.Sprint(testsbns[2]) {
		t.Fail()
	}
}

func Test_Mutate(t *testing.T) {
	bnp := testsbnp[0]
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
	bnp := testsbnp[0]
	bns := bnp.getDefaultSolution()
	newBns := bns.clone()
	if fmt.Sprint(bns) != fmt.Sprint(newBns) {
		t.Log("bns not equal")
		t.Fail()
	}
}
