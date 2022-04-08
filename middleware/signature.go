package middleware

import (
	"chujian-api/internal/controller"
	"chujian-api/internal/param"
	"chujian-api/pkg/app"
	"chujian-api/pkg/exception"
	"chujian-api/pkg/setting"
	"chujian-api/pkg/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Signature() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 本地环境不需要验签
		if setting.AppSetting.Environment != "local" {
			// 获取参数
			var (
				data interface{}
				code int
			)

			for i := 0; i < 1; i++ {
				/*checkSign := c.GetHeader("Signature")
				if checkSign == "" {
					code = exception.ErrorSignatureValidate
					break
				}*/

				body, _ := ioutil.ReadAll(c.Request.Body)
				var request param.RequestParam
				err := json.Unmarshal(body, &request)
				if err != nil {
					code = exception.ErrorSignatureValidate
					break
				}

				/*md5Str := util.StringToMd5(request.Request + "&" + setting.AppSetting.SignKey)
				logging.Info("checkSign:", checkSign, "newSign", md5Str)
				if strings.Compare(checkSign, md5Str) != 0 {
					code = exception.ErrorSignatureValidate
					break
				}*/

				result, err := util.PriDecrypt(request.Request)
				if err != nil || result == "" {
					code = exception.ErrorSignatureValidate
					break
				}
				controller.RequestData = result

				var bodyParam param.BodyParam
				err = json.Unmarshal([]byte(result), &bodyParam)
				if err != nil {
					code = exception.ErrorSignatureValidate
					break
				}

				/*if time.Now().Unix() > (bodyParam.Timestamp + 300) {
					code = exception.ErrorSignatureExpire
					break
				}*/
			}

			if code != 0 {
				app.Response(c, http.StatusBadRequest, code, data, "")
				c.Abort()
				return
			}
		} else {
			body, _ := ioutil.ReadAll(c.Request.Body)
			controller.RequestData = string(body)
		}
		c.Next()
	}
}
