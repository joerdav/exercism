package protein

import "fmt"

const (
	Methionine    = "Methionine"
	Phenylalanine = "Phenylalanine"
	Leucine       = "Leucine"
	Serine        = "Serine"
	Tyrosine      = "Tyrosine"
	Cysteine      = "Cysteine"
	Tryptophan    = "Tryptophan"
	STOP          = "STOP"
)

var (
	ErrStop        = fmt.Errorf("Stop CODON")
	ErrInvalidBase = fmt.Errorf("Invalid Base")
)

var codons = map[string]string{
	"AUG": Methionine,
	"UUU": Phenylalanine, "UUC": Phenylalanine,
	"UUA": Leucine, "UUG": Leucine,
	"UCU": Serine, "UCC": Serine, "UCA": Serine, "UCG": Serine,
	"UAU": Tyrosine, "UAC": Tyrosine,
	"UGU": Cysteine, "UGC": Cysteine,
	"UGG": Tryptophan,
	"UAA": STOP, "UAG": STOP, "UGA": STOP,
}

func FromRNA(rna string) ([]string, error) {
	var result []string
	for rna != "" {
		c, err := FromCodon(rna[:3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return nil, err
		}
		result = append(result, c)
		rna = rna[3:]
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	p, ok := codons[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if p == STOP {
		return "", ErrStop
	}
	return p, nil
}
