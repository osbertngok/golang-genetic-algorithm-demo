package main
import (
	"fmt"
	"os"
	"errors"
)

func readRobberShare(bnp *BankNoteProblem) {
	var robberCount int
	// Get robber Count as the first cin
    _, err := fmt.Scanf("%d", &robberCount)
    if err != nil {
    	os.Exit(1)
    }
    // Optimistically ignore err - demo only
    bnp.robberShare = make([]int, robberCount)
    // Read the share for each robber
    for i := 0; i < robberCount; i++ {
    	_, err := fmt.Scanf("%d", &bnp.robberShare[i])
		if err != nil {
		    	os.Exit(1)
		    }
    }
}

func readBankNoteDecks(bnp *BankNoteProblem) {
	var bankNoteDeckCount int
	// Get no. of bank note deck
	_, err := fmt.Scanf("%d", &bankNoteDeckCount)
	if err != nil {
    	os.Exit(1)
    }
	bnp.bankNoteDecks = make([]BankNoteDeck, bankNoteDeckCount)
	// Assuming the first value is the face value, the second is the quantity
	for i := 0; i < bankNoteDeckCount; i++ {
		_, err := fmt.Scanf("%f %d", bnp.bankNoteDecks[i].faceValue, bnp.bankNoteDecks[i].quantity)
		if err != nil {
	    	os.Exit(1)
	    }
	}
}

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