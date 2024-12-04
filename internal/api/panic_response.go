package api

import (
	stdHTTP "net/http"

	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

func WritePanicResponse(log logger.StructuredLogger) stdHTTP.HandlerFunc {
	return func(w stdHTTP.ResponseWriter, r *stdHTTP.Request) {
		w.WriteHeader(stdHTTP.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")

		err := errors.New("internal server error")
		if err != nil {
			log.Error(
				"error",
				logger.ErrorField(err),
			)
		}
	}
}
