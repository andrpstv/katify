package adapters

import (
	"fmt"

	domain "report/internal/domain/auth"
	dto "report/internal/dto/auth"
)

type AmoAuthMapperServiceImpl struct{}

func MapAuthDataToDomain(dto *dto.AuthData) (*domain.AccountData, error) {
	if dto.AccessToken == "" {
		return nil, fmt.Errorf("missing access token")
	}

	return &domain.AccountData{
		AmoUserID:    dto.ID,
		Name:         dto.Name,
		Email:        dto.Email,
		AccessToken:  dto.AccessToken,
		RefreshToken: dto.RefreshToken,
		ExpiresAt:    dto.ExpiresAt,
	}, nil
}
