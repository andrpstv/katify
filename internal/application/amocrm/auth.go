package application

import (
	"context"
	"fmt"

	dto "report/internal/dto/auth"
	ports "report/internal/ports/outboundPorts/api/services/amocrm/auth"
)

type AmoAuthServiceImpl struct {
	portsAuth ports.AuthClient
}

func (a *AmoAuthServiceImpl) Login(ctx context.Context, data *dto.AuthRequest) {
	_, err := a.portsAuth.GetCSRFtoken(ctx)
	if err != nil {
		fmt.Errorf("didn't find csrf token:", err)
	}
	authData, err := a.portsAuth.Login(ctx, data)
	if err != nil {
		fmt.Errorf("can't login with csrf token", err)
	}
}
