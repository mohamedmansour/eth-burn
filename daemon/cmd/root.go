package cmd

import (
	"fmt"

	"github.com/mohamedmansour/ethereum-burn-stats/daemon/hub"
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	var addr string
	var development bool
	var gethEndpointHTTP string
	var gethEndpointWebsocket string
	var dbPath string
	var ropsten bool

	rootCmd := &cobra.Command{
		// TODO:
		Use: "ethereum-burn-stats",
		// TODO:
		Short: "short",
		Long:  `long`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if gethEndpointHTTP == "" {
				cmd.Help()
				return fmt.Errorf("--geth-endpoint-http is required")
			}

			if gethEndpointWebsocket == "" {
				cmd.Help()
				return fmt.Errorf("--geth-endpoint-websocket is required")
			}

			if dbPath == "" {
				cmd.Help()
				return fmt.Errorf("--geth-endpoint-websocket is required")
			}

			return root(
				addr,
				development,
				gethEndpointHTTP,
				gethEndpointWebsocket,
				dbPath,
				ropsten,
			)
		},
	}

	rootCmd.Flags().StringVar(&addr, "addr", ":8080", "HTTP service address")
	rootCmd.Flags().BoolVar(&development, "development", true, "enable for development mode")
	rootCmd.Flags().StringVar(&gethEndpointHTTP, "geth-endpoint-http", "http://localhost:8545", "Endpoint to geth for http")
	rootCmd.Flags().StringVar(&gethEndpointWebsocket, "geth-endpoint-websocket", "ws://localhost:8546", "Endpoint to geth for websocket")
	rootCmd.Flags().StringVar(&dbPath, "db-path", "watchtheburn.db", "Path to the SQLite db")
	rootCmd.Flags().BoolVar(&ropsten, "ropsten", false, "Use ropsten block numbers")

	return rootCmd
}

func root(
	addr string,
	development bool,
	gethEndpointHTTP string,
	gethEndpointWebsocket string,
	dbPath string,
	ropsten bool,
) error {
	hub, err := hub.New(
		development,
		gethEndpointHTTP,
		gethEndpointWebsocket,
		dbPath,
		ropsten,
	)
	if err != nil {
		return err
	}

	err = hub.ListenAndServe(addr)
	if err != nil {
		return err
	}

	return nil
}

// Execute runs cli command
func Execute() error {
	return newRootCmd().Execute()
}
