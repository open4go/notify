package wx

import (
	"testing"
)

func TestDemo(t *testing.T) {

	LoadConfigForTest()

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
