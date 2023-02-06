package audit_event

import "gitlab-misconfig/internal/gitlab"

// 新建组没有日志

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
