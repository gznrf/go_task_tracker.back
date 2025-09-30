package service

import "github.com/gznrf/go_task_tracker.back.git/pkg/repo"

type Middleware interface {
}

type Service struct {
}

func NewService(repos *repo.Repo) *Service {
	return &Service{}
}
