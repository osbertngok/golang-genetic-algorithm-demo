package main
import "fmt"

type bankNoteDeck struct {
	faceValue float32
	quantity int
}

type bankNoteProblem struct {
	robberShare []int
	bankNoteDecks []bankNoteDeck
}

type robberAccount struct {
	bankNoteDecks []bankNoteDeck
}

type bankNoteSolution struct {
	robberAccounts []robberAccount
}

func readRobberShare(bnp *bankNoteProblem) {
	var robberCount int
	// Get robber Count as the first cin
    _, err := fmt.Scanf("%d", &robberCount)
    // Optimistically ignore err - demo only
    bnp.robberShare = make([]int, robberCount)
    // Read the share for each robber
    for i := 0; i < robberCount; i++ {
    	_, err := fmt.Scanf("%d", &bnp.robberShare[i])
    }
}

func readBankNoteDecks(bnp *bankNoteProblem) {
	var bankNoteDeckCount int
	// Get no. of bank note deck
	_, err := fmt.Scanf("%d", &bankNoteDeckCount)
	bnp.bankNoteDecks = make([]bankNoteDeck, bankNoteDeckCount)
	// Assuming the first value is the face value, the second is the quantity
	for i := 0; i < bankNoteDeckCount; i++ {
		_, err := fmt.Scanf("%f %d", bnp.bankNoteDecks[i].faceValue, bankNoteDecks[i].quantity)
	}
}

func main() {
    fmt.Println("hello world")
}