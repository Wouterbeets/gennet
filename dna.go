package gennet

import (
	"fmt"
	"sort"
)

type dna []gene

func (d dna) String() (s string) {
	for _, g := range d {
		s += fmt.Sprintln("sen:", g[0], "rec:", g[1], "w:", g[2], "b:", g[3])
	}
	return s
}

func (d dna) Len() int {
	return len(d)
}

func (d dna) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d dna) Less(i, j int) bool {
	if d[i][0] == d[j][0] {
		return d[i][1] < d[j][1]
	}
	return d[i][0] < d[j][0]
}

func (d dna) sort() {
	sort.Sort(d)
}
func (d dna) toFloat() []float64 {
	ret := make([]float64, 0, len(d)*4)
	for _, g := range d {
		ret = append(ret, g)
	}
	return ret
}

func floatToDNA(f []float64) dna {
	d := make(dna, 0, len(f)/4)
	for i := 0; i < len(f); i++ {
		d = append(d, gene{f[i], f[i+1], f[i+2], f[i+3]})
		i += 4
	}
	return d
}
