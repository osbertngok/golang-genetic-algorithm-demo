package main
import (
	"errors"
)

func validateBankNoteProblem(bnp *BankNoteProblem)(error) {
	sumOfRobberShare := 0
	for _, element := range bnp.robberShare {
		sumOfRobberShare += element
	}
	sumOfBankNotes := 0
	for _, element := range bnp.bankNoteDecks {
		sumOfBankNotes += element.quantity
	}

	if sumOfRobberShare != sumOfBankNotes {
		return errors.New("quantity does not match")
	}

	return nil
}

func getDefaultSolution(bnp *BankNoteProblem)(BankNoteSolution) {
	var bns BankNoteSolution
	var robberPointer int
	var robberRemaining int
	var deckPointer int
	var deckRemaining int
	// Initilization
	bns.robberAccounts = make([]RobberAccount, len(bnp.robberShare))
	for index, _ := range bns.robberAccounts {
		bns.robberAccounts[index] = RobberAccount{make([]BankNoteDeck, len(bnp.bankNoteDecks))}
		for index2, _ := range bns.robberAccounts[index].bankNoteDecks {
			bns.robberAccounts[index].bankNoteDecks[index2] = BankNoteDeck{0.0, 0}
		}
	}
	robberPointer = 0
	deckPointer = 0
	robberRemaining = bnp.robberShare[robberPointer]
	deckRemaining = bnp.bankNoteDecks[deckPointer].quantity
	for robberPointer < len(bnp.robberShare) {
		robberEmptyFlag := false
		deckEmptyFlag := false
		deck := bns.robberAccounts[robberPointer].bankNoteDecks[deckPointer]
		deck.faceValue = bnp.bankNoteDecks[deckPointer].faceValue
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

func main() {
    var bnp BankNoteProblem
    // input
    readRobberShare(&bnp)
    readBankNoteDecks(&bnp)
    _ = validateBankNoteProblem(&bnp)
    // output
}