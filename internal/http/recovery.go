package http

import (
	"net/http"
	"runtime/debug"

	"github.com/sumup-oss/go-pkgs/logger"
	"go.uber.org/zap"
)

func Recovery(
	log logger.StructuredLogger,
	responseWriter http.HandlerFunc,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				if err != nil && err != http.ErrAbortHandler {
					logErr, ok := err.(error)
					if ok {
						log.Error("panic", logger.ErrorField(logErr), zap.String("trace", string(debug.Stack())))
					} else {
						log.Error("Internal server error")
					}

					responseWriter(w, r)
				}
			}()

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
