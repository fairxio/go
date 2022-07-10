package cmd

import (
	"github.com/fairxio/go/applications/fairx/service/auth"
	"github.com/fairxio/go/log"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "FairX Authentication Service",
	Long:  `The FairX Authentication Service is responsible for authenticating callers to the FairX Node`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Starting FairX Authentication Service...")
		auth.Start()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
