package user

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
)

// 获取gitlab 所有用户信息
func getGitlabUsers(gitlabClient *gitlab.Client) []*gitlab.User {
	users, _, err := gitlabClient.Users.ListUsers(&gitlab.ListUsersOptions{})
	if err != nil {
		fmt.Println(err)
	}
	return users
}

// 统计不活跃用户数量
func countUnactiveNumbers(users []*gitlab.User) int {
	var totalNumberOfAuditor = 0
	for i := 0; i < len(users); i++ {
		if users[i].IsAuditor {
			totalNumberOfAuditor += 1
		}
	}
	return totalNumberOfAuditor
}

// 统计开启双因素认证用户数量
func countTwoFactorEnabled(users []*gitlab.User) int {
	var totalTwoFactorEnabled = 0
	for i := 0; i < len(users); i++ {
		if users[i].TwoFactorEnabled == true {
			totalTwoFactorEnabled += 1
		}
	}
	return totalTwoFactorEnabled
}
