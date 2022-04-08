package param

// RequestParam 签名必须参数
type RequestParam struct {
	Request string `json:"_request" valid:"Required"` // 必填项
}

type BodyParam struct {
	Timestamp int64 `json:"timestamp" valid:"Required"` // 时间戳
}
