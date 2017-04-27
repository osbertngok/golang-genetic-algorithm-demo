package main

import "math"
import "errors"
import "sort"
import "fmt"

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
		sumOfAverageFaceValueDifferenceSquare += math.Pow(averageFaceValue - totalAverageFaceValue, 2.0)
		count++
	}

	return sumOfAverageFaceValueDifferenceSquare, nil
}

func getMutatedBankNoteSolution(original *BankNoteSolution, noOfMutants int) ([]BankNoteSolution, error) {
	if noOfMutants <= 1 {
		return nil, errors.New("invalid no of mutants")
	}

	bankNoteSolutions := make([]BankNoteSolution, noOfMutants)
	for i := 0; i < noOfMutants; i++ {
		bankNoteSolution := original.clone()
		bankNoteSolution.mutate()
		bankNoteSolutions[i] = bankNoteSolution
	}
	return bankNoteSolutions, nil
}

// TODO: getMutatedBankNoteSolution using crossing over

type LessFunc func(*BankNoteSolution, *BankNoteSolution) bool

type bankNoteSolutionSorter struct {
	bankNoteSolutions []BankNoteSolution
	by LessFunc
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

func (bnp *BankNoteProblem) getGeneticAlgorithmSolution() BankNoteSolution {
	initialSolution := bnp.getDefaultSolution()
	maxGenerationCount := 100
	maxCandidateCount := 100
	noOfMutantForEachCandidate := 10
	candidateSolutionPool := make([]BankNoteSolution, 1)
	candidateSolutionPool[0] = initialSolution
	for generationCount := 0 ;; generationCount++ {
		fmt.Println(candidateSolutionPool[0])
		fmt.Println("")
		noOfCandidates := len(candidateSolutionPool)

		// Populate the offspring slice
		offspringSolutionPool := make([]BankNoteSolution, noOfCandidates, noOfCandidates * (noOfMutantForEachCandidate + 1))
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

				mutant.mutate()

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
		if generationCount >= maxGenerationCount {
			return offspringSolutionPool[0]
		}
		
		// Grab the best of them for next round
		nextGenerationCandidateCount := maxCandidateCount
		if length := len(offspringSolutionPool); length < maxCandidateCount {
			nextGenerationCandidateCount = length
		}

		candidateSolutionPool = make([]BankNoteSolution, nextGenerationCandidateCount)
		copy(candidateSolutionPool, offspringSolutionPool[0 : nextGenerationCandidateCount - 1])
	}
	// Not gonna happen
	return BankNoteSolution{}
}
