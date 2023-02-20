package project

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"
)

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {

	per_page_items := config.GetInt("projects.per_page_items.keywords")
	max_acquire_items := config.GetInt("projects.max_acquire_items.keywords")

	log.Info("[#] 项目配置分析检测开始")
	AllProjects, projectsService := ListAllProjects(gitlabClient, per_page_items, max_acquire_items)
	for _, single_project := range AllProjects {
		log.Info("[###] 当前检查project:", single_project.ID, single_project.Name)
		projectVisibility := ProjectVisibility(single_project)
		log.Info("[###] 可见性:", projectVisibility)
		projectSecurityAndCompliance := ProjectSecurityAndCompliance(single_project)
		log.Info("[###] 安全与合规:", projectSecurityAndCompliance)
		projectApprovalsBeforeMerge := ProjectApprovalsBeforeMerge(single_project)
		log.Info("[###] 批准合并请求人数:", projectApprovalsBeforeMerge)
		projectRejectUnsignedCommits := ProjectRejectUnsignedCommits(projectsService, single_project.ID)
		log.Info("[###] 未通过 GPG 签名时拒绝提交", projectRejectUnsignedCommits)
		projectCommitCommitterCheck := ProjectCommitCommitterCheck(projectsService, single_project.ID)
		log.Info("[###] 禁止未经验证的用户提交commit", projectCommitCommitterCheck)
	}
	log.Info("[#] 项目配置分析检测完毕")
}
