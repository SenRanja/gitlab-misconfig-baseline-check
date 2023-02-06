package user

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/types"
	"io"
	"net/http"
	"strings"
)

// 统计不活跃用户数量
func countUnactiveNumbers(users []*gitlab.User) int {
	var totalNumberOfAuditor = 0
	for i := 0; i < len(users); i++ {
		if users[i].IsAuditor {
			totalNumberOfAuditor += 1
		}
	}
	return totalNumberOfAuditor
}

// 统计开启双因素认证用户数量
func countTwoFactorEnabled(users []*gitlab.User) int {
	var totalTwoFactorEnabled = 0
	for i := 0; i < len(users); i++ {
		if users[i].TwoFactorEnabled == true {
			totalTwoFactorEnabled += 1
		}
	}
	return totalTwoFactorEnabled
}

// 查询是否禁用注册功能，返回true为开启，返回false为关闭
func JudgeIfDisableRegister(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) (bool, error) {
	gitlab_url := options.Url + "/users/sign_in"

	resp, err := http.Get(gitlab_url)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("检查是否允许注册权限，无法进行HTTP访问")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		if strings.Contains(string(body), "Register now") {
			return true, nil
		}
	} else {
		return false, errors.New("检查是否允许注册权限，访问注册页获得非200响应码")
	}

	//req, err := gitlabClient.NewRequestWithoutAPI(http.MethodGet, "/users/sign_in")
	//if err != nil {
	//	return false
	//}
	//resp, err := gitlabClient.Do(req, nil)
	//body, err := io.ReadAll(resp.Body)
	//log.Debug(body)
	return false, errors.New("检查是否允许注册权限，未进入预期循环")
}
