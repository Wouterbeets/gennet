package gennet

import (
	"fmt"
)

type nn struct {
	inp   []input
	neurs map[int]*neuron
	out   output
}

func (n *nn) In(input []float64) {
	for i, val := range input {
		fmt.Println("len inp", len(n.inp), i, n.inp)
		n.inp[i] <- signal{val: val, neuronID: i}
	}
}

func (n *nn) Out() []float64 {
	out := []float64{}
	for _, outChan := range n.out {
		out = append(out, (<-outChan).val)
	}
	return out
}

func newNN(nbIn, nbOut, maxSize int) *nn {
	n := new(nn)
	n.inp = make([]input, 0, 2)
	n.neurs = make(map[int]*neuron)
	middleNeur := newNeuron(nbIn)
	n.neurs[nbIn] = middleNeur
	for i := 0; i < nbIn; i++ {
		fmt.Println("adding neur", i)
		neur := newNeuron(i)
		neur.addOut(middleNeur.inp)
		n.inp = append(n.inp, neur.inp)
		n.neurs[i] = neur
	}
	n.out = make(output, nbOut)
	for i := 0; i < nbOut; i++ {
		fmt.Println(i)
		n.out[i] = make(input)
		neur := newNeuron(maxSize - (i + 1))
		neur.addOut(n.out[i])
		middleNeur.addOut(neur.inp)
		n.neurs[maxSize-(i+1)] = neur
	}
	for _, neurs := range n.neurs {
		go neurs.live()
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
		rec = &neuron{
			inp:     make(input),
			weights: newWeights(),
			id:      g.sender(),
		}
		n.neurs[g.sender()] = sen
	}
	sen.out = append(sen.out, rec.inp)
	rec.weights[g.sender()] = weight{g.weight(), g.bias()}
}

func (n *nn) addDNA(d dna) {
	for _, g := range d {
		n.addGene(g)
	}
}
