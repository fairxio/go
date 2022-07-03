/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vmCmd represents the vm command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "FairX Protocol Virtual Machine",
	Long:  `The FairX Protocol Virtual Machine executes FairX-protocol compatible WASI packages within a FairX-protocol compliant environment`,
	Run: func(cmd *cobra.Command, args []string) {
		// For now exit.
		fmt.Println("vm called")
	},
}

func init() {
	rootCmd.AddCommand(vmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
