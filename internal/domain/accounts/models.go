package domain

type Account struct {
	ID       int
	Name     string
	Projects []Project
	Data     AccountData
}

type Project struct {
	ID         int
	UUID       string
	Name       string
	Subdomain  string
	ShardType  int
	IsAdmin    bool
	Version    int
	IsKommo    bool
	Domain     string
	IsTrial    bool
	TrialEnded bool
	IsPayed    bool
	PayedEnded bool
	MFAEnabled bool
}
