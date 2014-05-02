package pairwise

import (
	"errors"
	"math"

	mat "github.com/skelterjohn/go.matrix"
)

type Euclidean struct{}

func NewEuclidean() *Euclidean {
	return &Euclidean{}
}

func (self *Euclidean) InnerProduct(vectorX *mat.DenseMatrix, vectorY *mat.DenseMatrix) (float64, error) {
	if !CheckDimMatch(vectorX, vectorY) {
		return 0, errors.New("Dimension mismatch")
	}

	result := mat.Product(mat.Transpose(vectorX), vectorY).Get(0, 0)

	return result, nil
}

// We may need to create Metrics / Vector interface for this
func (self *Euclidean) Distance(vectorX *mat.DenseMatrix, vectorY *mat.DenseMatrix) (float64, error) {
	difference, err := vectorY.MinusDense(vectorX)
	result, err := self.InnerProduct(difference, difference)

	return math.Sqrt(result), err
}
