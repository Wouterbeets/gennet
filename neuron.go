package gennet

import (
	"math"
	"math/rand"
	"time"
)

type neuron struct {
	inp     input
	weights map[int]weight
	out     output
	id      int
}

func newNeuron(id int) *neuron {
	return &neuron{
		inp:     make(input, 3),
		weights: newWeights(),
		id:      id,
	}
}

func (neur *neuron) genes() (d dna) {
	for k, v := range neur.weights {
		d = append(d, gene{float64(k), float64(neur.id), v.weight, v.bias})
	}
	return d
}

func (neur *neuron) live() {
	sum := float64(0)
	nbSig := 0
	for {
		select {
		case sig := <-neur.inp:
			nbSig++
			w, ok := neur.weights[sig.neuronID]
			if !ok {
				w = weight{rand.NormFloat64(), rand.NormFloat64()}
				neur.weights[sig.neuronID] = w
			}
			sum += sig.val*w.weight + w.bias
			if nbSig == len(neur.weights) {
				neur.send(sum)
				return
			}
		case <-time.After(time.Microsecond):
			neur.send(sum)
			return
		}
	}
}

func (neur *neuron) send(sum float64) {
	outVal := 1.0 / (1.0 + math.Exp(-sum))
	for _, outChan := range neur.out {
		outChan <- signal{val: outVal, neuronID: neur.id}
	}
}

func (neur *neuron) addOut(inp input) {
	neur.out = append(neur.out, inp)
}
