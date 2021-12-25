package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

const (
	PREFIX = "encrypt--->:"
)

// pKCS5Padding padding
func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// pKCS5UnPadding unpadding
func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// checkImportContent content is encrypted or not
func checkImportContent(content string) (string, error) {
	if len(content) < len(PREFIX) {
		return "", errors.New("illegal content")
	}
	if content[:len(PREFIX)] != PREFIX {
		return "", errors.New("illegal content")
	}
	return content[len(PREFIX):], nil
}

//CBC加密
func EncryptAES_CBC(src, key string) (string, error) {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	data = pKCS5Padding(data, block.BlockSize())
	//获取CBC加密模式
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%s%X", PREFIX, out), nil
}

//CBC解密
func DecryptAES_CBC(src, key string) (string, error) {
	var err error
	if src, err = checkImportContent(src); err != nil {
		return "", err
	}
	keyByte := []byte(key)
	data, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = pKCS5UnPadding(plaintext)
	return string(plaintext), nil
}

//ECB加密
func EncryptAES_ECB(src, key string) (string, error) {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	//对明文数据进行补码
	data = pKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//对明文按照blocksize进行分块加密
		//必要时可以使用go关键字进行并行加密
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%s%X", PREFIX, out), nil
}

//ECB解密
func DecryptAES_ECB(src, key string) (string, error) {
	var err error
	if src, err = checkImportContent(src); err != nil {
		return "", err
	}
	data, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = pKCS5UnPadding(out)
	return string(out), nil
}

func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
