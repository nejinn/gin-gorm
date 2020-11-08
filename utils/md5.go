package utils
import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密
func NewMD5(b []byte) string {
	newM := md5.New()
	newM.Write(b)
	return hex.EncodeToString(newM.Sum(nil))
}
