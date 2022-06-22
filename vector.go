// this file contains a bunch of probably redundent vector operations
package aoTools

import (
	"github.com/anon55555/mt"

	"math"
)

// Vec is a 3 dimensional float64 vector
type Vec [3]float64

// Pos2vec
func Pos2vec(p mt.Pos) (v Vec) {
	for k := range v {
		v[k] = float64(p[k])
	}

	return
}

func vecSub(a, b Vec) (v Vec) {
	for k := range v {
		v[k] = a[k] - b[k]
	}

	return
}

func vecLen2d(v Vec) float64 {
	return math.Sqrt(v[0]*v[0] + v[2]*v[2])
}

func vecAdd(a, b Vec) (v Vec) {
	for k := range v {
		v[k] = a[k] + b[k]
	}

	return
}

func (v Vec) Pos() (p mt.Pos) {
	for k := range p {
		p[k] = float32(v[k])
	}

	return
}

