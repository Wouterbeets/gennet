package gennet

type weight struct {
	weight float64
	bias   float64
}

func newWeights() map[int]weight {
	w := make(map[int]weight)
	return w
}

type output []input

type signal struct {
	val      float64
	neuronID int
}

type input chan signal
