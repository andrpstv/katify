package application

import (
	"context"
	"fmt"

	domain "report/internal/domain/auth"
	dto "report/internal/dto/auth"
	portsMapper "report/internal/ports/outboundPorts/api/mapper"
	portsParser "report/internal/ports/outboundPorts/api/parser"
	portsClient "report/internal/ports/outboundPorts/api/services/amocrm/auth"
)

type AuthUseCaseImpl struct {
	authClient portsClient.AuthClient
	parser     portsParser.AuthParserService
	mapper     portsMapper.AuthMapperService
}

func (a *AuthUseCaseImpl) Login(
	ctx context.Context,
	data *dto.AuthRequest,
) (*domain.AccountData, error) {
	resp, err := a.authClient.Login(ctx, data)
	if err != nil {
		fmt.Errorf("can't authorize in account with creds", err)
	}

	authData, err := a.parser.DecodeAuthData(resp)
	if err != nil {
		return nil, fmt.Errorf("can't parse authdata:", err)
	}

	// дополнить запись в бд
	return a.mapper.MapAuthDataToDomain(authData)
}
