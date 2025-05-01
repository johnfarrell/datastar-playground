package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Create a package-level logger here so we can initialize it all in one place
// and access it from subcommands.
// This is passed down to the actual logic of the commands in the recommended way.
var logger *zap.Logger

var rootCmd = &cobra.Command{
	Use:   "dsp",
	Short: "Datastar playground application",
	Long:  "Playground application for learning the Datastar framework.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := validateConfig(); err != nil {
			return err
		}
		if err := initLogger(); err != nil {
			return err
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

// init sets up bindings between Cobra and Viper
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().IntP("loglevel", "l", 1, "Log level (0 - Debug, 1 - Info, 2 - Warn, 3 - Error")
	_ = viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
}

// initConfig sets up the Viper configuration.
func initConfig() {
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// validateConfig keeps all of our validation logic into one place.
// TODO this could probably use a proper Viper struct or something?
func validateConfig() error {
	loglevel := viper.GetInt("loglevel")
	if loglevel < 0 || loglevel > 3 {
		return fmt.Errorf("invalid loglevel [%d], must be be in range [0,3]", loglevel)
	}
	return nil
}

// initLogger set up zap depending on what log level is specified.
// If the log level is 0 (DEBUG), we set a sugared zap.SugaredLogger.
// Otherwise, we use the standard zap.Logger.
func initLogger() error {
	loglevel := viper.GetInt("loglevel")

	var config zap.Config
	if loglevel == 0 {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		return fmt.Errorf("failed to build logger config: %w", err)
	}

	return nil
}
