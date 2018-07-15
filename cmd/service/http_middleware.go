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
		Str("server", "http1").
		Str("host", host).
		Logger()

	service.AddHttpMiddlware(hlog.NewHandler(_log))

	service.AddHttpMiddlware(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))

	service.AddHttpMiddlware(hlog.RemoteAddrHandler("ip"))
	service.AddHttpMiddlware(hlog.UserAgentHandler("user_agent"))
	service.AddHttpMiddlware(hlog.RefererHandler("referer"))
	service.AddHttpMiddlware(hlog.RequestIDHandler("req_id", "Request-Id"))

}
