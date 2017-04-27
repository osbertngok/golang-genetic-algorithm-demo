package main

func main() {
	var bnp BankNoteProblem
	// input
	readRobberShare(&bnp)
	readBankNoteDecks(&bnp)
	_ = bnp.validate()
	// output
}
