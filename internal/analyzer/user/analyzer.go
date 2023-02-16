package user

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/analyzer/audit_event"
	"gitlab-misconfig/internal/analyzer/project"
	"gitlab-misconfig/internal/analyzer/settings"
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

	//user
	users := getGitlabUsers(gitlabClient)
	// 【admin数量】
	log.Info("[#] admin权限检查")
	adminMaxNumbers := config.GetString("users.admin.admin_max_numbers.keywords")
	totalNumberOfAdmin := countAdminNumbers(users)
	log.Debug("total number of admin is ", totalNumberOfAdmin)
	reslut, err := rules.CheckRule(strconv.Itoa(totalNumberOfAdmin), ">", adminMaxNumbers)
	if reslut && err == nil {
		log.Info("gitlab管理员数量多于", totalNumberOfAdmin, "人")
	}

	// 【auditor数量】
	log.Info("[#] auditor权限检查")
	auditorMaxNumbers := config.GetString("users.auditor.auditor_max_numbers.keywords")
	totalNumberOfAuditor := countAuditorNumbers(users)
	log.Debug("total number of auditor is ", totalNumberOfAuditor)
	log.Info("gitlab 审计人员数量为：  ", totalNumberOfAuditor)

	reslut, err = rules.CheckRule(strconv.Itoa(totalNumberOfAuditor), "<", auditorMaxNumbers)
	if reslut && err == nil {
		log.Info("gitlab审计人员至少设置", auditorMaxNumbers, "人")
	}

	// user
	// 【用户双因素认证数量】
	log.Info("[#] user权限检查")
	countTwoFactorEnabledNum := countTwoFactorEnabled(users)
	log.Debug("开启双因素认证用户数量", countTwoFactorEnabledNum)

	// 日志审计
	log.Info("[#] 审计功能")
	// 获取近期审计日志
	aes, err := audit_event.GetAllEvents(gitlabClient)
	// 【新建用户】
	log.Info("[###] 新建用户检测开始")
	for _, userCreated := range audit_event.UserCreated(aes) {
		log.Info("新建用户", *userCreated)
	}
	log.Info("[###] 新建用户检测完毕")
	// 【删除用户】
	log.Info("[###] 被删除用户检测开始")
	for _, userDeleted := range audit_event.UserDeleted(aes) {
		log.Info("删除用户", *userDeleted)
	}
	log.Info("[###] 被删除用户检测完毕")
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
	// 【删除仓库】
	log.Info("[###] 删除仓库检测开始")
	for _, projectDeleted := range audit_event.ProjectDeleted(aes) {
		log.Info("删除仓库", *projectDeleted)
	}
	log.Info("[###] 新建仓库检测完毕")
	// 【仓库添加用户】
	log.Info("[###] 仓库添加用户检测开始")
	for _, userAdd2Project := range audit_event.AddUser2Project(aes) {
		log.Info("仓库添加用户 ", *userAdd2Project)
	}
	log.Info("[###] 仓库添加用户检测完毕")
	// 【仓库删除用户】
	log.Info("[###] 仓库删除用户检测开始")
	for _, userDelFromProject := range audit_event.DelUserFromProject(aes) {
		log.Info("仓库删除用户 ", *userDelFromProject)
	}
	// 【仓库异动用户权限】
	log.Info("[###] 库异动用户权限检测开始")
	for _, changeUserAccessFromProject := range audit_event.ChangeUserAccessFromProject(aes) {
		log.Info("库异动用户权限 ", *changeUserAccessFromProject)
	}
	log.Info("[###] 库异动用户权限检测完毕")
	log.Info("[###] 仓库删除用户检测完毕")
	// 【新建组】 新建组 行为没有日志
	// 【删除组】
	log.Info("[###] 删除组检测开始")
	for _, deleteGroup := range audit_event.DeletedGroup(aes) {
		log.Info("删除组 ", *deleteGroup)
	}
	log.Info("[###] 删除组检测完毕")
	// 【组添加用户】
	log.Info("[###] 组添加用户检测开始")
	for _, userAdd2Group := range audit_event.AddUser2Group(aes) {
		log.Info("组添加用户 ", *userAdd2Group)
	}
	log.Info("[###] 组添加用户检测完毕")
	// 【组删除用户】
	log.Info("[###] 组删除用户检测开始")
	for _, delUserFromProject := range audit_event.DelUserFromGroup(aes) {
		log.Info("组删除用户 ", *delUserFromProject)
	}
	log.Info("[###] 组删除用户检测完毕")
	// 【组异动用户权限】
	log.Info("[###] 组异动用户权限检测开始")
	for _, changeUserAccessFromGroup := range audit_event.ChangeUserAccessFromGroup(aes) {
		log.Info("组异动用户权限 ", *changeUserAccessFromGroup)
	}
	log.Info("[###] 组异动用户权限检测完毕")

	// 【Admin Settings】
	log.Info("[#] Admin Settings检测开始")
	setting, err := settings.GetAllEvents(gitlabClient)
	if err != nil {
		log.Error(err)
	}
	//允许git认证：通过HTTPS使用git时，使用密码认证
	passwordAuthenticationEnabledForGitDetect := settings.PasswordAuthenticationEnabledForGitDetect(setting)
	log.Info("[###] 允许git使用密码认证", passwordAuthenticationEnabledForGitDetect)
	//允许web认证：使用密码登录WEB
	passwordAuthenticationEnabledForWebDetect := settings.PasswordAuthenticationEnabledForWebDetect(setting)
	log.Info("[###] 允许Web使用密码认证", passwordAuthenticationEnabledForWebDetect)
	// 双因素认证
	requireTwoFactorAuthentication := settings.RequireTwoFactorAuthenticationDetect(setting)
	log.Info("[###] 是否允许开启双因素", requireTwoFactorAuthentication)
	// 是否允许用户注册
	signupEnabledDetect := settings.SignupEnabledDetect(setting)
	log.Info("[###] 是否允许用户注册", signupEnabledDetect)
	// 新用户注册后处于等待admin批准的待批准状态
	requireAdminApprovalAfterUserSignup := settings.RequireAdminApprovalAfterUserSignupDetect(setting)
	log.Info("[###] 新用户注册后处于等待admin批准的待批准状态", requireAdminApprovalAfterUserSignup)
	// 默认分支保护
	defaultBranchProtectionDetect := settings.DefaultBranchProtectionDetect(setting)
	log.Info("[###] 默认分支保护", defaultBranchProtectionDetect)
	// 用户登录会话默认维持时间
	sessionExpireDelayDetect := settings.SessionExpireDelayDetect(setting)
	log.Info("[###] 用户登录会话默认维持时间(秒)", sessionExpireDelayDetect)
	// 项目的默认可见度
	defaultProjectVisibilityDetect := settings.DefaultProjectVisibilityDetect(setting)
	log.Info("[###] 项目的默认可见度", defaultProjectVisibilityDetect)
	// 组的默认可见度
	defaultGroupVisibility := settings.DefaultGroupVisibilityDetect(setting)
	log.Info("[###] 组的默认可见度", defaultGroupVisibility)
	// 显示需要添加ssh密钥的警告信息
	userShowAddSSHKeyMessageDetect := settings.UserShowAddSSHKeyMessageDetect(setting)
	log.Info("[###] 显示需要添加ssh密钥的警告信息", userShowAddSSHKeyMessageDetect)

	// 密码复杂度检测
	log.Info("[###] 密码复杂度检测开始")
	PasswordNumberRequired, PasswordSymbolRequired, PasswordUppercaseRequired, PasswordLowercaseRequired, MinimumPasswordLength := settings.PasswordComplexDetect(setting)
	log.Info("是否要求数字", PasswordNumberRequired)
	log.Info("是否要求特殊字符", PasswordSymbolRequired)
	log.Info("是否要求大写字母", PasswordUppercaseRequired)
	log.Info("是否要求小写字母", PasswordLowercaseRequired)
	log.Info("密码最小长度", MinimumPasswordLength)
	log.Info("[###] 密码复杂度检测完毕")

	// 转发npm包请求
	NPMPackageRequestsForwarding := settings.NPMPackageRequestsForwardingDetect(setting)
	log.Info("[###] 转发npm包请求", NPMPackageRequestsForwarding)

	// 转发pypi包请求
	PyPIPackageRequestsForwarding := settings.PyPIPackageRequestsForwardingDetect(setting)
	log.Info("[###] 转发pypi包请求", PyPIPackageRequestsForwarding)

	// PAT前缀
	personalAccessTokenPrefix := settings.PersonalAccessTokenPrefixDetect(setting)
	log.Info("[###] PAT前缀", personalAccessTokenPrefix)

	// PAT最大生存时间
	MaxPersonalAccessTokenLifetime := settings.MaxPersonalAccessTokenLifetimeDetect(setting)
	log.Info("[###] PAT最大生存时间", MaxPersonalAccessTokenLifetime)

	// 默认项目创建保护
	defaultProjectCreation := settings.DefaultProjectCreationDetect(setting)
	log.Info("[###] 默认项目创建保护", defaultProjectCreation)

	// 禁用 OAuth 登录源
	disabledOauthSignInSources := settings.DisabledOauthSignInSourcesDetect(setting)
	log.Info("[###] 禁用 OAuth 登录源", disabledOauthSignInSources)

	log.Info("[#] Admin Settings检测完毕")

	// 【project】
	log.Info("[#] 项目配置分析检测开始")
	AllProjects, projectsService := project.ListAllProjects(gitlabClient)
	for _, single_project := range AllProjects {
		log.Info("[###] 当前检查project:", single_project.ID, single_project.Name)
		projectVisibility := project.ProjectVisibility(single_project)
		log.Info("[###] 可见性:", projectVisibility)
		projectSecurityAndCompliance := project.ProjectSecurityAndCompliance(single_project)
		log.Info("[###] 安全与合规:", projectSecurityAndCompliance)
		projectApprovalsBeforeMerge := project.ProjectApprovalsBeforeMerge(single_project)
		log.Info("[###] 批准合并请求人数:", projectApprovalsBeforeMerge)
		projectRejectUnsignedCommits := project.ProjectRejectUnsignedCommits(projectsService, single_project.ID)
		log.Info("[###] 未通过 GPG 签名时拒绝提交", projectRejectUnsignedCommits)
		projectCommitCommitterCheck := project.ProjectCommitCommitterCheck(projectsService, single_project.ID)
		log.Info("[###] 禁止未经验证的用户提交commit", projectCommitCommitterCheck)

	}

	log.Info("[#] 项目配置分析检测完毕")

	log.Debug("[#] 审计功能检测完毕")
}
