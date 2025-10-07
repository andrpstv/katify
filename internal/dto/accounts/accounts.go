package dto

type AccountsResponse struct {
	Links    LinksSection    `json:"_links"`
	Embedded EmbeddedSection `json:"_embedded"`
}

type LinksSection struct {
	Self Link `json:"self"`
}

type Link struct {
	Href   string `json:"href"`
	Method string `json:"method"`
}

type EmbeddedSection struct {
	Items []AccountDTO `json:"items"`
}

type AccountDTO struct {
	ID         int          `json:"id"`
	UUID       string       `json:"uuid"`
	Name       string       `json:"name"`
	Subdomain  string       `json:"subdomain"`
	ShardType  int          `json:"shard_type"`
	IsAdmin    bool         `json:"is_admin"`
	Version    int          `json:"account_version"`
	IsKommo    bool         `json:"is_kommo"`
	Domain     string       `json:"domain"`
	IsTrial    bool         `json:"is_trial"`
	TrialEnded bool         `json:"is_trial_expired"`
	IsPayed    bool         `json:"is_payed"`
	PayedEnded bool         `json:"is_payed_expired"`
	MFAEnabled bool         `json:"is_mandatory_mfa_enabled"`
	Links      LinksSection `json:"_links"`
}
