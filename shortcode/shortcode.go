package service

import (
	"strings"

	"github.com/yuexclusive/utils/snowflake"
)

var shortFactor = [62]string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
	"l", "m", "n", "o", "p", "q", "r", "s", "t", "u",
	"v", "w", "x", "y", "z", "A",
	"B", "C", "E", "F", "D", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

const (
	minSize = 8
)

// ShortCodeGenerator ShortCodeGenerator
type ShortCodeGenerator interface {
	Generate() string
}

// shortCodeGenerator 短链接
type shortCodeGenerator struct {
	LinkLen int
}

// genSnowflak slow number
func (s *shortCodeGenerator) genSnowflak() int64 {
	return snowflake.BaseNumber()
}

// GenShortLink 生成短code
func (s *shortCodeGenerator) Generate() string {
	return s.genCode(s.genSnowflak())
}

// 转62进制
func (s *shortCodeGenerator) baseCode(_code int64) string {
	var code = make([]string, 0, 32)
	id := _code
	binLen := int64(len(shortFactor))
	for ; id/binLen > 0; id /= binLen {
		index := id % binLen
		code = append(code, shortFactor[index])
	}
	if id > 0 && id < binLen {
		code = append(code, shortFactor[id])
	}
	return strings.Join(code, "")
}

// genCode genCode
func (s *shortCodeGenerator) genCode(_code int64) string {
	str := s.baseCode(_code)
	size := s.LinkLen
	if s.LinkLen == 0 || s.LinkLen < minSize || size > len(str) {
		size = len(str)
	}
	return str[:size]
}

// NewShortLinkGenerator NewShortLinkGenerator
func NewShortLinkGenerator() ShortCodeGenerator {
	return &shortCodeGenerator{
		LinkLen: minSize,
	}
}
