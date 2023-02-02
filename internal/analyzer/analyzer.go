package analyzer

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/types"
)

type Analyzer interface {
	// 初始化
	//Init(gitlabClient *gitlab.Client, options *types.Options)
	// 自动分析
	AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper)
}
