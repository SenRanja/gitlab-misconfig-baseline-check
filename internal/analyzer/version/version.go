package version

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
)

// 版本单独从 http://192.168.3.199:40080/api/v4/version 获取
func VersionDetect(gitlabClient *gitlab.Client) (gitlab.Version, error) {
	v := gitlabClient.Version
	version, _, err := v.GetVersion()
	if err != nil {
		fmt.Println(err)
	}
	return *version, err
}
