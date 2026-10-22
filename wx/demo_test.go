package wx

import (
	"os"
	"testing"
)

func TestDemo(t *testing.T) {

	// 该测试会调用真实的微信订阅消息接口（需要 APP_ID / SECRET_KEY 等真实配置），默认跳过
	if os.Getenv("WX_NOTIFY_E2E") != "1" {
		t.Skip("skip: set WX_NOTIFY_E2E=1 with APP_ID/SECRET_KEY to run against the real WeChat API")
	}

	type args struct {
		openId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test",
			args{
				"oV4HY5QrREP-qgVrsC3-vF7HvPv0",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Demo(tt.args.openId); (err != nil) != tt.wantErr {
				t.Errorf("Demo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
