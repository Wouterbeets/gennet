package gennet

type weight struct {
	weight float64
	bias   float64
}

func newWeights() map[int]weight {
	w := make(map[int]weight)
	w[0] = weight{1, 0.5}
	return w
}

type output []input

type signal struct {
	val      float64
	neuronID int
}

type input chan signal
