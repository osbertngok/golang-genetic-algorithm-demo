package main

import (
	"fmt"
	"io"
)

func readRobberShare(reader io.Reader, bnp *BankNoteProblem) error {
	var robberCount int
	// Get robber Count as the first cin
	if _, err := fmt.Fscanf(reader, "%d\n", &robberCount); err != nil {
		return err
	}
	// Optimistically ignore err - demo only
	bnp.robberShare = make([]int, robberCount)
	// Read the share for each robber
	for i := 0; i < robberCount; i++ {
		if _, err := fmt.Fscanf(reader, "%d", &bnp.robberShare[i]); err != nil {
			return err
		}
	}
	return nil
}

func readBankNoteDecks(reader io.Reader, bnp *BankNoteProblem) error {
	var bankNoteDeckCount int
	// Get no. of bank note deck
	if _, err := fmt.Fscanf(reader, "%d\n", &bankNoteDeckCount); err != nil {
		return err
	}
	bnp.bankNoteDecks = make([]BankNoteDeck, bankNoteDeckCount)
	// Assuming the first value is the face value, the second is the quantity
	for i := 0; i < bankNoteDeckCount; i++ {
		if _, err := fmt.Fscanf(reader, "%f %d\n", &bnp.bankNoteDecks[i].faceValue, &bnp.bankNoteDecks[i].quantity); err != nil {
			return err
		}
	}
	return nil
}
