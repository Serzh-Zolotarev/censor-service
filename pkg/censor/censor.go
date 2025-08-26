package censor

import (
	"regexp"
)

type ValidatorFunc func(content string) bool

func Validate(content string) bool {
	return !regexp.MustCompile(`(qwerty|йцукен|zxvbnm)`).Match([]byte(content))
}
