package regex

import (
	"errors"
	"regexp"
)

// match if match pattern of regexp return nil
func match(pattern string, s string) error {
	if ok, err := regexp.MatchString(pattern, s); err != nil {
		return err
	} else if !ok {
		return errors.New("Don't match regexp: " + s)
	}
	return nil
}

// MatchHan match Chinese
func MatchHan(s string) error {
	return match(`^\p{Han}+$`, s)
}

// MatchNick 中文、英文、数字及下划线
func MatchNick(s string) error {
	return match(`^[\p{Han}_a-zA-Z0-9]+$`, s)
}

// MatchEnglishNum 英文、数字
func MatchEnglishNum(s string) error {
	return match(`^[a-zA-Z0-9]+$`, s)
}

// MatchUserName 字母开头，允许5-16字节，允许字母数字下划线
func MatchUserName(s string) error {
	return match(`^[a-zA-Z][a-zA-Z0-9_]{4,63}$`, s)
}

// MatchTel
func MatchTel(s string) error {
	return match(`^((\(\d{2,3}\))|(\d{3}\-))?1\d{10}$`, s)
}

// MatchTel math if s is chinese telephone number
func MatchChineseTel(s string) error {
	return match(`\+861\d{10}$`, s)
}

// MatchInternationalTel math if s is international telephone number
func MatchInternationalTel(s string) error {
	return match(`\+(9[976]\d|8[987530]\d|6[987]\d|5[90]\d|42\d|3[875]\d|2[98654321]\d|`+
		`9[8543210]|8[6421]|6[6543210]|5[87654321]|4[987654310]|3[9643210]|2[70]|7|1)\d{1,14}$`, s)
}

// MatchTelCountryCode math if s is international country`s code
func MatchTelCountryCode(s string) error {
	return match(`^(\+)?(9[976]\d|8[987530]\d|6[987]\d|5[90]\d|42\d|3[875]\d|2[98654321]\d|`+
		`9[8543210]|8[6421]|6[6543210]|5[87654321]|4[987654310]|3[9643210]|2[70]|7|1)$`, s)
}

// MatchEmail math if s is email address
func MatchEmail(s string) error {
	return match(`^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$`, s)
}
