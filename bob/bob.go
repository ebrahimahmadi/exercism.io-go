package bob

import (
	"strings"
	"unicode"
)

const (
	resToQuestion        = "Sure."
	resToYelling         = "Whoa, chill out!"
	resToYellingQuestion = "Calm down, I know what I'm doing!"
	resToEmpty           = "Fine. Be that way!"
	defaultRes           = "Whatever."
)

type Remark struct {
	expression string
}

func Hey(remark string) string {
	boundRemark := newRemark(strings.TrimSpace(remark))

	switch {
	case boundRemark.isEmpty():
		return resToEmpty
	case boundRemark.isQuestion() && boundRemark.isYelling():
		return resToYellingQuestion
	case boundRemark.isQuestion():
		return resToQuestion
	case boundRemark.isYelling():
		return resToYelling
	default:
		return defaultRes
	}
}

func newRemark(expression string) Remark {
	return Remark{
		expression: expression,
	}
}

func (remark Remark) isEmpty() bool {
	return remark.expression == ""
}

func (remark Remark) isQuestion() bool {
	lastChr := remark.expression[len(remark.expression)-1 : len(remark.expression)]
	return lastChr == "?"
}

func (remark Remark) isYelling() bool {
	hasLetters := strings.IndexFunc(remark.expression, unicode.IsLetter) >= 0
	isUpper := strings.ToUpper(remark.expression) == remark.expression

	return hasLetters && isUpper
}
