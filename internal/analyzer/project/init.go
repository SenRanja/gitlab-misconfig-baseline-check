package project

import (
	"fmt"
	"gitlab-misconfig/internal/gitlab"
)

type Analyzer struct {
}

// 列出全部project
func ListAllProjects(gitlabClient *gitlab.Client) ([]*gitlab.Project, *gitlab.ProjectsService, *gitlab.MergeRequestApprovalsService, *gitlab.ProjectMembersService, *gitlab.ProtectedBranchesService) {

	var opt gitlab.ListOptions
	var listProjectsOptions *gitlab.ListProjectsOptions
	var projects []*gitlab.Project

	per_page_items := 2000
	i := 1
	for {
		opt = gitlab.ListOptions{
			Page:    i,
			PerPage: per_page_items,
		}
		i++
		listProjectsOptions = &gitlab.ListProjectsOptions{
			ListOptions: opt,
		}
		projects_tmp, _, err := gitlabClient.Projects.ListProjects(listProjectsOptions)
		if err != nil {
			fmt.Println(err)
		}
		projects = append(projects, projects_tmp...)

		if len(projects_tmp) < per_page_items {
			break
		}
	}

	projectsService := gitlab.ProjectsService{
		Client: gitlabClient,
	}

	approvalService := gitlab.MergeRequestApprovalsService{
		Client: gitlabClient,
	}

	projectMembersService := gitlab.ProjectMembersService{
		Client: gitlabClient,
	}

	protectedBranchesService := gitlab.ProtectedBranchesService{
		Client: gitlabClient,
	}

	return projects, &projectsService, &approvalService, &projectMembersService, &protectedBranchesService
}
