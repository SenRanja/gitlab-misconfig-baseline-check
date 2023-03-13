package analyzer

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/types"
)

type Analyzer interface {
	AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper, output *types.Output)
}
