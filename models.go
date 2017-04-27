package main

// BankNoteDeck is a model that represents a pile of cash.
// All the bank notes in the deck should share the same faceValue.
// Quantity of the struct indicates how many bank notes are there.
type BankNoteDeck struct {
	faceValue float64
	quantity  int
}

// BankNoteProblem defines the bank note quantity share of the robbers,
// as well as the available bank note decks.
type BankNoteProblem struct {
	robberShare   []int
	bankNoteDecks []BankNoteDeck
}

// RobberAccount is simply a pile of BankNoteDecks;
// Each of the BankNoteDeck should have different face value of bank notes.
type RobberAccount struct {
	bankNoteDecks []BankNoteDeck
}

// BankNoteSolution is a solution to its corresponding BankNoteProblem.
type BankNoteSolution struct {
	robberAccounts []RobberAccount
}
