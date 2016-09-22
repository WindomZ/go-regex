package regexutil

import "strings"

// StandardTelephoneNumber get International telephone number from country`s code and tel
func StandardTelephoneNumber(code, tel string) string {
	if len(tel) == 0 {
		return tel
	}
	code = strings.Replace(code, "+", "", -1)
	tel = strings.Replace(strings.TrimSpace(tel), " ", "", -1)
	if strings.HasPrefix(tel, "+"+code) {
	} else if len(code) != 0 {
		if strings.HasPrefix(tel, code) {
			tel = "+" + tel
		} else {
			tel = "+" + code + tel
		}
	}
	return strings.Replace(tel, "-", "", -1)
}

// BaseTelephoneNumber get base telephone number from country`s code and tel
func BaseTelephoneNumber(code, tel string) string {
	if len(tel) == 0 {
		return tel
	}
	code = strings.Replace(code, "+", "", -1)
	tel = strings.Replace(tel, "+", "", -1)
	if len(code) != 0 && strings.HasPrefix(tel, code) {
		tel = tel[len(code):]
	}
	tel = strings.Replace(strings.TrimSpace(tel), " ", "", -1)
	return strings.Replace(tel, "-", "", -1)
}
