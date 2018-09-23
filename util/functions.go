package util

import (
	"fmt"
	"github.com/pkg/errors"
	"crypto/md5"
)

func IsRequire(data map[string]string) (err error) {
	for k, v := range data {
		if (v == "" || len(v) == 0) {
			err = errors.New(fmt.Sprintf("%s is require", k))

			return err
		}
	}
	return nil
}

func JsonMsg(msg interface{}) (returnData map[string]interface{}) {
	errorMsg := fmt.Sprintf("%s",msg)
	returnData = map[string]interface{}{"ok": false, "msg": errorMsg, "error": ""}
	return returnData
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
