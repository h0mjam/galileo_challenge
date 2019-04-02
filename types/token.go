package types

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func NewToken() string {
	b := [32]byte{}
	io.ReadFull(rand.Reader, b[:])
	return base64.RawStdEncoding.EncodeToString(b[:])
}
