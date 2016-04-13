package regex

import (
	"errors"
	"regexp"
)

func toError(b bool, err error) error {
	if err != nil {
		return err
	} else if !b {
		return errors.New("No match!")
	}
	return nil
}

func match(pattern string, s string) error {
	return toError(regexp.MatchString(pattern, s))
}

// 中文
func MatchHan(s string) error {
	return match(`^\p{Han}+$`, s)
}

// 中文、英文、数字及下划线
func MatchNick(s string) error {
	return match(`^[\p{Han}_a-zA-Z0-9]+$`, s)
}

// 英文、数字
func MatchEnglishNum(s string) error {
	return match(`^[a-zA-Z0-9]+$`, s)
}

// 字母开头，允许5-16字节，允许字母数字下划线
func MatchUserName(s string) error {
	return match(`^[a-zA-Z][a-zA-Z0-9_]{4,63}$`, s)
}

func MatchTel(s string) error {
	return match(`^((\(\d{2,3}\))|(\d{3}\-))?1\d{10}$`, s)
}

func MatchEmail(s string) error {
	return match(`^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$`, s)
}
