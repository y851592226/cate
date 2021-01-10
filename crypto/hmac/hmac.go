package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Md5(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	_, err := hmac.Write([]byte(data))
	if err != nil {

	}
	return hex.EncodeToString(hmac.Sum(nil))
}

func Sha1(key, data string) string {
	hmac := hmac.New(sha1.New, []byte(key))
	_, err := hmac.Write([]byte(data))
	if err != nil {

	}
	return hex.EncodeToString(hmac.Sum(nil))
}
