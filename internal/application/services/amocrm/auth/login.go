package application

import (
	"context"
	"fmt"

	domain "katify/internal/domain/AuthUseCase"
	dto "katify/internal/dto/auth"
	portsMapper "katify/internal/ports/outboundPorts/api/mapper"
	portsParser "katify/internal/ports/outboundPorts/api/parser"
	portsClient "katify/internal/ports/outboundPorts/api/services/amocrm/auth"
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

	return a.mapper.MapAuthDataToDomain(authData)
}
