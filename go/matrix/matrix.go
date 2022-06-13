package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	ints  []int
	width int
}

func New(s string) (*Matrix, error) {
	var m Matrix
	rows := strings.Split(s, "\n")
	for _, r := range rows {
		f := strings.Fields(r)
		m.width = len(f)
		for _, n := range f {
			nn, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			m.ints = append(m.ints, nn)
		}
	}
	if len(rows)*m.width != len(m.ints) {
		return nil, errors.New("different length of rows")
	}
	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, m.width)
	for i, n := range m.ints {
		cols[i%m.width] = append(cols[i%m.width], n)
	}
	return cols
}

func (m *Matrix) Rows() [][]int {
	rows := make([][]int, len(m.ints)/m.width)
	for i, n := range m.ints {
		rows[i/m.width] = append(rows[i/m.width], n)
	}
	return rows
}

func (m *Matrix) Set(row, col, val int) bool {
	if col < 0 || row < 0 {
		return false
	}
	if col > m.width-1 || row > len(m.ints)/m.width-1 {
		return false
	}
	m.ints[row*m.width+col] = val
	return true
}
