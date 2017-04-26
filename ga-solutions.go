package main

import "math"
import "math/rand"
import "time"
import "errors"

func (bns *BankNoteSolution) clone() BankNoteSolution {
	robberAccounts := make([]RobberAccount, len(bns.robberAccounts))
	for _, robberAccount := range bns.robberAccounts {
		bankNoteDecks := make([]BankNoteDeck, len(robberAccount.bankNoteDecks))
		copy(bankNoteDecks, robberAccount.bankNoteDecks)
	}
	return BankNoteSolution{robberAccounts}
}

func (bns *BankNoteSolution) mutate() {
	rand.Seed(time.Now().UnixNano())
	maxMutateCount := 10
	maxAttemptCount := 100
	intensity := 0.1
	noOfRobbers := len(bns.robberAccounts)
	noOfBankNoteDecks := len(bns.robberAccounts[0].bankNoteDecks)
	// Edge cases
	if noOfRobbers < 2 {
		return
	}

	if noOfBankNoteDecks < 2 {
		return
	}

	for mutateCount, attemptCount := 0, 0; mutateCount < maxMutateCount && attemptCount < maxAttemptCount; attemptCount++ {
		// Pick two random robbers and a random deck
 		selectedRobber1 := rand.Intn(noOfRobbers)
 		selectedRobber2 := selectedRobber1
 		for ;selectedRobber2 != selectedRobber1; {
 			selectedRobber2 = rand.Intn(noOfRobbers)
 		}

 		selectedBankNoteDeck := rand.Intn(noOfBankNoteDecks)
 		remaining := bns.robberAccounts[selectedRobber1].bankNoteDecks[selectedBankNoteDeck].quantity
 		if remaining <= 0 {
 			continue
 		}

 		// Move at least 1 from Robber1 account to Robber2 account

 		moveQuantity := rand.Intn(int(math.Floor(float64(remaining) * intensity)))
 		if moveQuantity < 1 {
 			moveQuantity = 1
 		}

 		bns.robberAccounts[selectedRobber1].bankNoteDecks[selectedBankNoteDeck].quantity -= moveQuantity
 		bns.robberAccounts[selectedRobber2].bankNoteDecks[selectedBankNoteDeck].quantity += moveQuantity
 		mutateCount++
	}
}

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

func (bnp *BankNoteProblem) getGeneticAlgorithmSolution() BankNoteSolution {
	solution := bnp.getDefaultSolution()
	return solution
}