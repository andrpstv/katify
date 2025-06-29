package amocrm

import "context"

type AmoCRMService interface {
	GetCallsReport(ctx context.Context, filters models.AmoCRMFilters) (models.CallsReport, error)
}
