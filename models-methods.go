package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
)

// BankNoteDeck methods

func (bnd *BankNoteDeck) clone() BankNoteDeck {
	return BankNoteDeck{bnd.faceValue, bnd.quantity}
}

func (bnd BankNoteDeck) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	buffer.WriteString(fmt.Sprintf("%f", bnd.faceValue))
	buffer.WriteString(", ")
	buffer.WriteString(fmt.Sprintf("%d", bnd.quantity))
	buffer.WriteString("}")
	return buffer.String()
}

// BankNoteProblem methods

func (bnp BankNoteProblem) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	if bnp.robberShare != nil {
		buffer.WriteString("[")
		buffer.WriteString("]")
	} else {
		buffer.WriteString("nil")
	}
	buffer.WriteString(", ")
	if bnp.bankNoteDecks != nil {
		buffer.WriteString("[")
		bndCount := len(bnp.bankNoteDecks)
		for index, item := range bnp.bankNoteDecks {
			if item != nil {
				buffer.WriteString(item.String())
			} else {
				buffer.WriteString("nil")
			}
			if index < bndCount-1 {
				buffer.WriteString(", ")
			}
		}
		buffer.WriteString("]")
	} else {
		buffer.WriteString("nil")
	}
	buffer.WriteString("}")
	return buffer.String()
}

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

// RobberAccount methods

func (ra RobberAccount) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	if ra.bankNoteDecks != nil {
		buffer.WriteString("[")
		bndCount := len(ra.bankNoteDecks)
		for index, item := range ra.bankNoteDecks {
			if item != nil {
				buffer.WriteString(item.String())
			} else {
				buffer.WriteString("nil")
			}
			if index < bndCount-1 {
				buffer.WriteString(", ")
			}
		}
		buffer.WriteString("]")
	} else {
		buffer.WriteString("nil")
	}
	buffer.WriteString("}")
	return buffer.String()
}

// BankNoteSolution methods

func (bns *BankNoteSolution) clone() BankNoteSolution {
	robberAccounts := make([]*RobberAccount, len(bns.robberAccounts))
	for index, robberAccount := range bns.robberAccounts {
		robberAccounts[index] = &RobberAccount{}
		bankNoteDecks := make([]*BankNoteDeck, len(robberAccount.bankNoteDecks))
		for idx, bankNoteDeck := range robberAccount.bankNoteDecks {
			var tmpBankNoteDeck = bankNoteDeck.clone()
			bankNoteDecks[idx] = &tmpBankNoteDeck
		}
		robberAccounts[index].bankNoteDecks = bankNoteDecks
	}
	return BankNoteSolution{robberAccounts}
}

func (bns BankNoteSolution) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	if bns.robberAccounts != nil {
		buffer.WriteString("[")
		raCount := len(bns.robberAccounts)
		for index, item := range bns.robberAccounts {
			if item != nil {
				buffer.WriteString(item.String())
			} else {
				buffer.WriteString("nil")
			}
			if index < raCount-1 {
				buffer.WriteString(", ")
			}
		}
		buffer.WriteString("]")
	} else {
		buffer.WriteString("nil")
	}
	buffer.WriteString("}")
	return buffer.String()
}

func mutateFuncGenerator(maxMutateCount, maxAttemptCount int) func(*BankNoteSolution, float64) error {
	return func(bns *BankNoteSolution, intensity float64) error {
		// Edge cases
		noOfRobbers := len(bns.robberAccounts)
		if noOfRobbers < 2 {
			return fmt.Errorf("Only %d robber accounts found (<2)", noOfRobbers)
		}

		noOfBankNoteDecks := len(bns.robberAccounts[0].bankNoteDecks)
		if noOfBankNoteDecks < 2 {
			return fmt.Errorf("Only %d bank note decks found (<2)", noOfBankNoteDecks)
		}

		// Delegated to closure
		// maxMutateCount := 10
		// maxAttemptCount := 100

		for mutateCount, attemptCount := 0, 0; mutateCount < maxMutateCount && attemptCount < maxAttemptCount; attemptCount++ {
			// Pick two random robbers and a random deck
			selectedRobber1 := rand.Intn(noOfRobbers)
			selectedRobber2 := selectedRobber1
			for selectedRobber2 == selectedRobber1 {
				selectedRobber2 = rand.Intn(noOfRobbers)
			}

			selectedBankNoteDeck1 := rand.Intn(noOfBankNoteDecks)
			selectedBankNoteDeck2 := rand.Intn(noOfBankNoteDecks)
			for selectedBankNoteDeck2 == selectedBankNoteDeck1 {
				selectedBankNoteDeck2 = rand.Intn(noOfBankNoteDecks)
			}

			remaining1 := bns.robberAccounts[selectedRobber1].bankNoteDecks[selectedBankNoteDeck1].quantity
			remaining2 := bns.robberAccounts[selectedRobber2].bankNoteDecks[selectedBankNoteDeck2].quantity
			remaining := remaining1
			if remaining2 < remaining {
				remaining = remaining2
			}
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

			bns.robberAccounts[selectedRobber1].bankNoteDecks[selectedBankNoteDeck1].quantity -= moveQuantity
			bns.robberAccounts[selectedRobber2].bankNoteDecks[selectedBankNoteDeck1].quantity += moveQuantity
			bns.robberAccounts[selectedRobber2].bankNoteDecks[selectedBankNoteDeck2].quantity -= moveQuantity
			bns.robberAccounts[selectedRobber1].bankNoteDecks[selectedBankNoteDeck2].quantity += moveQuantity
			mutateCount++
		}
		return nil
	}
}

func (bns *BankNoteSolution) validate(bnp *BankNoteProblem) error {
	// No. of robber check
	noOfRobbersInBnp := len(bnp.robberShare)
	noOfBankNoteDecksInBnp := len(bnp.bankNoteDecks)

	if len(bns.robberAccounts) != noOfRobbersInBnp {
		bnpstr := bnp.String()
		bnsstr := bns.String()
		return fmt.Errorf("robber count mismatch, bnp: %d, bns: %d, p: %s, s: %s", noOfRobbersInBnp, len(bns.robberAccounts), bnpstr, bnsstr)
	}

	for index, element := range bns.robberAccounts {
		if len(element.bankNoteDecks) != noOfBankNoteDecksInBnp {
			bnpstr := bnp.String()
			bnsstr := bns.String()
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
			bnpstr := bnp.String()
			bnsstr := bns.String()
			return fmt.Errorf("robber share mismatch, index: %d, bnp: %d, bns: %d, p: %s, s: %s", i, bnp.robberShare[i], quantity, bnpstr, bnsstr)
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
