package crypto

import (
	"testing"
)

const (
	testContent = "hello world 321"
	testKey     = "89357601"
)

func TestEncryptDES_CBC(t *testing.T) {
	encryptStr, err := EncryptDES_CBC(testContent, testKey)

	t.Log(encryptStr)

	if err != nil {
		t.Error(err)
	}

	res, err := DecryptDES_CBC(encryptStr, testKey)

	if err != nil {
		t.Error(err)
	}

	if res != testContent {
		t.Errorf("got: %s,want: %s", res, testContent)
	}
}

func TestEncryptDES_ECB(t *testing.T) {
	encryptStr, err := EncryptDES_ECB(testContent, testKey)

	t.Log(encryptStr)

	if err != nil {
		t.Error(err)
	}

	res, err := DecryptDES_ECB(encryptStr, testKey)
	if err != nil {
		t.Error(err)
	}

	if res != testContent {
		t.Errorf("got: %s,want: %s", res, testContent)
	}
}
