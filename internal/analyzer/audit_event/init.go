package audit_event

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
	"time"
)

type Analyzer struct {
}

// 获取近期的审计事件。默认获取1000天内、2000条目的审计日志
// 这是审计类的基础类。其他基线检查项在获取此方法的返回值后进行遍历统计。
func GetAllEvents(gitlabClient *gitlab.Client, acquireTime int) ([]*gitlab.AuditEvent, error) {
	var aes_all []*gitlab.AuditEvent
	var err error

	CreatedAfterTime := time.Now().AddDate(0, 0, -acquireTime)
	CreatedBeforeTime := time.Now()

	var opt *gitlab.ListAuditEventsOptions

	i := 1
	for {
		opt = &gitlab.ListAuditEventsOptions{
			ListOptions: gitlab.ListOptions{
				Page:    i,
				PerPage: 2000,
			},
			CreatedAfter:  &CreatedAfterTime,
			CreatedBefore: &CreatedBeforeTime,
		}
		i++

		aes, _, err := gitlabClient.AuditEvents.ListInstanceAuditEvents(opt)
		if err != nil {
			fmt.Println(err)
		}
		aes_all = append(aes_all, aes...)

		if len(aes) < 2000 {
			break
		}
	}

	return aes_all, err
}
