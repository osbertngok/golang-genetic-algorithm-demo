package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func averageFaceValue(bankNoteDecks *[]BankNoteDeck) (float64, error) {
	sumOfCashValue := 0.0
	sumOfQuantity := 0
	for _, deck := range *bankNoteDecks {
		sumOfCashValue += deck.faceValue * float64(deck.quantity)
		sumOfQuantity += deck.quantity
	}
	if sumOfQuantity == 0 {
		return 0, errors.New("no bank note in deck")
	}
	return sumOfCashValue / float64(sumOfQuantity), nil
}

func (bnp *BankNoteProblem) evaluate(bns *BankNoteSolution) (float64, error) {
	sumOfAverageFaceValueDifferenceSquare := 0.0
	count := 0
	totalAverageFaceValue, err := averageFaceValue(&bnp.bankNoteDecks)
	if err != nil {
		return 0.0, err
	}

	for _, robberAccount := range bns.robberAccounts {
		averageFaceValue, err := averageFaceValue(&robberAccount.bankNoteDecks)
		if err != nil {
			continue
		}
		sumOfAverageFaceValueDifferenceSquare += math.Pow(averageFaceValue-totalAverageFaceValue, 2.0)
		count++
	}

	return sumOfAverageFaceValueDifferenceSquare, nil
}

type LessFunc func(*BankNoteSolution, *BankNoteSolution) bool

type bankNoteSolutionSorter struct {
	bankNoteSolutions []BankNoteSolution
	by                LessFunc
}

func (by LessFunc) Sort(bankNoteSolutions []BankNoteSolution) {
	ss := &bankNoteSolutionSorter{bankNoteSolutions, by}
	sort.Sort(ss)
}

// Need to define Len, Swap and Less func

func (s *bankNoteSolutionSorter) Len() int {
	return len(s.bankNoteSolutions)
}

func (s *bankNoteSolutionSorter) Swap(i, j int) {
	// Swapping pointer
	s.bankNoteSolutions[i], s.bankNoteSolutions[j] = s.bankNoteSolutions[j], s.bankNoteSolutions[i]
}

func (s *bankNoteSolutionSorter) Less(i, j int) bool {
	return s.by(&s.bankNoteSolutions[i], &s.bankNoteSolutions[j])
}

func sortBankNoteSolutionByEvaluationFunc(bankNoteProblem *BankNoteProblem, bankNoteSolutions []BankNoteSolution, evalFunc func(*BankNoteProblem, *BankNoteSolution) (float64, error)) {
	// Make the Less function using evalFunc and closure

	// It is essential that lessFunc is LessFunc, not func(s1, s2 *BankNoteSolution) bool, although we know logically they are the same
	// LessFunc is namedType while func(s1, s2 *BankNoteSolution) bool is unnamed type
	// but they are assignable to each other
	// We define the method Sort on the namedType, so lessFunc must be LessFunc
	// otherwise it would be an unnamed type and will not be able to find its named type method *Sort*
	var lessFunc LessFunc
	lessFunc = func(s1, s2 *BankNoteSolution) bool {
		// The less the better
		result1, err1 := evalFunc(bankNoteProblem, s1)
		result2, err2 := evalFunc(bankNoteProblem, s2)
		if err1 != nil {
			// we will prefer solution 2 in this case
			return false
		}

		if err2 != nil {
			return true
		}
		return result1 < result2
	}

	lessFunc.Sort(bankNoteSolutions)
}

func (bnp *BankNoteProblem) getGeneticAlgorithmSolution() (BankNoteSolution, error) {
	rand.Seed(time.Now().UnixNano())
	initialSolution := bnp.getDefaultSolution()

	// <Parameters>
	maxGenerationCount := 20
	maxCandidateCount := 200
	noOfMutantForEachCandidate := 200
	intensityFunc := func(generationCount, maxGenerationCount int) float64 {
		return 1.0
		/*
			intensity := 1.0 - float64(generationCount) / float64(maxGenerationCount)
			if intensity < 0.1 {
				intensity = 0.1
			}
			return intensity
		*/
	}
	maxMutateCount := 10
	maxAttemptCount := 100
	mutate := mutateFuncGenerator(maxMutateCount, maxAttemptCount)
	// </Parameters>

	candidateSolutionPool := make([]BankNoteSolution, 1)
	candidateSolutionPool[0] = initialSolution
	for generationCount := 0; ; generationCount++ {
		noOfCandidates := len(candidateSolutionPool)
		intensity := intensityFunc(generationCount, maxGenerationCount)

		// Populate the offspring slice
		offspringSolutionPool := make([]BankNoteSolution, noOfCandidates, noOfCandidates*(noOfMutantForEachCandidate+1))
		copy(offspringSolutionPool, candidateSolutionPool)

		// Reduce duplicates
		hashCodeMap := make(map[string]bool, 0)
		for i := 0; i < noOfCandidates; i++ {
			hashCode := fmt.Sprint(offspringSolutionPool[i])
			if _, ok := hashCodeMap[hashCode]; ok {
				// already exists, next; NOT EXPECTED TO BE RUN
				continue
			}
			hashCodeMap[hashCode] = true
		}
		for i := 0; i < noOfCandidates; i++ {
			for j := 0; j < noOfMutantForEachCandidate; j++ {
				mutant := candidateSolutionPool[i].clone()

				err := mutate(&mutant, intensity)
				if err != nil {
					return BankNoteSolution{}, err
				}

				hashCode := fmt.Sprint(mutant)
				if _, ok := hashCodeMap[hashCode]; ok {
					// already exists, next;
					continue
				}
				hashCodeMap[hashCode] = true
				offspringSolutionPool = append(offspringSolutionPool, mutant)
			}
		}

		sortBankNoteSolutionByEvaluationFunc(bnp, offspringSolutionPool, (*BankNoteProblem).evaluate)

		score, _ := bnp.evaluate(&offspringSolutionPool[0])
		fmt.Printf("Generation: %d, Score: %f\n", generationCount, score)
		if generationCount >= maxGenerationCount || score == 0 {
			return offspringSolutionPool[0], nil
		}

		// Grab the best of them for next round
		nextGenerationCandidateCount := maxCandidateCount
		length := len(offspringSolutionPool)
		if length < maxCandidateCount {
			nextGenerationCandidateCount = length
		}

		candidateSolutionPool = make([]BankNoteSolution, nextGenerationCandidateCount)
		copy(candidateSolutionPool, offspringSolutionPool[0:nextGenerationCandidateCount])
	}
	// Not gonna happen
	return BankNoteSolution{}, errors.New("unknown error")
}
