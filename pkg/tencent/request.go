package tencent

import (
	"chujian-api/pkg/setting"
	"chujian-api/pkg/util"
	"encoding/json"
	"fmt"
)

type Tencent struct{}

const (
	wechatApi  = "https://api.q.qq.com"
	checkParam = "appid=%s&secret=%s"
)

type Code2SessionResponseForm struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// Code2Session 小程序登录
func (*Tencent) Code2Session(code string) (Code2SessionResponseForm, error) {
	params := fmt.Sprintf(checkParam+"&grant_type=authorization_code&js_code=%s", setting.TencentSetting.AppId, setting.TencentSetting.AppSecret, code)
	requestUrl := wechatApi + "/sns/jscode2session?" + params
	requestData, err := util.Get(requestUrl, 5)
	var response Code2SessionResponseForm
	if err != nil {
		return response, err
	}
	_ = json.Unmarshal([]byte(requestData), &response)

	return response, nil
}
