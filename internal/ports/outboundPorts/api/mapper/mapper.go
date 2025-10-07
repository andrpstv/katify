package ports

import (
	domain "report/internal/domain/auth"
	dto "report/internal/dto/auth"
)

type AuthMapperService interface {
	MapAuthDataToDomain(
		*dto.AuthData,
	) (*domain.AccountData, error)
}

type AccountMapperService interface {
	MapAccountDataToDomain(
		*dto.AuthData,
	) (*domain.AccountInfo, error)
}
