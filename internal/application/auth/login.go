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
	auth   portsClient.AuthClient
	parser portsParser.AuthParserService
	mapper portsMapper.AuthMapperService
}

func (l *AuthUseCaseImpl) Login(
	ctx context.Context,
	data *dto.AuthRequest,
) (*domain.AccountData, error) {
	resp, err := l.auth.Login(ctx, data)
	if err != nil {
		fmt.Errorf("can't authorize in account with creds", err)
	}

	authData, err := l.parser.DecodeAuthData(resp)
	if err != nil {
		return nil, fmt.Errorf("can't parse authdata:", err)
	}

	return l.mapper.MapAuthDataToDomain(authData)
}
