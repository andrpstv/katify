package amocrm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	auth "report/internal/authAmoCrm"
	"report/internal/models"
)

type AmoCRMService interface {
	GetCallsReport(ctx context.Context, filters *models.FilterDate) (*models.AmoCrmCalls, error)
	GetAccountsList(ctx context.Context) (AccountsList []models.AccountInfo, err error)
}

type amocrmService struct {
	client *auth.AmocrmClient
}

func NewAmocrmService(client *auth.AmocrmClient) *amocrmService {
	return &amocrmService{
		client: client,
	}
}

func (a *amocrmService) GetCallsReport(ctx context.Context, filters *models.FilterDate) (*models.AmoCrmCalls, error) {
	req = models.FiltersCalls{
		CallStatus: []int{1,2,3,4,5,6,7,8},
		CallType: []int{10, 11},
		Entity: []int{1, 12, 2, 3},
		Filter_date_preset: string,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, )


}

func (a *amocrmService) GetAccountsList(ctx context.Context) (AccountsList []models.AccountInfo, err error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.client.Cfg.AccountsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	// Вставляем cookies из клиента
	for _, c := range a.client.Cookies {
		req.AddCookie(c)
	}

	resp, err := a.client.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	type accountsResponse struct {
		Embedded struct {
			Items []models.AccountInfo `json:"items"`
		} `json:"_embedded"`
	}

	var result accountsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}

	return result.Embedded.Items, err
}
