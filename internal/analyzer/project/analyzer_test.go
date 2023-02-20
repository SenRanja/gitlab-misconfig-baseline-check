package project

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/types"
	"testing"
)

func TestAnalyzer_AutoAnalysis(t *testing.T) {
	type args struct {
		gitlabClient *gitlab.Client
		options      *types.Options
		config       *viper.Viper
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			an := Analyzer{}
			an.AutoAnalysis(tt.args.gitlabClient, tt.args.options, tt.args.config)
		})
	}
}
