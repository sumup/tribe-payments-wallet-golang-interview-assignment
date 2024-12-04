package cmd

import (
	"context"
	"errors"
	stdOs "os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sumup-oss/go-pkgs/logger"
	"github.com/sumup-oss/go-pkgs/os"
	"go.uber.org/zap"
)

var (
	errShutdown = errors.New("shutdown")
	defSignals  = []stdOs.Signal{syscall.SIGINT, syscall.SIGTERM}
)

type shutdownTaskOptionsFunc = func(*shutdownTaskOptions)

type shutdownTaskOptions struct {
	signals []stdOs.Signal
}

func newShutdownTaskOptions(opts ...shutdownTaskOptionsFunc) *shutdownTaskOptions {
	o := &shutdownTaskOptions{
		signals: defSignals,
	}

	for _, optFn := range opts {
		optFn(o)
	}

	return o
}

// withSignals returns a function that sets the list stdOs.Signal to listen to.
func withSignals(signals []stdOs.Signal) shutdownTaskOptionsFunc {
	return func(o *shutdownTaskOptions) {
		o.signals = signals
	}
}

// shutdownTask is a task that listens for a list of stdOs.Signal and terminates after
// shutdownDeadline.
type shutdownTask struct {
	log              logger.StructuredLogger
	osExecutor       os.OsExecutor
	shutdownDeadline time.Duration
	signals          []stdOs.Signal
}

// newShutdownTask bootstrap a new instance of ShutdownTask.
func newShutdownTask(
	log logger.StructuredLogger,
	osExecutor os.OsExecutor,
	shutdownDeadline time.Duration,
	opts ...shutdownTaskOptionsFunc,
) *shutdownTask {
	taskOpts := newShutdownTaskOptions(opts...)

	return &shutdownTask{
		log:              log,
		osExecutor:       osExecutor,
		shutdownDeadline: shutdownDeadline,
		signals:          taskOpts.signals,
	}
}

// Run runs the ShutdownTask.
func (t *shutdownTask) Run(ctx context.Context) error {
	osSignalChan := make(chan stdOs.Signal, 1)
	signal.Notify(osSignalChan, t.signals...)

	select {
	case sig := <-osSignalChan:
		signal.Stop(osSignalChan)
		t.log.Info("going to try graceful shutdown", zap.String("signal", sig.String()))

		go func() {
			time.Sleep(t.shutdownDeadline)

			t.log.Error(
				"shutdown deadline exceeded, going to terminate forcibly",
				zap.Stringer("deadline", t.shutdownDeadline),
			)

			t.osExecutor.Exit(1)
		}()
	case <-ctx.Done():
	}

	return errShutdown
}
