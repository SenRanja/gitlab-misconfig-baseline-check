package project

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/types"
)

var opt gitlab.ListOptions
var listProjectsOptions *gitlab.ListProjectsOptions

func init() {
	opt = gitlab.ListOptions{
		Page:    1,
		PerPage: 2000,
	}

	listProjectsOptions = &gitlab.ListProjectsOptions{
		ListOptions: opt,
	}

}

type Analyzer struct {
}

// New 创建java解析器
func New() Analyzer {
	return Analyzer{}
}

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {
	//panic("implement me")
}
