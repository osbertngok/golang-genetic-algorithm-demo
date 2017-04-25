package main

func getDefaultSolution(bnp *BankNoteProblem)(BankNoteSolution) {
	var bns BankNoteSolution
	var robberPointer int
	var robberRemaining int
	var deckPointer int
	var deckRemaining int
	// Initilization
	bns.robberAccounts = make([]RobberAccount, len(bnp.robberShare))
	for index := 0; index < len(bns.robberAccounts); index++ {
		bns.robberAccounts[index] = RobberAccount{make([]BankNoteDeck, len(bnp.bankNoteDecks))}
		for index2, element := range bnp.bankNoteDecks {
			bns.robberAccounts[index].bankNoteDecks[index2] = BankNoteDeck{element.faceValue, 0}
		}
	}
	robberPointer = 0
	deckPointer = 0
	robberRemaining = bnp.robberShare[robberPointer]
	deckRemaining = bnp.bankNoteDecks[deckPointer].quantity
	for robberPointer < len(bnp.robberShare) {
		robberEmptyFlag := false
		deckEmptyFlag := false
		// Use pointer because it is a struct
		deck := &bns.robberAccounts[robberPointer].bankNoteDecks[deckPointer]
		if robberRemaining > deckRemaining {
			deck.quantity = deckRemaining
			robberRemaining -= deckRemaining
			deckRemaining = 0
			// rollover
			deckEmptyFlag = true
		} else if deckRemaining > robberRemaining {
			deck.quantity = robberRemaining
			deckRemaining -= robberRemaining
			robberRemaining = 0
			// rollover
			robberEmptyFlag = true
		} else {
			deck.quantity = deckRemaining
			deckRemaining = 0
			robberRemaining = 0
			deckEmptyFlag = true
			robberEmptyFlag = true
		}

		if deckEmptyFlag {
			deckEmptyFlag = true
			deckPointer++
			if deckPointer < len(bnp.bankNoteDecks) {
				deckRemaining = bnp.bankNoteDecks[deckPointer].quantity
			}
		}

		if robberEmptyFlag {
			robberPointer++
			if robberPointer < len(bnp.robberShare) {
				robberRemaining = bnp.robberShare[robberPointer]
			}
		}
	}
	return bns
}

func getGeneticAlgorithmSolution(bnp *BankNoteProblem)(BankNoteSolution) {
	return BankNoteSolution{}
}