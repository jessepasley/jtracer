package tuple

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTuple_IsVector(t *testing.T) {

	t.Run("is vector", func(t *testing.T) {
		testVector := NewVector(4.2, -4.7, 1.4)
		got := testVector.IsVector()
		want := true

		if got != want {
			t.Errorf(" this is not a vector")
		}
	})
}

func TestTuple_IsPoint(t *testing.T) {
	t.Run("is point", func(t *testing.T) {
		testPoint := NewPoint(4.2, -4.7, 1.4)
		got := testPoint.IsPoint()
		want := true

		if got != want {
			t.Errorf(" this is not a point")
		}
	})
}

func TestTuple_Add(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		testPoint := NewPoint(1.0, 1.0, 1.0)
		testVector := NewVector(2.0, 2.1, 2.2)
		got := Add(*testPoint, *testVector)
		want := NewFourTuple([]float64{3.0, 3.1, 3.2, 1.0})

		if !reflect.DeepEqual(got, want) {
			t.Errorf(" hmmm")
		}
	})
}

func TestTuple_Subtract(t *testing.T) {
	t1 := NewPoint(7, 1, 0)
	t2 := NewPoint(5, 6, 7)
	t3 := Subtract(*t1, *t2)
	assert.Equal(t, 2.0, t3.Get(0))
	assert.Equal(t, -5.0, t3.Get(1))
	assert.Equal(t, -7.0, t3.Get(2))
	assert.Equal(t, 0.0, t3.Get(3))
}

func TestSubtractVectorFromPoint(t *testing.T) {
	t1 := NewPoint(5, 1, 19)
	t2 := NewVector(5, 8, 2)

	t3 := Subtract(*t1, *t2)
	assert.Equal(t, 0.0, t3.Get(0))
	assert.Equal(t, -7.0, t3.Get(1))
	assert.Equal(t, 17.0, t3.Get(2))
	assert.Equal(t, 1.0, t3.Get(3))
}

func TestSubtractVectorFromVector(t *testing.T) {
	t1 := NewVector(5, 1, 19)
	t2 := NewVector(5, 8, 2)

	res := Subtract(*t1, *t2)
	assert.Equal(t, 0.0, res.Get(0))
	assert.Equal(t, -7.0, res.Get(1))
	assert.Equal(t, 17.0, res.Get(2))
	assert.Equal(t, 0.0, res.Get(3))
}

func TestNegateTuple(t *testing.T) {
	tup1 := FourTuple{[]float64{2, -1, 19, -4}}
	res := Negate(tup1)

	assert.Equal(t, -2.0, res.Get(0))
	assert.Equal(t, 1.0, res.Get(1))
	assert.Equal(t, -19.0, res.Get(2))
	assert.Equal(t, 4.0, res.Get(3))
}

func TestMultiplyByScalar(t *testing.T) {
	t1 := FourTuple{[]float64{1, -2, 3, -4}}
	t3 := MultiplyByScalar(t1, 3.5)
	assert.Equal(t, 3.5, t3.Get(0))
	assert.Equal(t, -7.0, t3.Get(1))
	assert.Equal(t, 10.5, t3.Get(2))
	assert.Equal(t, -14.0, t3.Get(3))
}

func TestMultiplyByScalarFraction(t *testing.T) {
	t1 := FourTuple{[]float64{1, -2, 3, -4}}
	t3 := MultiplyByScalar(t1, 0.5)
	assert.Equal(t, 0.5, t3.Get(0))
	assert.Equal(t, -1.0, t3.Get(1))
	assert.Equal(t, 1.5, t3.Get(2))
	assert.Equal(t, -2.0, t3.Get(3))
}

func TestDivideByScalar(t *testing.T) {
	t1 := FourTuple{[]float64{1, -2, 3, -4}}
	t3 := DivideByScalar(t1, 2)
	assert.Equal(t, 0.5, t3.Get(0))
	assert.Equal(t, -1.0, t3.Get(1))
	assert.Equal(t, 1.5, t3.Get(2))
	assert.Equal(t, -2.0, t3.Get(3))
}

func TestMagnitude(t *testing.T) {
	tc := []struct {
		tpl *FourTuple
		out float64
	}{
		{NewVector(1, 0, 0), 1.0},
		{NewVector(0, 1, 0), 1.0},
		{NewVector(0, 0, 1), 1.0},
		{NewVector(1, 2, 3), math.Sqrt(14)},
		{NewVector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, test := range tc {
		assert.Equal(t, test.out, Magnitude(*test.tpl))
	}
}

func TestNormalizeXOnly(t *testing.T) {
	t1 := NewVector(4, 0, 0)
	t3 := Normalize(*t1)
	assert.Equal(t, 1.0, t3.Get(0))
	assert.Equal(t, 0.0, t3.Get(1))
	assert.Equal(t, 0.0, t3.Get(2))
}

func TestNormalizeXYZ(t *testing.T) {
	t1 := NewVector(1, 2, 3)
	t3 := Normalize(*t1)
	assert.True(t, Eq(0.26726, t3.Get(0)))
	assert.True(t, Eq(0.53452, t3.Get(1)))
	assert.True(t, Eq(0.80178, t3.Get(2)))
}

func TestNormalizedMagnitudeIsOne(t *testing.T) {
	t1 := NewVector(1, 2, 3)
	t3 := Normalize(*t1)
	assert.Equal(t, 1.0, Magnitude(*t3))
}

func TestDot(t *testing.T) {
	t1 := NewVector(1, 2, 3)
	t2 := NewVector(2, 3, 4)
	dotProduct := Dot(*t1, *t2)
	assert.Equal(t, 20.0, dotProduct)
}

func TestCross(t *testing.T) {
	t1 := NewVector(1, 2, 3)
	t2 := NewVector(2, 3, 4)
	crossT1 := Cross(*t1, *t2)
	crossT2 := Cross(*t2, *t1)
	assert.True(t, TupleEquals(*crossT1, *NewVector(-1, 2, -1)))
	assert.True(t, TupleEquals(*crossT2, *NewVector(1, -2, 1)))
}
