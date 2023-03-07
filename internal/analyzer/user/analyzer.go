package user

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"
)

// 本`user`目录主要是围绕`用户`项的代码
// admin 是`管理员`权限相关，auditor是`审计员`权限相关，user是`用户`权限相关
// 而 `analyzer.go` 是和自动扫描相关，
// 	如engine/engine.go中 analyzer.AutoAnalysis(gitlabClient, options, config)
// 	会出发`analyzer.go`中的主要代码

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {

	// 【user】
	// getGitlabUsers()的返回值的第二个值是返回 usersService，目前没有用到
	users, _ := getGitlabUsers(gitlabClient)

	// admin数量
	adminMaxNumbers := config.GetInt("users.admin.admin_max_numbers.keywords")
	totalNumberOfAdmin := countAdminNumbers(users)
	if len(totalNumberOfAdmin) > adminMaxNumbers {
		log.Info("[-] 管理员人数过多，共有", len(totalNumberOfAdmin), "个管理员用户")
	} else {
		log.Info("[+] 管理员人数合理")
	}
	log.Info("[#] 管理员用户列表", ListPrint(totalNumberOfAdmin))

	// auditor数量
	auditorLeastNumbers := config.GetInt("users.auditor.auditor_least_numbers.keywords")
	totalNumberOfAuditor := countAuditorNumbers(users)
	if len(totalNumberOfAuditor) < auditorLeastNumbers {
		log.Info("[-] 审计员人数过少，共有", len(totalNumberOfAuditor), "个审计员")
	} else {
		log.Info("[+] 审计员人数合理")
	}
	log.Info("[#] 审计员用户列表", ListPrint(totalNumberOfAuditor))

	// 用户双因素认证数量
	countTwoFactorEnabledNum := countTwoFactorEnabled(users)
	log.Info("[ ] 开启双因素认证用户数量", len(countTwoFactorEnabledNum))
	log.Info("[#] 开启双因素认证用户列表", ListPrint(countTwoFactorEnabledNum))

	// 统计未启用用户
	StatUnactiveUsers := getUnstatUsers(users)
	log.Info("[-] 未启用用户数量", len(StatUnactiveUsers))
	log.Info("[#] 未启用用户名称列表", ListPrint(StatUnactiveUsers))

	// 统计不活跃用户
	unActiveUserNoLoginTime := config.GetInt("users.users.unactive_time.keywords")
	UnactiveUsers := getUnactiveUsers(users, unActiveUserNoLoginTime)
	log.Info("[-] 不活跃用户数量", len(UnactiveUsers))
	log.Info("[#] 不活跃用户名称列表", ListPrint(UnactiveUsers))

	// 统计外部用户
	countExternal := countExternalDetect(users)
	log.Info("[ ] 统计外部用户数量", len(countExternal))
	log.Info("[#] 外部用户列表", ListPrint(countExternal))

}
