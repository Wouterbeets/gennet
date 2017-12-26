package gennet

import (
	"fmt"
	"math/rand"

	"github.com/MaxHalford/gago"
)

type Nn struct {
	inp     []input
	neurs   map[int]*neuron
	out     output
	maxSize int
	Eval    func() float64
	rand    *rand.Rand
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
	n.inp = make([]input, 0, nbIn)
	n.neurs = make(map[int]*neuron)
	for i := 0; i < nbIn; i++ {
		neur := newNeuron(i)
		neur.weights[-1] = weight{1, 0}
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
				n.neurs[id].weights[i] = weight{1, 1}
			}
		}
	}
	n.Eval = func() float64 { return rpc(n) }
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
		rec = newNeuron(g.receiver())
		n.neurs[g.receiver()] = rec
	}
	sen, ok := n.neurs[g.sender()]
	if !ok {
		sen = newNeuron(g.sender())
		n.neurs[g.sender()] = sen
	}
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
		gago.MutNormalFloat64(d[i][:2], 0.3, rng)
		if d[i].sender() < 0 || d[i].sender() >= n.maxSize {
			d[i][0] = float64(rng.Intn(n.maxSize))
		}
		if d[i].receiver() < 0 || d[i].receiver() >= n.maxSize {
			d[i][1] = float64(rng.Intn(n.maxSize))
		}
		gago.MutNormalFloat64(d[i][2:], 0.8, rng)
	}
	if rng.Int()%10 == 0 {
		d = append(d, gene{
			float64(rng.Intn(n.maxSize)),
			float64(rng.Intn(n.maxSize)),
			rng.NormFloat64(),
			rng.NormFloat64()})
	}
	*n = *newNN(len(n.inp), len(n.out), n.maxSize, d)
}

func (n *Nn) Crossover(cross gago.Genome, rng *rand.Rand) {
	d := n.DNA()
	d2 := cross.(*Nn).DNA()
	for len(d) < len(d2) {
		d = append(d, gene{
			float64(rng.Intn(n.maxSize)),
			float64(rng.Intn(n.maxSize)),
			rng.NormFloat64(),
			rng.NormFloat64()})
	}
	for len(d2) < len(d) {
		d2 = append(d2, gene{
			float64(rng.Intn(n.maxSize)),
			float64(rng.Intn(n.maxSize)),
			rng.NormFloat64(),
			rng.NormFloat64()})
	}
	for i := range d {
		if i < len(d2) {
			if rng.Int()%2 == 0 {
				copy(d[i], d2[i])
			}
		} else if len(d2) > 0 {
			if rng.Int()%2 == 0 {
				copy(d[i], d2[rng.Intn(len(d2))])
			}
		}
	}
	*n = *newNN(len(n.inp), len(n.out), n.maxSize, d)
}

func (n *Nn) Clone() gago.Genome {
	d := n.DNA()
	d2 := make(dna, 0, len(d))
	for _, g := range d {
		g2 := make([]float64, 4)
		copy(g2, g)

		d2 = append(d2, g2)
	}
	n2 := newNN(len(n.inp), len(n.out), n.maxSize, d2)
	return n2
}

func (n *Nn) Evaluate() float64 {
	if n.Eval != nil {
		return n.Eval()
	} else {
		panic("no eval func")
	}
}

type fitnessFunc func(n *Nn) float64

func makeGenomeMaker(inp, out, max int, d ...dna) func(*rand.Rand) gago.Genome {
	return func(r *rand.Rand) gago.Genome {
		var n *Nn
		if len(d) == 1 {
			for i := range d[0] {
				d[0][i][2] = r.NormFloat64()
				d[0][i][3] = r.NormFloat64()
			}
			n = newNN(inp, out, max, d[0])
		} else {
			n = newNN(inp, out, max)
		}
		return n
	}
}

func orGate(n *Nn) float64 {
	n.In([]float64{1, 1})
	out3 := (n.Out()[0])
	n.In([]float64{1, 0})
	out2 := (n.Out()[0])
	n.In([]float64{0, 0})
	out4 := (n.Out()[0])
	n.In([]float64{0, 1})
	out := (n.Out()[0])

	out3 = 1 - out3
	out4 = 1 - out4
	score := (out + out2 + out3 + out4) / 4
	//score = score - (float64(len(n.neurs)) / 1000.0)
	return -score
}

func rpc(n *Nn) float64 {
	n.In([]float64{1, 0, 0})
	out := n.Out()
	var score float64
	score += out[1] - (out[2] + out[0])
	n.In([]float64{0, 1, 0})
	out = n.Out()
	score += out[2] - (out[1] + out[0])
	n.In([]float64{0, 0, 1})
	out = n.Out()
	score += out[0] - (out[1] + out[2])
	return -(score / 3)
}
