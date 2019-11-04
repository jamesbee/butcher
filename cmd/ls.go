package cmd

import (
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "Show all provider and consumers",
	Run: func(cmd *cobra.Command, args []string) {
		Client.Cmd("ls")
		waitGlobal()
	},
}
