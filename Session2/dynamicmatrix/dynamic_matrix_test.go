package dynamicmatrix

import "testing"

/*
implement the following function which ALWAYS sets the value of matrix_ij:
func SetElement(matrix [][]int, i int, j int, value int) (result [][]int)
if i or j are out of range, matrix must be extended with '0' value
*/

func TestSetElement(t *testing.T) {

	var matrix [][]int

	/* --------------- */

	matrix = SetElement(matrix, 0, 0, 2)

	expected := [][]int{
		{2},
	}

	if !matrixEqual(matrix, expected) {
		t.Fatalf("expected %v; got %v", expected, matrix)
	}

	/* --------------- */

	matrix = SetElement(matrix, 1, 3, 1)

	expected = [][]int{
		{2, 0, 0, 0},
		{0, 0, 0, 1},
	}

	if !matrixEqual(matrix, expected) {
		t.Fatalf("expected %v; got %v", expected, matrix)
	}

	/* --------------- */

	matrix = SetElement(matrix, 1, 1, 3)

	expected = [][]int{
		{2, 0, 0, 0},
		{0, 3, 0, 1},
	}

	if !matrixEqual(matrix, expected) {
		t.Fatalf("expected %v; got %v", expected, matrix)
	}

	/* --------------- */

	matrix = SetElement(matrix, 0, 3, 4)

	expected = [][]int{
		{2, 0, 0, 4},
		{0, 3, 0, 1},
	}

	if !matrixEqual(matrix, expected) {
		t.Fatalf("expected %v; got %v", expected, matrix)
	}

	/* --------------- */

	matrix = SetElement(matrix, 3, 3, 6)

	expected = [][]int{
		{2, 0, 0, 4},
		{0, 3, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 6},
	}

	if !matrixEqual(matrix, expected) {
		t.Fatalf("expected %v; got %v", expected, matrix)
	}

	/* --------------- */

	matrix = SetElement(matrix, 3, 0, 7)

	expected = [][]int{
		{2, 0, 0, 4},
		{0, 3, 0, 1},
		{0, 0, 0, 0},
		{7, 0, 0, 6},
	}

	if !matrixEqual(matrix, expected) {
		t.Fatalf("expected %v; got %v", expected, matrix)
	}

	/* --------------- */

	matrix = SetElement(matrix, 2, 5, 8)

	expected = [][]int{
		{2, 0, 0, 4, 0, 0},
		{0, 3, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 8},
		{7, 0, 0, 6, 0, 0},
	}

	if !matrixEqual(matrix, expected) {
		t.Fatalf("expected %v; got %v", expected, matrix)
	}

	/* --------------- */

}

/* ---------- helpers ---------- */

func matrixEqual(m1, m2 [][]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for i := 0; i < len(m1); i++ {
		if len(m1[i]) != len(m2[i]) {
			return false
		}

		for j := 0; j < len(m1[i]); j++ {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}

	return true
}
