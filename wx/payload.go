package wx

import (
	rtime "github.com/r2day/base/time"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

// NewPayMessage 下单支付通知
func NewPayMessage(queueNumber string, amount string, location string, orderTime int64) map[string]*subscribe.DataItem {

	return map[string]*subscribe.DataItem{
		"thing5":   {Value: queueNumber},
		"amount12": {Value: amount},
		"thing7":   {Value: location},
		"date4":    {Value: rtime.FomratTimeAsReader(orderTime)},
	}
}

// NewTakeGoodsMessage 创建待取餐消息模板
// 订单状态:{{phrase16.DATA}}
// 商品名:{{thing6.DATA}}
// 订单金额:{{amount13.DATA}}
// 餐厅名称:{{thing17.DATA}}
// 取餐时间:{{time21.DATA}}
func NewTakeGoodsMessage(queueNumber string, amount string, location string) map[string]*subscribe.DataItem {
	// 创建
	return map[string]*subscribe.DataItem{
		"phrase16": {
			Value: "待领取",
		},
		"thing6": {
			Value: queueNumber,
		},
		"amount13": {
			Value: amount,
		},
		"thing17": {
			Value: location,
		},
		"time21": {
			Value: rtime.GetCurrentTime(),
		},
	}
}
