package gennet

import (
	"math"
	"time"
)

type neuron struct {
	inp     input
	weights map[int]weight
	out     output
	id      int
}

func newNeuron() *neuron {
	return &neuron{
		inp:     make(input),
		weights: map[int]weight{},
		id:      newId(),
	}
}

func (neur *neuron) live() {
	for {
		sum := float64(0)
		for {
			select {
			case sig := <-neur.inp:
				w := neur.weights[sig.neuronID].weight
				sum += sig.val*w + sig.bias
			case <-time.After(time.Millisecond):
				outVal := 1.0 / (1.0 + math.Exp(-sum))
				for _, outChan := range output {
					outChan <- signal{val: outVal, neuronID: neur.id}
				}
				break
			}
		}
	}
}

func (neur *neuron) addOut(inp input) {
	neur.out = append(neur.out, inp)
}
