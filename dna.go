package gennet

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/MaxHalford/gago"
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

func (d dna) Mutate(rng *rand.Rand) {
	for _, g := range d {
		gago.MutNormalFloat64(g[2:], 0.8, rng)
	}
}
