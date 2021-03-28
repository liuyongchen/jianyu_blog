package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	m := md5.New()
	salt := "goodLuck"
	m.Write([]byte(value + salt))
	return hex.EncodeToString(m.Sum(nil))
}
