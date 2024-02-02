package httpfx

import (
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	provideHttpClient,
)

func provideHttpClient() *http.Client {
	return &http.Client{}
}
