package user

import (
	"gitlab-misconfig/internal/gitlab"
)

func countAdminNumbers(users []*gitlab.User) int {
	var totalNumberOfAdmin = 0
	for i := 0; i < len(users); i++ {
		if users[i].IsAdmin {
			totalNumberOfAdmin += 1
		}
	}
	return totalNumberOfAdmin
}
