package main

import ()

func main() {
	nn := &nn{
		inp: make(input),
	}

}
func NewNn() *nn {
	n := &nn{
		inp: make(input),
	}
	n.neurs = []neuron{
		{
			inp:     n.inp,
			weights: newWeights(),
		}, {
			inp:     n.inp,
			weights: newWeights(),
		},
	}
	n.output = make(output, 1)
	n.output[0] = make(input)
	outneurs := []neurons{
		{
			inp:     make(input),
			output:  output{n.output[0]},
			weights: newWeights(),
		}, {
			inp:     make(input),
			output:  output{n.output[0]},
			weights: newWeights(),
		},
	}
}

type NN interface {
	In([]float64) error
	Out() []float64
}

type weight struct {
	weight float64
	bias   float64
}

type neuron struct {
	inp     input
	weights map[int]weight
	out     output
}

type output []input

type signal struct {
	val      float64
	neuronID int
}

type input chan signal

type nn struct {
	inp   input
	neurs []neuron
	out   output
}

func (n *nn) In(input []float64) error {
	for i, val := range input {
		n.inp <- signal{val: val, neuronID: i}
	}
}

func (n *nn) Out() []float64 {
	out := []float64{}
	for i, outChan := range n.output {
		out := append(out, (<-outChan).val)
	}
	return out
}
