package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	if err := validate(records); err != nil {
		return nil, err
	}
	var root *Node
	var remaining []Record
	for _, r := range records {
		if r.ID != 0 {
			remaining = append(remaining, r)
			continue
		}
		if root != nil {
			return nil, errors.New("duplicate root")
		}
		if r.Parent != 0 {
			return nil, errors.New("root node has parent")
		}
		root = &Node{
			ID: r.ID,
		}
	}
	if root == nil {
		return nil, errors.New("no root found")
	}
	cs, err := children(root.ID, remaining)
	root.Children = cs
	return root, err
}

func validate(records []Record) error {
	m := make(map[int]*Record)
	for _, r := range records {
		if m[r.ID] != nil {
			return errors.New("duplicate node")
		}
		if r.ID < r.Parent {
			return errors.New("record higher than parent, possible cycle")
		}
		m[r.ID] = &r
	}
	for i := 0; i < len(records); i++ {
		if m[i] == nil {
			return errors.New("non-continuous")
		}
	}
	return nil
}

func children(parentID int, records []Record) ([]*Node, error) {
	var remaining []Record
	var rs []*Node
	ids := make(map[int]bool)
	for _, r := range records {
		if r.Parent == r.ID {
			return nil, errors.New("cyclical")
		}
		if r.Parent != parentID {
			remaining = append(remaining, r)
			continue
		}
		rs = append(rs, &Node{
			ID: r.ID,
		})
		ids[r.ID] = true
		continue
	}
	for _, n := range rs {
		cs, err := children(n.ID, remaining)
		if err != nil {
			return nil, err
		}
		n.Children = cs
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].ID < rs[j].ID
	})
	return rs, nil
}
