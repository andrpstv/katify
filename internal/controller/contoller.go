package controller

import (
	"report/internal/logger"
	"report/internal/report"
)

type Contoller struct {
	log           *logger.Logger
	reportService *report.ReportService
}

