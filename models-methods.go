package main

import ("math"
	"math/rand"
	"time"
	"errors"
	"fmt"
)

// BankNoteProblem methods
func (bnp *BankNoteProblem) validate() error {
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

// BankNoteSolution methods

func (bns *BankNoteSolution) clone() BankNoteSolution {
	robberAccounts := make([]RobberAccount, len(bns.robberAccounts))
	for index, robberAccount := range bns.robberAccounts {
		bankNoteDecks := make([]BankNoteDeck, len(robberAccount.bankNoteDecks))
		copy(bankNoteDecks, robberAccount.bankNoteDecks)
		robberAccounts[index].bankNoteDecks = bankNoteDecks
	}
	return BankNoteSolution{robberAccounts}
}

func (bns *BankNoteSolution) hashCode() int64 {
	return 0
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
 		moveQuantity := 1
 		intnParameter := int(math.Floor(float64(remaining) * intensity))
 		if intnParameter == 0 {
 			moveQuantity = 1
		} else {
			moveQuantity = rand.Intn(moveQuantity)
	 		if moveQuantity < 1 {
	 			moveQuantity = 1
			}
		}
 		

 		bns.robberAccounts[selectedRobber1].bankNoteDecks[selectedBankNoteDeck].quantity -= moveQuantity
 		bns.robberAccounts[selectedRobber2].bankNoteDecks[selectedBankNoteDeck].quantity += moveQuantity
 		mutateCount++
	}
}

func (bns *BankNoteSolution) validate(bnp *BankNoteProblem) error {
	// No. of robber check
	noOfRobbersInBnp := len(bnp.robberShare)
	noOfBankNoteDecksInBnp := len(bnp.bankNoteDecks)

	if len(bns.robberAccounts) != noOfRobbersInBnp {
		bnpstr := fmt.Sprint(*bnp)
		bnsstr := fmt.Sprint(*bns)
		return fmt.Errorf("robber count mismatch, bnp: %d, bns: %d, p: %s, s: %s", noOfRobbersInBnp, len(bns.robberAccounts), bnpstr, bnsstr)
	}

	for index, element := range bns.robberAccounts {
		if len(element.bankNoteDecks) != noOfBankNoteDecksInBnp {
			bnpstr := fmt.Sprint(*bnp)
			bnsstr := fmt.Sprint(*bns)
			return fmt.Errorf("banknotedeck count mismatch, index: %d, bnp: %d, bns: %d, p: %s, s: %s", index, noOfBankNoteDecksInBnp, len(element.bankNoteDecks), bnpstr, bnsstr)
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
