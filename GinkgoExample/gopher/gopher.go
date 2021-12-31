package gopher

import (
	"errors"
	"unicode/utf8"
)

type Gopher struct {
	Name   string
	Gender string
	Age    int
}

func Validate(g Gopher) error {
	if utf8.RuneCountInString(g.Name) < 3 {
		return errors.New("名字太短，不能小于3")
	}

	if g.Gender != "男" {
		return errors.New("抱歉！不是男生，目前只面向男生")
	}

	if g.Age < 18 {
		return errors.New("年龄太小，不能小于18")
	}
	return nil
}
