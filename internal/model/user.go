package model

import "gorm.io/gorm"

type UserModel struct{}

type User struct {
	BaseModel
	UserId          int    `gorm:"primary_key" column:"user_id" json:"userId"`
	Username        string `column:"username" json:"username"`
	Password        string `column:"password" json:"password"`
	Nickname        string `column:"nickname" json:"nickname"`
	Avatar          string `column:"avatar" json:"avatar"`
	Sex             int    `column:"sex" json:"sex"`
	BgImage         string `column:"bg_image" json:"bgImage"`
	Phone           string `column:"phone" json:"phone"`
	Email           string `column:"email" json:"email"`
	State           string `column:"state" json:"state"`
	Country         string `column:"country" json:"country"`
	Province        string `column:"province" json:"province"`
	City            string `column:"city" json:"city"`
	Distinct        string `column:"distinct" json:"distinct"`
	Address         string `column:"address" json:"address"`
	Remark          string `column:"remark" json:"remark"`
	WechatCode      string `column:"wechat_code" json:"wechatCode"`
	WechatOpenid    string `column:"wechat_openid" json:"wechatOpenid"`
	UnionId         string `column:"union_id" json:"unionId"`
	SessionKey      string `column:"session_key" json:"sessionKey"`
	Profession      string `column:"profession" json:"profession"`
	Charge          string `column:"charge" json:"charge"`
	Introduction    string `column:"introduction" json:"introduction"`
	DeleteTimestamp int    `column:"delete_timestamp" json:"deleteTimestamp"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "user"
}

// CheckAuth 验证用户
func (*UserModel) CheckAuth(username, password string) (*User, bool) {
	var auth User
	err := dbHandle.Select(
		"user_id", "username", "nickname", "avatar", "bg_image", "phone", "email", "state", "province", "city", "distinct", "address", "create_time",
	).Where(User{Username: username, Password: password}).First(&auth).Error
	if err != nil {
		return nil, false
	}

	return &auth, auth.UserId > 0
}

func (*UserModel) GetByUserId(id int) (User, error) {
	var user User
	err := dbHandle.Select(
		"user_id", "username", "nickname", "password", "avatar", "bg_image", "phone", "email", "state", "province", "city", "distinct", "address", "create_time",
	).Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// AddUser 验证用户
func (*UserModel) AddUser(data *User) error {
	if err := dbHandle.Select(
		"username",
		"password",
		"nickname",
	).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

// AddUser 验证用户
func (*UserModel) AddWechatUser(data *User) int {
	if err := dbHandle.Select(
		"wechat_openid",
		"union_id",
		"nickname",
		"sex",
		"country",
		"province",
		"avatar",
		"city",
	).Create(&data).Error; err != nil {
		return 0
	}

	return data.UserId
}

func (*UserModel) EditUser(id int, data User) error {
	if err := dbHandle.Model(&User{}).Where("user_id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// GetByCode 通过 code 获取用户信息
func (*UserModel) GetByCode(code string) (User, error) {
	var user User
	err := dbHandle.Select(
		"user_id", "username", "nickname", "password", "avatar", "bg_image", "phone", "email", "state", "province", "city", "distinct", "address", "create_time",
	).Where("wechat_openid = ?", code).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return User{}, err
	}

	return user, nil
}
