package wx

import (
	"github.com/open4go/notify/utils"
	"os"
)

type TokenHandler struct{}

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{}
}

// GetAccessToken 查询小程序的接口调用凭据
func (t *TokenHandler) GetAccessToken() (accessToken string, err error) {
	a, err := utils.GetWxaAccessToken(os.Getenv("APP_ID"), os.Getenv("SECRET_KEY"))
	return a.AccessToken, nil
}

type tokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}
