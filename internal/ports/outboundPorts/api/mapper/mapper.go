package ports

import (
	domain "katify/internal/domain/accounts"
	dto "katify/internal/dto/auth"
)

type AuthMapperService interface {
	MapAuthDataToDomain(
		*dto.AuthData,
	) (*domain.AccountData, error)
}

type AccountMapperService interface {
	MapAccountDataToDomain(
		*dto.AuthData,
	) (*domain.AccountData, error)
}
