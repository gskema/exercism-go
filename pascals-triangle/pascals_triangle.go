// Package pascal prints Pascale triangle
package pascal

// Triangle computes Pascal's triangle number up to a given row number
func Triangle(n int) [][]int {
	var matrix = [][]int{
		{1},
	}

	for i := 0; i < n-1; i++ {
		row := []int{}
		for j := range matrix[i] {
			if j == 0 {
				row = append(row, matrix[i][j])
			} else {
				row = append(row, matrix[i][j-1]+matrix[i][j])
			}
		}
		row = append(row, 1)
		matrix = append(matrix, row)
	}

	return matrix
}
