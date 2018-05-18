package crypt

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func Sha256(msg string) string {
	h := sha256.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}
