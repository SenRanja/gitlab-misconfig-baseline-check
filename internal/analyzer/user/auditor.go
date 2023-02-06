package user

import "gitlab-misconfig/internal/gitlab"

// 统计 Auditor 数量
func countAuditorNumbers(users []*gitlab.User) int {
	var totalNumberOfAuditor = 0
	for i := 0; i < len(users); i++ {
		if users[i].IsAuditor {
			totalNumberOfAuditor += 1
		}
	}
	return totalNumberOfAuditor
}
