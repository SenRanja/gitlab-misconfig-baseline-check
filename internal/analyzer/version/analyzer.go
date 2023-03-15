package version

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"
	"strings"

	"github.com/spf13/viper"
)

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper, output *types.Output) {

	// 【版本 version】
	version, err := VersionDetect(gitlabClient)
	if err != nil {
		fmt.Println(err)
	}
	log.Info("版本: ", version.Version, "校订版本: ", version.Revision)

	output.Version.Version = version.Version
	output.Version.Revision = version.Revision
	if strings.Contains(output.Version.Version, "ee") {
		output.Version.VersionIsEE = true
	} else {
		output.Version.VersionIsEE = false
	}

	// 版本是否存在风险，此处暂设置为不存在风险
	output.Version.CheckRule = "版本"
	output.Version.SecondCheckRule = "版本风险检测"
	output.Version.Result = output.Version.Version + " " + output.Version.Revision
	output.Version.VersionExistRisk = false
	if !output.Version.VersionExistRisk {
		output.Version.Complaince = true
	} else {
		output.Version.Complaince = false
	}
	output.Version.Description = "未检测到当前gitlab版本存在风险"
	output.Version.Advice = "建议"

}
