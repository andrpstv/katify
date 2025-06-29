type AmoCRMRepository interface {
    GetCallsCount(ctx context.Context, filters AmoFilters) (int, error)
    GetOutgoingMessagesCount(ctx context.Context, filters AmoFilters) (int, error)
}