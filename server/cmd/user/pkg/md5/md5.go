package md5

import (
	"crypto/md5"
	"fmt"
	"strings"
)

type EncryptManager struct {
	Salt string
}

func (e *EncryptManager) EncryptPassword(code string) string {
	return Crypt(code, e.Salt)
}

func Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
