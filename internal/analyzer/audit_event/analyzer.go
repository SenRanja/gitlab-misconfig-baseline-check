package audit_event

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"
)

func (Analyzer) AutoAnalysis(gitlabClient *gitlab.Client, options *types.Options, config *viper.Viper) {

	acquireTime := config.GetInt("audit_events.time.keywords")
	acquireItems := config.GetInt("audit_events.items.keywords")
	acquireItemsPerPage := config.GetInt("audit_events.per_page_items.keywords")

	// 【日志审计】
	log.Info("[#] 审计功能")
	// 获取近期审计日志
	aes, err := GetAllEvents(gitlabClient, acquireTime, acquireItems, acquireItemsPerPage)
	if err != nil {
		fmt.Println(err)
	}
	// 【新建用户】
	for _, userCreated := range UserCreated(aes) {
		log.Info("新建用户", *userCreated)
	}
	// 【删除用户】
	for _, userDeleted := range UserDeleted(aes) {
		log.Info("删除用户", *userDeleted)
	}
	// 【登陆失败】
	for _, userFailedLogin := range UserLoginFailed(aes) {
		log.Info("登陆失败", *userFailedLogin)
	}
	//// 【登陆正常】
	//log.Info("[###] 登陆正常检测开始")
	//for _, userLoginCorrectly := range UserLoginCorrectly(aes) {
	//	log.Info("登陆正常", *userLoginCorrectly)
	//}
	//log.Info("[###] 登陆正常检测完毕")

	// 【新建仓库】
	for _, projectCreated := range ProjectCreated(aes) {
		log.Info("新建仓库", *projectCreated)
	}
	// 【删除仓库】
	for _, projectDeleted := range ProjectDeleted(aes) {
		log.Info("删除仓库", *projectDeleted)
	}
	// 【仓库添加用户】
	for _, userAdd2Project := range AddUser2Project(aes) {
		log.Info("仓库添加用户 ", *userAdd2Project)
	}
	// 【仓库删除用户】
	for _, userDelFromProject := range DelUserFromProject(aes) {
		log.Info("仓库删除用户 ", *userDelFromProject)
	}
	// 【仓库异动用户权限】
	for _, changeUserAccessFromProject := range ChangeUserAccessFromProject(aes) {
		log.Info("库异动用户权限 ", *changeUserAccessFromProject)
	}
	// 【新建组】 新建组 行为没有日志
	// 【删除组】
	for _, deleteGroup := range DeletedGroup(aes) {
		log.Info("删除组 ", *deleteGroup)
	}
	// 【组添加用户】
	for _, userAdd2Group := range AddUser2Group(aes) {
		log.Info("组添加用户 ", *userAdd2Group)
	}
	// 【组删除用户】
	for _, delUserFromProject := range DelUserFromGroup(aes) {
		log.Info("组删除用户 ", *delUserFromProject)
	}
	// 【组异动用户权限】
	for _, changeUserAccessFromGroup := range ChangeUserAccessFromGroup(aes) {
		log.Info("组异动用户权限 ", *changeUserAccessFromGroup)
	}
}
