package gennet

func newNn(dna dna) *nn {
}

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
