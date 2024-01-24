package wx

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Demo(openId string) error {
	appId := os.Getenv("APP_ID")
	secretKey := os.Getenv("SECRET_KEY")
	log.WithField("appId", appId).WithField("key", secretKey).Info("test")
	saber := NewWxMiniSubscriber(appId, secretKey)
	payload := SendPayload{
		TakeOrderTmpId,
		NewTakeGoodsMessage("A001", "11.00", "HK.001. LS"),
	}
	err := saber.SendApi(openId, payload)
	if err != nil {
		return err
	}
	return nil
}
