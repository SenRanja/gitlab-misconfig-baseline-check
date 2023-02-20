package project

import (
	"gitlab-misconfig/internal/gitlab"
	"reflect"
	"testing"
)

func TestListAllProjects(t *testing.T) {
	type args struct {
		gitlabClient      *gitlab.Client
		per_page_items    int
		max_acquire_items int
	}
	tests := []struct {
		name  string
		args  args
		want  []*gitlab.Project
		want1 *gitlab.ProjectsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				got, got1 := ListAllProjects(tt.args.gitlabClient, tt.args.per_page_items, tt.args.max_acquire_items)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ListAllProjects() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					t.Errorf("ListAllProjects() got1 = %v, want %v", got1, tt.want1)
				}
			})
	}
}

func TestProjectApprovalsBeforeMerge(t *testing.T) {
	type args struct {
		p *gitlab.Project
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProjectApprovalsBeforeMerge(tt.args.p); got != tt.want {
				t.Errorf("ProjectApprovalsBeforeMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectCommitCommitterCheck(t *testing.T) {
	type args struct {
		gitlabClientProjectService *gitlab.ProjectsService
		pid                        int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProjectCommitCommitterCheck(tt.args.gitlabClientProjectService, tt.args.pid); got != tt.want {
				t.Errorf("ProjectCommitCommitterCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectRejectUnsignedCommits(t *testing.T) {
	type args struct {
		gitlabClientProjectService *gitlab.ProjectsService
		pid                        int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProjectRejectUnsignedCommits(tt.args.gitlabClientProjectService, tt.args.pid); got != tt.want {
				t.Errorf("ProjectRejectUnsignedCommits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectSecurityAndCompliance(t *testing.T) {
	type args struct {
		p *gitlab.Project
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProjectSecurityAndCompliance(tt.args.p); got != tt.want {
				t.Errorf("ProjectSecurityAndCompliance() = %v, want %v", got, tt.want)
			}
		})
	}
}
