package service

import "testing"

func Test_ShortCode(t *testing.T) {
	generator := NewShortLinkGenerator()
	code1 := generator.Generate()
	code2 := generator.Generate()

	t.Logf("code1: %s, code2: %s", code1, code2)

	if code1 == "" || code2 == "" || code1 == code2 {
		t.Errorf("generate wrong code: code1: %s, code2: %s", code1, code2)
	}
}
