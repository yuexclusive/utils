package crypto

import (
	"testing"
)

const (
	testContent = "hello world 321"
	testKey     = "8935760112796693"
)

func TestEncryptAES_CBC(t *testing.T) {
	encryptStr, err := EncryptAES_CBC(testContent, testKey)

	t.Log(encryptStr)

	if err != nil {
		t.Error(err)
	}

	res, err := DecryptAES_CBC(encryptStr, testKey)

	if err != nil {
		t.Error(err)
	}

	if res != testContent {
		t.Errorf("got: %s,want: %s", res, testContent)
	}
}

func TestEncryptAES_ECB(t *testing.T) {
	encryptStr, err := EncryptAES_ECB(testContent, testKey)

	t.Log(encryptStr)

	if err != nil {
		t.Error(err)
	}

	res, err := DecryptAES_ECB(encryptStr, testKey)
	if err != nil {
		t.Error(err)
	}

	if res != testContent {
		t.Errorf("got: %s,want: %s", res, testContent)
	}
}
