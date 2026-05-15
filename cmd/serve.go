package cmd

import (
	"fmt"
	"log/slog"

	"github.com/aliamerj/aircup/api"
	"github.com/aliamerj/aircup/config"
	"github.com/aliamerj/aircup/network"
	"github.com/skip2/go-qrcode"
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

	slog.Info("start mDNS", "addr", cfg.Addr)
	adv, err := network.StartMDNS(cfg.Addr)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer adv.Close()

	sharedUrl, err := network.ShareURL(cmd.Context())
	if err != nil {
		slog.Error(err.Error())
		return
	}

	if err := printQRCode(sharedUrl); err != nil {
		slog.Error(err.Error())
		return
	}

	slog.Info("starting server local", "addr", cfg.Addr, "root", cfg.Root)
	if err := api.Run(*cfg); err != nil {
		slog.Error(err.Error())
		return
	}
}

func printQRCode(url string) error {
	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Scan to open Aircup:")
	fmt.Println()

	fmt.Println(qr.ToSmallString(false))

	fmt.Println(url)
	fmt.Println()

	return nil
}
