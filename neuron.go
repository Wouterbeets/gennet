package gennet

type neuron struct {
	inp     input
	weights map[int]weight
	out     output
}
