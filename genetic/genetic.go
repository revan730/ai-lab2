package main

import (
	"math"
)
import "math/rand"
import "time"
import "fmt"

type boardPosition struct {
	sequence []int
	fitness int
	survival float64
}

const MUTATION_CHANCE = 0.00001

func makeUnique(a []int, b []int) []int {

	check := make(map[int]int)
	d := append(a, b...)
	res := make([]int,0)
	for _, val := range d {
		check[val] = 1
	}

	for letter := range check {
		res = append(res,letter)
	}

	return res
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func fitness(chromosome []int) int {
	clashes := 0
	diff := float64(len(chromosome) - len(makeUnique(chromosome, chromosome)))
	rowColClashes := math.Abs(diff)
	clashes += int(rowColClashes)

	for i := 0;i < len(chromosome);i++ {
		for j := 0;j < len(chromosome);j++ {
			if i != j {
				dx := math.Abs(float64(i - j))
				dy := math.Abs(float64(chromosome[i] - chromosome[j]))
				if dx == dy {
					clashes++
				}
			}	
		}
	}

	return 28 - clashes
}

func generateChromosome() []int {
	chromosome := []int{0,1,2,3,4,5,6,7}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(chromosome), func(i, j int) { chromosome[i], chromosome[j] = chromosome[j], chromosome[i] })
	return chromosome
}

func generatePopulation(size int) []boardPosition {
	population := make([]boardPosition, size, size)
	for i := 0;i < size;i++ {
		population[i].sequence = generateChromosome()
		population[i].fitness = fitness(population[i].sequence)
	}
	return population
}

func generatePopulation1(size int) []boardPosition {
	population := make([]boardPosition, 0)
	for i := 0;i < size; {
		var genome boardPosition
		genome.sequence = generateChromosome()
		genome.fitness = fitness(genome.sequence)
		for j := 0;j < len(population);j++ {
			if Equal(population[j].sequence, genome.sequence) {
				continue
			}
		}
		population = append(population, genome)
		i++
	}
	return population
}

func getSequnceBelowSurvival(population []boardPosition, survival float64) []boardPosition {
	seq := make([]boardPosition, 0)
	for i := 0;i < len(population);i++ {
		if population[i].survival <= survival {
			seq = append(seq, population[i])
		}
	}

	return seq
}

func getParents(population []boardPosition) (boardPosition, boardPosition) {
	summationFitness := 0.0
	var parent1, parent2 boardPosition
	for i := 0;i < len(population);i++ {
		summationFitness += float64(population[i].fitness)
	}

	for i := 0;i < len(population);i++ {
		population[i].survival = float64(population[i].fitness) / summationFitness
	}

	for {
		parent1Random := rand.Float64()
		parent1Seq := getSequnceBelowSurvival(population, parent1Random)
		if len(parent1Seq) == 0 {
			continue
		}
		parent1 = parent1Seq[0]
		if len(parent1.sequence) == 0 {
			continue
		}
		break
	}

	for {
		parent2Random := rand.Float64()
		parent2Seq := getSequnceBelowSurvival(population, parent2Random)
		if len(parent2Seq) == 0 {
			continue
		}

		index := rand.Intn(len(parent2Seq))
		parent2 = parent2Seq[index]
		if Equal(parent1.sequence, parent2.sequence) == false {
			break
		}
	}
	return parent1, parent2
}

func mutate(child boardPosition) boardPosition {
	chance := rand.Float64()
	if chance < MUTATION_CHANCE {
		index1 := rand.Intn(len(child.sequence))
		index2 := rand.Intn(len(child.sequence))
		newSeq := make([]int, len(child.sequence))
		copy(newSeq, child.sequence)
		mutated := boardPosition{
			fitness:  child.fitness,
			survival: child.survival,
			sequence: newSeq,
		}
		tmp := mutated.sequence[index1]
		mutated.sequence[index1] = mutated.sequence[index2]
		mutated.sequence[index2] = tmp
		mutated.fitness = fitness(mutated.sequence)
		return mutated
	}
	return child
}

func reproduceCrossover(parent1 boardPosition, parent2 boardPosition) boardPosition {
	child := boardPosition{}
	childSequence := make([]int, 0)
	index := rand.Intn(len(parent1.sequence))
	childSequence = append(childSequence, parent1.sequence[0:index]...)
	childSequence = append(childSequence, parent2.sequence[index:]...)
	child.sequence = childSequence
	child.fitness = fitness(childSequence)
	return child
}

func geneticAlgorithm(population []boardPosition) []boardPosition {
	newPopulation := make([]boardPosition, 0)
	for i := 0;i < len(population);i++ {
		parent1, parent2 := getParents(population)
		child := reproduceCrossover(parent1, parent2)
		child = mutate(child)
		newPopulation = append(newPopulation, child)
	}
	return newPopulation
}

func main() {
	controlSeq := []int{7,3,0,2,5,1,6,4}
	fmt.Printf("Test subject fitness %d\n", fitness(controlSeq))
	population := generatePopulation1(2500)
	for i := 0;i < 1000;i++ {
		population = geneticAlgorithm(population)
		for j := 0;j < len(population);j++ {
			if population[j].fitness == 28 {
				fmt.Printf("fitness %d sequence %v\n", population[j].fitness, population[j].sequence)
				//os.Exit(1)
			}
		}
	}
}