package ledger

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type translations struct {
	Date           string
	Description    string
	Change         string
	DateFormat     func(day, month, year string) string
	CurrencyFormat func(currency string, cents int) string
}

var locales = map[string]translations{
	"nl-NL": {
		Date:        "Datum",
		Description: "Omschrijving",
		Change:      "Verandering",
		DateFormat: func(day, month, year string) string {
			return fmt.Sprintf("%s-%s-%s", day, month, year)
		},
		CurrencyFormat: func(currency string, cents int) string {
			negative := false
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			d := float32(cents) / 100
			p := message.NewPrinter(language.English)
			if negative {
				return p.Sprintf("%v %.2f-", currency, d)
			}
			return p.Sprintf("%v %.2f ", currency, d)
		},
	},
	"en-US": {
		Date:        "Date",
		Description: "Description",
		Change:      "Change",
		DateFormat: func(day, month, year string) string {
			return fmt.Sprintf("%s/%s/%s", month, day, year)
		},
		CurrencyFormat: func(currency string, cents int) string {
			negative := false
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			d := float32(cents) / 100
			p := message.NewPrinter(language.English)
			if negative {
				return p.Sprintf("(%v%.2f)", currency, d)
			}
			return p.Sprintf("%v%.2f ", currency, d)
		},
	},
}

func writeWithPadding(w io.Writer, text string, pad int) {
	fmt.Fprintf(w, "%s%s", text, strings.Repeat(" ", pad-len(text)))
}
func writeRow(w io.Writer, s1, s2, s3 string) {
	writeWithPadding(w, s1, 10)
	fmt.Fprint(w, " | ")
	writeWithPadding(w, s2, 25)
	fmt.Fprint(w, " | ")
	fmt.Fprint(w, s3)
	fmt.Fprintln(w)
}
func writeHeader(w io.Writer, t translations) {
	writeRow(w, t.Date, t.Description, t.Change)
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	t, ok := locales[locale]
	switch {
	case currency != "EUR" && currency != "USD":
		return "", errors.New("unsupported currency")
	case !ok:
		return "", errors.New("locale is required")
	}
	es := make([]Entry, len(entries))
	copy(es, entries)
	sort.Slice(es, func(i, j int) bool {
		if es[i].Date != es[j].Date {
			return es[i].Date < es[j].Date
		}
		if es[i].Description != es[j].Description {
			return es[i].Description < es[j].Description
		}
		return es[i].Change < es[j].Change
	})
	bu := new(strings.Builder)
	writeHeader(bu, t)
	for _, et := range es {
		if len(et.Date) != 10 {
			return "", errors.New("date too long")
		}
		d1, d2, d3, d4, d5 := et.Date[0:4], et.Date[4], et.Date[5:7], et.Date[7], et.Date[8:10]
		if d2 != '-' || d4 != '-' {
			return "", errors.New("date in wrong format")
		}
		de := et.Description
		if len(de) > 25 {
			de = de[:22] + "..."
		}
		d := t.DateFormat(d5, d3, d1)
		cents := et.Change
		var a string
		c := "$"
		if currency == "EUR" {
			c = "€"
		}
		a += t.CurrencyFormat(c, cents)
		writeRow(bu, d, de, strings.Repeat(" ", 13-len(a))+a)
	}
	return bu.String(), nil
}
