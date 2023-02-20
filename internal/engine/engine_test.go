package engine

import (
	"github.com/spf13/viper"
	"gitlab-misconfig/cmd"
	"gitlab-misconfig/internal/analyzer"
	"gitlab-misconfig/internal/analyzer/audit_event"
	"gitlab-misconfig/internal/analyzer/project"
	"gitlab-misconfig/internal/gitlab"
	"gitlab-misconfig/internal/types"
	"reflect"
	"testing"
)

func TestEngine_Analysis(t *testing.T) {
	type fields struct {
		Analyzers []analyzer.Analyzer
	}
	type args struct {
		gitlabClient *gitlab.Client
		options      *types.Options
	}
	gitlabClient := cmd.GitlabAuthClientInit()
	options := 
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "asd",
			fields: fields{
				Analyzers: []analyzer.Analyzer{
					project.New(),
					audit_event.New(),
				},
			},
			args: {
				gitlabClient: gitlabClient,
				options:      options,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				Analyzers: tt.fields.Analyzers,
			}
			e.Analysis(tt.args.gitlabClient, tt.args.options)
		})
	}
}

func TestNewEngine(t *testing.T) {
	tests := []struct {
		name string
		want *Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEngine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initConfig(t *testing.T) {
	type args struct {
		rulePath string
	}
	tests := []struct {
		name string
		args args
		want *viper.Viper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initConfig(tt.args.rulePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
