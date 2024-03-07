package ctxlogger

import (
	"bytes"
	"context"
	"io"
	"log/slog"

	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/lithammer/shortuuid"
)

// Добавляет в контекст хэндлера логгер
func Logger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			// Встраиваем свой ReqID в контекст
			// Для получения reqID := ctxlogger.GetReqID(ctx)
			ctx := CtxWithNewReqID(r.Context())

			// Встраиваем в логгер поле request_id
			l := logger.With(
				slog.String("req_id", GetReqID(ctx)),
			)

			//Считываем body, что бы вывести его в лог
			rawBody, err := io.ReadAll(r.Body)
			if err != nil {
				l.Error("Ошбика чтения body", "error", err, "func",
					"ctxlogger.Logger")
			}
			// Restore the io.ReadCloser to it's original state
			r.Body = io.NopCloser(bytes.NewBuffer(rawBody))

			l.Info("new request",
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("method", r.Method),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("path", r.URL.Path),
				slog.String("req_body", string(rawBody)),
			)

			// Встраиваем логгер в контекст
			ctx = ContextWithLogger(ctx, l)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				if time.Since(t1) > 10*time.Millisecond {
					l.Warn("Slow request complete",
						slog.Int("status", ww.Status()),
						slog.String("duration", time.Since(t1).String()),
					)
				}

				if ww.Status() != http.StatusOK {
					l.Warn("Result not 200 OK",
						slog.Int("status", ww.Status()),
						slog.String("duration", time.Since(t1).String()),
					)
				} else {
					l.Info("request completed",
						slog.Int("status", ww.Status()),
						slog.String("duration", time.Since(t1).String()),
					)
				}

			}()

			next.ServeHTTP(ww, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

type ctxLogger struct{}

// ContextWithLogger добавляет логгер в контекст
func ContextWithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext извлекает логгер из контекста
func LoggerFromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*slog.Logger); ok {
		return l
	}
	slog.Error("Отсутсвует логгер в контексте", ctx)
	return slog.Default()
}

// key типа для уникального идентификатора запроса
type ctxKeyReqID string

const reqIDKey ctxKeyReqID = "reqID"

// withRequestID добавляет уникальный идентификатор запроса в контекст
func CtxWithNewReqID(ctx context.Context) context.Context {
	reqID := shortuuid.New()
	return context.WithValue(ctx, reqIDKey, reqID)
}

func GetReqID(ctx context.Context) string {
	requestID, ok := ctx.Value(reqIDKey).(string)
	if !ok {
		return "N/A"
	}
	return requestID
}
