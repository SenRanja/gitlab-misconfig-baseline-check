package engine

import (
	"bytes"
	"github.com/spf13/viper"
	"gitlab-misconfig/bindata"
	"gitlab-misconfig/internal/analyzer"
	"gitlab-misconfig/internal/analyzer/project"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/log"
	"gitlab-misconfig/internal/types"
)

type Engine struct {
	Analyzers []analyzer.Analyzer
}

func NewEngine() *Engine {
	return &Engine{
		Analyzers: []analyzer.Analyzer{
			//user.New(),
			//project.New(),
			new(project.Analyzer),
			//audit_event.New(),
			//settings.New(),
		},
	}
}

func (e *Engine) Analysis(gitlabClient *gitlab.Client, options *types.Options) {
	// 加载规则
	config := initConfig(options.RulePath)
	// 扫描逻辑
	for _, analyzer := range e.Analyzers {
		analyzer.AutoAnalysis(gitlabClient, options, config)
	}
}

func initConfig(rulePath string) *viper.Viper {
	// 设定值，如果用户没有指定，就使用内置的base.toml的配置
	config := viper.New()
	if rulePath != "" {
		config.SetConfigFile(rulePath)
		config.SetConfigType("toml")
		log.Info("using rules is " + rulePath)
		if err := config.ReadInConfig(); err != nil {
			log.Error("unable to load  config")
			log.Error(err)
		}
	} else {
		bindataDefaultToml, _ := bindata.Asset("base.toml")
		config.SetConfigType("toml")
		if err := config.ReadConfig(bytes.NewBuffer(bindataDefaultToml)); err != nil {
			log.Error("unable to load config")
			log.Error(err)
		}
	}
	return config
}
