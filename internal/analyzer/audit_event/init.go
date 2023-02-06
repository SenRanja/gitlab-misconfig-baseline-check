package audit_event

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
	"time"
)

// 获取近期的审计事件。默认获取1000天内、2000条目的审计日志
// 这是审计类的基础类。其他基线检查项在获取此方法的返回值后进行遍历统计。
func GetAllEvents(gitlabClient *gitlab.Client) ([]*gitlab.AuditEvent, error) {
	CreatedAfterTime := time.Now().AddDate(0, 0, -1000)
	CreatedBeforeTime := time.Now()
	opt := &gitlab.ListAuditEventsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 2000,
		},
		CreatedAfter:  &CreatedAfterTime,
		CreatedBefore: &CreatedBeforeTime,
	}

	aes, _, err := gitlabClient.AuditEvents.ListInstanceAuditEvents(opt)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("请求审计事件数量：", len(aes))
	//for ae, ae_v := range aes {
	//	fmt.Println(ae, ae_v.ID, ae_v.AuthorID, ae_v.EntityID, ae_v.EntityType, ae_v.Details, ae_v.CreatedAt)
	//}
	return aes, err
}
