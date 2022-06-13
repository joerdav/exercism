package robotname

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

const maximumNames = 676000
const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Define the Robot type here.
type Robot struct {
	namer namer
	name  string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	name, err := r.namer.name()
	r.name = name
	return r.name, err
}

func (r *Robot) Reset() {
	r.name = ""
}

type namer struct {
	names map[string]bool
}

func (n *namer) name() (string, error) {
	na := fmt.Sprintf(
		"%s%d%d%d",
		string([]byte{characters[n.randInt(len(characters))],
			characters[n.randInt(len(characters))]}),
		n.randInt(10),
		n.randInt(10),
		n.randInt(10),
	)
	if len(n.names) == maximumNames {
		return "", errors.New("max names reached")
	}
	if n.names[na] {
		return n.name()
	}
	n.names[na] = true
	return na, nil
}

func (n *namer) randInt(max int) int {
	if n.names == nil {
		n.names = map[string]bool{}
	}
	i, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(i.Int64())
}
