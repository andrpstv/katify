package application

import (
	accountsClient "report/internal/ports/outboundPorts/api/services/amocrm/accounts"
	authClient "report/internal/ports/outboundPorts/api/services/amocrm/auth"
)

type AmoAccountServiceImpl struct {
	authAdapter     authClient.AuthClient
	accountsAdapter accountsClient.AmoAccountClient
}

func (a *AmoAccountServiceImpl) GetAccountProjects() {
	// TODO: нужно проверять залогинились ли/не истек ли токен
	a.accountsAdapter.FetchAccounts()
}
