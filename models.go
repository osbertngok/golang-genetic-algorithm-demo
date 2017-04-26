package main

type BankNoteDeck struct {
	faceValue float64
	quantity int
}

type BankNoteProblem struct {
	robberShare []int
	bankNoteDecks []BankNoteDeck
}

type RobberAccount struct {
	bankNoteDecks []BankNoteDeck
}

type BankNoteSolution struct {
	robberAccounts []RobberAccount
}