package configfx

import (
	"go.uber.org/fx"

	"sending-svc/pkg/config"
)

func Initialize(configFile string) fx.Option {
	return fx.Invoke(func() {
		config.InitConfig(configFile)
	})
}
