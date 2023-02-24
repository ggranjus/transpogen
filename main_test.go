package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrganismSize(t *testing.T) {
	size := 13
	organism := NewOrganism([]byte("HELLO WORLD !"), size)
	assert.Equal(t, size, len(organism.DNA))
}

func TestNewPopulationSize(t *testing.T) {
	blockSize := 13
	cypher := addPadding([]byte("test test test"), blockSize)
	population := NewPopulation(cypher, 10, blockSize)
	assert.Equal(t, 10, len(population))
	for _, organism := range population {
		assert.Equal(t, blockSize, len(organism.DNA))
	}
}

func TestPopulationBestIndividual(t *testing.T) {
	population := NewPopulation([]byte("DH OE!ROLW LL"), 10, 13)
	best := population.GetBest()
	for _, organism := range population {
		assert.True(t, best.Fitness >= organism.Fitness)
	}
	assert.Greater(t, best.Fitness, float64(0.0))
}

func TestPermute(t *testing.T) {
	o := Organism{
		DNA: []int{10, 0, 11, 7, 1, 12, 8, 4, 3, 6, 5, 2, 9},
	}
	cyphertext := o.GetSolution([]byte("HELLO WORLD !"))
	assert.Equal(t, "DH OE!ROLW LL", cyphertext)
	o = Organism{
		DNA: []int{1, 4, 11, 12, 3, 2, 9, 7, 6, 8, 0, 10, 5},
	}
	plain := o.GetSolution([]byte(cyphertext))
	assert.Equal(t, "HELLO WORLD !", plain)
}

func TestFit(t *testing.T) {
	badOrganism := Organism{DNA: []int{10, 0, 11, 7, 1, 12, 8, 4, 3, 6, 5, 2, 9}}
	badOrganism.calcFitness([]byte("HELLO WORLD !"))

	goodOrganism := Organism{DNA: []int{1, 4, 11, 12, 3, 2, 9, 7, 6, 8, 0, 10, 5}}
	goodOrganism.calcFitness([]byte("DH OE!ROLW LL"))

	assert.Greater(t, goodOrganism.Fitness, badOrganism.Fitness)
}

func TestIntegration(t *testing.T) {
	blockSize := 13
	cypher, _ := os.ReadFile("cipher.txt")
	cypher = addPadding(cypher, blockSize)
	assert.Equal(t, 0, len(cypher)%blockSize)

	popSize := 200
	population := NewPopulation(cypher, popSize, blockSize)
	assert.Equal(t, popSize, len(population))

	for _, organism := range population {
		assert.Equal(t, blockSize, len(organism.DNA))
	}
}

func TestCrossover(t *testing.T) {
	blockSize := 13
	parent1 := Organism{DNA: []int{1, 4, 11, 12, 3, 2, 9, 7, 6, 8, 0, 10, 5}}
	parent2 := Organism{DNA: []int{10, 0, 11, 7, 1, 12, 8, 4, 3, 6, 5, 2, 9}}

	child := crossover(parent1, parent2)
	assert.Equal(t, blockSize, len(child.DNA))
	for i := 0; i < blockSize; i++ {
		assert.Contains(t, child.DNA, i)
	}
}
