package types

import (
	"bytes"
	"github.com/spf13/viper"
	"gitlab-misconfig/bindata"
)

type Options struct {
	Output    string
	Version   string
	ProjectId string
	Mode      string
	Token     string
	Url       string
	Check     string
	Log       string
	RulePath  string
}

type Version struct {
	Version          string
	Revision         string
	VersionIsEE      bool
	VersionExistRisk bool
	CheckRule        string `mapstructure:"check_rule" json:"check_rule"`
	SecondCheckRule  string `mapstructure:"second_check_rule" json:"second_check_rule"`
	Result           string `mapstructure:"result" json:"result"`
	Keyword          string `mapstructure:"keyword" json:"keyword"`
	Complaince       bool   `mapstructure:"complaince" json:"complaince"`
	Description      string `mapstructure:"description" json:"description"`
	Advice           string `mapstructure:"advice" json:"advice"`
}

type Output struct {
	Version     Version     `mapstructure:"version" json:"version"`
	Projects    Projects    `mapstructure:"projects" json:"projects"`
	Settings    Settings    `mapstructure:"settings" json:"settings"`
	User        User        `mapstructure:"user" json:"user"`
	OutputTitle OutputTitle `mapstructure:"output_title" json:"output_title"`
}

type Projects struct {
	Projects []Project `mapstructure:"project" json:"project"`
}

type Project struct {
	Id                                        int
	NameSpace                                 string
	ProjectName                               string
	VisibilityNoPublic                        bool
	TagMediaLinkAuth                          bool
	SecurityAndCompliance                     bool
	MergeRequestsAuthorApproval               bool
	MergeRequestsDisableCommittersApproval    bool
	DisableOverridingApproversPerMergeRequest bool
	RequirePasswordToApprove                  bool
	ResetApprovalsOnPush                      bool
	MegerRequestApprovalsRulesRequireNumber   bool
	MergePipelinesEnabled                     bool
	MergeTrainsEnabled                        bool
	DefaultBranchProtected                    bool
	DefaultBranchProtectedDescription         string
	RejectUnsignCommit                        bool
	RejectUnverifiedEmailPush                 bool
	RejectCommitUnverifiedPush                bool
	AccessTokenExpire                         string
	CICD                                      bool
	CICDYamlStageNum                          int
	RunnersDescription                        string
}

type Runner struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Paused      bool   `json:"paused"`
	IsShared    bool   `json:"is_shared"`
	IPAddress   string `json:"ip_address"`
	RunnerType  string `json:"runner_type"`
	Name        string `json:"name"`
	Online      bool   `json:"online"`
	Status      string `json:"status"`
}

type OutputTitle struct {
	CheckRule       string `mapstructure:"check_rule" json:"check_rule"`
	SecondCheckRule string `mapstructure:"second_check_rule" json:"second_check_rule"`
	Result          string `mapstructure:"result" json:"result"`
	Keyword         string `mapstructure:"keyword" json:"keyword"`
	Complaince      string `mapstructure:"complaince" json:"complaince"`
	Description     string `mapstructure:"description" json:"description"`
	Advice          string `mapstructure:"advice" json:"advice"`
}

type User struct {
	Inactive struct {
		CheckRule   string `mapstructure:"check_rule" json:"check_rule"`
		Result      int    `mapstructure:"result" json:"result"`
		Keyword     int    `mapstructure:"keyword" json:"keyword"`
		Complaince  bool   `mapstructure:"complaince" json:"complaince"`
		Description string `mapstructure:"description" json:"description"`
		Advice      string `mapstructure:"advice" json:"advice"`
	} `mapstructure:"inactive"`
	Unactive struct {
		CheckRule   string `mapstructure:"check_rule" json:"check_rule"`
		Result      int    `mapstructure:"result" json:"result"`
		Keyword     int    `mapstructure:"keyword" json:"keyword"`
		Complaince  bool   `mapstructure:"complaince" json:"complaince"`
		Description string `mapstructure:"description" json:"description"`
		Advice      string `mapstructure:"advice" json:"advice"`
	} `mapstructure:"unactive"`
	TwoFactorAuth struct {
		CheckRule   string `mapstructure:"check_rule"`
		Result      int
		Keyword     int
		Complaince  bool
		Description string `mapstructure:"description"`
		Advice      string `mapstructure:"advice"`
	} `mapstructure:"two_factor_auth"`
	Admin struct {
		CheckRule   string `mapstructure:"check_rule"`
		Result      int
		Keyword     int `mapstructure:"check_rule"`
		Complaince  bool
		Description string `mapstructure:"description"`
		Advice      string `mapstructure:"advice"`
	} `mapstructure:"admin"`
	Auditor struct {
		CheckRule   string `mapstructure:"check_rule"`
		Result      int
		Keyword     int `mapstructure:"check_rule"`
		Complaince  bool
		Description string `mapstructure:"description"`
		Advice      string `mapstructure:"advice"`
	} `mapstructure:"auditor"`
	External struct {
		CheckRule   string `mapstructure:"check_rule"`
		Result      int
		Keyword     int `mapstructure:"check_rule"`
		Complaince  bool
		Description string `mapstructure:"description"`
		Advice      string `mapstructure:"advice"`
	} `mapstructure:"external"`
}

type Settings struct {
	Password struct {
		CheckRule   string `mapstructure:"check_rule"`
		Result      bool
		Keyword     int `mapstructure:"keyword"`
		Complaince  bool
		Description string `mapstructure:"description"`
		Advice      string `mapstructure:"advice"`
		Length      struct {
			CheckRule   string `mapstructure:"check_rule"`
			Result      int
			Keyword     int `mapstructure:"keyword"`
			Complaince  bool
			Description string `mapstructure:"description"`
			Advice      string `mapstructure:"advice"`
		} `mapstructure:"length"`
		Num struct {
			CheckRule   string `mapstructure:"check_rule"`
			Result      bool
			Keyword     bool `mapstructure:"keyword"`
			Complaince  bool
			Description string `mapstructure:"description"`
			Advice      string `mapstructure:"advice"`
		} `mapstructure:"num"`
		Upper struct {
			CheckRule   string `mapstructure:"check_rule"`
			Result      bool
			Keyword     bool `mapstructure:"keyword"`
			Complaince  bool
			Description string `mapstructure:"description"`
			Advice      string `mapstructure:"advice"`
		} `mapstructure:"upper"`
		Lower struct {
			CheckRule   string `mapstructure:"check_rule"`
			Result      bool
			Keyword     bool `mapstructure:"keyword"`
			Complaince  bool
			Description string `mapstructure:"description"`
			Advice      string `mapstructure:"advice"`
		} `mapstructure:"lower"`
		Special struct {
			CheckRule   string `mapstructure:"check_rule"`
			Result      bool
			Keyword     bool `mapstructure:"keyword"`
			Complaince  bool
			Description string `mapstructure:"description"`
			Advice      string `mapstructure:"advice"`
		} `mapstructure:"special"`
	} `mapstructure:"password"`
	TwoFactorAuth struct {
		CheckRule string `mapstructure:"check_rule"`
		Result    bool
		//Keyword     bool `mapstructure:"keyword"`
		Complaince  bool
		Description string `mapstructure:"description"`
		Advice      string `mapstructure:"advice"`
	} `mapstructure:"two_factor_auth"`
	Register struct {
		RegisterEnable struct {
			CheckRule   string `mapstructure:"check_rule"`
			Result      bool
			Keyword     bool `mapstructure:"keyword"`
			Complaince  bool
			Description string `mapstructure:"description"`
			Advice      string `mapstructure:"advice"`
		} `mapstructure:"register_enable"`
		EmailConfirmation struct {
			CheckRule   string `mapstructure:"check_rule"`
			Result      bool
			Keyword     bool `mapstructure:"keyword"`
			Complaince  bool
			Description string `mapstructure:"description"`
			Advice      string `mapstructure:"advice"`
		} `mapstructure:"email_confirmation"`
		AdminApproval struct {
			CheckRule   string `mapstructure:"check_rule"`
			Result      bool
			Keyword     bool `mapstructure:"keyword"`
			Complaince  bool
			Description string `mapstructure:"description"`
			Advice      string `mapstructure:"advice"`
		} `mapstructure:"admin_approval"`
		External struct {
			CheckRule   string `mapstructure:"check_rule"`
			Result      bool
			Keyword     bool `mapstructure:"keyword"`
			Complaince  bool
			Description string `mapstructure:"description"`
			Advice      string `mapstructure:"advice"`
		} `mapstructure:"external"`
		EmailRegexpInternal struct {
			CheckRule   string `mapstructure:"check_rule"`
			Result      bool
			Keyword     bool `mapstructure:"keyword"`
			Complaince  bool
			Description string `mapstructure:"description"`
			Advice      string `mapstructure:"advice"`
		} `mapstructure:"email_regexp_internal"`
	} `mapstructure:"register"`
	InitProjectEnableDefaultBranchProtection struct {
		CheckRule   string `mapstructure:"check_rule"`
		Result      string
		Keyword     string `mapstructure:"keyword"`
		Complaince  bool
		Description string `mapstructure:"description"`
		Advice      string `mapstructure:"advice"`
	} `mapstructure:"init_project_enable_default_branch_protection"`
	ForbidCreatePublicProject struct {
		CheckRule   string `mapstructure:"check_rule"`
		Result      string
		Keyword     string `mapstructure:"keyword"`
		Complaince  bool
		Description string `mapstructure:"description"`
		Advice      string `mapstructure:"advice"`
	} `mapstructure:"forbid_create_public_project"`
	WebSessionExpire struct {
		CheckRule   string `mapstructure:"check_rule"`
		Result      int
		Keyword     int
		Complaince  bool
		Description string `mapstructure:"description"`
		Advice      string `mapstructure:"advice"`
	} `mapstructure:"web_session_expire"`
	PATExpire struct {
		CheckRule   string `mapstructure:"check_rule"`
		Result      string
		Keyword     int
		Complaince  bool
		Description string `mapstructure:"description"`
		Advice      string `mapstructure:"advice"`
	} `mapstructure:"pat_expire"`
}

func (o *Output) GetDefault() {
	outputConfig := viper.New()
	bindataOutputDefaultToml, _ := bindata.Asset("output.toml")
	outputConfig.SetConfigType("toml")
	if err := outputConfig.ReadConfig(bytes.NewBuffer(bindataOutputDefaultToml)); err != nil {
		panic("unable to load output config")
		panic(err)
	}
	outputConfig.Unmarshal(o)
}
