package project

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
)

// 列出全部project
func ListAllProjects(gitlabClient *gitlab.Client) ([]*gitlab.Project, *gitlab.ProjectsService) {
	projectsService := gitlab.ProjectsService{
		Client: gitlabClient,
	}
	projects, _, err := gitlabClient.Projects.ListProjects(listProjectsOptions)
	if err != nil {
		fmt.Println(err)
	}
	return projects, &projectsService
}

// project_visibility
// 项目可见性，只有三种： "private"、"internal"、"public"
func ProjectVisibility(p *gitlab.Project) string {
	return string(p.Visibility)
}

// 安全与合规
// 总共四种状态: "disabled" "enabled" "private" "public"
// ? 我推测，对于项目而言，应该是 `非 "disable"`， 则json中直接返回bool的true
func ProjectSecurityAndCompliance(p *gitlab.Project) bool {
	if p.SecurityAndComplianceAccessLevel != "disabled" {
		return true
	}
	return false
}

// project_pages_access_level
// ?页面访问级别  什么是页面访问级别，我还不太清楚

// 权限

// 获取项目近期审计事件

// 列举项目用户

// 合并前的批准
// 默认情况下应该有多少批准人批准合并请求。要配置批准规则
func ProjectApprovalsBeforeMerge(p *gitlab.Project) int {
	return p.ApprovalsBeforeMerge
}

// 未通过 GPG 签名时拒绝提交
func ProjectRejectUnsignedCommits(gitlabClientProjectService *gitlab.ProjectsService, pid int) bool {
	// ppr 是 ProjectPushRule 的简写
	ppr, _, err := gitlabClientProjectService.GetProjectPushRules(pid)
	if err != nil {
		fmt.Println("ProjectRejectUnsignedCommits err")
	}
	return ppr.RejectUnsignedCommits
}

// 只允许验证过的committer进行commit
func ProjectCommitCommitterCheck(gitlabClientProjectService *gitlab.ProjectsService, pid int) bool {
	ppr, _, err := gitlabClientProjectService.GetProjectPushRules(pid)
	if err != nil {
		fmt.Println("ProjectRejectUnsignedCommits err")
	}
	return ppr.CommitCommitterCheck
}
