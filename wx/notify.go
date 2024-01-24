package wx

import (
	"context"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	// 支付订单模板
	PayTmpId = "XDH55V-yVYDu4_OhIWvYwXDLX5fvjDBUv2IgY6MDPD4"
	// 取餐模板
	TakeOrderTmpId = "bI6fgdcs9lK9YtFHz00NVAinh6hX_0bRFaRl2N6Dl1w"
)

// MiniSubscribeMessage 小程序订阅消息提醒
type MiniSubscribeMessage struct {
	miniProgram *miniprogram.MiniProgram // 小程序
}

type SendPayload struct {
	TemplateId string
	Data       map[string]*subscribe.DataItem
}

func NewWxMiniSubscriber(appId string, secretKey string) *MiniSubscribeMessage {
	// 小程序配置
	config := &miniConfig.Config{
		AppID:     appId,
		AppSecret: secretKey,
	}

	r := os.Getenv("REDIS_ADDR")
	if r == "" {
		panic("please set REDIS_ADDR env")
	}
	cacheRedis := cache.NewRedis(context.Background(), &cache.RedisOpts{
		Host:        r,
		Password:    "",
		Database:    0,
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 0,
	})

	// 初始化小程序模块
	wx := wechat.NewWechat()
	wx.SetCache(cacheRedis)
	miniProgram := wx.GetMiniProgram(config)
	miniProgram.SetAccessTokenHandle(NewTokenHandler()) //  设置微信小程序AccessToken的函数
	return &MiniSubscribeMessage{
		miniProgram: miniProgram,
	}
}

// SendApi 内部接口调用的接口
func (w *MiniSubscribeMessage) SendApi(openId string, p SendPayload) error {
	// 发送通知
	err := w.Send(openId, p.TemplateId, p.Data)
	if err != nil {
		return err
	}
	return nil
}

// Send 发送订阅消息给微信用户
func (w *MiniSubscribeMessage) Send(
	openId string, // 用户openid
	templateId string, // 模版id
	data map[string]*subscribe.DataItem, // 内容数据
) error {

	msg := &subscribe.Message{
		ToUser:           openId,
		TemplateID:       templateId,
		Page:             "",
		Data:             data, // 模板信息
		MiniprogramState: "",
		Lang:             "",
	}

	if err := w.miniProgram.GetSubscribe().Send(msg); err != nil {
		log.WithFields(log.Fields{
			"openId":     openId,
			"templateId": templateId,
			"data":       data,
			"err":        err,
		}).Error("发送订阅消息失败")
		return err
	}

	log.WithFields(log.Fields{"openId": openId, "templateId": templateId, "data": data}).Info("订阅消息发送成功")
	return nil
}

// ListTemplates 查询小程序的模板列表
func (w *MiniSubscribeMessage) ListTemplates() *subscribe.TemplateList {
	listTemplates, err := w.miniProgram.GetSubscribe().ListTemplates()
	if err != nil {
		return nil
	}
	return listTemplates
}
