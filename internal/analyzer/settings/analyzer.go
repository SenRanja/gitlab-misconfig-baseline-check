package settings

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"
)

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {

	// 【Admin Settings】
	log.Debug("[#] Admin Settings检测开始")
	setting, err := GetAllEvents(gitlabClient)
	if err != nil {
		log.Error(err)
	}
	//允许git认证：通过HTTPS使用git时，使用密码认证
	passwordAuthenticationEnabledForGitDetect := PasswordAuthenticationEnabledForGitDetect(setting)
	log.Info("[###] 允许git使用密码认证", passwordAuthenticationEnabledForGitDetect)
	//允许web认证：使用密码登录WEB
	passwordAuthenticationEnabledForWebDetect := PasswordAuthenticationEnabledForWebDetect(setting)
	log.Info("[###] 允许Web使用密码认证", passwordAuthenticationEnabledForWebDetect)
	// 双因素认证
	requireTwoFactorAuthentication := RequireTwoFactorAuthenticationDetect(setting)
	log.Info("[###] 是否允许开启双因素", requireTwoFactorAuthentication)
	// 是否允许用户注册
	signupEnabledDetect := SignupEnabledDetect(setting)
	log.Info("[###] 是否允许用户注册", signupEnabledDetect)
	// 新用户注册后处于等待admin批准的待批准状态
	requireAdminApprovalAfterUserSignup := RequireAdminApprovalAfterUserSignupDetect(setting)
	log.Info("[###] 新用户注册后处于等待admin批准的待批准状态", requireAdminApprovalAfterUserSignup)
	// 默认分支保护
	defaultBranchProtectionDetect := DefaultBranchProtectionDetect(setting)
	log.Info("[###] 默认分支保护", defaultBranchProtectionDetect)
	// 用户登录会话默认维持时间
	sessionExpireDelayDetect := SessionExpireDelayDetect(setting)
	log.Info("[###] 用户登录会话默认维持时间(秒)", sessionExpireDelayDetect)
	// 受限的能见度
	RestrictedVisibilityLevels := RestrictedVisibilityLevelsDetect(setting)
	log.Info("[###] 受限的能见度", RestrictedVisibilityLevels)
	// 项目的默认可见度
	defaultProjectVisibilityDetect := DefaultProjectVisibilityDetect(setting)
	log.Info("[###] 项目的默认可见度", defaultProjectVisibilityDetect)
	// 组的默认可见度
	defaultGroupVisibility := DefaultGroupVisibilityDetect(setting)
	log.Info("[###] 组的默认可见度", defaultGroupVisibility)
	// 显示需要添加ssh密钥的警告信息
	userShowAddSSHKeyMessageDetect := UserShowAddSSHKeyMessageDetect(setting)
	log.Info("[###] 显示需要添加ssh密钥的警告信息", userShowAddSSHKeyMessageDetect)

	// 密码复杂度检测
	log.Info("[###] 密码复杂度检测开始")
	PasswordNumberRequired, PasswordSymbolRequired, PasswordUppercaseRequired, PasswordLowercaseRequired, MinimumPasswordLength := PasswordComplexDetect(setting)
	log.Info("是否要求数字", PasswordNumberRequired)
	log.Info("是否要求特殊字符", PasswordSymbolRequired)
	log.Info("是否要求大写字母", PasswordUppercaseRequired)
	log.Info("是否要求小写字母", PasswordLowercaseRequired)
	log.Info("密码最小长度", MinimumPasswordLength)
	log.Info("[###] 密码复杂度检测完毕")

	// 转发npm包请求
	NPMPackageRequestsForwarding := NPMPackageRequestsForwardingDetect(setting)
	log.Info("[###] 转发npm包请求", NPMPackageRequestsForwarding)

	// 转发pypi包请求
	PyPIPackageRequestsForwarding := PyPIPackageRequestsForwardingDetect(setting)
	log.Info("[###] 转发pypi包请求", PyPIPackageRequestsForwarding)

	// PAT前缀
	personalAccessTokenPrefix := PersonalAccessTokenPrefixDetect(setting)
	log.Info("[###] PAT前缀", personalAccessTokenPrefix)

	// PAT最大生存时间
	MaxPersonalAccessTokenLifetime := MaxPersonalAccessTokenLifetimeDetect(setting)
	log.Info("[###] PAT最大生存时间", MaxPersonalAccessTokenLifetime)

	// 默认项目创建保护
	defaultProjectCreation := DefaultProjectCreationDetect(setting)
	log.Info("[###] 默认项目创建保护", defaultProjectCreation)

	// 禁用 OAuth 登录源
	disabledOauthSignInSources := DisabledOauthSignInSourcesDetect(setting)
	log.Info("[###] 禁用 OAuth 登录源", disabledOauthSignInSources)

	log.Info("[#] Admin Settings检测完毕")

	log.Debug("[#] 审计功能检测完毕")
}
