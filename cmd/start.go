package cmd

import (
	"github.com/johnfarrell/datastar-playground/internal/router"
	"github.com/johnfarrell/datastar-playground/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the DSP server",
	Long:  "Main startup command for the DSP.",
	RunE:  startRunE,
}

// init adds necessary flags for the start command.
func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().IntP("port", "p", 8080, "DSP server port")
	_ = viper.BindPFlag("port", startCmd.PersistentFlags().Lookup("port"))
}

func startRunE(cmd *cobra.Command, args []string) error {
	serv := server.New(logger)
	rout := router.New(logger)

	return serv.Run(cmd.Context(), rout.Routes(cmd.Context()))
}
