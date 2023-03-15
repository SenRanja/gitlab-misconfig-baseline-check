package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gitlab-misconfig/internal/types"
	"strconv"
	"time"
)

func ExportExcelFromCE(o *types.Output) {

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

	// 版本风险检测
	row++
	versionCheckName, err := excelize.JoinCellName("A", row)
	versionCheck := []interface{}{
		o.Version.CheckRule,
		o.Version.SecondCheckRule,
		o.Version.Result,
		"",
		ConvertBool2StrIfComplaince(o.Version.Complaince),
		o.Version.Description,
		o.Version.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, versionCheckName, &versionCheck)
	if err != nil {
		fmt.Println(err)
	}

	// 密码复杂度-最小长度
	row++
	password_len_name, err := excelize.JoinCellName("A", row)
	password_len := []interface{}{
		o.Settings.Password.CheckRule,
		o.Settings.Password.Length.CheckRule,
		o.Settings.Password.Length.Result,
		o.Settings.Password.Length.Keyword,
		ConvertBool2StrIfComplaince(o.Settings.Password.Length.Complaince),
		o.Settings.Password.Length.Description,
		o.Settings.Password.Length.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, password_len_name, &password_len)
	if err != nil {
		fmt.Println(err)
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
		ConvertBool2StrIfEnable(!o.Settings.Register.RegisterEnable.Result),
		ConvertBool2StrIfEnable(o.Settings.Register.RegisterEnable.Keyword),
		ConvertRegisterEnable2ForbidenRegisterBool(o.Settings.Register.RegisterEnable.Complaince),
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
		RegisterEnableCheckComplainceOutput(o, o.Settings.Register.EmailConfirmation.Complaince),
		o.Settings.Register.EmailConfirmation.Description,
		o.Settings.Register.EmailConfirmation.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, email_confirmation_name, &email_confirmation)
	if err != nil {
		fmt.Println(err)
	}

	// 注册-新用户注册后需等待admin批准
	row++
	register_merge_end := "A" + strconv.Itoa(row)
	admin_approval_name, err := excelize.JoinCellName("A", row)
	admin_approval := []interface{}{
		"",
		o.Settings.Register.AdminApproval.CheckRule,
		ConvertBool2StrIfEnable(!o.Settings.Register.RegisterEnable.Result),
		ConvertBool2StrIfEnable(o.Settings.Register.RegisterEnable.Keyword),
		RegisterEnableCheckComplainceOutput(o, o.Settings.Register.RegisterEnable.Complaince),
		o.Settings.Register.AdminApproval.Description,
		o.Settings.Register.AdminApproval.Advice,
	}
	err = f.SetSheetRow(sheetName_SettingsAndUser, admin_approval_name, &admin_approval)
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
	var unComplainceCellLocation string
	headStyle, err := f.NewStyle(SetHeadStyle())
	titleStyle, err := f.NewStyle(SetTitleStyle())
	textStyle, err := f.NewStyle(SetTextStyle())
	uncomplainceStyle, err := f.NewStyle(SetUncomplainceStyle())
	//uncomplainceStyle, err := f.NewStyle(SetUncomplainceStyle())
	if err != nil {
		fmt.Println(err)
	}

	// 自动找表格的最大行数  SettingsSheetRow
	SettingsSheetRow := 1
	var GetSettingsValueCellLocation, tmp_v string
	for {
		GetSettingsValueCellLocation = "A" + strconv.Itoa(SettingsSheetRow+1)
		tmp_v, err = f.GetCellValue(sheetName_SettingsAndUser, GetSettingsValueCellLocation)
		if tmp_v == "" {
			break
		}
		SettingsSheetRow++
	}

	AEnd := "A" + strconv.Itoa(SettingsSheetRow)
	GEnd := "G" + strconv.Itoa(SettingsSheetRow)

	f.SetColWidth(sheetName_SettingsAndUser, "A", "A", 29.07)
	f.SetColWidth(sheetName_SettingsAndUser, "B", "B", 27.67)
	f.SetColWidth(sheetName_SettingsAndUser, "C", "C", 9.07)
	f.SetColWidth(sheetName_SettingsAndUser, "D", "D", 21.33)
	f.SetColWidth(sheetName_SettingsAndUser, "E", "E", 9)
	f.SetColWidth(sheetName_SettingsAndUser, "F", "F", 47)
	f.SetColWidth(sheetName_SettingsAndUser, "G", "G", 37)
	f.SetCellStyle(sheetName_SettingsAndUser, "A1", "G1", headStyle)
	f.SetCellStyle(sheetName_SettingsAndUser, "A2", AEnd, titleStyle)
	f.SetCellStyle(sheetName_SettingsAndUser, "B2", GEnd, textStyle)
	// 获取不合规的单元格并给予醒目颜色
	rows, err := f.GetRows(sheetName_SettingsAndUser)
	for row_id, row := range rows {
		for col_id, colCell := range row {
			if colCell == "不合规" {
				// (row_id+1, col_id+1) 即 (toChar(row_id+1), col_id+1)
				unComplainceCellLocation = toChar(col_id) + strconv.Itoa(row_id+1)
				f.SetCellStyle(sheetName_SettingsAndUser, unComplainceCellLocation, unComplainceCellLocation, uncomplainceStyle)
			}
		}
	}

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
		"合规性",
		"不合规描述",
		"项目访问令牌的有效期是否过长",
		"CICD管道",
		"共享runners",
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

	//var MR_flag_start, MR_flag_end, PR_flag_start, PR_flag_end, DBP_flag_start, DBP_flag_end string
	var DBP_flag_start, DBP_flag_end string

	// Title实际是两行
	for i, v := range title {
		if v == "合规性" {
			DBP_flag_start = toChar(i) + "1"
		}
		if v == "不合规描述" {
			DBP_flag_end = toChar(i) + "1"
		}

		if v == "合规性" || v == "不合规描述" {
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

	f.SetCellValue(sheetName_Project, DBP_flag_start, "默认分支保护")
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
			v.DefaultBranchProtected,
			v.DefaultBranchProtectedDescription,
			v.AccessTokenExpire,
			ConvertBool2StrIfEnable(v.CICD),
			v.RunnersDescription,
		}

		single_project = ConvertBool2Str(single_project)

		err = f.SetSheetRow(sheetName_Project, name, &single_project)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 美化外观
	//获取项目共多少行，获取有内容的行的行数
	ProjectSheetRow := 1
	var GetValueCellLocation string
	for {
		GetValueCellLocation = "A" + strconv.Itoa(ProjectSheetRow+1)
		tmp_v, err = f.GetCellValue(sheetName_Project, GetValueCellLocation)
		if tmp_v == "" {
			break
		}
		ProjectSheetRow++
	}

	headStyleProjectSheet, err := f.NewStyle(SetHeadProjectStyle())
	titleStyleProjectSheet, err := f.NewStyle(SetTitleStyle())
	textStyleProjectSheet, err := f.NewStyle(SetTextStyle())
	//uncomplainceStyle, err := f.NewStyle(SetUncomplainceStyle())
	if err != nil {
		fmt.Println(err)
	}
	titleEndCellLocation := "A" + strconv.Itoa(ProjectSheetRow)
	textEndCellLocation := "V" + strconv.Itoa(ProjectSheetRow)
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
	f.SetCellStyle(sheetName_Project, "A1", "V2", headStyleProjectSheet)
	f.SetCellStyle(sheetName_Project, "A3", titleEndCellLocation, titleStyleProjectSheet)
	f.SetCellStyle(sheetName_Project, "B3", textEndCellLocation, textStyleProjectSheet)
	// 获取不合规的单元格并给予醒目颜色
	rows, err = f.GetRows(sheetName_Project)
	for row_id, row := range rows {
		for col_id, colCell := range row {
			if colCell == "不合规" {
				// (row_id+1, col_id+1) 即 (toChar(row_id+1), col_id+1)
				unComplainceCellLocation = toChar(col_id) + strconv.Itoa(row_id+1)
				f.SetCellStyle(sheetName_Project, unComplainceCellLocation, unComplainceCellLocation, uncomplainceStyle)
			}
		}
	}

	// ----------------------Save to a xlsx----------------------------------
	currentTime := time.Now().Format("2006-01-02-15-04-05")

	xlsxFileName := "gitlab基线检查-" + currentTime + ".xlsx"
	if err := f.SaveAs(xlsxFileName); err != nil {
		fmt.Println(err)
	}
}

// 如果注册功能关闭，则注册的几项的合规性不输出
// 如果注册功能开启，则输出
func RegisterEnableCheckComplainceOutput(o *types.Output, s bool) string {
	if o.Settings.Register.RegisterEnable.Result {
		return ConvertBool2StrIfComplaince(s)
	} else {
		return ""
	}
}
