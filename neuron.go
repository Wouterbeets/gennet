package gennet

import (
	"math"
	"math/rand"
)

type neuron struct {
	inp     input
	weights map[int]weight
	out     output
	id      int
	die     chan bool
}

func newNeuron(id int) *neuron {
	return &neuron{
		inp:     make(input, 1),
		weights: newWeights(),
		id:      id,
		die:     make(chan bool),
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
	kill := make(chan bool)
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
				neur.send(sum, kill)
				nbSig = 0
			}
		case <-neur.die:
			kill <- true
			return
		}
	}
}

func (neur *neuron) send(sum float64, kill <-chan bool) {
	outVal := 1.0 / (1.0 + math.Exp(-sum))
	go func(id int, out output, val float64) {
		for _, outChan := range out {
			select {
			case outChan <- signal{val, id}:
			case <-kill:
				return
			}
		}
	}(neur.id, neur.out, outVal)
}

func (neur *neuron) addOut(inp input) {
	neur.out = append(neur.out, inp)
}
