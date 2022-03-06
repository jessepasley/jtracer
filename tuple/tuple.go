package tuple

// import "math"

type FourTuple struct {
	Elems []float64
}

func NewVector(x, y, z float64) *FourTuple {
	return &FourTuple{[]float64{x, y, z, 0.0}}
}

func (this FourTuple) IsVector() bool {
	return this.Elems[3] == 0.0
}
