package cmd

import (
	"log/slog"

	"github.com/aliamerj/aircup/api"
	"github.com/aliamerj/aircup/config"
	"github.com/spf13/cobra"
)

var (
	serveAddr string
	serveRoot string
)

var serveCmd = &cobra.Command{
	Use:   "serve [config-file]",
	Short: "Start the Aircup server",
	Run:   runServe,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&serveAddr, "addr", "a", "", "Server listen address")
	serveCmd.Flags().StringVarP(&serveRoot, "root", "r", "", "Root directory to share")
}

func runServe(cmd *cobra.Command, args []string) {
	importPath := ""
	if len(args) == 1 {
		importPath = args[0]
	}

	cfg, err := config.Parse(importPath, config.Config{
		Addr: serveAddr,
		Root: serveRoot,
	})
	if err != nil {
		slog.Error(err.Error())
		return
	}

	slog.Info("starting server", "addr", cfg.Addr, "root", cfg.Root)

	if err := api.Run(*cfg); err != nil {
		slog.Error(err.Error())
		return
	}

}
