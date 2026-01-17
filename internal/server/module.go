package server

import (
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Module("server",
	fx.Provide(
		NewGinEngine,
		NewHTTPServer,
	),
	fx.Invoke(RegisterRoutes,
		func(_ *http.Server) {},
	),
)
