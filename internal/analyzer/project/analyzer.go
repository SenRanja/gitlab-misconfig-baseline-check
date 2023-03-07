package project

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"
)

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {

	log.Debug("[#] 项目配置分析检测开始")
	AllProjects, projectsService, approvalService, projectMembersService, protectedBranchesService := ListAllProjects(gitlabClient)
	//AllProjects, _, approvalService, projectMembersService := ListAllProjects(gitlabClient)
	for _, single_project := range AllProjects {
		log.Info("目前正在检查项目 id:", single_project.ID, "Name:", single_project.NameWithNamespace)
		projectVisibility := ProjectVisibility(single_project)
		if projectVisibility == "public" {
			log.Info("违规，禁止建立public项目")
		}

		require_authentication_to_view_media_files := RequireAuthenticationToViewMediaFiles(single_project)
		if require_authentication_to_view_media_files != true {
			log.Info("违规，访问tag媒体文件链接时未验证登录，会导致tag内容的媒体文件连接被直接访问")
		}

		projectSecurityAndCompliance := ProjectSecurityAndCompliance(single_project)
		if projectSecurityAndCompliance != true {
			log.Info("违规，未开启安全与合规")
		}

		// 项目的审批

		projectApprovals, _, _ := approvalService.GetProjectApprovals(single_project.ID, nil)
		log.Debug("MergeRequest 审批规则")
		if projectApprovals.MergeRequestsAuthorApproval {
			log.Debug("阻止发起MR的人（作者）审批")
		} else {
			log.Info("不合规，不阻止发起MR的人（作者）审批")
		}
		if projectApprovals.MergeRequestsDisableCommittersApproval {
			log.Debug("阻止有commit的人审批")
		} else {
			log.Info("不合规，不阻止有commit的人审批")
		}
		if projectApprovals.DisableOverridingApproversPerMergeRequest {
			log.Debug("阻止MR中可以编辑审批规则")
		} else {
			log.Info("不合规，不阻止MR中可以编辑审批规则")
		}
		if projectApprovals.RequirePasswordToApprove {
			log.Debug("批准时需要密码 ")
		} else {
			log.Debug("批准时不需要密码 ")
		}
		//发起MR时若有新commit被push，是否撤销批准
		if projectApprovals.ResetApprovalsOnPush == false && projectApprovals.SelectiveCodeOwnerRemovals == false {
			log.Info("不合规，发起MR后若有新的commit被提交，则已审批人数不变")
		} else if projectApprovals.ResetApprovalsOnPush == true && projectApprovals.SelectiveCodeOwnerRemovals == false {
			log.Debug("发起MR后若有新的commit被提交，则已审批人数清零")
		} else if projectApprovals.ResetApprovalsOnPush == false && projectApprovals.SelectiveCodeOwnerRemovals == true {
			log.Debug("发起MR后若有新的commit被提交，如果存在批准者的代码被改变，那么该批准者的审批被撤销")
		}

		// 项目的审批规则
		// 获取要求的最小审批MR的人数
		MinimumMergeRequestApprovalNumber := config.GetInt("projects.project.project_approvals_before_merge_approval_number.keywords")
		// 获取MR的多条规则
		ApprovalsRules, err := approvalService.GetProjectApprovalsRules(single_project.ID, nil)
		if err != nil {
			log.Error(err)
		}
		// 获取项目的member数量
		// 单人项目不进行此项检查，两人项目由于一人发起MR但是无法审批，所以两人项目也不进行检查
		// 如果 >= 3则开始进行MR规则审批的检查项
		members := ProjectMembers(projectMembersService, single_project.ID)
		if len(members) >= 3 {
			if len(ApprovalsRules) == 0 {
				log.Info("不合规，未设置审批规则，maintainer及owner一个人就可以完成MR的审批是不合规的")
			} else {
				approval_num := MegerRequestApprovalsRulesRequireNumber(ApprovalsRules)
				if approval_num >= MinimumMergeRequestApprovalNumber {
					log.Debug("合规，项目审批人数合格")
				} else {
					log.Info("不合规，项目审批人数不合规")
				}
			}
		}

		// 启用或禁用合并管道
		mergePipelinesEnabled := single_project.MergePipelinesEnabled
		if !mergePipelinesEnabled {
			log.Info("不合规，合并管道被禁用")
		}

		// 启用或禁用合并队列
		mergeTrainsEnabled := single_project.MergeTrainsEnabled
		if !mergeTrainsEnabled {
			log.Info("不合规，合并队列被禁用")
		}

		// 检查受到保护的分支中是否有默认分支
		// 理想的合规返回值 true true true false
		default_branch_exist_flag, push_access_level_flag, merge_access_level_flag, allow_to_force_push_flag := BranchProtected(protectedBranchesService, single_project.ID, single_project.DefaultBranch)

		if !default_branch_exist_flag {
			log.Info("不合规，不存在默认分支的分支保护规则")
		} else {
			if !push_access_level_flag {
				log.Info("不合规，默认分支存在dev和maintainer以下权限角色可以push")
			}
			if !merge_access_level_flag {
				log.Info("不合规，默认分支存在maintainer以下权限角色角色可以准许merge")
			}
			if allow_to_force_push_flag {
				log.Info("不合规，默认分支被设置为允许force push")
			}
		}

		// push rule
		// 均仅作提示，理想下为 t t t
		reject_unsign_commit, reject_unverified_email_push, reject_commit_unverified_push := PushRule(projectsService, single_project.ID)
		if reject_unsign_commit {
			log.Info("加分项：拒绝未签名的提交")
		}
		if reject_unverified_email_push {
			log.Info("加分项：拒绝未验证邮箱用户的push")
		}
		if reject_commit_unverified_push {
			log.Info("加分项：拒绝存在未验证用户的commit的push")
		}

	}
}
