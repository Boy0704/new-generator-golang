package cmd

import (
	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/cmd/http"
	"github.com/spf13/cobra"
	"os"
)

func Start() {
	var rootCmd = &cobra.Command{Use: "4dw command"}

	var allCmd = &cobra.Command{
		Use:   "http",
		Short: "Run HTTP",
		Run: func(cmd *cobra.Command, args []string) {
			http.Start()
		},
	}

	rootCmd.AddCommand(allCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
		os.Exit(1)
	}
}
