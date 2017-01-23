// Package pascal calculates a pascal triangle
package pascal

// Triangle returns a pascal triangle with the requested number of rows
func Triangle(rows int) [][]int {
	triangle := make([][]int, rows, rows)
	for r := 0; r < rows; r++ {
		row := make([]int, r+1, r+1)
		for c := 0; c <= r; c++ {
			if c == 0 || c == r {
				row[c] = 1
			} else {
				row[c] = triangle[r-1][c-1] + triangle[r-1][c]
			}
		}
		triangle[r] = row
	}
	return triangle
}
