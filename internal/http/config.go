package http

import (
	"crypto/tls"
	"time"
)

const (
	defaultName              = "api"
	defaultPath              = "/"
	defaultReadTimeout       = 60 * time.Second
	defaultReadHeaderTimeout = 60 * time.Second
	defaultWriteTimeout      = 60 * time.Second
	defaultIdleTimeout       = 60 * time.Second
	defaultMaxHeaderBytes    = 1 << 20 // 1 MB
	defaultShutdownTimeout   = 60 * time.Second
)

type config struct {
	name              string
	path              string
	tlsConfig         *tls.Config
	readTimeout       time.Duration
	readHeaderTimeout time.Duration
	writeTimeout      time.Duration
	idleTimeout       time.Duration
	maxHeaderBytes    int
	shutdownTimeout   time.Duration
}

type option func(*config)

func WithName(name string) option {
	return func(cfg *config) {
		cfg.name = name
	}
}

func WithReadTimeout(readTimeout time.Duration) option {
	return func(cfg *config) {
		cfg.readTimeout = readTimeout
	}
}

func WithReadHeaderTimeout(readHeaderTimeout time.Duration) option {
	return func(cfg *config) {
		cfg.readHeaderTimeout = readHeaderTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) option {
	return func(cfg *config) {
		cfg.writeTimeout = writeTimeout
	}
}

func WithMaxHeaderBytes(maxHeaderBytes int) option {
	return func(cfg *config) {
		cfg.maxHeaderBytes = maxHeaderBytes
	}
}

func WithServerShutdownTimeout(shutdownTimeout time.Duration) option {
	return func(cfg *config) {
		cfg.shutdownTimeout = shutdownTimeout
	}
}
