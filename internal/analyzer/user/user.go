package user

import "gitlab-misconfig/internal/gitlab"

func countUnactiveNumbers(users []*gitlab.User) int {
	var totalNumberOfAuditor = 0
	for i := 0; i < len(users); i++ {
		if users[i].IsAuditor {
			totalNumberOfAuditor += 1
		}
	}
	return totalNumberOfAuditor
}
