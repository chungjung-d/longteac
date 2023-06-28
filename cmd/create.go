/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/chungjung-d/longteac/config"
	"github.com/chungjung-d/longteac/process"
	"github.com/spf13/cobra"
)

var socketPath string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the longtea container root process",
	Long:  `Create the longtea container root process`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()

		ctx = context.WithValue(ctx, config.SocketPath, socketPath)

		process.StartContainerRootProcess(ctx)
	},

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if socketPath == "" {
			return fmt.Errorf("the 'socket' flag is required")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&socketPath, "socket", "s", "", "The socket address (required)")
}
