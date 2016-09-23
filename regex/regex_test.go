package regex

import (
	"errors"
	"testing"
)

var ErrTest error = errors.New("There are some test errors")

func TestMatchHan(t *testing.T) {
	if err := MatchHan("这是中文"); err != nil {
		t.Fatal(err)
	}
	if err := MatchHan("这是中文abc"); err == nil {
		t.Fatal(ErrTest)
	}
	if err := MatchNick("这是中文_abc_123"); err != nil {
		t.Fatal(err)
	}
	if err := MatchNick("这是中文@abc*123"); err == nil {
		t.Fatal(ErrTest)
	}
	if err := MatchUserName("abc_123"); err != nil {
		t.Fatal(err)
	}
	if err := MatchUserName("abc@123*abc"); err == nil {
		t.Fatal(ErrTest)
	}
}

func TestMatchTel(t *testing.T) {
	if err := MatchTel("13088880808"); err != nil {
		t.Fatal(err)
	}
	if err := MatchTel("13088880808"); err != nil {
		t.Fatal(err)
	}
	if err := MatchTel("+86 13088880808"); err == nil {
		t.Fatal(ErrTest)
	}
	if err := MatchTel("13088880808122"); err == nil {
		t.Fatal(ErrTest)
	}
}

func TestMatchChineseTel(t *testing.T) {
	if err := MatchChineseTel("+8613088880808"); err != nil {
		t.Fatal(err)
	}
	if err := MatchChineseTel("8613088880808"); err != nil {
		t.Fatal(err)
	}
	if err := MatchChineseTel("13088880808"); err != nil {
		t.Fatal(err)
	}
	if err := MatchChineseTel("+13088880808"); err == nil {
		t.Fatal(ErrTest)
	}
	if err := MatchChineseTel("+86 13088880808"); err == nil {
		t.Fatal(ErrTest)
	}
}

func TestMatchInternationalTel(t *testing.T) {
	if err := MatchInternationalTel("+8613088880808"); err != nil {
		t.Fatal(err)
	}
	if err := MatchInternationalTel("+86 13088880808"); err == nil {
		t.Fatal(ErrTest)
	}
}

func TestMatchTelCountryCode(t *testing.T) {
	if err := MatchTelCountryCode("86"); err != nil {
		t.Fatal(err)
	}
	if err := MatchTelCountryCode("+86"); err != nil {
		t.Fatal(err)
	}
	if err := MatchTelCountryCode("86 "); err == nil {
		t.Fatal(ErrTest)
	}
}

func TestMatchEmail(t *testing.T) {
	if err := MatchEmail("admin@163.com"); err != nil {
		t.Fatal(err)
	}
	if err := MatchEmail("admi2*n@a163a.com"); err == nil {
		t.Fatal(ErrTest)
	}
}
