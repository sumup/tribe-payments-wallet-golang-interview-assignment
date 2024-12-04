package http

import (
	"context"
	"net/http"
	"sync"

	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
	"go.uber.org/zap"
)

// Server is a HTTP server.
type Server struct {
	server *http.Server
	log    logger.StructuredLogger
	cfg    *config
}

// NewServer creates a new Server.
func NewServer(
	log logger.StructuredLogger,
	listenAddr string,
	mux http.Handler,
	options ...option,
) *Server {
	cfg := &config{
		name:              defaultName,
		path:              defaultPath,
		tlsConfig:         nil,
		readTimeout:       defaultReadTimeout,
		readHeaderTimeout: defaultReadHeaderTimeout,
		writeTimeout:      defaultWriteTimeout,
		idleTimeout:       defaultIdleTimeout,
		maxHeaderBytes:    defaultMaxHeaderBytes,
		shutdownTimeout:   defaultShutdownTimeout,
	}

	for _, opt := range options {
		opt(cfg)
	}

	httpServer := &http.Server{
		Handler:           mux,
		Addr:              listenAddr,
		TLSConfig:         cfg.tlsConfig,
		ReadTimeout:       cfg.readTimeout,
		ReadHeaderTimeout: cfg.readHeaderTimeout,
		WriteTimeout:      cfg.writeTimeout,
		IdleTimeout:       cfg.idleTimeout,
		MaxHeaderBytes:    cfg.maxHeaderBytes,
	}

	server := &Server{
		server: httpServer,
		log:    log,
		cfg:    cfg,
	}

	return server
}

// Run starts the server.
func (s *Server) Run(ctx context.Context) error {
	s.log = s.log.With(zap.String("name", s.cfg.name))

	var wg sync.WaitGroup

	wg.Add(1)

	doneCh := make(chan struct{})

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			s.log.Info(
				"Received shutdown signal, going to shutdown HTTP server",
				zap.String("address", s.server.Addr),
			)

			ctxTimeout, cancel := context.WithTimeout(context.Background(), s.cfg.shutdownTimeout)
			defer cancel()

			_ = s.server.Shutdown(ctxTimeout) //nolint:contextcheck
		case <-doneCh:
		}
	}()

	defer wg.Wait()
	defer close(doneCh)

	s.log.Info("Starting HTTP server", zap.String("address", s.server.Addr))

	err := s.server.ListenAndServe()

	// NOTE: ListenAndServe always returns an error.
	s.log.Info(
		"HTTP server shutdown",
		zap.String("address", s.server.Addr),
		logger.ErrorField(err),
	)

	return errors.Wrap(err, "server stopped")
}
