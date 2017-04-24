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

func main() {
    fmt.Println("hello world")
}