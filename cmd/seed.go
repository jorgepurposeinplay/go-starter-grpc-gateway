package cmd

import (
	"github.com/purposeinplay/go-commons/logs"
	"github.com/purposeinplay/go-starter-grpc-gateway/internal/adapters/psql"
	"github.com/purposeinplay/go-starter-grpc-gateway/internal/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var seedCmd = &cobra.Command{
	Use:  "seed",
	Long: "Seed database strucutures. This will create new tables and add missing columns and indexes.",
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := logs.NewLogger()
		defer func() {
			_ = logger.Sync()
		}()

		cfg, err := config.LoadConfig(cmd)

		if err != nil {
			logger.Fatal("Unable to read config", zap.Error(err))
		}

		_, err = psql.Connect(cfg)

		if err != nil {
			logger.Fatal("error opening database", zap.Error(err))
		}

	},
}

func init() {
	RootCmd.AddCommand(seedCmd)
}
