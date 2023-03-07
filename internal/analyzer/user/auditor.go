package user

import "gitlab-misconfig/internal/gitlab"

// 统计 Auditor 数量
func countAuditorNumbers(users []*gitlab.User) []string {
	r := make([]string, 0, len(users))
	for _, v := range users {
		if v.IsAuditor == true {
			r = append(r, v.Username)
		}
		//if v.Name == "zhangxiaoming" {
		//	r = append(r, v.Name)
		//}
	}
	return r
}
