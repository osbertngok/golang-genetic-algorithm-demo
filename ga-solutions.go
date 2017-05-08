package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GeneticAlgorithmConfig includes all available parameters to be used in func getGeneticAlgorithmSolution
type GeneticAlgorithmConfig struct {
	maxGenerationCount         int
	maxCandidateCount          int
	noOfMutantForEachCandidate int
	maxMutateCount             int
	maxAttemptCount            int
	intensityFunc              func(generationCount, maxGenerationCount int) float64
}

func populateOffspringSolutionPool(candidateSolutionPool *[]BankNoteSolution, config GeneticAlgorithmConfig, generationCount int) ([]BankNoteSolution, error) {
	intensity := config.intensityFunc(generationCount, config.maxGenerationCount)
	mutate := mutateFuncGenerator(config.maxMutateCount, config.maxAttemptCount)
	noOfCandidates := len(*candidateSolutionPool)
	offspringSolutionPool := make([]BankNoteSolution, noOfCandidates, noOfCandidates*(config.noOfMutantForEachCandidate+1))
	copy(offspringSolutionPool, *candidateSolutionPool)

	// Reduce duplicates
	hashCodeMap := make(map[string]bool, 0)
	for i := 0; i < noOfCandidates; i++ {
		hashCode := offspringSolutionPool[i].String()
		if _, ok := hashCodeMap[hashCode]; ok {
			// already exists, next; NOT EXPECTED TO BE RUN
			continue
		}
		hashCodeMap[hashCode] = true
	}
	for i := 0; i < noOfCandidates; i++ {
		for j := 0; j < config.noOfMutantForEachCandidate; j++ {
			mutant := (*candidateSolutionPool)[i].clone()

			if err := mutate(&mutant, intensity); err != nil {
				return nil, err
			}

			hashCode := mutant.String()
			if _, ok := hashCodeMap[hashCode]; ok {
				// already exists, next;
				continue
			}
			hashCodeMap[hashCode] = true
			offspringSolutionPool = append(offspringSolutionPool, mutant)
		}
	}
	return offspringSolutionPool, nil
}

func getAndPrintScore(bnp *BankNoteProblem, offspringSolutionPool *[]BankNoteSolution, generationCount int) float64 {
	// <for-information-only>
	score, _ := bnp.evaluate(&(*offspringSolutionPool)[0])
	fmt.Printf("Generation: %d, Score: %f\n", generationCount, score)
	// </for-information-only>
	return score
}

func (bnp *BankNoteProblem) getGeneticAlgorithmSolution(config GeneticAlgorithmConfig) (BankNoteSolution, error) {
	rand.Seed(time.Now().UnixNano())
	initialSolution := bnp.getDefaultSolution()
	candidateSolutionPool := make([]BankNoteSolution, 1)
	candidateSolutionPool[0] = initialSolution
	for generationCount := 0; ; generationCount++ {

		// Populate the offspring slice
		offspringSolutionPool, _ := populateOffspringSolutionPool(&candidateSolutionPool, config, generationCount)

		sortBankNoteSolutionByEvaluationFunc(bnp, offspringSolutionPool, (*BankNoteProblem).evaluate)

		if score := getAndPrintScore(bnp, &offspringSolutionPool, generationCount); generationCount >= config.maxGenerationCount || score == 0 {
			return offspringSolutionPool[0], nil
		}

		// Grab the best of them for next round
		nextGenerationCandidateCount := min(config.maxCandidateCount, len(offspringSolutionPool))
		candidateSolutionPool = make([]BankNoteSolution, nextGenerationCandidateCount)
		copy(candidateSolutionPool, offspringSolutionPool[0:nextGenerationCandidateCount])
	}
}
