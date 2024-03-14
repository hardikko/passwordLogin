package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func StringToInt64(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func GetMd5(input string) string {
	password_hash := md5.New()
	defer password_hash.Reset()
	password_hash.Write([]byte(input))
	return hex.EncodeToString(password_hash.Sum(nil))
}
