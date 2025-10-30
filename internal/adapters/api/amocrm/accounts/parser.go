package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	domain "katify/internal/domain/accounts"
	dto "katify/internal/dto/accounts"
)

type AccountParserServiceImpl struct{}

func (a *AccountParserServiceImpl) ParseAccounts(resp *http.Response) (*domain.Account, error) {
	var dto dto.AccountsResponse
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(dto.Embedded.Items) == 0 {
		return nil, fmt.Errorf("no accounts found")
	}

	var projects []domain.Project
	for _, item := range dto.Embedded.Items {
		projects = append(projects, domain.Project{
			ID:         item.ID,
			UUID:       item.UUID,
			Name:       item.Name,
			Subdomain:  item.Subdomain,
			Domain:     item.Domain,
			IsAdmin:    item.IsAdmin,
			IsKommo:    item.IsKommo,
			MFAEnabled: item.MFAEnabled,
		})
	}

	result := &domain.Account{
		Projects: projects,
	}

	return result, nil
}
