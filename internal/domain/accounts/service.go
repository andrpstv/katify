package domain

type AccountsService struct{}

func NewAccountsService() *AccountsService {
	return &AccountsService{}
}

func (s *AccountsService) GetAccountsList(raw []Account) ([]Account, error) {
}
