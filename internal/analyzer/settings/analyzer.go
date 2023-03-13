package settings

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"
	"strconv"
)

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper, output *types.Output) {

	// 【Admin Settings】
	setting, err := GetAllEvents(gitlabClient)
	if err != nil {
		log.Error(err)
	}

	// 是否允许用户注册
	output.Settings.Register.RegisterEnable.Result = SignupEnabledDetect(setting)

	if output.Settings.Register.RegisterEnable.Result {
		output.Settings.Register.RegisterEnable.Complaince = false

		// 获取是否开启邮箱验证
		output.Settings.Register.EmailConfirmation.Result = SendConfirmationEmailDetect(setting)
		if output.Settings.Register.EmailConfirmation.Result {
			output.Settings.Register.EmailConfirmation.Complaince = true
		} else {
			output.Settings.Register.EmailConfirmation.Complaince = false
		}

		// 是否 用户注册等待admin批准
		output.Settings.Register.AdminApproval.Result = RequireAdminApprovalAfterUserSignupDetect(setting)
		if output.Settings.Register.AdminApproval.Result {
			output.Settings.Register.AdminApproval.Complaince = true
			log.Info("[+] 用户注册功能开放，且开启admin审核")
		} else {
			output.Settings.Register.AdminApproval.Complaince = false
			log.Info("[-] 用户注册功能开放，但未开启admin审核")
		}

		output.Settings.Register.External.Result = SignUpDefaultExternalDetect(setting)
		if output.Settings.Register.External.Result {
			output.Settings.Register.External.Complaince = true
			log.Info("[+] 用户注册默认为外部用户")
		} else {
			output.Settings.Register.External.Complaince = false
			log.Info("[-] 用户注册默认为内部用户")
		}

		signupInternalEmailRegDetect := SignupInternalEmailRegDetect(setting)

		if signupInternalEmailRegDetect == "" {
			output.Settings.Register.EmailRegexpInternal.Result = true
			output.Settings.Register.EmailRegexpInternal.Complaince = true
			log.Info("[-] 未配置内部用户邮箱匹配")
		} else {
			output.Settings.Register.EmailRegexpInternal.Result = false
			output.Settings.Register.EmailRegexpInternal.Complaince = false
			log.Info("[+] 已配置内部用户邮箱匹配")
		}

	} else {
		output.Settings.Register.RegisterEnable.Complaince = true
		log.Info("[+] 用户注册功能关闭")
	}

	// 密码复杂度检测
	// 获取gitlab的关于密码的配置
	output.Settings.Password.Num.Result, output.Settings.Password.Special.Result, output.Settings.Password.Upper.Result, output.Settings.Password.Lower.Result, output.Settings.Password.Length.Result = PasswordComplexDetect(setting)
	// 获取阈值
	output.Settings.Password.Length.Keyword = config.GetInt("settings.password.minimum_length.keywords")
	output.Settings.Password.Keyword = config.GetInt("settings.password.least_complex.keywords")
	// 如果密码长度合规，长度>=阈值
	if output.Settings.Password.Length.Result >= output.Settings.Password.Length.Keyword {
		output.Settings.Password.Length.Complaince = true

		ComplexRate := 0
		for _, v := range []bool{output.Settings.Password.Num.Result, output.Settings.Password.Special.Result, output.Settings.Password.Upper.Result, output.Settings.Password.Lower.Result} {
			if v {
				ComplexRate++
			}
		}
		if ComplexRate >= output.Settings.Password.Keyword {
			output.Settings.Password.Result = true
			output.Settings.Password.Complaince = true
			log.Info("[+] 密码复杂度要求合格")
		} else {
			output.Settings.Password.Result = false
			output.Settings.Password.Complaince = false
			log.Info("[-] 密码复杂度要求不合格")
		}
	} else {
		// 如果密码长度不合规，过短
		output.Settings.Password.Length.Complaince = false
		log.Info("[-] 密码复杂度过低，长度要求低于", output.Settings.Password.Length.Keyword)
	}

	if output.Settings.Password.Num.Result {
		output.Settings.Password.Num.Complaince = true
		log.Info("要求密码有数字")
	} else {
		output.Settings.Password.Num.Complaince = false
		log.Info("不要求密码有数字")
	}
	if output.Settings.Password.Special.Result {
		output.Settings.Password.Special.Complaince = true
		log.Info("要求密码有特殊字符")
	} else {
		output.Settings.Password.Special.Complaince = false
		log.Info("不要求密码有特殊字符")
	}
	if output.Settings.Password.Upper.Result {
		output.Settings.Password.Upper.Complaince = true
		log.Info("要求密码有大写字母")
	} else {
		output.Settings.Password.Upper.Complaince = false
		log.Info("不要求密码有大写字母")
	}
	if output.Settings.Password.Lower.Result {
		output.Settings.Password.Lower.Complaince = true
		log.Info("要求密码有小写字母")
	} else {
		output.Settings.Password.Lower.Complaince = false
		log.Info("不要求密码有小写字母")
	}

	// 双因素认证
	output.Settings.TwoFactorAuth.Result = RequireTwoFactorAuthenticationDetect(setting)
	if output.Settings.TwoFactorAuth.Result {
		output.Settings.TwoFactorAuth.Complaince = true
		log.Info("[ ] 系统开启全局双因素验证")
	} else {
		output.Settings.TwoFactorAuth.Complaince = false
		log.Info("[ ] 系统未开启全局双因素验证")
	}

	// 默认分支保护

	DBPD := DefaultBranchProtectionDetect(setting)

	switch DBPD {
	case 0:
		// 0 未受保护，具有 Developer 角色或 Maintainer 角色的用户都可以推送新提交和强制推送）
		output.Settings.InitProjectEnableDefaultBranchProtection.Result = "未受保护"
		output.Settings.InitProjectEnableDefaultBranchProtection.Complaince = false
		log.Info("[-] 分支保护：未受保护")
	case 1:
		// 1 （部分保护，具有 Developer 角色或 Maintainer 角色的用户可以推送新提交，但不能强制推送）
		output.Settings.InitProjectEnableDefaultBranchProtection.Result = "部分保护"
		output.Settings.InitProjectEnableDefaultBranchProtection.Complaince = false
		log.Info("[-] 分支保护：部分保护")
	case 2:
		// 2 （完全protected，具有 Developer 或 Maintainer 角色的用户不能推送新提交，但具有 Developer 或 Maintainer 角色的用户可以；没有人可以强制推送）
		output.Settings.InitProjectEnableDefaultBranchProtection.Result = "完全保护"
		output.Settings.InitProjectEnableDefaultBranchProtection.Complaince = true
		log.Info("[+] 分支保护：完全保护")
	}

	// 用户登录会话默认维持时间
	output.Settings.WebSessionExpire.Keyword = config.GetInt("settings.session_expire_delay.keywords")
	output.Settings.WebSessionExpire.Result = SessionExpireDelayDetect(setting)
	if output.Settings.WebSessionExpire.Result >= output.Settings.WebSessionExpire.Keyword {
		output.Settings.WebSessionExpire.Complaince = false
		log.Info("[-] 用户登录的session存活时间过长:", output.Settings.WebSessionExpire.Result, "分钟")
	} else {
		output.Settings.WebSessionExpire.Complaince = true
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
		output.Settings.ForbidCreatePublicProject.Result = "未禁止public"
		output.Settings.ForbidCreatePublicProject.Complaince = false
		log.Info("[-] 未禁止更改项目可见性为public")
	} else {
		output.Settings.ForbidCreatePublicProject.Result = "禁止public"
		output.Settings.ForbidCreatePublicProject.Complaince = true
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
	output.Settings.PATExpire.Result = strconv.Itoa(MaxPersonalAccessTokenLifetimeDetect(setting))
	if output.Settings.PATExpire.Result == "0" {
		output.Settings.PATExpire.Result = "永久"
		output.Settings.PATExpire.Complaince = false
		log.Info("[-] PAT的最大生存时间没有限制")
	} else {
		output.Settings.PATExpire.Complaince = true
		log.Info("[ ] PAT最大生存时间(单位:天)", output.Settings.PATExpire.Result)
	}

}
