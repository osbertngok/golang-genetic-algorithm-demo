package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func getReaderFromString(str string) io.Reader {
	buffer := bytes.NewBufferString(str)
	return buffer
}

func Test_bnp_reader(t *testing.T) {
	var bnp BankNoteProblem
	inputShareString := testsbnpshareinput[0]
	inputBndString := testsbnpbndinput[0]
	reader := getReaderFromString(inputShareString)
	if err := readRobberShare(reader, &bnp); err != nil {
		t.Log(err)
		t.Fail()
	}
	reader = getReaderFromString(inputBndString)
	if err := readBankNoteDecks(reader, &bnp); err != nil {
		t.Log(err)
		t.Fail()
	}

	if bnp.String() != testsbnp[0].String() {
		fmt.Println(bnp)
		fmt.Println(testsbnp[0])
		t.Log("bnp not equal")
		t.Fail()
	}
}
