package gennet

import (
	"fmt"
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

func Test_add_gene(t *testing.T) {
	nn := newNN(2, 2, 6)
	nn.addGene(gene{2, 3, 1, 1})
	nn.addGene(gene{3, 4, 1, 1})
	nn.In([]float64{0, 1})
	out := nn.Out()
	if len(out) != 2 {
		t.Fail()
	}
	fmt.Println("out", out)
}

//3-1:1.12445:1.3234
//4-1
//2-3
//2-4
//3-4
//3-6
//4-5
//4-6
