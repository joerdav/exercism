package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
	r, err := opener()
	var tr TransientError
	if errors.As(err, &tr) {
		return Use(opener, input)
	}
	if err != nil {
		return err
	}
	defer r.Close()
	defer func() {
		p := recover()
		if p == nil {
			return
		}
		err = p.(error)
		if fr, ok := err.(FrobError); ok {
			r.Defrob(fr.defrobTag)
		}
		return
	}()
	r.Frob(input)
	return nil
}
