package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
	"github.com/sumup-oss/go-pkgs/os"
	"github.com/sumup-oss/go-pkgs/task"
	"github.com/sumup/payments-bank-account-golang-assignment/internal/config"
)

//nolint:gocognit
func NewApiCmd(osExecutor os.OsExecutor) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Run application server",
		Long:  "Run application server",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			cfg, err := config.NewServerConfig()
			if err != nil {
				return errors.Wrap(err, "failed to create runtime config")
			}

			log, err := logger.NewZapLogger(
				logger.Configuration{
					Level:         cfg.Log.Level,
					Encoding:      logger.EncodingJSON,
					StdoutEnabled: cfg.Log.StdoutEnabled,
				},
			)
			if err != nil {
				return errors.Wrap(err, "failed to create logger")
			}

			defer log.Sync() //nolint:errcheck

			shutdownTask := newShutdownTask(
				log,
				osExecutor,
				cfg.GracefulShutdownTimeout,
			)

			taskGroup := task.NewGroup()

			taskGroup.Go(
				shutdownTask.Run,
			)

			err = taskGroup.Wait(ctx)
			if err != nil {
				return errors.Wrap(err, "task group exits with error")
			}

			log.Info("taskGroup successfully shutdown")

			return nil
		},
	}
}
