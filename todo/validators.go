package todo

import "errors"

var (
	ErrTitleTooLong  = errors.New("todo: title too long")
	ErrTitleTooShort = errors.New("todo: title too short")
	ErrTitleEmpty    = errors.New("todo: title empty")
)

const maxTitleLength = 254
const minTitleLength = 6

func validateTitle(title string) error {
	l := len(title)

	switch {
	case l == 0:
		return ErrTitleEmpty
	case l < minTitleLength:
		return ErrTitleTooShort
	case l > maxTitleLength:
		return ErrTitleTooLong
	default:
		return nil
	}
}
