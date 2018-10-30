package util

import (
	"crypto/md5"
	"encoding/hex"
)

func PasswordHashing(original string, salt string) string {
	originalPlusSalt := []byte(original + salt)
	md5converted := md5.Sum(originalPlusSalt)
	return hex.EncodeToString(md5converted[:])
}
