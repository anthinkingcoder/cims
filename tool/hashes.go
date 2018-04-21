package tool

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
)

func Md5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	rs := hex.EncodeToString(hash.Sum(nil))
	return rs
}

func Sha1(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	rs := hex.EncodeToString(hash.Sum(nil))
	return rs
}
