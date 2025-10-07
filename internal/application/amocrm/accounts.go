package application

import (
	accountsPort "report/internal/ports/amocrm/accounts"
	authPort "report/internal/ports/amocrm/auth"
)

type AmoAccountServiceImpl struct {
	authAdapter     authPort.AmoAuthClient
	accountsAdapter accountsPort.AmoAccountClient
}

func (a *AmoAccountService) GetAccountProjects() {
	// TODO: нужно проверять залогинились ли/не истек ли токен
}
