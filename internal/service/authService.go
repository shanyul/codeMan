package service

import (
	"chujian-api/internal/model"
	"chujian-api/pkg/exception"
	"chujian-api/pkg/gredis"
	"encoding/json"
	"time"
)

type AuthService struct {
	UserModel model.UserModel
}

// CheckByCode 检查用户通过 openid
func (service *AuthService) CheckByCode(openid string, sessionKey string, unionId string) (info map[string]interface{}, responseCode int) {
	authInfo, err := service.UserModel.GetByCode(openid)
	responseCode = exception.Success
	if err != nil {
		responseCode = exception.ErrorLoginParams
		return
	}
	user := model.User{}
	if authInfo.UserId == 0 {
		user.WechatOpenid = openid
		user.SessionKey = sessionKey
		user.UnionId = unionId
		authInfo.UserId = service.UserModel.AddWechatUser(&user)
	} else if sessionKey != "" {
		user.SessionKey = sessionKey
		_ = service.UserModel.EditUser(authInfo.UserId, user)
	}

	if authInfo.UserId == 0 {
		responseCode = exception.ErrorLoginParams
		return
	}

	info = make(map[string]interface{})
	info["userId"] = authInfo.UserId
	info["username"] = authInfo.Username
	info["nickname"] = authInfo.Nickname
	info["avatar"] = authInfo.Avatar
	info["sex"] = authInfo.Sex
	info["bgImage"] = authInfo.BgImage
	info["phone"] = authInfo.Phone
	info["email"] = authInfo.Email
	info["state"] = authInfo.State
	info["createTime"] = authInfo.CreateTime
	info["introduction"] = authInfo.Introduction
	info["token"] = sessionKey

	key := sessionKey
	// 小程序登录保存一个星期过期时间
	ttl := 7 * 24 * time.Hour
	err = gredis.Set(key, info, int(ttl))
	if err != nil {
		responseCode = exception.ErrorLoginParams
		return
	}

	return
}

func (service *AuthService) GetUserInfo(key string) (userInfo model.User) {
	cacheData, err := gredis.Get(key)
	if cacheData != nil && err == nil {
		_ = json.Unmarshal(cacheData, &userInfo)
	}

	return userInfo
}

func (service *AuthService) GetUserById(id int) (userInfo model.User) {
	userInfo, _ = service.UserModel.GetByUserId(id)

	return userInfo
}

func (service *AuthService) IsLogin(key string) model.User {
	var user model.User
	cacheData, err := gredis.Get(key)
	if cacheData != nil && err == nil {
		_ = json.Unmarshal(cacheData, &user)
		return user
	}

	return user
}

func (service *AuthService) Logout(key string) bool {
	result, _ := gredis.Delete(key)

	return result
}
