package adapters

import (
	"fmt"

	domainAuth "report/internal/domain/AuthUseCase"
	domainAccount "report/internal/domain/accounts"
	dto "report/internal/dto/auth"
)

type AmoAuthMapperServiceImpl struct{}

func MapAuthDataToDomain(
	dto *dto.AuthData,
) (*domainAccount.Account, *domainAuth.AccountData, error) {
	if dto.AccessToken == "" {
		return nil, nil, fmt.Errorf("missing access token")
	}

	acc := &domainAccount.Account{
		Name:  dto.Name,
		Email: dto.Email,
	}

	accData := &domainAuth.AccountData{
		AmoUserID:    dto.ID,
		Name:         dto.Name,
		Email:        dto.Email,
		AccessToken:  dto.AccessToken,
		RefreshToken: dto.RefreshToken,
		ExpiresAt:    dto.ExpiresAt,
	}

	return acc, accData, nil
}
