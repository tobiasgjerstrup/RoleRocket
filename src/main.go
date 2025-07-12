package main

import (
	"log/slog"
	"net/http"
	"os"
	"rolerocket/routes"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

var logger *slog.Logger

func main() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})).WithGroup("info")
	router := routes.Routes()

	server := http.Server{
		//Addr: ":8080",
		Addr:    "localhost:8080", // ? use localhost:xxxx to make it not ask for admin permissions
		Handler: logging(router),
	}
	server.ListenAndServe()
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		logger.Debug("Http Call",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("durationUs", int(time.Since(start).Microseconds())),
			slog.Int("statusCode", wrapped.statusCode),
		)
	})
}
