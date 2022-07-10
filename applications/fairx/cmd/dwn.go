package cmd

import (
	"github.com/fairxio/go/applications/fairx/service/dwn"
	"github.com/fairxio/go/log"

	"github.com/spf13/cobra"
)

// dwnCmd represents the dwn command
var dwnCmd = &cobra.Command{
	Use:   "dwn",
	Short: "FairX Decentralized Web Node",
	Long:  `The FairX Decentralized Web Node implements the identity.foundations spec`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Starting FairX Decentralized Web Node...")
		dwn.Start()
	},
}

func init() {
	rootCmd.AddCommand(dwnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dwnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dwnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
