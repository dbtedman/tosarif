package osvscanner

import "errors"

type Source struct{}

var ErrorNilSource = errors.New("cannot convert nil source")

func (my *Source) Validate() error {
	return nil
}
