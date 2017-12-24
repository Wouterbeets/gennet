package gennet

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_newNN(t *testing.T) {
	nn := newNN(2, 2, 6)
	nn.In([]float64{0, 1})
	out := nn.Out()
	if len(out) != 2 {
		t.Fail()
	}
	fmt.Println("out", out)
}

func Test_newNN_with_dna(t *testing.T) {
	nn := newNN(2, 2, 6, dna{gene{1, 2, 1, 1}, gene{2, 4, 1, 1}})
	nn.In([]float64{0, 1})
	out := nn.Out()
	if len(out) != 2 {
		t.Fail()
	}
	fmt.Println("out", out)
}

func Test_add_gene(t *testing.T) {
	nn := newNN(2, 2, 6)
	nn.addGene(gene{1, 2, 1, 1})
	nn.addGene(gene{2, 3, 1, 1})
	nn.In([]float64{0, 1})
	out := nn.Out()
	if len(out) != 2 {
		t.Fail()
	}
	fmt.Println("out", out)
	fmt.Println(nn.DNA())
}

func Test_DNA(t *testing.T) {
	nn := newNN(2, 2, 6)
	d := nn.DNA()
	fmt.Println(d)
}

func Test_mutate(t *testing.T) {
	nn := newNN(2, 2, 6)
	fmt.Println("before", nn.DNA())
	nn.Mutate(rand.New(rand.NewSource(1)))
	fmt.Println("after", nn.DNA())
}

func Test_mem(t *testing.T) {
	nn := newNN(2, 1, 5)
	nn.addGene(gene{2, 4, 1, 1})
	nn.addGene(gene{4, 1, 1, 1})
	nn.In([]float64{0, 1})
	out := nn.Out()
	fmt.Println("out", out)
	fmt.Println(nn.DNA())
}
