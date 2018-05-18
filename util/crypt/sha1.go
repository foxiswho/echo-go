package crypt

import (
	"crypto/sha1"
	"fmt"
	"io"
	"crypto/hmac"
	"encoding/base64"
)

func Sha1(msg string) string {
	h := sha1.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func HamSha1(data string, key []byte) string {
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(data))

	return base64.StdEncoding.EncodeToString(hmac.Sum(nil))
}
