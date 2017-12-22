package gennet

type gene []float64

func (g gene) sender() int {
	return int(g[0])
}

func (g gene) receiver() int {
	return int(g[1])
}

func (g gene) weight() float64 {
	return g[2]
}

func (g gene) bias() float64 {
	return g[3]
}
