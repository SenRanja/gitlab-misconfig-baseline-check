/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab-misconfig/internal/types"
	"os"
)

const banner = "       _ _   _       _                 _                      __ _       \n" +
	"      (_) | | |     | |               (_)                    / _(_)      \n  " +
	"__ _ _| |_| | __ _| |__    _ __ ___  _ ___  ___ ___  _ __ | |_ _  __ _ \n " +
	"/ _` | | __| |/ _` | '_ \\  | '_ ` _ \\| / __|/ __/ _ \\| '_ \\|  _| |/ _` |\n" +
	"| (_| | | |_| | (_| | |_) | | | | | | | \\__ \\ (_| (_) | | | | | | | (_| |\n " +
	"\\__, |_|\\__|_|\\__,_|_.__/  |_| |_| |_|_|___/\\___\\___/|_| |_|_| |_|\\__, |\n  __/ |                                                             __/ |\n |___/                                                             |___/ \n"

var (
	options = &types.Options{}
)

var (
	token     string
	url       string
	projectId string
	mode      string
	check     string
	//logLevel  string
	rulePath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitlab-misconfig",
	Short: "gitlab ee misconfig detection tools",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Fprint(os.Stderr, banner)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "gitlab access token.")
	// token 为必填项
	rootCmd.MarkFlagRequired("token")
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "gitlab url address")
	// url 为必填项
	rootCmd.MarkFlagRequired("url")
	rootCmd.PersistentFlags().StringVar(&projectId, "id", "all", "scan gitlab project,projectid or all")
	//  模式默认为baseline，也可以是inventory，详细模式
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "baseline", "sacn mode")
	rootCmd.PersistentFlags().StringVarP(&check, "check", "c", "project", "is check project")
	rootCmd.PersistentFlags().StringVarP(&rulePath, "rule", "r", "", "rule")
}
