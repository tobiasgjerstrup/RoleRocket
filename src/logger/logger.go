package logger

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

var Logger *slog.Logger

func Main() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})).WithGroup("info")
}

func (writer wrappedWriter) WriteHeader(statusCode int) {
	writer.ResponseWriter.WriteHeader(statusCode)
	writer.statusCode = statusCode
}

func Middlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		Logger.Debug("Http Call",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("durationUs", int(time.Since(start).Microseconds())),
			slog.Int("statusCode", wrapped.statusCode),
		)
	})
}
