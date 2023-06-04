package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func CalculateETag(content []byte) string {

	hash := md5.New()
	hash.Write(content)
	etag := hex.EncodeToString(hash.Sum(nil))

	return etag
}
