package user

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/rules"
	"gitlab-misconfig/internal/types"
	"strconv"
)

type Analyzer struct {
}

// New 创建User分析模块
func New() Analyzer {
	return Analyzer{}
}

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {
	users := getGitlabUsers(gitlabClient)
	// admin
	adminMaxNumbers := config.GetString("users.admin.admin_max_numbers.keywords")
	totalNumberOfAdmin := countAdminNumbers(users)
	log.Debug("total number of admin is ", totalNumberOfAdmin)
	reslut, err := rules.CheckRule(strconv.Itoa(totalNumberOfAdmin), ">", adminMaxNumbers)
	if reslut && err == nil {
		log.Info("gitlab管理员数量多于", totalNumberOfAdmin, "人")
	}

	// auditor
	auditorMaxNumbers := config.GetString("users.auditor.auditor_max_numbers.keywords")
	totalNumberOfAuditor := countAuditorNumbers(users)
	log.Debug("total number of auditor is ", totalNumberOfAuditor)
	log.Info("gitlab 审计人员数量为：  ", totalNumberOfAuditor)

	reslut, err = rules.CheckRule(strconv.Itoa(totalNumberOfAuditor), "<", auditorMaxNumbers)
	if reslut && err == nil {
		log.Info("gitlab审计人员至少设置", auditorMaxNumbers, "人")
	}

	// 是否设置禁止注册功能

	// 用户密码复杂度是否配置

}

// 获取gitlab 所有用户信息
func getGitlabUsers(gitlabClient *gitlab.Client) []*gitlab.User {
	users, _, err := gitlabClient.Users.ListUsers(&gitlab.ListUsersOptions{})
	if err != nil {
		log.Error("get gitlab users error")
		log.Error(err)
	}
	return users
}

func checkUsers(users []*gitlab.User) {

}
