package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Scores struct {
	name                string
	wins, draws, losses int
}

func (s Scores) Points() int {
	return s.wins*3 + s.draws
}

func Tally(reader io.Reader, writer io.Writer) error {
	scores := map[string]Scores{}
	s := bufio.NewScanner(reader)
	for s.Scan() {
		if strings.TrimSpace(s.Text()) == "" || strings.HasPrefix(s.Text(), "#") {
			continue
		}
		r := strings.Split(s.Text(), ";")
		if len(r) != 3 {
			return fmt.Errorf("invalid row: %v", s.Text())
		}
		t1, t2 := scores[r[0]], scores[r[1]]
		t1.name = r[0]
		t2.name = r[1]
		switch r[2] {
		case "win":
			t1.wins += 1
			t2.losses += 1
		case "loss":
			t2.wins += 1
			t1.losses += 1
		case "draw":
			t2.draws += 1
			t1.draws += 1
		default:
			return fmt.Errorf("invalid match status: %v", r[2])
		}
		scores[r[0]], scores[r[1]] = t1, t2
	}
	err := writeRow(writer, "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}
	var ss []Scores
	for _, s := range scores {
		ss = append(ss, s)
	}
	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Points() == ss[j].Points() {
			return ss[i].name < ss[j].name
		}
		return ss[i].Points() > ss[j].Points()
	})
	for _, s := range ss {
		err := writeRow(writer, s.name,
			fmt.Sprint(s.draws+s.losses+s.wins),
			fmt.Sprint(s.wins),
			fmt.Sprint(s.draws),
			fmt.Sprint(s.losses),
			fmt.Sprint(s.Points()))
		if err != nil {
			return err
		}
	}
	return nil
}

func writeRow(w io.Writer, name string, cols ...string) error {
	_, err := fmt.Fprint(w, name)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, strings.Repeat(" ", 30-len(name)))
	if err != nil {
		return err
	}
	for _, c := range cols {
		_, err = fmt.Fprintf(w, " | %v%v", strings.Repeat(" ", 2-len(c)), c)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintln(w)
	return err
}
