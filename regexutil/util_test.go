package regexutil

import (
	"github.com/WindomZ/go-regex/regex"
	"testing"
)

func TestStandardTelephoneNumber(t *testing.T) {
	tel := StandardTelephoneNumber("86", "13088880808")
	if err := regex.MatchChineseTel(tel); err != nil {
		t.Fatal(err)
	}
	tel = StandardTelephoneNumber("86", "8613088880808")
	if err := regex.MatchChineseTel(tel); err != nil {
		t.Fatal(err)
	}
	tel = StandardTelephoneNumber("+86", "+86 - 13088880808")
	if err := regex.MatchChineseTel(tel); err != nil {
		t.Fatal(err)
	}
}

func TestBaseTelephoneNumber(t *testing.T) {
	if tel := BaseTelephoneNumber("+86",
		"+86 - 13088880808"); tel != "13088880808" {
		t.Fatal("Error tel:", tel)
	}
}
