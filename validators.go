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

func validateBankNoteSolution(bnp *BankNoteProblem, bns *BankNoteSolution)(error) {
	// No. of robber check
	noOfRobbersInBnp := len(bnp.robberShare)
	noOfBankNoteDecksInBnp := len(bnp.bankNoteDecks)

	if len(bns.robberAccounts) != noOfRobbersInBnp {
		return errors.New("robber count mismatch")
	}

	for _, element := range bns.robberAccounts {
		if len(element.bankNoteDecks) != noOfBankNoteDecksInBnp {
			return errors.New("banknotedeck count mismatch")
		}
	}

	// face value check
	for i, ra := range bns.robberAccounts {
		quantity := 0
		for j, bnd := range ra.bankNoteDecks {
			if bnd.faceValue != bnp.bankNoteDecks[j].faceValue {
				return errors.New("faceValue mismatch")
			}
			if bnd.quantity < 0 {
				return errors.New("share less than 0")
			}
			quantity += bnd.quantity
		}
		if quantity != bnp.robberShare[i] {
			return errors.New("robber share mismatch")
		}
	}

	for j := 0; j < len(bnp.bankNoteDecks); j++ {
		quantity := 0
		for _, ra := range bns.robberAccounts {
			quantity += ra.bankNoteDecks[j].quantity
		}
		if quantity != bnp.bankNoteDecks[j].quantity {
			return errors.New("banknotedeck quantity mismatch")
		}
	}
	return nil
}