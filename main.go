package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func subMain() {
	var bnp BankNoteProblem
	input, err := os.Open("./data/bnp1.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	readRobberShare(input, &bnp)
	readBankNoteDecks(input, &bnp)
	_ = bnp.validate()
	config := GeneticAlgorithmConfig{
		maxGenerationCount:         20,
		maxCandidateCount:          100,
		noOfMutantForEachCandidate: 100,
		maxMutateCount:             10,
		maxAttemptCount:            100,
		intensityFunc: func(generationCount, maxGenerationCount int) float64 {
			return 1.0
			/*
				intensity := 1.0 - float64(generationCount) / float64(maxGenerationCount)
				if intensity < 0.1 {
					intensity = 0.1
				}
				return intensity
			*/
		}}
	bns, err := bnp.getGeneticAlgorithmSolution(config)
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	if err = bns.validate(&bnp); err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	fmt.Println(bns)
	if bns.String() != testsbns[2].String() {
		fmt.Print(err)
		os.Exit(-1)
	}
}

func profileWrapper(func1 func()) {

	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		func1()
		defer pprof.StopCPUProfile()
	} else {
		func1()
	}
}

func main() {
	profileWrapper(subMain)
}
