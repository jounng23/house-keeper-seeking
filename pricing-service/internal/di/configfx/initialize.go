package configfx

import (
	"go.uber.org/fx"

	"pricing-svc/pkg/config"
)

func Initialize(configFile string) fx.Option {
	return fx.Invoke(func() {
		config.InitConfig(configFile)
	})
}
