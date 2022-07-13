package cmd

import (
	"github.com/fairxio/go/applications/fairx/service/didsrv"
	"github.com/fairxio/go/log"

	"github.com/spf13/cobra"
)

// didCmd represents the dwn command
var didCmd = &cobra.Command{
	Use:   "did",
	Short: "FairX Decentralized Identity Document Service",
	Long:  `The FairX DID Service stores DID Documents for the did:fairx method`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Starting FairX DID Service...")
		didsrv.Start()
	},
}

func init() {
	rootCmd.AddCommand(didCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
