package user

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/rules"
	"gitlab-misconfig/internal/types"
	"strconv"
)

// 本`user`目录主要是围绕`用户`项的代码
// admin 是`管理员`权限相关，auditor是`审计员`权限相关，user是`用户`权限相关
// 而 `analyzer.go` 是和自动扫描相关，
// 	如engine/engine.go中 analyzer.AutoAnalysis(gitlabClient, options, config)
// 	会出发`analyzer.go`中的主要代码

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {

	// 【user】
	users, usersService := getGitlabUsers(gitlabClient)
	// admin数量
	log.Info("[#] admin权限检查")
	adminMaxNumbers := config.GetString("users.admin.admin_max_numbers.keywords")
	totalNumberOfAdmin := countAdminNumbers(users)
	log.Debug("total number of admin is ", totalNumberOfAdmin)
	reslut, err := rules.CheckRule(strconv.Itoa(totalNumberOfAdmin), ">", adminMaxNumbers)
	if reslut && err == nil {
		log.Info("gitlab管理员数量多于", adminMaxNumbers, "人")
	}

	// auditor数量
	log.Info("[#] auditor权限检查")
	auditorMaxNumbers := config.GetString("users.auditor.auditor_max_numbers.keywords")
	totalNumberOfAuditor := countAuditorNumbers(users)
	log.Debug("total number of auditor is ", totalNumberOfAuditor)
	log.Info("gitlab 审计人员数量为：  ", totalNumberOfAuditor)

	reslut, err = rules.CheckRule(strconv.Itoa(totalNumberOfAuditor), "<", auditorMaxNumbers)
	if reslut && err == nil {
		log.Info("gitlab审计人员至少设置", auditorMaxNumbers, "人")
	}

	log.Info("[#] user权限检查")

	// 用户双因素认证数量
	countTwoFactorEnabledNum := countTwoFactorEnabled(users)
	log.Debug("开启双因素认证用户数量", countTwoFactorEnabledNum)

	// 统计不活跃用户数量
	unActiveUserNoLoginTime := config.GetInt("users.users.unactive_time.keywords")
	countUnactiveUsers := getUnactiveUsers(users, unActiveUserNoLoginTime)
	log.Debug("统计不活跃用户数量", len(countUnactiveUsers))

	// 统计外部用户数量
	countExternal := countExternalDetect(users)
	log.Debug("统计外部用户数量", countExternal)

	// 遍历单个用户，输出group和project信息
	for _, v := range users {
		tmpRes, _, _ := usersService.GetUserMemberships(v.ID, nil, nil)
		groupMembership := GetUserMembership(tmpRes, "group")
		projectMembership := GetUserMembership(tmpRes, "project")
		log.Info("组关系", groupMembership)
		log.Info("项目关系", projectMembership)
	}

}
