package main

import (
	"net/http"
	"time"

	// "github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func httpMiddleware() {

	_log := log.With().
		Str("server", "http").
		Str("host", host).
		Logger()

	service.AddHttpHandler(hlog.NewHandler(_log))

	service.AddHttpHandler(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))

	service.AddHttpHandler(hlog.RemoteAddrHandler("ip"))
	service.AddHttpHandler(hlog.UserAgentHandler("user_agent"))
	service.AddHttpHandler(hlog.RefererHandler("referer"))
	service.AddHttpHandler(hlog.RequestIDHandler("req_id", "Request-Id"))

}
