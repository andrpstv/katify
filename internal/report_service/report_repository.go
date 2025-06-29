package report_service

import "context"

type AmoCRMRepository interface {
	GetCallsCount(ctx context.Context, filters AmoFilters) (int, error)
	GetOutgoingMessagesCount(ctx context.Context, filters AmoFilters) (int, error)
}
