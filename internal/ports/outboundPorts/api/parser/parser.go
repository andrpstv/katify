package ports

import (
	"net/http"

	domain "report/internal/domain/accounts"
	dto "report/internal/dto/auth"
)

type AuthParserService interface {
	ParseCSRF(html *http.Response) (string, error)
	DecodeAuthData(
		resp *http.Response,
	) (*dto.AuthData, error)
}
type AccountParserService interface {
	ParseAccounts(resp *http.Response) (*domain.Account, error)
}
