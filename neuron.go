package gennet

import (
	"fmt"
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
		inp:     make(input),
		weights: newWeights(),
		id:      id,
	}
}

func (neur *neuron) live() {
	sum := float64(0)
	nbSig := 0
	for {
		fmt.Println("looping ", neur.id, "expecting", len(neur.weights), "received", nbSig)
		select {
		case sig := <-neur.inp:
			fmt.Println("received inp", neur.id)
			nbSig++
			w, ok := neur.weights[sig.neuronID]
			if !ok {
				w = weight{rand.NormFloat64(), 0.5}
				neur.weights[sig.neuronID] = w
			}
			sum += sig.val*w.weight + w.bias
			if nbSig == len(neur.weights) {
				neur.send(sum)
				return
			}
		case <-time.After(time.Millisecond):
			neur.send(sum)
			return
		}
	}
}

func (neur *neuron) send(sum float64) {
	outVal := 1.0 / (1.0 + math.Exp(-sum))
	for _, outChan := range neur.out {
		fmt.Println("neur", neur.id, "sending")
		outChan <- signal{val: outVal, neuronID: neur.id}
	}
}

func (neur *neuron) addOut(inp input) {
	neur.out = append(neur.out, inp)
}
