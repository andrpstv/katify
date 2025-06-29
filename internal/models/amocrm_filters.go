package models

type FilterDate struct {
	FilterDateFrom string `json:"filter[date_from]"`
	FilterDateTo   string `json:"filter[date_to]"`
}

type FiltersCalls struct {
	CallStatus         []int  `json:"callstatus"`
	CallType           []int  `json:"calltype"`
	Entity             []int  `json:"entity"`
	Filter_date_preset string `json:"filter[date_preset]"`
	FilterDate         *FilterDate
	Filter_main_user   int    `json:"filter[main_user]"`
	UseFilter          string `json:"useFilter"`
}

type AmoCrmCalls struct {
	CallsCount int `json:"CallsCount"`
	Inbound    int `json:"Inbound"`
	Outbound   int `json:"Outbound"`
}
