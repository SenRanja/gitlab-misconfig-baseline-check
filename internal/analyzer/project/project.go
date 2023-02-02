package project

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/types"
)

type Analyzer struct {
}

// New 创建java解析器
func New() Analyzer {
	return Analyzer{}
}

//func (Analyzer) Init(gitlabClient *gitlab.Client, options *types.Options) {
//	//users := getGitlabUsers(gitlabClient)
//	//log.Info("gitlab total users is :" + string(len(users)))
//	log.Info("project analyzer")
//}

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {
	//panic("implement me")
}
