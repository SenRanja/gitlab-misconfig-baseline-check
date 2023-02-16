package audit_event

import "gitlab-misconfig/internal/gitlab"

// 新建组没有日志

// 删除组
func DeletedGroup(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	deletedGroup := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) && ae_v.EntityType == "Group" {
			if ae_v.Details.Remove == "group" && ae_v.Details.TargetType == "Group" {
				deletedGroup = append(deletedGroup, ae_v)
			}
		}

	}
	return deletedGroup
}

// 向组添加用户
// `entity_id`是 组id，`entity_type`是 组，`target_id` 是用户id，`target_details` 是用户的名字
func AddUser2Group(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	addUser2Group := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) && ae_v.EntityType == "Group" {
			if ae_v.Details.Add == "user_access" && ae_v.Details.TargetType == "User" {
				addUser2Group = append(addUser2Group, ae_v)
			}
		}

	}
	return addUser2Group
}

// 组删除用户
// `entity_id`是 组id，`entity_type`是 组，`target_id` 是用户id，`target_details` 是用户的名字
func DelUserFromGroup(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	delUserFromGroup := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) && ae_v.EntityType == "Group" {
			if ae_v.Details.Remove == "user_access" && ae_v.Details.TargetType == "User" {
				delUserFromGroup = append(delUserFromGroup, ae_v)
			}
		}

	}
	return delUserFromGroup
}

// 组异动用户
func ChangeUserAccessFromGroup(aes []*gitlab.AuditEvent) []*gitlab.AuditEvent {
	changeUserAccessFromGroup := []*gitlab.AuditEvent{}
	for _, ae_v := range aes {
		if ae_v.Details != (gitlab.AuditEventDetails{}) && ae_v.EntityType == "Group" {
			if ae_v.Details.Change == "access_level" && ae_v.Details.TargetType == "User" {
				changeUserAccessFromGroup = append(changeUserAccessFromGroup, ae_v)
			}
		}

	}
	return changeUserAccessFromGroup
}
