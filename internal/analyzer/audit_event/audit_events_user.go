package audit_event

import (
	"gitlab-misconfig/internal/gitlab"
)

// 被删除的用户
func UserDeleted(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	userDeleted := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) {
			if ae_v.Details.Remove == "user" {
				userDeleted = append(userDeleted, ae_v)
			}
		}

	}
	return userDeleted
}

// 创建的用户
func UserCreated(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	userCreated := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) {
			if ae_v.Details.Add == "user" {
				userCreated = append(userCreated, ae_v)
			}
		}

	}
	return userCreated
}

// 登陆失败
func UserLoginFailed(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	userLoginFailed := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) {
			if ae_v.Details.FailedLogin != "" {
				userLoginFailed = append(userLoginFailed, ae_v)
			}
		}

	}
	return userLoginFailed
}

// 正常登录的用户
func UserLoginCorrectly(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	userLoginCorrectly := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) {
			if ae_v.Details.With != "" {
				userLoginCorrectly = append(userLoginCorrectly, ae_v)
			}
		}

	}
	return userLoginCorrectly
}

// 用户被添加到项目中
// 没看懂 member_id 是干啥的，推测是每个仓库都有独立的`member`编号，用来记录仓库成员
// `entity_id` 是仓库id
// `entity_path` 是仓库名称
func AddUser2Project(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	userAdd2Project := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) && ae_v.EntityType == "Project" {
			if ae_v.Details.Add == "user_access" {
				userAdd2Project = append(userAdd2Project, ae_v)
			}
		}

	}
	return userAdd2Project
}
