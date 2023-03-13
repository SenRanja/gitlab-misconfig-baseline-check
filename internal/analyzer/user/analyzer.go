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

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper, output *types.Output) {

	// 【user】
	// getGitlabUsers()的返回值的第二个值是返回 usersService，目前没有用到
	users, _ := getGitlabUsers(gitlabClient)

	// admin数量
	adminMaxNumbers := config.GetInt("users.admin.admin_max_numbers.keywords")
	totalNumberOfAdmin := countAdminNumbers(users)
	output.User.Admin.Keyword = adminMaxNumbers
	output.User.Admin.Result = len(totalNumberOfAdmin)
	if len(totalNumberOfAdmin) > adminMaxNumbers {
		output.User.Admin.Complaince = false
		log.Info("[-] 管理员人数过多，共有", len(totalNumberOfAdmin), "个管理员用户")
	} else {
		output.User.Admin.Complaince = true
		log.Info("[+] 管理员人数合理")
	}
	log.Info("[#] 管理员用户列表", ListPrint(totalNumberOfAdmin))

	// auditor数量
	auditorLeastNumbers := config.GetInt("users.auditor.auditor_least_numbers.keywords")
	output.User.Auditor.Keyword = auditorLeastNumbers
	totalNumberOfAuditor := countAuditorNumbers(users)
	output.User.Auditor.Result = len(totalNumberOfAuditor)
	if len(totalNumberOfAuditor) < auditorLeastNumbers {
		output.User.Auditor.Complaince = false
		log.Info("[-] 审计员人数过少，共有", len(totalNumberOfAuditor), "个审计员")
	} else {
		output.User.Auditor.Complaince = true
		log.Info("[+] 审计员人数合理")
	}
	output.User.Auditor.Description += ListPrint(totalNumberOfAuditor)
	log.Info("[#] 审计员用户列表", ListPrint(totalNumberOfAuditor))

	// 用户双因素认证数量
	countTwoFactorEnabledNum := countTwoFactorEnabled(users)
	output.User.TwoFactorAuth.Result = len(countTwoFactorEnabledNum)
	output.User.TwoFactorAuth.Keyword = 0
	output.User.TwoFactorAuth.Complaince = true
	output.User.TwoFactorAuth.Description += ListPrint(countTwoFactorEnabledNum)
	log.Info("[ ] 开启双因素认证用户数量", len(countTwoFactorEnabledNum))
	log.Info("[#] 开启双因素认证用户列表", ListPrint(countTwoFactorEnabledNum))

	// 统计未启用用户
	StatUnactiveUsers := getUnstatUsers(users)
	output.User.Unactive.Result = len(StatUnactiveUsers)
	output.User.Unactive.Keyword = 0
	output.User.Unactive.Complaince = true
	output.User.Unactive.Description += ListPrint(StatUnactiveUsers)
	log.Info("[-] 未启用用户数量", len(StatUnactiveUsers))
	log.Info("[#] 未启用用户名称列表", ListPrint(StatUnactiveUsers))

	// 统计不活跃用户
	unActiveUserNoLoginTime := config.GetInt("users.users.unactive_time.keywords")
	UnactiveUsers := getUnactiveUsers(users, unActiveUserNoLoginTime)
	output.User.Inactive.Result = len(UnactiveUsers)
	output.User.Inactive.Keyword = 0
	output.User.Inactive.Complaince = true
	output.User.Inactive.Description += ListPrint(UnactiveUsers)
	log.Info("[-] 不活跃用户数量", len(UnactiveUsers))
	log.Info("[#] 不活跃用户名称列表", ListPrint(UnactiveUsers))

	// 统计外部用户
	countExternal := countExternalDetect(users)
	output.User.External.Result = len(countExternal)
	output.User.External.Keyword = 0
	output.User.External.Complaince = true
	output.User.External.Description += ListPrint(countExternal)
	log.Info("[ ] 统计外部用户数量", len(countExternal))
	log.Info("[#] 外部用户列表", ListPrint(countExternal))

}
