package regexutil

import "strings"

// StandardTelephoneNumber get International telephone number from country and tel
func StandardTelephoneNumber(country, tel string) string {
	if len(tel) == 0 {
		return tel
	}
	country = strings.Replace(country, "+", "", -1)
	tel = strings.Replace(strings.TrimSpace(tel), " ", "", -1)
	if strings.HasPrefix(tel, "+"+country) {
	} else if len(country) != 0 {
		if strings.HasPrefix(tel, country) {
			tel = "+" + tel
		} else {
			tel = "+" + country + tel
		}
	}
	return strings.Replace(tel, "-", "", -1)
}

// BaseTelephoneNumber get base telephone number from country`s country and tel
func BaseTelephoneNumber(country, tel string) string {
	if len(tel) == 0 {
		return tel
	}
	country = strings.Replace(country, "+", "", -1)
	tel = strings.Replace(tel, "+", "", -1)
	if len(country) != 0 && strings.HasPrefix(tel, country) {
		tel = tel[len(country):]
	}
	tel = strings.Replace(strings.TrimSpace(tel), " ", "", -1)
	return strings.Replace(tel, "-", "", -1)
}
