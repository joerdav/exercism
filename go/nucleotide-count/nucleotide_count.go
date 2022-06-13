package dna

import "errors"

var nucleotides = []rune{'A', 'C', 'G', 'T'}

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{}
	for _, n := range nucleotides {
		h[n] = 0
	}
	for _, l := range d {
		if _, ok := h[l]; !ok {
			return nil, errors.New("invalud nucleotide")
		}
		h[l]++
	}
	return h, nil
}
