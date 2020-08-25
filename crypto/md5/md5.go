package md5

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) //nolint
	return hex.EncodeToString(h.Sum(nil))
}

func UppercaseMD5(str string) string {
	return strings.ToUpper(MD5(str))
}
