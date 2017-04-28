package main

import (
	"errors"
	"math"
	"sort"
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
	totalAverageFaceValue, _ := averageFaceValue(&bnp.bankNoteDecks)

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

// LessFunc is a func type that takes two pointers of BankNoteSolution and
// returns a bool that indicates whether the evaluation result
// of the first solution is less than the second one.
type LessFunc func(*BankNoteSolution, *BankNoteSolution) bool

type bankNoteSolutionSorter struct {
	bankNoteSolutions []BankNoteSolution
	by                LessFunc
}

// Sort is a LessFunc method that uses the LessFunc
// to sort a BankNoteSolution slice.
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
