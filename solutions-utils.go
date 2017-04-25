package main

import (
	"testing"
)

func validateBnsEqual(t *testing.T, bnp *BankNoteProblem, bns1 *BankNoteSolution, bns2 *BankNoteSolution) {
	for i := 0; i < len(bns1.robberAccounts); i++ {
		robberAccount1 := bns1.robberAccounts[i]
		robberAccount2 := bns2.robberAccounts[i]
		for j := 0; j < len(bnp.bankNoteDecks); j++ {
			bankNoteDeck1 := robberAccount1.bankNoteDecks[j]
			bankNoteDeck2 := robberAccount2.bankNoteDecks[j]
			if (bankNoteDeck1.faceValue != bankNoteDeck2.faceValue) {
				t.Logf("face value mismatched, %d:%d - %f, %f", i, j, bankNoteDeck1.faceValue, bankNoteDeck2.faceValue)
				t.Fail()
			}

			if (bankNoteDeck1.quantity != bankNoteDeck2.quantity) {
				t.Logf("quantity mismatched, %d:%d - %d, %d", i, j, bankNoteDeck1.quantity, bankNoteDeck2.quantity)
				t.Fail()
			}
		}
	}
}