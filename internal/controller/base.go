package controller

var RequestData string
var UserInfo = &UserData{}

type UserData struct {
	UserId     int
	Username   string
	Nickname   string
	Avatar     string
	Sex        int
	BgImage    string
	Phone      string
	Email      string
	State      string
	SessionKey string
}

// Base 注册所有控制器
type Base struct {
	IndexApi Index
}
