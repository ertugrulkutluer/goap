package utility

import (
	"crypto/md5"
	"fmt"
)

func ToMd5(r string) string {
	data := []byte(r)
	return fmt.Sprintf("%x", md5.Sum(data))
}
