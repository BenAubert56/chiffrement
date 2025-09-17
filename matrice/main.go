package main

import (
	"errors"
	"fmt"
)

func MatVecMul(A [][]float64, x []float64) ([]float64, error) {
	m := len(A)
	if m == 0 {
		return nil, errors.New("matrice avec 0 lignes")
	}
	n := len(A[0])
	if n == 0 {
		return nil, errors.New("matrice avec 0 colonnes")
	}
	if len(x) != n {
		return nil, errors.New("nombre de colonnes de A doit être égal à la taille de x")
	}
	for i := 1; i < m; i++ {
		if len(A[i]) != n {
			return nil, errors.New("toutes les lignes de A doivent avoir la même taille")
		}
	}

	y := make([]float64, m)
	for i := 0; i < m; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += A[i][j] * x[j]
		}
		y[i] = sum
	}
	return y, nil
}

func main() {
	A := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	x := []float64{1, 2, 3}
	y, err := MatVecMul(A, x)
	if err != nil {
		panic(err)
	}
	fmt.Println(y)
}
