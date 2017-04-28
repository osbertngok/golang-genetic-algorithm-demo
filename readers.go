package main

import (
	"fmt"
	"os"
)

func readRobberShare(bnp *BankNoteProblem) {
	var robberCount int
	// Get robber Count as the first cin
	if _, err := fmt.Scanf("%d", &robberCount); err != nil {
		os.Exit(1)
	}
	// Optimistically ignore err - demo only
	bnp.robberShare = make([]int, robberCount)
	// Read the share for each robber
	for i := 0; i < robberCount; i++ {
		if _, err := fmt.Scanf("%d", &bnp.robberShare[i]); err != nil {
			os.Exit(1)
		}
	}
}

func readBankNoteDecks(bnp *BankNoteProblem) {
	var bankNoteDeckCount int
	// Get no. of bank note deck
	if _, err := fmt.Scanf("%d", &bankNoteDeckCount); err != nil {
		os.Exit(1)
	}
	bnp.bankNoteDecks = make([]BankNoteDeck, bankNoteDeckCount)
	// Assuming the first value is the face value, the second is the quantity
	for i := 0; i < bankNoteDeckCount; i++ {
		if _, err := fmt.Scanf("%f %d", bnp.bankNoteDecks[i].faceValue, bnp.bankNoteDecks[i].quantity); err != nil {
			os.Exit(1)
		}
	}
}
