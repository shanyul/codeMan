package service

import qrcode "github.com/skip2/go-qrcode"

type IndexService struct{}

func (service *IndexService) GetCode(codeString string) (code []byte, err error) {
	code, err = qrcode.Encode(codeString, qrcode.Medium, 256)
	return
}
