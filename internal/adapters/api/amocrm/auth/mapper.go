package adapters

import (
	domain "report/internal/domain/auth"
	dto "report/internal/dto/auth"
)

type AuthMapperServiceImpl struct{}

func MapAuthDataToDomain(*dto.AuthData) (*domain.AccountData, error) {
}
