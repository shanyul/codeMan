package exception

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrorLoginParams = 10000
	ErrorAddData     = 20000
	ErrorEditData    = 20001
	ErrorDeleteData  = 20002

	ErrorTokenValidate     = 30000
	ErrorTokenExpire       = 30001
	ErrorSignatureValidate = 30002
	ErrorSignatureExpire   = 30003

	ErrorGetDataFail = 40000
)

var MsgFlags = map[int]string{
	Success:                "ok",
	Error:                  "fail",
	InvalidParams:          "请求参数错误",
	ErrorLoginParams:       "登录失败",
	ErrorAddData:           "新增失败",
	ErrorEditData:          "修改失败",
	ErrorDeleteData:        "删除失败",
	ErrorTokenValidate:     "登录验证失败",
	ErrorTokenExpire:       "登录已过期",
	ErrorSignatureValidate: "签名验证失败",
	ErrorSignatureExpire:   "签名已过期",
	ErrorGetDataFail:       "获取数据失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
