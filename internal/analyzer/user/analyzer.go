package user

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/analyzer/audit_event"
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

type Analyzer struct {
}

// New 创建User分析模块
func New() Analyzer {
	return Analyzer{}
}

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {
	users := getGitlabUsers(gitlabClient)
	// admin
	log.Info("[#] admin权限检查")
	adminMaxNumbers := config.GetString("users.admin.admin_max_numbers.keywords")
	totalNumberOfAdmin := countAdminNumbers(users)
	log.Debug("total number of admin is ", totalNumberOfAdmin)
	reslut, err := rules.CheckRule(strconv.Itoa(totalNumberOfAdmin), ">", adminMaxNumbers)
	if reslut && err == nil {
		log.Info("gitlab管理员数量多于", totalNumberOfAdmin, "人")
	}

	// auditor
	log.Info("[#] auditor权限检查")
	auditorMaxNumbers := config.GetString("users.auditor.auditor_max_numbers.keywords")
	totalNumberOfAuditor := countAuditorNumbers(users)
	log.Debug("total number of auditor is ", totalNumberOfAuditor)
	log.Info("gitlab 审计人员数量为：  ", totalNumberOfAuditor)

	reslut, err = rules.CheckRule(strconv.Itoa(totalNumberOfAuditor), "<", auditorMaxNumbers)
	if reslut && err == nil {
		log.Info("gitlab审计人员至少设置", auditorMaxNumbers, "人")
	}

	//user
	// 【test】
	log.Info("[#] user权限检查")
	countTwoFactorEnabledNum := countTwoFactorEnabled(users)
	log.Debug("开启双因素认证用户数量", countTwoFactorEnabledNum)

	// 是否设置禁止注册功能
	log.Info("[#] 用户注册功能状态")
	RegisterIfEnable, err := JudgeIfDisableRegister(gitlabClient, options, config)
	if err != nil {
		log.Error(err)
	}
	log.Debug("用户注册功能状态", RegisterIfEnable)

	// 用户密码复杂度是否配置

	// 审计
	log.Info("[#] 审计功能")
	// 获取近期审计日志
	aes, err := audit_event.GetAllEvents(gitlabClient)
	// 【删除用户】
	log.Info("[###] 被删除用户检测开始")
	for _, userDeleted := range audit_event.UserDeleted(aes) {
		log.Info("删除用户", *userDeleted)
	}
	log.Info("[###] 被删除用户检测完毕")
	// 【新建用户】
	log.Info("[###] 新建用户检测开始")
	for _, userCreated := range audit_event.UserCreated(aes) {
		log.Info("新建用户", *userCreated)
	}
	log.Info("[###] 新建用户检测完毕")
	// 【登陆失败】
	log.Info("[###] 登陆失败检测开始")
	for _, userFailedLogin := range audit_event.UserLoginFailed(aes) {
		log.Info("登陆失败", *userFailedLogin)
	}
	log.Info("[###] 登陆失败检测完毕")
	//// 【登陆正常】
	//log.Info("[###] 登陆正常检测开始")
	//for _, userLoginCorrectly := range audit_event.UserLoginCorrectly(aes) {
	//	log.Info("登陆正常", *userLoginCorrectly)
	//}
	//log.Info("[###] 登陆正常检测完毕")
	// 【登陆失败】
	log.Info("[###] 登陆失败检测开始")
	for _, userFailedLogin := range audit_event.UserLoginFailed(aes) {
		log.Info("登陆失败", *userFailedLogin)
	}
	log.Info("[###] 登陆失败检测完毕")
	// 【新建仓库】
	log.Info("[###] 新建仓库检测开始")
	for _, projectCreated := range audit_event.ProjectCreated(aes) {
		log.Info("新建仓库", *projectCreated)
	}
	log.Info("[###] 新建仓库检测完毕")
	// 【仓库添加用户】
	log.Info("[###] 仓库添加用户检测开始")
	for _, userAdd2Project := range audit_event.AddUser2Project(aes) {
		log.Info("仓库添加用户 ", *userAdd2Project)
	}
	log.Info("[###] 仓库添加用户检测完毕")
	// 【组添加用户】
	log.Info("[###] 组添加用户检测开始")
	for _, userAdd2Group := range audit_event.AddUser2Group(aes) {
		log.Info("组添加用户 ", *userAdd2Group)
	}
	log.Info("[###] 组添加用户检测完毕")

	log.Debug("[#] 审计功能检测完毕")
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
