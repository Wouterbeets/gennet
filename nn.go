package gennet

import (
	"fmt"
	"github.com/MaxHalford/gago"
	"math/rand"
)

type Nn struct {
	inp     []input
	neurs   map[int]*neuron
	out     output
	maxSize int
	Eval func () float64
}

func (n *Nn) In(input []float64) {
	for _, neur := range n.neurs {
		go neur.live()
	}
	for i, val := range input {
		n.inp[i] <- signal{val: val, neuronID: -1}
	}
}

func (n *Nn) Out() []float64 {
	out := []float64{}
	for _, outChan := range n.out {
		out = append(out, (<-outChan).val)
	}
	return out
}

func newNN(nbIn, nbOut, maxSize int, d ...dna) *Nn {
	n := new(Nn)
	n.maxSize = maxSize
	n.inp = make([]input, 0, 2)
	n.neurs = make(map[int]*neuron)
	for i := 0; i < nbIn; i++ {
		fmt.Println("adding neur", i)
		neur := newNeuron(i)
		neur.weights[-1] = weight{1, 0.0}
		n.inp = append(n.inp, neur.inp)
		n.neurs[i] = neur
	}

	n.out = make(output, nbOut)
	outIds := []int{}
	for i := 0; i < nbOut; i++ {
		id := maxSize - (i + 1)
		outIds = append(outIds, id)
	}
	for i, id := range outIds {
		fmt.Println("adding", id)
		n.out[i] = make(input)
		neur := newNeuron(id)
		neur.addOut(n.out[i])
		n.neurs[id] = neur
	}

	if len(d) == 1 {
		n.addDNA(d[0])
	} else {
		for i := 0; i < nbIn; i++ {
			for _, id := range outIds {
				n.neurs[i].addOut(n.neurs[id].inp)
				n.neurs[id].weights[i] = weight{1, 0}
			}
		}
	}
	return n
}

func (n *Nn) String() string {
	s := ""
	for id, neur := range n.neurs {
		s += fmt.Sprintln(id, neur.weights)
	}
	return s
}

func (n *Nn) DNA() (d dna) {
	for _, neur := range n.neurs {
		if neur.id >= len(n.inp) {
			d = append(d, neur.genes()...)
		}
	}
	d.sort()
	return d
}

func (n *Nn) addGene(g gene) {
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

func (n *Nn) addDNA(d dna) {
	for _, g := range d {
		n.addGene(g)
	}
}

func (n *Nn) Mutate(rng *rand.Rand) {
	d := n.DNA()
	for i := range d {
		gago.MutNormalFloat64(d[i][:2], 0.1, rng)
		if d[i].sender() < 0 || d[i].sender() > n.maxSize {
			d[i][0] = float64(rand.Intn(n.maxSize))
		}
		if d[i].receiver() < 0 || d[i].receiver() > n.maxSize {
			d[i][1] = float64(rand.Intn(n.maxSize))
		}
		gago.MutNormalFloat64(d[i][2:], 0.8, rng)
	}
	*n = *newNN(len(n.inp), len(n.out), n.maxSize, d)
}

func (n *Nn) Crossover(cross gago.Genome, rng *rand.Rand) (gago.Genome, gago.Genome){
	d := n.DNA().toFloat()
	d2 := cross.(*Nn).DNA().toFloat()
	gago.CrossUniformFloat64(d, d2, rng)
	ret := newNN(len(n.inp), len(n.out), n.maxSize, floatToDNA(d))
	ret2 := newNN(len(n.inp), len(n.out), n.maxSize, floatToDNA(d2))
	return ret, ret2
}

func (n *Nn) Clone() gago.Genome {
	d := n.DNA()
	d2 := make(dna, 0, len(d))
	for _, g := range d {
		g2 := make([]float64, 0, 4)
		copy(g2, g)
		d2 = append(d2, g2)
	}
	n2 := newNN(len(n.inp), len(n.out), n.maxSize, d2)
	return n2
}

func (n *Nn) Evaluate() float64 {
	if n.Eval != nil {
		return n.Eval()
	}else  {
		return 1
	}
}

func makeGenome(rng *rand.Rand) gago.Genome {
	return newNN(2,2,6)
}
