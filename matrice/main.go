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

// AddMatrices additionne deux ou plusieurs matrices.
// Retourne une erreur si les dimensions ne correspondent pas.
func AddMatrices(matrices ...[][]int) ([][]int, error) {
	if len(matrices) < 2 {
		return nil, errors.New("au moins deux matrices sont requises")
	}

	rows := len(matrices[0])
	cols := len(matrices[0][0])

	// Vérifier que toutes les matrices ont la même taille
	for _, m := range matrices {
		if len(m) != rows || len(m[0]) != cols {
			return nil, errors.New("toutes les matrices doivent avoir la même taille")
		}
	}

	// Créer la matrice résultat
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}

	// Addition
	for _, m := range matrices {
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				result[i][j] += m[i][j]
			}
		}
	}

	return result, nil
}

// det2x2 calcule le déterminant d'une matrice 2x2 [[a,b],[c,d]]
// Formule : det = ad - bc
func det2x2(a, b, c, d int64) int64 {
	return a*d - b*c
}

// gcd calcule le plus grand commun diviseur (PGCD) avec l'algorithme d'Euclide
func gcd(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// modNorm normalise un entier x dans l'intervalle [0, n-1]
func modNorm(x, n int64) int64 {
	if n <= 0 {
		panic("le module n doit être > 0")
	}
	r := x % n
	if r < 0 {
		r += n
	}
	return r
}

// DetAndInvertibleModN calcule :
// - le déterminant d'une matrice 2x2 [[a,b],[c,d]]
// - indique si la matrice est inversible modulo n
//
// Une matrice est inversible dans Z/nZ si et seulement si PGCD(det, n) = 1
func DetAndInvertibleModN(a, b, c, d, n int64) (det int64, invertible bool) {
	if n <= 0 {
		panic("le module n doit être > 0")
	}
	// Calcul du déterminant "classique"
	det = det2x2(a, b, c, d)

	// Condition d'inversibilité modulo n : PGCD(det, n) = 1
	invertible = gcd(det, n) == 1

	return det, invertible
}

func main() {

	// multiplication matrice
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

	// Somme de matrices
	A2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	B := [][]int{
		{7, 8, 9},
		{10, 11, 12},
	}
	C := [][]int{
		{1, 1, 1},
		{1, 1, 1},
	}

	sum, err := AddMatrices(A2, B, C)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	fmt.Println("Résultat:")
	for _, row := range sum {
		fmt.Println(row)
	}

	// Déterminant et inversibilité modulo n
	// Exemple 1 : matrice inversible modulo 7
	m11, m12, m21, m22, modulo := int64(1), int64(2), int64(3), int64(4), int64(7)
	determinant, estInversible := DetAndInvertibleModN(m11, m12, m21, m22, modulo)
	fmt.Printf("Matrice [[%d,%d],[%d,%d]] : det=%d, det mod %d=%d, inversible mod %d ? %v\n",
		m11, m12, m21, m22, determinant, modulo, modNorm(determinant, modulo), modulo, estInversible)

	// Exemple 2 : matrice NON inversible modulo 6
	m11, m12, m21, m22, modulo = 2, 4, 1, 2, 6
	determinant, estInversible = DetAndInvertibleModN(m11, m12, m21, m22, modulo)
	fmt.Printf("Matrice [[%d,%d],[%d,%d]] : det=%d, det mod %d=%d, inversible mod %d ? %v\n",
		m11, m12, m21, m22, determinant, modulo, modNorm(determinant, modulo), modulo, estInversible)
}
