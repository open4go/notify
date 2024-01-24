package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/r2day/db"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

// 获取微信小程序的AccessToken
const (
	wxaAccessTokenKey    = "wxa_access_token_%s"
	getWxaAccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

type GetWxaAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// getWxaAccessToken 获取微信小程序的AccessToken
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
func getWxaAccessToken(appId, appSecret string) (*GetWxaAccessTokenResp, error) {
	urlStr := fmt.Sprintf(getWxaAccessTokenUrl, appId, appSecret)
	wxResp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer wxResp.Body.Close()

	b, err := ioutil.ReadAll(wxResp.Body)
	if err != nil {
		return nil, err
	}

	var resp GetWxaAccessTokenResp
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	if resp.ErrCode != 0 {
		return nil, fmt.Errorf("GetWxaAccessToken, errcode=%d, errmsg=%s", resp.ErrCode, resp.ErrMsg)
	}
	return &resp, nil
}

// GetWxaAccessToken 获取小程序AccessToken
func GetWxaAccessToken(appId, appSecret string) (*GetWxaAccessTokenResp, error) {
	// 缓存的token无效的话，则重新发起请求
	for i := 0; i < 3; i++ {
		resp, err := getWxaAccessToken(appId, appSecret)
		if err != nil {
			log.Errorf("GetWxaAccessToken, err=%v", err)
			continue
		}
		return resp, nil
	}
	return nil, fmt.Errorf("获取微信小程序AccessToken失败")
}

// GetWxaAccessTokenWithCache 缓存中获取AccessToken
func GetWxaAccessTokenWithCache(ctx context.Context, appId, appSecret string) (string, error) {
	// 读取缓存
	token := db.RDB.Get(ctx, fmt.Sprintf(wxaAccessTokenKey, appId)).Val()
	if token != "" {
		return token, nil
	}

	// 缓存中没有，则重新获取
	resp, err := GetWxaAccessToken(appId, appSecret)
	if err != nil {
		return "", err
	}

	// 设置缓存，缓存提前5分钟失效
	db.RDB.Set(ctx, fmt.Sprintf(wxaAccessTokenKey, appId),
		resp.AccessToken, time.Duration(resp.ExpiresIn-5*60)*time.Second)

	return resp.AccessToken, err
}
