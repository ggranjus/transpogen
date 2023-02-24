package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

type Organism struct {
	DNA     []int
	Fitness float64
}

func NewOrganism(target []byte, size int) (organism Organism) {
	seq := make([]int, size)
	for i := 0; i < size; i++ {
		seq[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(seq), func(i, j int) { seq[i], seq[j] = seq[j], seq[i] })
	organism = Organism{
		DNA:     seq,
		Fitness: 0,
	}
	organism.calcFitness(target)
	return
}

type Indicator struct {
	Pattern     string
	Coefficient float64
}

func (o *Organism) calcFitness(target []byte) {
	indicators := []Indicator{
		{" FLAG ", 60},
		{"FLAG", 50},
		{"THE", 22.5},
		{"AND", 9.0},
		{"ING", 6.6},
		{"ENT", 5.4},
		{"ION", 4.5},
		{"NTH", 4.2},
		{"TER", 3.9},
		{"INT", 3.9},
		{"OFT", 3.9},
		{"THA", 3.9},
		{"ERE", 3.9},
		{"TIO", 3.6},
		{"HER", 3.6},
		{"FTH", 3.6},
		{"ETH", 3.3},
		{"ATI", 3.3},
		{"HAT", 3},
		{"ATE", 3},
		{"STH", 3},
		{"EST", 3},
		{"TH", 3.2},
		{"HE", 2.6},
		{"IN", 2.2},
		{"ER", 1.9},
		{"AN", 1.8},
		{"RE", 1.5},
		{"ES", 1.4},
		{"ON", 1.4},
		{"ST", 1.4},
		{"NT", 1.3},
		{"EN", 1.3},
		{"ED", 1.3},
		{"ND", 1.2},
		{"AT", 1.2},
		{"TI", 1.2},
		{"TE", 1.1},
		{"OR", 1.1},
		{"AR", 1},
		{"HA", 1},
		{"OF", 1},
		{"TX", -1.0},
		{"TZ", -1.0},
		{"OZ", -1.0},
		{"IJ", -1.0},
		{"IY", -1.0},
	}
	solution := o.GetSolution(target)
	var score float64
	for _, indicator := range indicators {
		score += indicator.Coefficient * float64(strings.Count(string(solution), indicator.Pattern))
	}
	o.Fitness = score
}

func (o *Organism) size() int {
	return len(o.DNA)
}

func (o *Organism) mutate() {
	for i := 0; i < len(o.DNA); i++ {
		if rand.Float64() < 0.01 {
			j := rand.Intn(len(o.DNA))
			o.DNA[i], o.DNA[j] = o.DNA[j], o.DNA[i]
		}
	}
}

func permuteBlock(block []byte, permutation []int) (result []byte) {
	for i := 0; i < len(block); i++ {
		result = append(result, block[permutation[i]])
	}
	return
}

func (o *Organism) GetSolution(cypher []byte) string {
	solution := []byte{}
	for i := 0; i < len(cypher); i += o.size() {
		block := cypher[i : i+o.size()]
		solution = append(solution, permuteBlock(block, o.DNA)...)
	}
	return string(solution)
}

type Population []Organism

func (p Population) GetBest() Organism {
	sort.Slice(p, func(i, j int) bool {
		return p[i].Fitness > p[j].Fitness
	})
	return p[0]
}

func NewPopulation(target []byte, size, organismSize int) (population Population) {
	population = make([]Organism, size)
	for i := 0; i < size; i++ {
		population[i] = NewOrganism(target, organismSize)
	}
	sort.Slice(population, func(i, j int) bool {
		return population[i].Fitness > population[j].Fitness
	})
	return population
}

func createPool(population []Organism, target []byte, maxFitness float64) (pool []Organism) {
	pool = make([]Organism, 0)
	for i := 0; i < len(population); i++ {
		population[i].calcFitness(target)
		num := int((population[i].Fitness / maxFitness) * 100)
		for n := 0; n < num; n++ {
			pool = append(pool, population[i])
		}
	}
	return
}

func (p Population) NaturalSelection(target []byte) Population {
	pool := createPool(p, target, p.GetBest().Fitness)
	next := make([]Organism, len(p))

	for i := 0; i < len(p); i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := pool[r1]
		b := pool[r2]

		child := crossover(a, b)
		child.mutate()
		child.calcFitness(target)

		next[i] = child
	}

	// Ensure that Best Organism is always in the next generation
	next[rand.Intn(len(next))] = p.GetBest()

	return next
}

func crossover(o1, o2 Organism) Organism {
	child := Organism{
		DNA:     []int{},
		Fitness: 0,
	}
	mid := rand.Intn(len(o1.DNA))
	for i := 0; i < mid; i++ {
		child.DNA = append(child.DNA, o1.DNA[i])
	}
	for _, j := range o2.DNA {
		if !contains(child.DNA, j) {
			child.DNA = append(child.DNA, j)
		}
	}
	return child
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func addPadding(s []byte, blockSize int) []byte {
	for len(s)%blockSize != 0 {
		s = append(s, ' ')
	}
	return s
}

func main() {
	cypher, err := os.ReadFile("cipher.txt")
	if err != nil {
		panic(err)
	}

	blockSize := 13
	cypher = addPadding(cypher, blockSize)

	start := time.Now()
	population := NewPopulation(cypher, 100, blockSize)

	var bestSolution string

	for generation := 1; generation <= 50; generation++ {
		bestOrganism := population.GetBest()
		bestSolution = bestOrganism.GetSolution(cypher)
		fmt.Printf("\rGeneration: %d | %s | fitness: %.2f", generation, bestSolution[:blockSize*2], bestOrganism.Fitness)
		population = population.NaturalSelection(cypher)
	}

	elapsed := time.Since(start)
	fmt.Printf("\nTime elapsed: %s\n", elapsed)

	err = os.WriteFile("result.txt", []byte(bestSolution), 0644)
	if err != nil {
		panic(err)
	}
}
