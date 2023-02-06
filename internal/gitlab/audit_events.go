package gitlab

import (
	"fmt"
	"net/http"
	"time"
)

// GitLab API docs: https://docs.gitlab.com/ee/api/audit_events.html
type AuditEvent struct {
	ID int `json:"id"`
	// 日志id
	AuthorID int `json:"author_id"`
	// 操作者的user id
	// 如果是登陆失败，AuthorID 和 EntityID 都是 -1
	EntityID int `json:"entity_id"`
	// 操作对象的id
	EntityType string `json:"entity_type"`
	// 操作对象的类型：Project、Group或User
	Details   AuditEventDetails `json:"details"`
	CreatedAt *time.Time        `json:"created_at"`
	// 什么时间操作的
}

// GitLab API docs: https://docs.gitlab.com/ee/api/audit_events.html
// 访问 http://192.168.3.199:40080/api/v4/audit_events 会获得[{xx}]，xx.details就是这里的数据结构
type AuditEventDetails struct {
	With string `json:"with"`
	// with如果是"standard"，通常是正常登录的用户
	Add string `json:"add"`
	// add 类型。如添加用户，"add": "user",
	As string `json:"as"`
	// 指添加某用户具有的权限，比如是审计，运维者，访客
	Change string `json:"change"`
	From   string `json:"from"`
	To     string `json:"to"`
	Remove string `json:"remove"`
	// remove字段为User或者Group等，删除用户或项目
	CustomMessage string `json:"custom_message"`
	AuthorName    string `json:"author_name"`
	// 操作者的名字
	TargetID interface{} `json:"target_id"`
	// 操作对象的id，删除用户则是用户的id
	TargetType    string `json:"target_type"`
	TargetDetails string `json:"target_details"`
	// 目标细节，但具体不详
	// 如果是删除用户，这里就是 用户名
	// 如果是新建用户，这里就是用户名
	IPAddress string `json:"ip_address"`
	// 要么是操作者的ip，要么是登录失败者的ip
	EntityPath string `json:"entity_path"`
	// 实体路径，但具体不详
	// 如果是删除用户，这里就是 用户名
	// 如果是新建用户，这里就是用户名
	FailedLogin string `json:"failed_login"`
	// 登陆失败，此处通常为 "STANDARD"。 登陆失败这里不会存用户输入的密码。
}

// GitLab API docs: https://docs.gitlab.com/ee/api/audit_events.html
// 创建专门用于审计Audit的类
type AuditEventsService struct {
	client *Client
}

// GitLab API docs: https://docs.gitlab.com/ee/api/audit_events.html
// 此处限定时间范围，参考audit_event/audit_events.go中代码，对此处赋值
type ListAuditEventsOptions struct {
	ListOptions
	CreatedAfter  *time.Time `url:"created_after,omitempty" json:"created_after,omitempty"`
	CreatedBefore *time.Time `url:"created_before,omitempty" json:"created_before,omitempty"`
}

// GitLab API docs: https://docs.gitlab.com/ee/api/audit_events.html#retrieve-all-instance-audit-events
// 顺序打印其返回值，即 `AuditEvent` 类型，如下：
// 注意`AuditEvent`.`CreatedAt`字段是 *time.Time 类型
// fmt.Println(ae_v.ID, ae_v.AuthorID, ae_v.EntityID, ae_v.EntityType, ae_v.Details, ae_v.CreatedAt)
// 36 36 36 User {standard        syj 36 User syj 192.168.3.87 syj } 2023-02-06 02:41:27.414 +0000 UTC
func (s *AuditEventsService) ListInstanceAuditEvents(opt *ListAuditEventsOptions, options ...RequestOptionFunc) ([]*AuditEvent, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "audit_events", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var aes []*AuditEvent
	resp, err := s.client.Do(req, &aes)
	if err != nil {
		return nil, resp, err
	}

	return aes, resp, err
}

// GitLab API docs: https://docs.gitlab.com/ee/api/audit_events.html#retrieve-all-group-audit-events
// 获取某组的全部日志，传参组id
func (s *AuditEventsService) ListGroupAuditEvents(gid interface{}, opt *ListAuditEventsOptions, options ...RequestOptionFunc) ([]*AuditEvent, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/audit_events", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var aes []*AuditEvent
	resp, err := s.client.Do(req, &aes)
	if err != nil {
		return nil, resp, err
	}

	return aes, resp, err
}

// GitLab API docs: https://docs.gitlab.com/ee/api/audit_events.html#retrieve-all-project-audit-events
// 获取某仓库的全部日志，传参仓库的id，如：http://192.168.3.199:40080/api/v4/projects/35/audit_events
func (s *AuditEventsService) ListProjectAuditEvents(pid interface{}, opt *ListAuditEventsOptions, options ...RequestOptionFunc) ([]*AuditEvent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/audit_events", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var aes []*AuditEvent
	resp, err := s.client.Do(req, &aes)
	if err != nil {
		return nil, resp, err
	}

	return aes, resp, err
}
