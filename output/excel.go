package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gitlab-misconfig/internal/types"
	"strconv"
)

func ExportExcel(o *types.Output) {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName_SettingsAndUser := "系统+用户"
	f.SetSheetName("Sheet1", sheetName_SettingsAndUser)

	title := []string{
		o.OutputTitle.CheckRule,
		o.OutputTitle.SecondCheckRule,
		o.OutputTitle.Result,
		o.OutputTitle.Keyword,
		o.OutputTitle.Complaince,
		o.OutputTitle.Description,
		o.OutputTitle.Advice,
	}

	// 行 自增
	row := 0

	// Title 设置标题
	row++
	name, err := excelize.JoinCellName("A", row)
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, name, &title)
	if err != nil {
		fmt.Println(err)
	}

	// 密码复杂度-最小长度
	row++
	password_checkrule_merge_start := "A" + strconv.Itoa(row)
	password_merge_start := "E" + strconv.Itoa(row)
	password_len_name, err := excelize.JoinCellName("A", row)
	password_len := []interface{}{
		o.Settings.Password.CheckRule,
		o.Settings.Password.Length.CheckRule,
		o.Settings.Password.Length.Result,
		o.Settings.Password.Length.Keyword,
		o.Settings.Password.Length.Complaince,
		o.Settings.Password.Length.Description,
		o.Settings.Password.Length.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, password_len_name, &password_len)
	if err != nil {
		fmt.Println(err)
	}

	// 密码复杂度-要求数字
	row++
	password_num_name, err := excelize.JoinCellName("A", row)
	password_num := []interface{}{
		"",
		o.Settings.Password.Num.CheckRule,
		ConvertBool2StrIfEnable(o.Settings.Password.Num.Result),
		ConvertBool2StrIfEnable(o.Settings.Password.Num.Keyword),
		ConvertBool2StrIfComplaince(o.Settings.Password.Num.Complaince),
		o.Settings.Password.Num.Description,
		o.Settings.Password.Num.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, password_num_name, &password_num)
	if err != nil {
		fmt.Println(err)
	}

	// 密码复杂度-要求大写
	row++
	password_upper_name, err := excelize.JoinCellName("A", row)
	password_upper := []interface{}{
		"",
		o.Settings.Password.Upper.CheckRule,
		ConvertBool2StrIfEnable(o.Settings.Password.Upper.Result),
		ConvertBool2StrIfEnable(o.Settings.Password.Upper.Keyword),
		ConvertBool2StrIfComplaince(o.Settings.Password.Upper.Complaince),
		o.Settings.Password.Upper.Description,
		o.Settings.Password.Upper.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, password_upper_name, &password_upper)
	if err != nil {
		fmt.Println(err)
	}

	// 密码复杂度-要求小写
	row++
	password_lower_name, err := excelize.JoinCellName("A", row)
	password_lower := []interface{}{
		"",
		o.Settings.Password.Lower.CheckRule,
		ConvertBool2StrIfEnable(o.Settings.Password.Lower.Result),
		ConvertBool2StrIfEnable(o.Settings.Password.Lower.Keyword),
		ConvertBool2StrIfComplaince(o.Settings.Password.Lower.Complaince),
		o.Settings.Password.Lower.Description,
		o.Settings.Password.Lower.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, password_lower_name, &password_lower)
	if err != nil {
		fmt.Println(err)
	}

	// 密码复杂度-要求特殊
	row++
	password_checkrule_merge_end := "A" + strconv.Itoa(row)
	password_merge_end := "E" + strconv.Itoa(row)
	password_special_name, err := excelize.JoinCellName("A", row)
	password_special := []interface{}{
		"",
		o.Settings.Password.Special.CheckRule,
		ConvertBool2StrIfEnable(o.Settings.Password.Special.Result),
		ConvertBool2StrIfEnable(o.Settings.Password.Special.Keyword),
		ConvertBool2StrIfComplaince(o.Settings.Password.Special.Complaince),
		o.Settings.Password.Special.Description,
		o.Settings.Password.Special.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, password_special_name, &password_special)
	if err != nil {
		fmt.Println(err)
	}

	// 密码检查项合并单元格
	// 合并`合规情况`并赋值总合规情况 password_merge_start
	err = f.MergeCell(sheetName_SettingsAndUser, password_merge_start, password_merge_end)
	if err != nil {
		fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
		return
	}
	f.SetCellValue(sheetName_SettingsAndUser, password_merge_start, ConvertBool2StrIfComplaince(o.Settings.Password.Complaince))
	// 合并检查项名称 password_checkrule_merge_end
	err = f.MergeCell(sheetName_SettingsAndUser, password_checkrule_merge_start, password_checkrule_merge_end)
	if err != nil {
		fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
		return
	}

	// 双因素认证
	row++
	two_factor_auth_name, err := excelize.JoinCellName("A", row)
	two_factor_auth := []interface{}{
		o.Settings.TwoFactorAuth.CheckRule,
		"",
		ConvertBool2StrIfEnable(o.Settings.TwoFactorAuth.Result),
		"",
		"",
		o.Settings.TwoFactorAuth.Description,
		o.Settings.TwoFactorAuth.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, two_factor_auth_name, &two_factor_auth)
	if err != nil {
		fmt.Println(err)
	}

	// 禁用注册功能
	row++
	register_enable_name, err := excelize.JoinCellName("A", row)
	register_enable := []interface{}{
		o.Settings.Register.RegisterEnable.CheckRule,
		"",
		ConvertBool2StrIfEnable(o.Settings.Register.RegisterEnable.Result),
		ConvertBool2StrIfEnable(o.Settings.Register.RegisterEnable.Keyword),
		ConvertBool2StrIfComplaince(o.Settings.Register.RegisterEnable.Complaince),
		o.Settings.Register.RegisterEnable.Description,
		o.Settings.Register.RegisterEnable.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, register_enable_name, &register_enable)
	if err != nil {
		fmt.Println(err)
	}

	// 注册-启用邮箱验证
	row++
	register_merge_start := "A" + strconv.Itoa(row)
	email_confirmation_name, err := excelize.JoinCellName("A", row)
	email_confirmation := []interface{}{
		o.Settings.Register.EmailConfirmation.CheckRule,
		o.Settings.Register.EmailConfirmation.CheckRule,
		ConvertBool2StrIfEnable(o.Settings.Register.EmailConfirmation.Result),
		ConvertBool2StrIfEnable(o.Settings.Register.EmailConfirmation.Keyword),
		ConvertBool2StrIfComplaince(o.Settings.Register.EmailConfirmation.Complaince),
		o.Settings.Register.EmailConfirmation.Description,
		o.Settings.Register.EmailConfirmation.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, email_confirmation_name, &email_confirmation)
	if err != nil {
		fmt.Println(err)
	}

	// 注册-新用户注册后需等待admin批准
	row++
	admin_approval_name, err := excelize.JoinCellName("A", row)
	admin_approval := []interface{}{
		"",
		o.Settings.Register.AdminApproval.CheckRule,
		ConvertBool2StrIfEnable(o.Settings.Register.AdminApproval.Result),
		ConvertBool2StrIfEnable(o.Settings.Register.AdminApproval.Keyword),
		ConvertBool2StrIfComplaince(o.Settings.Register.AdminApproval.Complaince),
		o.Settings.Register.AdminApproval.Description,
		o.Settings.Register.AdminApproval.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, admin_approval_name, &admin_approval)
	if err != nil {
		fmt.Println(err)
	}

	// 注册-注册区分外部、内部用户
	row++
	external_name, err := excelize.JoinCellName("A", row)
	external := []interface{}{
		"",
		o.Settings.Register.External.CheckRule,
		ConvertBool2StrIfEnable(o.Settings.Register.External.Result),
		ConvertBool2StrIfEnable(o.Settings.Register.External.Keyword),
		ConvertBool2StrIfComplaince(o.Settings.Register.External.Complaince),
		o.Settings.Register.External.Description,
		o.Settings.Register.External.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, external_name, &external)
	if err != nil {
		fmt.Println(err)
	}

	// 注册-邮箱正则匹配内部用户
	row++
	register_merge_end := "A" + strconv.Itoa(row)
	email_reg_external_name, err := excelize.JoinCellName("A", row)
	email_reg_external := []interface{}{
		"",
		o.Settings.Register.EmailRegexpInternal.CheckRule,
		ConvertBool2StrIfEnable(o.Settings.Register.EmailRegexpInternal.Result),
		ConvertBool2StrIfEnable(o.Settings.Register.EmailRegexpInternal.Keyword),
		ConvertBool2StrIfComplaince(o.Settings.Register.EmailRegexpInternal.Complaince),
		o.Settings.Register.EmailRegexpInternal.Description,
		o.Settings.Register.EmailRegexpInternal.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, email_reg_external_name, &email_reg_external)
	if err != nil {
		fmt.Println(err)
	}

	// 注册功能检查项合并单元格
	err = f.MergeCell(sheetName_SettingsAndUser, register_merge_start, register_merge_end)
	if err != nil {
		fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
		return
	}
	f.SetCellValue(sheetName_SettingsAndUser, register_merge_start, "注册功能")

	// 初始化的项目启用默认分支保护
	row++
	init_project_enable_default_branch_protection_name, err := excelize.JoinCellName("A", row)
	o.Settings.InitProjectEnableDefaultBranchProtection.Keyword = "完全保护"
	init_project_enable_default_branch_protection := []interface{}{
		o.Settings.InitProjectEnableDefaultBranchProtection.CheckRule,
		"",
		o.Settings.InitProjectEnableDefaultBranchProtection.Result,
		o.Settings.InitProjectEnableDefaultBranchProtection.Keyword,
		ConvertBool2StrIfComplaince(o.Settings.InitProjectEnableDefaultBranchProtection.Complaince),
		o.Settings.InitProjectEnableDefaultBranchProtection.Description,
		o.Settings.InitProjectEnableDefaultBranchProtection.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, init_project_enable_default_branch_protection_name, &init_project_enable_default_branch_protection)
	if err != nil {
		fmt.Println(err)
	}

	// 用户登录web页面会话存活时间
	row++
	web_session_expire_name, err := excelize.JoinCellName("A", row)
	web_session_expire := []interface{}{
		o.Settings.WebSessionExpire.CheckRule,
		"",
		o.Settings.WebSessionExpire.Result,
		o.Settings.WebSessionExpire.Keyword,
		ConvertBool2StrIfComplaince(o.Settings.WebSessionExpire.Complaince),
		o.Settings.WebSessionExpire.Description,
		o.Settings.WebSessionExpire.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, web_session_expire_name, &web_session_expire)
	if err != nil {
		fmt.Println(err)
	}

	// 禁止项目的可见度设置为public
	row++
	forbid_create_public_project_name, err := excelize.JoinCellName("A", row)
	forbid_create_public_project := []interface{}{
		o.Settings.ForbidCreatePublicProject.CheckRule,
		"",
		o.Settings.ForbidCreatePublicProject.Result,
		o.Settings.ForbidCreatePublicProject.Keyword,
		ConvertBool2StrIfComplaince(o.Settings.ForbidCreatePublicProject.Complaince),
		o.Settings.ForbidCreatePublicProject.Description,
		o.Settings.ForbidCreatePublicProject.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, forbid_create_public_project_name, &forbid_create_public_project)
	if err != nil {
		fmt.Println(err)
	}

	// PAT(个人访问令牌)最大生存时间
	row++
	pat_expire_name, err := excelize.JoinCellName("A", row)
	pat_expire := []interface{}{
		o.Settings.PATExpire.CheckRule,
		"",
		o.Settings.PATExpire.Result,
		"",
		"",
		o.Settings.PATExpire.Description,
		o.Settings.PATExpire.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, pat_expire_name, &pat_expire)
	if err != nil {
		fmt.Println(err)
	}

	// 统计不活跃用户数量
	row++
	inactive_name, err := excelize.JoinCellName("A", row)
	inactive := []interface{}{
		o.User.Inactive.CheckRule,
		"",
		o.User.Inactive.Result,
		"",
		"",
		o.User.Inactive.Description,
		o.User.Inactive.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, inactive_name, &inactive)
	if err != nil {
		fmt.Println(err)
	}

	// 统计未启用账户
	row++
	unactive_name, err := excelize.JoinCellName("A", row)
	unactive := []interface{}{
		o.User.Unactive.CheckRule,
		"",
		o.User.Unactive.Result,
		"",
		"",
		o.User.Unactive.Description,
		o.User.Unactive.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, unactive_name, &unactive)
	if err != nil {
		fmt.Println(err)
	}

	// 统计开启双因素认证用户数量
	row++
	two_factor_auth_user_name, err := excelize.JoinCellName("A", row)
	two_factor_auth_user := []interface{}{
		o.User.TwoFactorAuth.CheckRule,
		"",
		o.User.TwoFactorAuth.Result,
		"",
		"",
		o.User.TwoFactorAuth.Description,
		o.User.TwoFactorAuth.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, two_factor_auth_user_name, &two_factor_auth_user)
	if err != nil {
		fmt.Println(err)
	}

	// 统计admin数量并打印名称
	row++
	admin_name, err := excelize.JoinCellName("A", row)
	admin := []interface{}{
		o.User.Admin.CheckRule,
		"",
		o.User.Admin.Result,
		"",
		"",
		o.User.Admin.Description,
		o.User.Admin.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, admin_name, &admin)
	if err != nil {
		fmt.Println(err)
	}

	// 统计审计人员数量
	row++
	auditor_name, err := excelize.JoinCellName("A", row)
	auditor := []interface{}{
		o.User.Auditor.CheckRule,
		"",
		o.User.Auditor.Result,
		"",
		"",
		o.User.Auditor.Description,
		o.User.Auditor.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, auditor_name, &auditor)
	if err != nil {
		fmt.Println(err)
	}

	// 统计外部用户数量
	row++
	external_user_name, err := excelize.JoinCellName("A", row)
	external_user := []interface{}{
		o.User.External.CheckRule,
		"",
		o.User.External.Result,
		"",
		"",
		o.User.External.Description,
		o.User.External.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, external_user_name, &external_user)
	if err != nil {
		fmt.Println(err)
	}

	// 美化外观
	headStyle, err := f.NewStyle(SetHeadStyle())
	titleStyle, err := f.NewStyle(SetTitleStyle())
	textStyle, err := f.NewStyle(SetTextStyle())
	//uncomplainceStyle, err := f.NewStyle(SetUncomplainceStyle())
	if err != nil {
		fmt.Println(err)
	}
	f.SetColWidth(sheetName_SettingsAndUser, "A", "A", 29.07)
	f.SetColWidth(sheetName_SettingsAndUser, "B", "B", 27.67)
	f.SetColWidth(sheetName_SettingsAndUser, "C", "C", 9.07)
	f.SetColWidth(sheetName_SettingsAndUser, "D", "D", 21.33)
	f.SetColWidth(sheetName_SettingsAndUser, "E", "E", 9)
	f.SetColWidth(sheetName_SettingsAndUser, "F", "F", 47)
	f.SetColWidth(sheetName_SettingsAndUser, "G", "G", 37)
	f.SetCellStyle(sheetName_SettingsAndUser, "A1", "G1", headStyle)
	f.SetCellStyle(sheetName_SettingsAndUser, "A2", "A22", titleStyle)
	f.SetCellStyle(sheetName_SettingsAndUser, "B2", "G22", textStyle)

	//res, err := f.GetCols(sheetName_SettingsAndUser)
	//fmt.Println(res)

	// ----- 项目维度 -------------------------------------------

	sheetName_Project := "项目"
	f.NewSheet(sheetName_Project)
	row = 0

	title = []string{
		"项目ID",
		"项目命名空间",
		"项目名称",
		"禁止建立public权限项目",
		"访问tag媒体文件链接时是否需要验证登录",
		"安全与合规",
		"阻止发起MR的人审批",
		"阻止push过commit的人审批",
		"阻止MR中可以编辑审批规则",
		"批准时需要密码",
		"发起MR时若有新commit被push，是否撤销批准",
		"审批MR的人数合规",
		"合规性",
		"不合规描述",
		"项目访问令牌的有效期是否过长",
		"启用或禁用合并管道",
		"启用或禁用合并队列",
		"推送规则：未GPG签名的提交",
		"推送规则：拒绝未验证邮箱的用户的push",
		"推送规则：commit作者和push者需要是经过邮箱验证的gitlab用户",
		"xxx",
		"xxxxx",
		"xxxx",
		"xxxx",
	}

	row++
	name, err = excelize.JoinCellName("A", row)
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetSheetRow(sheetName_Project, name, &title)
	if err != nil {
		fmt.Println(err)
	}

	row++
	name, err = excelize.JoinCellName("A", row)
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetSheetRow(sheetName_Project, name, &title)
	if err != nil {
		fmt.Println(err)
	}

	var MR_flag_start, MR_flag_end, PR_flag_start, PR_flag_end, DBP_flag_start, DBP_flag_end string

	// Title实际是两行
	for i, v := range title {
		if v == "阻止发起MR的人审批" {
			MR_flag_start = toChar(i) + "1"
		}
		if v == "发起MR时若有新commit被push，是否撤销批准" {
			MR_flag_end = toChar(i) + "1"
		}
		if v == "推送规则：未GPG签名的提交" {
			PR_flag_start = toChar(i) + "1"
		}
		if v == "推送规则：commit作者和push者需要是经过邮箱验证的gitlab用户" {
			PR_flag_end = toChar(i) + "1"
		}
		if v == "合规性" {
			DBP_flag_start = toChar(i) + "1"
		}
		if v == "不合规描述" {
			DBP_flag_end = toChar(i) + "1"
		}

		if v == "阻止发起MR的人审批" || v == "阻止push过commit的人审批" || v == "阻止MR中可以编辑审批规则" || v == "批准时需要密码" || v == "发起MR时若有新commit被push，是否撤销批准" {
			continue
		} else if v == "推送规则：未GPG签名的提交" || v == "推送规则：拒绝未验证邮箱的用户的push" || v == "推送规则：commit作者和push者需要是经过邮箱验证的gitlab用户" {
			continue
		} else if v == "合规性" || v == "不合规描述" {
			continue
		} else {
			project_title_merge_start := toChar(i) + strconv.Itoa(1)
			project_title_merge_end := toChar(i) + strconv.Itoa(2)
			err = f.MergeCell(sheetName_Project, project_title_merge_start, project_title_merge_end)
			if err != nil {
				fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
				return
			}
		}
	}
	// 合并title的推送规则和push rule和默认分支保护
	err = f.MergeCell(sheetName_Project, MR_flag_start, MR_flag_end)
	if err != nil {
		fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
		return
	}
	f.SetCellValue(sheetName_Project, MR_flag_start, "合并批准")
	err = f.MergeCell(sheetName_Project, PR_flag_start, PR_flag_end)
	if err != nil {
		fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
		return
	}
	f.SetCellValue(sheetName_Project, PR_flag_start, "推送规则")
	err = f.MergeCell(sheetName_Project, DBP_flag_start, DBP_flag_end)
	if err != nil {
		fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
		return
	}
	f.SetCellValue(sheetName_Project, DBP_flag_start, "默认分支保护")

	var single_project []interface{}
	for _, v := range o.Projects.Projects {

		row++
		name, err = excelize.JoinCellName("A", row)
		if err != nil {
			fmt.Println(err)
		}

		single_project = []interface{}{
			strconv.Itoa(v.Id),
			v.NameSpace,
			v.ProjectName,
			v.VisibilityNoPublic,
			v.TagMediaLinkAuth,
			v.SecurityAndCompliance,
			v.MergeRequestsAuthorApproval,
			v.MergeRequestsDisableCommittersApproval,
			v.DisableOverridingApproversPerMergeRequest,
			v.RequirePasswordToApprove,
			v.ResetApprovalsOnPush,
			v.MegerRequestApprovalsRulesRequireNumber,
			v.DefaultBranchProtected,
			v.DefaultBranchProtectedDescription,
			v.AccessTokenExpire,
			v.MergePipelinesEnabled,
			v.MergeTrainsEnabled,
			v.RejectUnsignCommit,
			v.RejectUnverifiedEmailPush,
			v.RejectCommitUnverifiedPush,
		}

		single_project = ConvertBool2Str(single_project)

		err = f.SetSheetRow(sheetName_Project, name, &single_project)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 美化外观
	headStyleProjectSheet, err := f.NewStyle(SetHeadProjectStyle())
	titleStyleProjectSheet, err := f.NewStyle(SetTitleStyle())
	textStyleProjectSheet, err := f.NewStyle(SetTextStyle())
	//uncomplainceStyle, err := f.NewStyle(SetUncomplainceStyle())
	if err != nil {
		fmt.Println(err)
	}
	f.SetColWidth(sheetName_Project, "A", "A", 8.4)
	f.SetColWidth(sheetName_Project, "B", "B", 22)
	f.SetColWidth(sheetName_Project, "C", "C", 25)
	f.SetColWidth(sheetName_Project, "D", "D", 11)
	f.SetColWidth(sheetName_Project, "E", "E", 11)
	f.SetColWidth(sheetName_Project, "F", "F", 11)
	f.SetColWidth(sheetName_Project, "G", "G", 11)
	f.SetColWidth(sheetName_Project, "H", "H", 11)
	f.SetColWidth(sheetName_Project, "I", "I", 11)
	f.SetColWidth(sheetName_Project, "J", "J", 11)
	f.SetColWidth(sheetName_Project, "K", "K", 11)
	f.SetColWidth(sheetName_Project, "L", "L", 11)
	f.SetColWidth(sheetName_Project, "M", "M", 11)
	f.SetColWidth(sheetName_Project, "N", "N", 11)
	f.SetColWidth(sheetName_Project, "O", "O", 11)
	f.SetColWidth(sheetName_Project, "P", "P", 11)
	f.SetColWidth(sheetName_Project, "Q", "Q", 11)
	f.SetColWidth(sheetName_Project, "R", "R", 11)
	f.SetColWidth(sheetName_Project, "S", "S", 11)
	f.SetColWidth(sheetName_Project, "T", "T", 11)
	f.SetColWidth(sheetName_Project, "U", "U", 11)
	f.SetColWidth(sheetName_Project, "V", "V", 11)
	f.SetColWidth(sheetName_Project, "W", "W", 11)
	f.SetColWidth(sheetName_Project, "X", "X", 11)
	f.SetCellStyle(sheetName_Project, "A1", "T2", headStyleProjectSheet)
	f.SetCellStyle(sheetName_Project, "A3", "A30", titleStyleProjectSheet)
	f.SetCellStyle(sheetName_Project, "B3", "T30", textStyleProjectSheet)

	// ----------------------Save to a xlsx----------------------------------

	if err := f.SaveAs("gitlab基线检查.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func ConvertBool2Str(single_project []interface{}) []interface{} {
	for k, v := range single_project {
		switch v.(type) {
		case bool:
			if v == true {
				single_project[k] = "合规"
			} else {
				single_project[k] = "不合规"
			}
		}
	}
	return single_project
}

func ConvertBool2StrIfComplaince(b bool) string {
	if b {
		return "合规"
	} else {
		return "不合规"
	}
}

func ConvertBool2StrIfEnable(b bool) string {
	if b {
		return "启用"
	} else {
		return "未启用"
	}
}

func toChar(i int) string {
	return string('A' + i)
}

func SetHeadStyle() *excelize.Style {
	return &excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "top",     // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "bottom",  // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "left",    // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "right",   // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
		},

		Fill: excelize.Fill{
			Type:    "pattern", // gradient 渐变色    pattern   填充图案
			Pattern: 1,         // 填充样式  当类型是 pattern 0-18 填充图案  1 实体填充
			// Color:   []string{"#FF0000"}, // 当Type = pattern 时，只有一个
			Color: []string{"#1E90FF"},
		},

		Font: &excelize.Font{
			Bold: true,
			// Italic: false,
			// Underline: "single",
			Size:   12,
			Family: "微软雅黑",
			// Strike:    true, // 删除线
			Color: "#FFFFFF",
		},

		Alignment: &excelize.Alignment{
			Horizontal: "center", // 水平对齐方式 center left right fill(填充) justify(两端对齐)  centerContinuous(跨列居中) distributed(分散对齐)
			Vertical:   "center", // 垂直对齐方式 center top  justify distributed
			// Indent:     1,        // 缩进  只要有值就变成了左对齐 + 缩进
			// TextRotation: 30, // 旋转
			// RelativeIndent:  10,   // 好像没啥用
			// ReadingOrder:    0,    // 不知道怎么设置
			// JustifyLastLine: true, // 两端分散对齐，只有 水平对齐 为 distributed 时 设置true 才有效
			WrapText: true, // 自动换行
			// ShrinkToFit:     true, // 缩小字体以填充单元格
		},

		Protection: &excelize.Protection{
			Hidden: true, // 貌似没啥用
			Locked: true, // 貌似没啥用
		},

		NumFmt:        0,       // 内置的数字格式样式   0-638  常用的 0-58  配合lang使用，因为语言不同样式不同 具体的样式参照文档
		Lang:          "zh-cn", // zh-cn 中文
		DecimalPlaces: 2,       // 小数位数  只有NumFmt是 2-11 有效
		// CustomNumFmt: "",// 自定义样式  是指针，只能通过变量的方式
		NegRed: true, // 不知道具体的含义
	}
}

func SetHeadProjectStyle() *excelize.Style {
	return &excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "top",     // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "bottom",  // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "left",    // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "right",   // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
		},

		Fill: excelize.Fill{
			Type:    "pattern", // gradient 渐变色    pattern   填充图案
			Pattern: 1,         // 填充样式  当类型是 pattern 0-18 填充图案  1 实体填充
			// Color:   []string{"#FF0000"}, // 当Type = pattern 时，只有一个
			Color: []string{"#1E90FF"},
		},

		Font: &excelize.Font{
			Bold: false,
			// Italic: false,
			// Underline: "single",
			Size:   12,
			Family: "微软雅黑",
			// Strike:    true, // 删除线
			Color: "#FFFFFF",
		},

		Alignment: &excelize.Alignment{
			Horizontal: "center", // 水平对齐方式 center left right fill(填充) justify(两端对齐)  centerContinuous(跨列居中) distributed(分散对齐)
			Vertical:   "center", // 垂直对齐方式 center top  justify distributed
			// Indent:     1,        // 缩进  只要有值就变成了左对齐 + 缩进
			// TextRotation: 30, // 旋转
			// RelativeIndent:  10,   // 好像没啥用
			// ReadingOrder:    0,    // 不知道怎么设置
			// JustifyLastLine: true, // 两端分散对齐，只有 水平对齐 为 distributed 时 设置true 才有效
			WrapText: false, // 自动换行
			// ShrinkToFit:     true, // 缩小字体以填充单元格
		},

		Protection: &excelize.Protection{
			Hidden: true, // 貌似没啥用
			Locked: true, // 貌似没啥用
		},

		NumFmt:        0,       // 内置的数字格式样式   0-638  常用的 0-58  配合lang使用，因为语言不同样式不同 具体的样式参照文档
		Lang:          "zh-cn", // zh-cn 中文
		DecimalPlaces: 2,       // 小数位数  只有NumFmt是 2-11 有效
		// CustomNumFmt: "",// 自定义样式  是指针，只能通过变量的方式
		NegRed: true, // 不知道具体的含义
	}
}

func SetTitleStyle() *excelize.Style {
	return &excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "top",     // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "bottom",  // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "left",    // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "right",   // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
		},

		Fill: excelize.Fill{
			Type:    "pattern", // gradient 渐变色    pattern   填充图案
			Pattern: 1,         // 填充样式  当类型是 pattern 0-18 填充图案  1 实体填充
			// Color:   []string{"#FF0000"}, // 当Type = pattern 时，只有一个
			Color: []string{"#C4DBB9"},
		},

		Font: &excelize.Font{
			Bold: false,
			// Italic: false,
			// Underline: "single",
			Size:   12,
			Family: "黑体",
			// Strike:    true, // 删除线
			Color: "#000000",
		},

		Alignment: &excelize.Alignment{
			Horizontal: "center", // 水平对齐方式 center left right fill(填充) justify(两端对齐)  centerContinuous(跨列居中) distributed(分散对齐)
			Vertical:   "center", // 垂直对齐方式 center top  justify distributed
			// Indent:     1,        // 缩进  只要有值就变成了左对齐 + 缩进
			// TextRotation: 30, // 旋转
			// RelativeIndent:  10,   // 好像没啥用
			// ReadingOrder:    0,    // 不知道怎么设置
			// JustifyLastLine: true, // 两端分散对齐，只有 水平对齐 为 distributed 时 设置true 才有效
			WrapText: false, // 自动换行
			// ShrinkToFit:     true, // 缩小字体以填充单元格
		},

		Protection: &excelize.Protection{
			Hidden: true, // 貌似没啥用
			Locked: true, // 貌似没啥用
		},

		NumFmt:        0,       // 内置的数字格式样式   0-638  常用的 0-58  配合lang使用，因为语言不同样式不同 具体的样式参照文档
		Lang:          "zh-cn", // zh-cn 中文
		DecimalPlaces: 2,       // 小数位数  只有NumFmt是 2-11 有效
		// CustomNumFmt: "",// 自定义样式  是指针，只能通过变量的方式
		NegRed: true, // 不知道具体的含义
	}
}

func SetTextStyle() *excelize.Style {
	return &excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "top",     // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "bottom",  // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "left",    // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "right",   // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
		},

		Fill: excelize.Fill{
			Type:    "pattern", // gradient 渐变色    pattern   填充图案
			Pattern: 1,         // 填充样式  当类型是 pattern 0-18 填充图案  1 实体填充
			// Color:   []string{"#FF0000"}, // 当Type = pattern 时，只有一个
			Color: []string{"#FFFFFF"},
		},

		Font: &excelize.Font{
			Bold: false,
			// Italic: false,
			// Underline: "single",
			Size:   12,
			Family: "黑体",
			// Strike:    true, // 删除线
			Color: "#000000",
		},

		Alignment: &excelize.Alignment{
			Horizontal: "left",   // 水平对齐方式 center left right fill(填充) justify(两端对齐)  centerContinuous(跨列居中) distributed(分散对齐)
			Vertical:   "center", // 垂直对齐方式 center top  justify distributed
			// Indent:     1,        // 缩进  只要有值就变成了左对齐 + 缩进
			// TextRotation: 30, // 旋转
			// RelativeIndent:  10,   // 好像没啥用
			// ReadingOrder:    0,    // 不知道怎么设置
			// JustifyLastLine: true, // 两端分散对齐，只有 水平对齐 为 distributed 时 设置true 才有效
			WrapText: false, // 自动换行
			// ShrinkToFit:     true, // 缩小字体以填充单元格
		},

		Protection: &excelize.Protection{
			Hidden: true, // 貌似没啥用
			Locked: true, // 貌似没啥用
		},

		NumFmt:        0,       // 内置的数字格式样式   0-638  常用的 0-58  配合lang使用，因为语言不同样式不同 具体的样式参照文档
		Lang:          "zh-cn", // zh-cn 中文
		DecimalPlaces: 2,       // 小数位数  只有NumFmt是 2-11 有效
		// CustomNumFmt: "",// 自定义样式  是指针，只能通过变量的方式
		NegRed: true, // 不知道具体的含义
	}
}

func SetUncomplainceStyle() *excelize.Style {
	return &excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "top",     // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "bottom",  // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "left",    // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
			{
				Type:  "right",   // top bottom left right diagonalDown diagonalUp 中的一个
				Color: "#000000", // 十六进制的颜色编码
				Style: 2,         // 0-13 有对应的样式
			},
		},

		Fill: excelize.Fill{
			Type:    "pattern", // gradient 渐变色    pattern   填充图案
			Pattern: 1,         // 填充样式  当类型是 pattern 0-18 填充图案  1 实体填充
			// Color:   []string{"#FF0000"}, // 当Type = pattern 时，只有一个
			Color: []string{"#FF0101"},
		},

		Font: &excelize.Font{
			Bold: false,
			// Italic: false,
			// Underline: "single",
			Size:   12,
			Family: "黑体",
			// Strike:    true, // 删除线
			Color: "#FFFFFF",
		},

		Alignment: &excelize.Alignment{
			Horizontal: "left",   // 水平对齐方式 center left right fill(填充) justify(两端对齐)  centerContinuous(跨列居中) distributed(分散对齐)
			Vertical:   "center", // 垂直对齐方式 center top  justify distributed
			// Indent:     1,        // 缩进  只要有值就变成了左对齐 + 缩进
			// TextRotation: 30, // 旋转
			// RelativeIndent:  10,   // 好像没啥用
			// ReadingOrder:    0,    // 不知道怎么设置
			// JustifyLastLine: true, // 两端分散对齐，只有 水平对齐 为 distributed 时 设置true 才有效
			WrapText: false, // 自动换行
			// ShrinkToFit:     true, // 缩小字体以填充单元格
		},

		Protection: &excelize.Protection{
			Hidden: true, // 貌似没啥用
			Locked: true, // 貌似没啥用
		},

		NumFmt:        0,       // 内置的数字格式样式   0-638  常用的 0-58  配合lang使用，因为语言不同样式不同 具体的样式参照文档
		Lang:          "zh-cn", // zh-cn 中文
		DecimalPlaces: 2,       // 小数位数  只有NumFmt是 2-11 有效
		// CustomNumFmt: "",// 自定义样式  是指针，只能通过变量的方式
		NegRed: true, // 不知道具体的含义
	}
}
