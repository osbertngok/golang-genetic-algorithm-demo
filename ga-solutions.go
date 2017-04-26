package main

import "math"
import "math/rand"
import "time"

func bankNoteSolutionClone(bns *BankNoteSolution)(BankNoteSolution) {
	robberAccounts := make([]RobberAccount, len(bns.robberAccounts))
	for _, robberAccount := range bns.robberAccounts {
		bankNoteDecks := make([]BankNoteDeck, len(robberAccount.bankNoteDecks))
		copy(bankNoteDecks, robberAccount.bankNoteDecks)
	}
	return BankNoteSolution{robberAccounts}
}

func mutateBankNoteSolution(bns *BankNoteSolution) {
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

func evaluateBankNoteSolution(bns *BankNoteSolution)(float32) {
	return 0
}

func getGeneticAlgorithmSolution(bnp *BankNoteProblem)(BankNoteSolution) {
	solution := getDefaultSolution(bnp)
	return solution
}