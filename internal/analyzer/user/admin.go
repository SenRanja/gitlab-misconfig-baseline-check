package user

import (
	"gitlab-misconfig/internal/gitlab"
)

// 统计 Admin 数量
func countAdminNumbers(users []*gitlab.User) []string {
	r := make([]string, 0, len(users))
	for i := 0; i < len(users); i++ {
		if users[i].IsAdmin {
			r = append(r, users[i].Username)
		}
	}
	return r
}
