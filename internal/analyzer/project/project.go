package project

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
	"strconv"
	"strings"
)

// project_visibility
// 项目可见性，只有三种： "private"、"internal"、"public"
func ProjectVisibility(p *gitlab.Project) string {
	return string(p.Visibility)
}

// 访问tag媒体文件链接时是否需要验证登录
func RequireAuthenticationToViewMediaFiles(p *gitlab.Project) bool {
	return p.EnforceAuthChecksOnUploads
}

// 安全与合规 true false
func ProjectSecurityAndCompliance(p *gitlab.Project) bool {
	return p.SecurityAndComplianceEnabled
}

// Merge Request 审批
// 默认情况下应该有多少批准人批准合并请求。要配置批准规则

// 获取项目成员
// 由于project如果设置访问令牌则会增加project{project_id}_bot@noreply.{Gitlab.config.gitlab.host}机器人，此处需要过滤掉机器人账户
// 过滤方法为匹配 email 和 username
func ProjectMembers(p *gitlab.ProjectMembersService, pid int) []*gitlab.ProjectMember {
	members, _, err := p.ListProjectMembers(pid, nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	bot_keyword := "project_" + strconv.Itoa(pid) + "_bot"

	m := make([]*gitlab.ProjectMember, 0, len(members))
	for _, v := range members {
		if !strings.Contains(v.Email, "@noreply") && !strings.Contains(v.Username, bot_keyword) {
			m = append(m, v)
		}
	}
	return members
}

// Merge Request 看默认规则中审批人数
// n为计数，检查目标分支是否包含 默认分支
func MegerRequestApprovalsRulesRequireNumber(ApprovalsRules []gitlab.MergeRequestApprovalRule) int {
	n := 0
	for _, v := range ApprovalsRules {
		if v.Name == "All Members" {
			n += v.ApprovalsRequired
		}
	}
	return n
}

// 遍历受到保护的分支中，是否包含默认分支
func BranchProtected(ps *gitlab.ProtectedBranchesService, pid int, default_branch string) (bool, bool, bool, bool) {
	// 返回项:
	// 1 是否包含默认分支
	// 2 默认分支merger权限仅maintainer，全部40及以上返回true
	// 3 默认分支push仅dev+maintainer，全部30及以上返回true
	// 4 默认分支allow_to_force_push
	// 理想的合规返回值 true true true false
	protected_branches, _, err := ps.ListProtectedBranches(pid, nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	default_branch_exist_flag := false
	push_access_level_flag := true
	merge_access_level_flag := true
	var allow_to_force_push_flag bool
	for _, v := range protected_branches {
		if v.Name == default_branch {
			default_branch_exist_flag = true
			for _, vv := range v.PushAccessLevels {
				if vv.AccessLevel < 30 {
					push_access_level_flag = false
				}
			}
			for _, vv := range v.MergeAccessLevels {
				if vv.AccessLevel < 40 {
					merge_access_level_flag = false
				}
			}
			allow_to_force_push_flag = v.AllowForcePush

			//v.AllowForcePush
		}
	}
	if default_branch_exist_flag == false {
		return false, false, false, false
	}
	return default_branch_exist_flag, push_access_level_flag, merge_access_level_flag, allow_to_force_push_flag
}

// 未通过 GPG 签名时拒绝提交
func PushRule(gitlabClientProjectService *gitlab.ProjectsService, pid int) (bool, bool, bool) {
	// ppr 是 ProjectPushRule 的简写
	// 返回值：1. 拒绝未GPG签名的push
	// 2. 拒绝未验证邮箱的用户的push
	// 3. 拒绝commit中存在未验证用户的push
	ppr, _, err := gitlabClientProjectService.GetProjectPushRules(pid)
	if err != nil {
		fmt.Println("ProjectRejectUnsignedCommits err")
	}

	return ppr.RejectUnsignedCommits, ppr.CommitCommitterCheck, ppr.MemberCheck
}
