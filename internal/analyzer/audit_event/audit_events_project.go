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
