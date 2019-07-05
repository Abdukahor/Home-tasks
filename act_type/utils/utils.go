package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(login *string, pass *string) string {
	hasher := md5.New()
	hasher.Write([]byte(*login + *pass))
	return hex.EncodeToString(hasher.Sum(nil))
}
