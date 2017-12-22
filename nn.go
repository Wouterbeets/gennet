package gennet

import "fmt"

type nn struct {
	inp   []input
	neurs map[int]*neuron
	out   output
}

func (n *nn) In(input []float64) {
	for _, neur := range n.neurs {
		go neur.live()
	}
	for i, val := range input {
		n.inp[i] <- signal{val: val, neuronID: -1}
	}
}

func (n *nn) Out() []float64 {
	out := []float64{}
	for _, outChan := range n.out {
		out = append(out, (<-outChan).val)
	}
	return out
}

func newNN(nbIn, nbOut, maxSize int, d ...dna) *nn {
	n := new(nn)
	n.inp = make([]input, 0, 2)
	n.neurs = make(map[int]*neuron)
	for i := 0; i < nbIn; i++ {
		fmt.Println("adding neur", i)
		neur := newNeuron(i)
		neur.weights[-1] = weight{1, 0.5}
		n.inp = append(n.inp, neur.inp)
		n.neurs[i] = neur
	}

	n.out = make(output, nbOut)
	for i := 0; i < nbOut; i++ {
		id := maxSize - (i + 1)
		fmt.Println("adding", id)
		n.out[i] = make(input)
		neur := newNeuron(id)
		neur.addOut(n.out[i])
		n.neurs[id] = neur
	}
	if len(d) == 1 {
		n.addDNA(d[0])
	}
	fmt.Println("neural", n)
	return n
}

func (n *nn) String() string {
	s := ""
	for id, neur := range n.neurs {
		s += fmt.Sprintln(id, neur.weights)
	}
	return s
}

func (n *nn) DNA() (d dna) {
	for _, neur := range n.neurs {
		if neur.id >= len(n.inp) {
			d = append(d, neur.genes()...)
		}
	}
	d.sort()
	return d
}

func (n *nn) addGene(g gene) {
	rec, ok := n.neurs[g.receiver()]
	if !ok {
		rec = &neuron{
			inp:     make(input),
			weights: newWeights(),
			id:      g.receiver(),
		}
		n.neurs[g.receiver()] = rec
	}
	sen, ok := n.neurs[g.sender()]
	if !ok {
		sen = &neuron{
			inp:     make(input),
			weights: newWeights(),
			id:      g.sender(),
		}
		n.neurs[g.sender()] = sen
	}
	fmt.Println(sen)
	sen.out = append(sen.out, rec.inp)
	rec.weights[g.sender()] = weight{g.weight(), g.bias()}
}

func (n *nn) addDNA(d dna) {
	for _, g := range d {
		n.addGene(g)
	}
}
