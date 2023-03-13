package project

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"
	"regexp"
	"strings"
)

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper, output *types.Output) {

	log.Debug("[#] 项目配置分析检测开始")
	AllProjects, projectsService, approvalService, projectMembersService, protectedBranchesService, projectAccessTokensService, projectValidateService := ListAllProjects(gitlabClient)

	output.Projects.Projects = make([]types.Project, 0, len(AllProjects))
	var o_single_project types.Project

	for _, single_project := range AllProjects {

		o_single_project.Id = single_project.ID
		posNameSpace := strings.Index(single_project.NameWithNamespace, "/")
		o_single_project.NameSpace = single_project.NameWithNamespace[:posNameSpace]
		o_single_project.ProjectName = single_project.Name

		log.Info("目前正在检查项目 id:", single_project.ID, "Name:", single_project.NameWithNamespace)
		projectVisibility := ProjectVisibility(single_project)
		if projectVisibility == "public" {
			o_single_project.VisibilityNoPublic = false
			log.Info("违规，禁止建立public项目")
		} else {
			o_single_project.VisibilityNoPublic = true
		}

		require_authentication_to_view_media_files := RequireAuthenticationToViewMediaFiles(single_project)
		if require_authentication_to_view_media_files != true {
			o_single_project.TagMediaLinkAuth = false
			log.Info("违规，访问tag媒体文件链接时未验证登录，会导致tag内容的媒体文件连接被直接访问")
		} else {
			o_single_project.TagMediaLinkAuth = true
		}

		projectSecurityAndCompliance := ProjectSecurityAndCompliance(single_project)
		if projectSecurityAndCompliance != true {
			o_single_project.SecurityAndCompliance = false
			log.Info("违规，未开启安全与合规")
		} else {
			o_single_project.SecurityAndCompliance = true
		}

		// 项目的审批
		projectApprovals, _, _ := approvalService.GetProjectApprovals(single_project.ID, nil)
		log.Debug("MergeRequest 审批规则")
		if projectApprovals.MergeRequestsAuthorApproval {
			o_single_project.MergeRequestsAuthorApproval = true
			log.Debug("阻止发起MR的人（作者）审批")
		} else {
			o_single_project.MergeRequestsAuthorApproval = false
			log.Info("不合规，不阻止发起MR的人（作者）审批")
		}
		if projectApprovals.MergeRequestsDisableCommittersApproval {
			o_single_project.MergeRequestsDisableCommittersApproval = true
			log.Debug("阻止有commit的人审批")
		} else {
			o_single_project.MergeRequestsDisableCommittersApproval = false
			log.Info("不合规，不阻止有commit的人审批")
		}
		if projectApprovals.DisableOverridingApproversPerMergeRequest {
			o_single_project.DisableOverridingApproversPerMergeRequest = true
			log.Debug("阻止MR中可以编辑审批规则")
		} else {
			o_single_project.DisableOverridingApproversPerMergeRequest = false
			log.Info("不合规，不阻止MR中可以编辑审批规则")
		}
		if projectApprovals.RequirePasswordToApprove {
			o_single_project.RequirePasswordToApprove = true
			log.Debug("批准时需要密码 ")
		} else {
			o_single_project.RequirePasswordToApprove = false
			log.Debug("批准时不需要密码 ")
		}

		//发起MR时若有新commit被push，是否撤销批准
		if projectApprovals.ResetApprovalsOnPush == false && projectApprovals.SelectiveCodeOwnerRemovals == false {
			o_single_project.ResetApprovalsOnPush = false
			log.Info("不合规，发起MR后若有新的commit被提交，则已审批人数不变")
		} else if projectApprovals.ResetApprovalsOnPush == true && projectApprovals.SelectiveCodeOwnerRemovals == false {
			o_single_project.ResetApprovalsOnPush = true
			log.Debug("发起MR后若有新的commit被提交，则已审批人数清零")
		} else if projectApprovals.ResetApprovalsOnPush == false && projectApprovals.SelectiveCodeOwnerRemovals == true {
			o_single_project.ResetApprovalsOnPush = true
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
				o_single_project.MegerRequestApprovalsRulesRequireNumber = false
				log.Info("不合规，未设置审批规则，maintainer及owner一个人就可以完成MR的审批是不合规的")
			} else {
				approval_num := MegerRequestApprovalsRulesRequireNumber(ApprovalsRules)
				if approval_num >= MinimumMergeRequestApprovalNumber {
					o_single_project.MegerRequestApprovalsRulesRequireNumber = true
					log.Debug("合规，项目审批人数合格")
				} else {
					o_single_project.MegerRequestApprovalsRulesRequireNumber = false
					log.Info("不合规，项目审批人数不合规")
				}
			}
		}

		// 启用或禁用合并管道
		mergePipelinesEnabled := single_project.MergePipelinesEnabled
		if !mergePipelinesEnabled {
			o_single_project.MergePipelinesEnabled = false
			log.Info("不合规，合并管道被禁用")
		} else {
			o_single_project.MergePipelinesEnabled = true
		}

		// 启用或禁用合并队列
		mergeTrainsEnabled := single_project.MergeTrainsEnabled
		if !mergeTrainsEnabled {
			o_single_project.MergeTrainsEnabled = false
			log.Info("不合规，合并队列被禁用")
		} else {
			o_single_project.MergeTrainsEnabled = true
		}

		// 检查受到保护的分支中是否有默认分支
		// 理想的合规返回值 true true true false
		default_branch_exist_flag, push_access_level_flag, merge_access_level_flag, allow_to_force_push_flag := BranchProtected(protectedBranchesService, single_project.ID, single_project.DefaultBranch)
		DefaultBranchProtectedComplaince_flag := true
		o_single_project.DefaultBranchProtectedDescription = ""
		var o_DefaultBranchProtectedDescription []string
		if !default_branch_exist_flag {
			DefaultBranchProtectedComplaince_flag = false
			log.Info("不合规，不存在默认分支的分支保护规则")
			o_DefaultBranchProtectedDescription = append(o_DefaultBranchProtectedDescription, "不存在默认分支的分支保护规则")
		} else {
			if !push_access_level_flag {
				DefaultBranchProtectedComplaince_flag = false
				log.Info("不合规，默认分支存在dev和maintainer以下权限角色可以push")
				o_DefaultBranchProtectedDescription = append(o_DefaultBranchProtectedDescription, "默认分支存在dev和maintainer以下权限角色可以push")
			}
			if !merge_access_level_flag {
				DefaultBranchProtectedComplaince_flag = false
				log.Info("不合规，默认分支存在maintainer以下权限角色角色可以准许merge")
				o_DefaultBranchProtectedDescription = append(o_DefaultBranchProtectedDescription, "默认分支存在maintainer以下权限角色角色可以准许merge")
			}
			if allow_to_force_push_flag {
				DefaultBranchProtectedComplaince_flag = false
				log.Info("不合规，默认分支被设置为允许force push")
				o_DefaultBranchProtectedDescription = append(o_DefaultBranchProtectedDescription, "默认分支被设置为允许force push")
			}
		}
		if DefaultBranchProtectedComplaince_flag {
			o_single_project.DefaultBranchProtected = true
		} else {
			o_single_project.DefaultBranchProtected = false
		}
		o_single_project.DefaultBranchProtectedDescription = strings.Join(o_DefaultBranchProtectedDescription, "\n")

		// 项目令牌的有效期是否过长
		o_single_project.AccessTokenExpire = ""
		projectAccessTokens, _, err := projectAccessTokensService.ListProjectAccessTokens(single_project.ID, nil, nil)
		if err != nil {
			fmt.Println(err)
		}
		var AccessTokenDes []string
		var current_access_token_des string
		for _, v := range projectAccessTokens {
			if v.ExpiresAt != nil {
				current_access_token_des = v.Name + ": " + v.ExpiresAt.String()
			} else {
				current_access_token_des = v.Name + ": 永久"
			}
			AccessTokenDes = append(AccessTokenDes, current_access_token_des)
		}
		o_single_project.AccessTokenExpire = strings.Join(AccessTokenDes, "\n")

		// push rule
		// 均仅作提示，理想下为 t t t
		reject_unsign_commit, reject_unverified_email_push, reject_commit_unverified_push := PushRule(projectsService, single_project.ID)
		if reject_unsign_commit {
			o_single_project.RejectUnsignCommit = true
			log.Info("加分项：拒绝未签名的提交")
		} else {
			o_single_project.RejectUnsignCommit = false
		}
		if reject_unverified_email_push {
			o_single_project.RejectUnverifiedEmailPush = true
			log.Info("加分项：拒绝未验证邮箱用户的push")
		} else {
			o_single_project.RejectUnverifiedEmailPush = false
		}
		if reject_commit_unverified_push {
			o_single_project.RejectCommitUnverifiedPush = true
			log.Info("加分项：拒绝存在未验证用户的commit的push")
		} else {
			o_single_project.RejectCommitUnverifiedPush = false
		}

		// 获取CICD
		cicdContent, _, err := projectValidateService.ProjectLint(single_project.ID, nil, nil)
		if err != nil {
			fmt.Println(err)
		}
		if cicdContent.Valid {
			tmp := ([]byte(cicdContent.MergedYaml))
			stageRegexp := regexp.MustCompile(`stage:`)
			tmp_res := stageRegexp.FindAllIndex(tmp, -1)

			o_single_project.CICD = true
			o_single_project.CICDYamlStageNum = len(tmp_res)

		} else {
			o_single_project.CICD = false
			o_single_project.CICDYamlStageNum = 0
		}

		// 这部分放最后的
		// 并入输出列表
		output.Projects.Projects = append(output.Projects.Projects, o_single_project)
	}

}
