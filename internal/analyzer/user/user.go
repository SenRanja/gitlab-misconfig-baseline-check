package user

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
	"time"
)

// 获取gitlab 所有用户信息
func getGitlabUsers(gitlabClient *gitlab.Client) ([]*gitlab.User, gitlab.UsersService) {
	users, _, err := gitlabClient.Users.ListUsers(&gitlab.ListUsersOptions{})
	if err != nil {
		fmt.Println(err)
	}

	usersService := gitlab.UsersService{
		Client: gitlabClient,
	}
	return users, usersService
}

// 不活跃用户数量
func getUnactiveUsers(users []*gitlab.User, unactive_time int) []*gitlab.User {

	// 设置多长时间不登陆，则视为 不活跃
	noLoginTime := time.Now().AddDate(0, 0, -unactive_time)

	var unActiveUsers []*gitlab.User
	for _, user := range users {
		if user.LastSignInAt == nil || user.LastSignInAt.Before(noLoginTime) {
			unActiveUsers = append(unActiveUsers, user)
		}
	}

	return unActiveUsers
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

// 统计外部用户数量
func countExternalDetect(users []*gitlab.User) int {
	var n = 0
	for i := 0; i < len(users); i++ {
		if users[i].External == true {
			n += 1
		}
	}
	return n
}

type projectIdNameAccesslevel struct {
	ID          int
	Accesslevel int
}

// 获取membership
func GetUserMembership(m []*gitlab.UserMembership, t string) []projectIdNameAccesslevel {

	var source_type string
	switch t {
	case "project":
		source_type = "Project"
	case "group":
		source_type = "Namespace"
	}

	var r []projectIdNameAccesslevel
	for _, memberShip := range m {
		if memberShip.SourceType == source_type {
			r = append(
				r,
				projectIdNameAccesslevel{
					ID:          memberShip.SourceID,
					Accesslevel: int(memberShip.AccessLevel),
				},
			)
		}
	}

	return r

}
