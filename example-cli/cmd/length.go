/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lengthCmd represents the length command
var lengthCmd = &cobra.Command{
	Use:   "length",
	Short: "Retrieves length of a string input.",
	Long: `Retrieves length of a string input. For example:

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("length called")
	},
}

func init() {
	rootCmd.AddCommand(lengthCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lengthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lengthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
