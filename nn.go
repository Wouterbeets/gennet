package gennet

func newNn(dna dna) *nn {
}

type nn struct {
	inp   []input
	neurs map[int]*neuron
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

func newNN(nbIn, nbOut, maxSize int) *nn {
	n := new(nn)
	n.inp = make([]input, nbIn)
	for i := 0; i < nbIn; i++ {
		neur := newNeuron()
		n.inp = append(n.inp, neur.inp)
		n.neurs = append(n.neurs, neur)
	}
	n.out = make(output, nbOut)
	for i := 0; i < nbOut; i++ {
		n.out[i] = make(input)
		neur := newNeuron()
		neur.addOut(n.out[i])
		n.neurs = append(n.neurs, neur)
	}
}

func (n *nn) addGene(g gene) {
	rec, ok := n.neurs[g.receiver()]
	if !ok {
		rec = &neuron{
			inp:     make(input),
			weights: make(weight),
		}
		n.neurs[g.receiver()] = rec
	}
	sen, ok := n.neurs[g.sender()]
	if !ok {
		rec = &neuron{
			inp:     make(input),
			weights: make(weight),
		}
		n.neurs[g.sender()] = sen
	}
	sen.out = append(sen.out, rec.inp)
	rec.weights[g.sender()] = weight{gene.weight, gene.bias}
	rec.nbInputs++
}

func (n *nn) addDNA(dna) {
	for _, g := range dna {
		n.addGene(g)
	}
}
