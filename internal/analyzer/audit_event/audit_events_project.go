package audit_event

import "gitlab-misconfig/internal/gitlab"

// 新建仓库
func ProjectCreated(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	projectCreated := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) && ae_v.EntityType == "Project" {
			if ae_v.Details.Add == "project" && ae_v.Details.TargetType == "Project" {
				projectCreated = append(projectCreated, ae_v)
			}
		}

	}
	return projectCreated
}

// 删除仓库
func ProjectDeleted(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	projectDeleted := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) && ae_v.EntityType == "Project" {
			if ae_v.Details.Remove == "project" && ae_v.Details.TargetType == "Project" {
				projectDeleted = append(projectDeleted, ae_v)
			}
		}

	}
	return projectDeleted
}

// 仓库添加用户
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

// 仓库删除用户
func DelUserFromProject(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	userDelFromProject := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) && ae_v.EntityType == "Project" {
			if ae_v.Details.Remove == "user_access" {
				userDelFromProject = append(userDelFromProject, ae_v)
			}
		}

	}
	return userDelFromProject
}

// 仓库异动用户权限
func ChangeUserAccessFromProject(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	changeUserAccessFromProject := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) && ae_v.EntityType == "Project" {
			if ae_v.Details.Change == "access_level" && ae_v.Details.TargetType == "User" {
				changeUserAccessFromProject = append(changeUserAccessFromProject, ae_v)
			}
		}

	}
	return changeUserAccessFromProject
}
