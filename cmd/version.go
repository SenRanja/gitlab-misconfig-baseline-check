/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gitlab-misconfig/internal/config"

	"github.com/spf13/cobra"
)

var Version = config.Version

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本",
	Run:   runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Println(Version)
}
