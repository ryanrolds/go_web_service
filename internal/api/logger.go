package api

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fields := log.Fields{
			"remote_addr": r.RemoteAddr,
			"method":      r.Method,
			"uri":         r.RequestURI,
		}
		requestLogger := log.WithFields(fields)
		requestLogger.Debug("request received")

		next.ServeHTTP(w, r)

		duration := time.Since(start).Microseconds()
		requestLogger.WithField("duration_us", duration).Info("request served")
	})
}
