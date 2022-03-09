package tuple

import "math"

// import "math"

type FourTuple struct {
	Elements []float64
}

func NewVector(x, y, z float64) *FourTuple {
	return &FourTuple{[]float64{x, y, z, 0.0}}
}

func (t FourTuple) IsVector() bool {
	return t.Elements[3] == 0.0
}

func NewPoint(x, y, z float64) *FourTuple {
	return &FourTuple{[]float64{x, y, z, 1.0}}
}

func (t FourTuple) IsPoint() bool {
	return t.Elements[3] == 1.0
}

func NewFourTuple(elements []float64) *FourTuple {
	return &FourTuple{Elements: elements}
}

func (t FourTuple) Get(row int) float64 {
	return t.Elements[row]
}

func Add(tup1, tup2 FourTuple) *FourTuple {
	res := NewFourTuple(make([]float64, 4))
	for i := range res.Elements {
		res.Elements[i] = tup1.Get(i) + tup2.Get(i)
	}
	return res
}

func Subtract(tup1, tup2 FourTuple) *FourTuple {
	res := NewFourTuple(make([]float64, 4))
	for i := range res.Elements {
		res.Elements[i] = tup1.Get(i) - tup2.Get(i)
	}
	return res
}

func Negate(tup FourTuple) *FourTuple {
	res := NewFourTuple(make([]float64, 4))
	for i := range res.Elements {
		res.Elements[i] = 0 - tup.Get(i)
	}
	return res
}

func MultiplyByScalar(tup1 FourTuple, scalar float64) *FourTuple {
	res := &FourTuple{Elements: make([]float64, 4)}
	for i := 0; i < 4; i++ {
		res.Elements[i] = tup1.Get(i) * scalar
	}
	return res
}

func DivideByScalar(tup1 FourTuple, scalar float64) *FourTuple {
	res := &FourTuple{Elements: make([]float64, 4)}
	for i := 0; i < 4; i++ {
		res.Elements[i] = tup1.Get(i) / scalar
	}
	return res
}

func Magnitude(tup1 FourTuple) float64 {
	return math.Sqrt(tup1.Elements[0]*tup1.Elements[0] +
		tup1.Elements[1]*tup1.Elements[1] +
		tup1.Elements[2]*tup1.Elements[2])
}

//NormalLize
func Normalize(tup1 FourTuple) *FourTuple {
	res := &FourTuple{Elements: make([]float64, 4)}
	magnitude := Magnitude(tup1)
	for i := 0; i < 4; i++ {
		res.Elements[i] = tup1.Get(i) / magnitude
	}
	return res
}

//Dot
func Dot(tup1 FourTuple, tup2 FourTuple) float64 {
	sum := 0.0
	for i := 0; i < 4; i++ {
		sum += tup1.Get(i) * tup2.Get(i)
	}
	return sum
}

//Cross
func Cross(tup1 FourTuple, tup2 FourTuple) *FourTuple {
	res := &FourTuple{Elements: make([]float64, 4)}

	res.Elements[0] = tup1.Get(1)*tup2.Get(2) - tup1.Get(2)*tup2.Get(1)
	res.Elements[1] = tup1.Get(2)*tup2.Get(0) - tup1.Get(0)*tup2.Get(2)
	res.Elements[2] = tup1.Get(0)*tup2.Get(1) - tup1.Get(1)*tup2.Get(0)
	res.Elements[3] = 0
	return res
}

//Eq
func TupleEquals(tup1, tup2 FourTuple) bool {
	return Eq(tup1.Get(0), tup2.Get(0)) &&
		Eq(tup1.Get(1), tup2.Get(1)) &&
		Eq(tup1.Get(2), tup2.Get(2)) &&
		Eq(tup1.Get(3), tup2.Get(3))
}
