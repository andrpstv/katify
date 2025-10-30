package application

import (
	accountsClient "katify/internal/ports/outboundPorts/api/services/amocrm/accounts"
	authClient "katify/internal/ports/outboundPorts/api/services/amocrm/auth"
)

type AmoAccountServiceImpl struct {
	authAdapter     authClient.AuthClient
	accountsAdapter accountsClient.AmoAccountClient
}

func (a *AmoAccountServiceImpl) GetAccountProjects() {
	// TODO: нужно проверять залогинились ли/не истек ли токен
	a.accountsAdapter.FetchAccounts()
}
