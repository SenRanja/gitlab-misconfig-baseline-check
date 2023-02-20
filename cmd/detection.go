/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"gitlab-misconfig/internal/config"
	"gitlab-misconfig/internal/engine"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"
)

// detectionCmd represents the detection command
var detectionCmd = &cobra.Command{
	Use:   "detection",
	Short: "detect gitlab misconfig ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ParseOptions(options)
		//fmt.Println("detection called")
		gitlabClient := GitlabAuthClientInit()
		//analyzer.Analysis(gitlabClient, options)
		engine.NewEngine().Analysis(gitlabClient, options)
	},
}

func init() {
	rootCmd.AddCommand(detectionCmd)
}

func ParseOptions(options *types.Options) {
	options.Version = config.Version
	options.Url = url
	options.Token = token
	options.ProjectId = projectId
	options.Mode = mode
	options.Check = check
	options.RulePath = rulePath
}

func GitlabAuthClientInit() *gitlab.Client {
	//git, err := gitlab.NewClient(token,gitlab.WithBaseURL("https://gitlab.example.com/api/v4"), gitlab.WithHTTPClient(httpClient))
	git, err := gitlab.NewClient(options.Token, gitlab.WithBaseURL(options.Url))
	if err != nil {
		log.Error("Failed to create client")
		log.Error(err)
	}
	return git
}
