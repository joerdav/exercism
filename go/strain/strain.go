package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var result Ints
	for _, ss := range i {
		if filter(ss) {
			result = append(result, ss)
		}
	}
	return result
}

func (i Ints) Discard(filter func(int) bool) Ints {
	return i.Keep(func(i int) bool { return !filter(i) })
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var result Lists
	for _, ss := range l {
		if filter(ss) {
			result = append(result, ss)
		}
	}
	return result
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var result Strings
	for _, ss := range s {
		if filter(ss) {
			result = append(result, ss)
		}
	}
	return result
}
