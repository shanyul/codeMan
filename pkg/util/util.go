package util

import (
	"chujian-api/pkg/app"
	"chujian-api/pkg/setting"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/validation"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
	// set valid
	validation.SetDefaultMessage(app.MessageTmp)
}

func StringToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
