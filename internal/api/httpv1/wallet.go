package httpv1

import (
	"net/http"

	"github.com/sumup-oss/go-pkgs/logger"
)

func NewCreateWalletHandler(log logger.StructuredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Info("CreateWalletHandler")

		w.WriteHeader(http.StatusOK)
	}
}
