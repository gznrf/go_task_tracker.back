package h_column

import (
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
)

type HColumn struct {
	service *service.Service
}

func NewHColumn(service *service.Service) *HColumn {
	return &HColumn{service: service}
}
