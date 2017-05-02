package main

import "os"

func main() {
	var bnp BankNoteProblem
	input := os.Stdin
	readRobberShare(input, &bnp)
	readBankNoteDecks(input, &bnp)
	_ = bnp.validate()
	// output
}
