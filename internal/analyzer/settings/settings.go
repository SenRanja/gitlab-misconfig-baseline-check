package settings

import (
	"gitlab-misconfig/internal/gitlab"
)

func GetAllEvents(gitlabClient *gitlab.Client) (*gitlab.Settings, error) {
	settings, _, err := gitlabClient.Settings.GetSettings()
	if err != nil {
		return nil, err
	}
	return settings, err
}

// 用户注册时密码复杂度检测
func PasswordComplexDetect(settings *gitlab.Settings) (bool, bool, bool, bool, int) {
	return settings.PasswordNumberRequired, settings.PasswordSymbolRequired, settings.PasswordUppercaseRequired, settings.PasswordLowercaseRequired, settings.MinimumPasswordLength
}

// 允许git认证：通过HTTPS使用git时，使用密码认证
func PasswordAuthenticationEnabledForGitDetect(settings *gitlab.Settings) bool {
	return settings.PasswordAuthenticationEnabledForGit
}

// 允许web认证：使用密码登录WEB
func PasswordAuthenticationEnabledForWebDetect(settings *gitlab.Settings) bool {
	return settings.PasswordAuthenticationEnabledForWeb
}

// 是否允许双因素认证
func RequireTwoFactorAuthenticationDetect(settings *gitlab.Settings) bool {
	return settings.RequireTwoFactorAuthentication
}

// 是否允许用户注册
func SignupEnabledDetect(settings *gitlab.Settings) bool {
	return settings.SignupEnabled
}

// 新用户注册后处于等待admin批准的待批准状态
func RequireAdminApprovalAfterUserSignupDetect(settings *gitlab.Settings) bool {
	return settings.RequireAdminApprovalAfterUserSignup
}

// 默认分支保护
func DefaultBranchProtectionDetect(settings *gitlab.Settings) int {
	return settings.DefaultBranchProtection
}

// 用户登录会话默认维持时间
func SessionExpireDelayDetect(settings *gitlab.Settings) int {
	return settings.SessionExpireDelay
}

// 项目的默认可见度
func DefaultProjectVisibilityDetect(settings *gitlab.Settings) string {
	return string(settings.DefaultProjectVisibility)
}

// 组的默认可见度
func DefaultGroupVisibilityDetect(settings *gitlab.Settings) string {
	return string(settings.DefaultGroupVisibility)
}

// 显示需要添加ssh密钥的警告信息
func UserShowAddSSHKeyMessageDetect(settings *gitlab.Settings) bool {
	return settings.UserShowAddSSHKeyMessage
}

// 转发mvn包请求，没有找到该项设置

// 转发npm包请求
func NPMPackageRequestsForwardingDetect(settings *gitlab.Settings) bool {
	return settings.NPMPackageRequestsForwarding
}

// 转发pypi包请求
func PyPIPackageRequestsForwardingDetect(settings *gitlab.Settings) bool {
	return settings.PyPIPackageRequestsForwarding
}

// PAT前缀
func PersonalAccessTokenPrefixDetect(settings *gitlab.Settings) string {
	return settings.PersonalAccessTokenPrefix
}

// 默认项目创建保护
func DefaultProjectCreationDetect(settings *gitlab.Settings) int {
	return settings.DefaultProjectCreation
}

// 禁用 OAuth 登录源
func DisabledOauthSignInSourcesDetect(settings *gitlab.Settings) []string {
	return settings.DisabledOauthSignInSources
}

// PAT最大生存时间
func MaxPersonalAccessTokenLifetimeDetect(settings *gitlab.Settings) int {
	return settings.MaxPersonalAccessTokenLifetime
}
