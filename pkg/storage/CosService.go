package storage

import (
	"chujian-api/pkg/setting"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const CosUrl = "https://s1-1309644651.cos.ap-shanghai.myqcloud.com"

func setClient() *cos.Client {
	u, _ := url.Parse(CosUrl)
	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  setting.TencentSetting.CosKey,
			SecretKey: setting.TencentSetting.CosSecret,
		},
	})
}

// UploadFile 上传文件
func UploadFile(path string, file *multipart.FileHeader) (url string, err error) {
	// 获取客户端实例
	client := setClient()
	fd, err := file.Open()
	if err != nil {
		return
	}
	defer fd.Close()
	_, err = client.Object.Put(context.Background(), path, fd, nil)
	if err != nil {
		return
	}

	return CosUrl + "/" + path, nil
}

func GetPresignedURL(name string) string {
	if strings.HasPrefix(name, "/") {
		name = fmt.Sprintf("assert%s", name)
	} else {
		name = fmt.Sprintf("assert/%s", name)
	}
	client := setClient()
	ak := setting.TencentSetting.CosKey
	sk := setting.TencentSetting.CosSecret
	ctx := context.Background()
	presignedURL, err := client.Object.GetPresignedURL(ctx, http.MethodGet, name, ak, sk, time.Hour, nil)
	if err != nil {
		return ""
	}

	return presignedURL.String()
}
