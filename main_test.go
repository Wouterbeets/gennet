package gennet

import (
	"fmt"
	"testing"

	"github.com/MaxHalford/gago"
)

//func Test_newNN(t *testing.T) {
//	nn := newNN(2, 2, 6)
//	nn.In([]float64{0, 1})
//	out, _ := nn.Out()
//	if len(out) != 2 {
//		t.Fail()
//	}
//}
//
//func Test_newNN_with_dna(t *testing.T) {
//	nn := newNN(2, 2, 6, dna{gene{1, 2, 1, 1}, gene{0, 2, 1, 1}, gene{2, 4, 1, 1}, gene{2, 5, 1, 1}})
//	nn.In([]float64{0, 1})
//	out, _ := nn.Out()
//	if len(out) != 2 {
//		t.Fail()
//	}
//}
//
//func Test_add_gene(t *testing.T) {
//	nn := newNN(2, 2, 6)
//	nn.addGene(gene{1, 2, 1, 1})
//	nn.addGene(gene{2, 3, 1, 1})
//	nn.In([]float64{0, 1})
//	out, _ := nn.Out()
//	if len(out) != 2 {
//		t.Fail()
//	}
//}
//
//func Test_DNA(t *testing.T) {
//	nn := newNN(2, 2, 6)
//	_ = nn.DNA()
//}
//
//func Test_mutate(t *testing.T) {
//	nn := newNN(2, 2, 6)
//	d := nn.DNA()
//	nn.Mutate(rand.New(rand.NewSource(1)))
//	d2 := nn.DNA()
//	if dnaEqual(d, d2) {
//		t.Error("dna equal")
//	}
//}
//
//func dnaEqual(d, d2 dna) bool {
//	equal := true
//	for i := range d {
//		for j := range d[i] {
//			if d[i][j] != d2[i][j] {
//				equal = false
//			}
//		}
//	}
//	return equal
//}
//
//func Test_crossover(t *testing.T) {
//	n := newNN(2, 2, 6)
//	ndna := n.DNA()
//	rng := rand.New(rand.NewSource(0))
//	n.Mutate(rng)
//	n2 := newNN(2, 2, 6)
//	n2dna := n2.DNA()
//	n.Crossover(gago.Genome(n2), rng)
//	ndna2 := n.DNA()
//	n2dna2 := n2.DNA()
//	if dnaEqual(ndna, ndna2) {
//		t.Error("dna equal")
//	}
//	if dnaEqual(n2dna, n2dna2) {
//		t.Error("dna2 equal")
//	}
//}

func Test_eval(t *testing.T) {
	model := gago.Generational(makeGenomeMaker(2, 1, 10, dna{
		gene{0, 9, 0.5, 0.1},
		gene{1, 9, 0.5, 0.1},
	}))
	model.Initialize()
	fmt.Println("start", model.Populations[0].Individuals[0].Genome.(*Nn).DNA())
	for i := 0; i < 10000; i++ {
		model.Evolve()
		fmt.Printf("Best fitness at generation %d: %f\n", i, model.HallOfFame[0].Fitness)
		fmt.Println(model.HallOfFame[0].Genome.(*Nn).DNA())
	}
}
