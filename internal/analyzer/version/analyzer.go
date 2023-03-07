package version

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"

	"github.com/spf13/viper"
)

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {

	// 【版本 version】
	version, err := VersionDetect(gitlabClient)
	if err != nil {
		fmt.Println(err)
	}
	log.Info("版本: ", version.Version, "校订版本: ", version.Revision)

}
