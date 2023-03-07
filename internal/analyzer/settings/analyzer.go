package settings

import (
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"

	"github.com/spf13/viper"
)

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {

	// 【Admin Settings】
	setting, err := GetAllEvents(gitlabClient)
	if err != nil {
		log.Error(err)
	}

	// 是否允许用户注册
	signupEnabledDetect := SignupEnabledDetect(setting)
	if signupEnabledDetect {
		requireAdminApprovalAfterUserSignup := RequireAdminApprovalAfterUserSignupDetect(setting)
		if requireAdminApprovalAfterUserSignup {
			log.Info("[+] 用户注册功能开放，且开启admin审核")
		} else {
			log.Info("[-] 用户注册功能开放，但未开启admin审核")
		}

		signUpDefaultExternalDetect := SignUpDefaultExternalDetect(setting)
		if signUpDefaultExternalDetect {
			log.Info("[+] 用户注册默认为外部用户")

			signupInternalEmailRegDetect := SignupInternalEmailRegDetect(setting)
			if signupInternalEmailRegDetect == "" {
				log.Info("[-] 未配置内部用户邮箱匹配")
			} else {
				log.Info("[+] 已配置内部用户邮箱匹配")
			}
		} else {
			log.Info("[-] 用户注册默认为内部用户")
		}

	} else {
		log.Info("[+] 用户注册功能关闭")
	}

	// 密码复杂度检测
	PasswordNumberRequired, PasswordSymbolRequired, PasswordUppercaseRequired, PasswordLowercaseRequired, MinimumPasswordLength := PasswordComplexDetect(setting)
	MinimumLength := config.GetInt("settings.password.minimum_length.keywords")
	leastPasswdComplexItems := config.GetInt("settings.password.least_complex.keywords")
	if MinimumPasswordLength >= MinimumLength {
		ComplexRate := 0
		for _, v := range []bool{PasswordNumberRequired, PasswordSymbolRequired, PasswordUppercaseRequired, PasswordLowercaseRequired} {
			if v {
				ComplexRate++
			}
		}
		if ComplexRate >= leastPasswdComplexItems {
			log.Info("[+] 密码复杂度要求合格")
		} else {
			log.Info("[-] 密码复杂度要求不合格")
		}
	} else {
		log.Info("[-] 密码复杂度过低，长度要求低于", MinimumLength)
	}
	if PasswordNumberRequired {
		log.Info("要求密码有数字")
	} else {
		log.Info("不要求密码有数字")
	}
	if PasswordSymbolRequired {
		log.Info("要求密码有特殊字符")
	} else {
		log.Info("不要求密码有特殊字符")
	}
	if PasswordUppercaseRequired {
		log.Info("要求密码有大写字母")
	} else {
		log.Info("不要求密码有大写字母")
	}
	if PasswordLowercaseRequired {
		log.Info("要求密码有小写字母")
	} else {
		log.Info("不要求密码有小写字母")
	}
	log.Info("要求密码长度", MinimumPasswordLength)

	// 双因素认证
	requireTwoFactorAuthentication := RequireTwoFactorAuthenticationDetect(setting)
	if requireTwoFactorAuthentication {
		log.Info("[ ] 系统开启全局双因素验证")
	} else {
		log.Info("[ ] 系统未开启全局双因素验证")
	}

	// 默认分支保护
	defaultBranchProtectionDetect := DefaultBranchProtectionDetect(setting)
	switch defaultBranchProtectionDetect {
	case 0:
		// 0 未受保护，具有 Developer 角色或 Maintainer 角色的用户都可以推送新提交和强制推送）
		log.Info("[-] 分支保护：未受保护")
	case 1:
		// 1 （部分保护，具有 Developer 角色或 Maintainer 角色的用户可以推送新提交，但不能强制推送）
		log.Info("[-] 分支保护：部分保护")
	case 2:
		// 2 （完全protected，具有 Developer 或 Maintainer 角色的用户不能推送新提交，但具有 Developer 或 Maintainer 角色的用户可以；没有人可以强制推送）
		log.Info("[+] 分支保护：完全保护")
	}

	// 用户登录会话默认维持时间
	SessionExpireDelayTime := config.GetInt("settings.session_expire_delay.keywords")
	sessionExpireDelayDetect := SessionExpireDelayDetect(setting)
	if sessionExpireDelayDetect >= SessionExpireDelayTime {
		log.Info("[-] 用户登录的session存活时间过长:", sessionExpireDelayDetect, "分钟")
	} else {
		log.Info("[+] 用户登录的session存活时间合规")
	}

	// 受限的能见度
	RestrictedVisibilityLevels_Contains_Private := false
	RestrictedVisibilityLevels := RestrictedVisibilityLevelsDetect(setting)
	for _, v := range RestrictedVisibilityLevels {
		if v == "public" {
			RestrictedVisibilityLevels_Contains_Private = true
			break
		}
	}
	if RestrictedVisibilityLevels_Contains_Private {
		log.Info("[-] 未禁止更改项目可见性为public")
	} else {
		log.Info("[+] 禁止更改项目可见性为public")
	}

	// 项目的默认可见度
	defaultProjectVisibilityDetect := DefaultProjectVisibilityDetect(setting)
	if defaultProjectVisibilityDetect == "private" {
		log.Info("[+] 项目默认可见度合规")
	} else {
		log.Info("[-] 项目默认可见度不合规")
	}

	// 组的默认可见度
	defaultGroupVisibility := DefaultGroupVisibilityDetect(setting)
	if defaultGroupVisibility == "private" {
		log.Info("[+] 组默认可见度合规")
	} else {
		log.Info("[-] 组默认可见度不合规")
	}

	// PAT最大生存时间
	MaxPersonalAccessTokenLifetime := MaxPersonalAccessTokenLifetimeDetect(setting)
	if MaxPersonalAccessTokenLifetime == 0 {
		log.Info("[-] PAT的最大生存时间没有限制")
	} else {
		log.Info("[ ] PAT最大生存时间(单位:天)", MaxPersonalAccessTokenLifetime)
	}

}
