package logger

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type DBHandler struct {
	db       *sql.DB
	minLevel slog.Level
}

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

var Slog *slog.Logger

type MultiHandler struct {
	handlers []slog.Handler
}

type correlationIDKeyType struct{}

var correlationIDKey = correlationIDKeyType{}

// region multi handlers

func (h *DBHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.minLevel
}

func (h *DBHandler) Handle(ctx context.Context, record slog.Record) error {
	msg := record.Message

	// Convert attributes to a map
	attrMap := make(map[string]any)
	record.Attrs(func(attr slog.Attr) bool {
		if attr.Key == "error" {
			attrMap[attr.Key] = fmt.Sprintf("%v", attr.Value.Any())
		} else {
			attrMap[attr.Key] = attr.Value.Any()
		}
		return true
	})

	// Marshal the map to JSON
	jsonBytes, err := json.Marshal(attrMap)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return err
	}

	cid := ctx.Value(correlationIDKey)

	_, err = h.db.Exec("INSERT INTO logs (log, level, info, correlationId) VALUES ($1, $2, $3, $4)", msg, record.Level.String(), string(jsonBytes), cid)
	return err
}

func (h *DBHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h // Could extend this to store attrs in DB
}

func (h *DBHandler) WithGroup(name string) slog.Handler {
	return h
}

// endregion

// region multi handlers

func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m *MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, h := range m.handlers {
		if h.Enabled(ctx, record.Level) {
			if err := h.Handle(ctx, record); err != nil {
				fmt.Println("Handler error:", err)
			}
		}
	}
	return nil
}

func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	updated := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		updated[i] = h.WithAttrs(attrs)
	}
	return &MultiHandler{handlers: updated}
}

func (m *MultiHandler) WithGroup(name string) slog.Handler {
	updated := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		updated[i] = h.WithGroup(name)
	}
	return &MultiHandler{handlers: updated}
}

//endregion

func Main(db *sql.DB) {
	dbHandler := &DBHandler{db: db, minLevel: slog.LevelDebug}
	stdoutHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn})

	multi := &MultiHandler{handlers: []slog.Handler{stdoutHandler, dbHandler}}
	Slog = slog.New(multi)
}

func (writer *wrappedWriter) WriteHeader(statusCode int) {
	writer.statusCode = statusCode
	writer.ResponseWriter.WriteHeader(statusCode)
}

func (writer *wrappedWriter) Write(b []byte) (int, error) {
	return writer.ResponseWriter.Write(b)
}

func Middlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		cid := uuid.New()
		ctx := context.WithValue(r.Context(), correlationIDKey, cid)
		r = r.WithContext(ctx)

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		if wrapped.statusCode >= 500 {
			Error(r.Context(), "Http Call",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("durationUs", int(time.Since(start).Microseconds())),
				slog.Int("statusCode", wrapped.statusCode),
			)
		} else {
			Info(r.Context(), "Http Call",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("durationUs", int(time.Since(start).Microseconds())),
				slog.Int("statusCode", wrapped.statusCode),
			)
		}
	})
}

// ? 	LogLevels
// ?	LevelDebug Level = -4
// ? 	LevelInfo  Level = 0
// ?	LevelWarn  Level = 4
// ?	LevelError Level = 8

func Debug(ctx context.Context, msg string, args ...any) {
	Slog.Log(ctx, -4, msg, args...)
}

func Info(ctx context.Context, msg string, args ...any) {
	Slog.Log(ctx, 0, msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	Slog.Log(ctx, 4, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	Slog.Log(ctx, 8, msg, args...)
}
