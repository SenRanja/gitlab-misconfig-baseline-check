package user

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
	"regexp"
	"strings"
	"time"
)

// 获取gitlab 所有用户信息
func getGitlabUsers(gitlabClient *gitlab.Client) ([]*gitlab.User, gitlab.UsersService) {

	// 需要剔除project 的 部署令牌 产生的bot用户
	PAT_bot_re, err := regexp.Compile(`project\d+_bot\d*@noreply`)
	if err != nil {
		fmt.Println(err)
	}

	users_tmp, _, err := gitlabClient.Users.ListUsers(&gitlab.ListUsersOptions{})
	if err != nil {
		fmt.Println(err)
	}
	users := make([]*gitlab.User, 0, len(users_tmp))

	for _, v := range users_tmp {
		if !PAT_bot_re.MatchString(v.Email) {
			users = append(users, v)
		}
	}

	usersService := gitlab.UsersService{
		Client: gitlabClient,
	}
	return users, usersService
}

func ListPrint(in []string) string {
	return "[ " + strings.Join(in, ", ") + " ]"
}

func FilterNonPeopleUser(in []string) []string {
	length := len(in)
	r := make([]string, 0, length)
	//r := []string{}
	for _, v := range in {
		if v != "ghost" && v != "ghost" && v != "support-bot" && v != "alert-bot" && v != "GitLab Support Bot" && v != "GitLab Alert Bot" {
			r = append(r, v)
		}
	}
	return r
}

// 未启用用户数量
// 指激活了，ActivityOn == nil( LastSignInAt在admin注册后用户登录却没有接受条款激活的时候，可能是有值的 )
func getUnstatUsers(users []*gitlab.User) []string {

	var unActiveUsers []string
	for _, user := range users {
		if user.LastActivityOn == nil {
			unActiveUsers = append(unActiveUsers, user.Username)
		}
	}

	return FilterNonPeopleUser(unActiveUsers)
}

// 不活跃用户数量
// 指激活了，ActivityOn != nil && LastSignInAt.Before(noLoginTime)
func getUnactiveUsers(users []*gitlab.User, unactive_time int) []string {
	// 设置多长时间不登陆，则视为 不活跃
	noLoginTime := time.Now().AddDate(0, 0, -unactive_time)

	var unActiveUsers []string
	for _, user := range users {
		if user.LastActivityOn != nil && user.LastSignInAt.Before(noLoginTime) {
			unActiveUsers = append(unActiveUsers, user.Username)
		}
	}
	return FilterNonPeopleUser(unActiveUsers)
}

// 统计开启双因素认证用户数量
func countTwoFactorEnabled(users []*gitlab.User) []string {
	var r []string
	for i := 0; i < len(users); i++ {
		if users[i].TwoFactorEnabled == true {
			r = append(r, users[i].Username)
		}
	}
	return FilterNonPeopleUser(r)
}

// 统计外部用户数量
func countExternalDetect(users []*gitlab.User) []string {
	r := make([]string, 0, len(users))
	for i := 0; i < len(users); i++ {
		if users[i].External == true {
			r = append(r, users[i].Username)
		}
	}
	return FilterNonPeopleUser(r)
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
