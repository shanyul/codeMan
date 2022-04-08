package util

import (
	"chujian-api/pkg/setting"
	"github.com/wenzhenxi/gorsa"
)

// PubEncrypt 加密
func PubEncrypt(data string) (encrypt string, err error) {
	encrypt, err = gorsa.PublicEncrypt(data, setting.RsaSetting.PublicKey)
	return
}

// PriDecrypt 解密
func PriDecrypt(data string) (decrypt string, err error) {
	decrypt, err = gorsa.PriKeyDecrypt(data, setting.RsaSetting.PrivateKey)
	return
}
