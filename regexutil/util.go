package regexutil

import "strings"

// StandardTelephoneNumber get international telephone number from country and tel
func StandardTelephoneNumber(country, tel string) string {
	if len(tel) == 0 {
		return tel
	}
	tel = strings.Replace(strings.Replace(strings.TrimSpace(tel),
		" ", "", -1), "+", "", -1)
	if country = BaseCountryCode(country); len(country) != 0 &&
		!strings.HasPrefix(tel, country) {
		tel = country + tel
	}
	return "+" + strings.Replace(tel, "-", "", -1)
}

// BaseTelephoneNumber get base telephone number from country and tel
func BaseTelephoneNumber(country, tel string) string {
	if len(tel) == 0 {
		return tel
	}
	country = BaseCountryCode(country)
	tel = strings.Replace(tel, "+", "", -1)
	if len(country) != 0 && strings.HasPrefix(tel, country) {
		tel = tel[len(country):]
	}
	tel = strings.Replace(strings.TrimSpace(tel), " ", "", -1)
	return strings.Replace(tel, "-", "", -1)
}

// BaseCountryCode get base country`s code from country
func BaseCountryCode(country string) string {
	if len(country) == 0 {
		return country
	}
	return strings.Replace(strings.Replace(strings.TrimSpace(country),
		" ", "", -1), "+", "", -1)
}
