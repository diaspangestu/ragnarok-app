package dto

type BrimoExpeditionRequest struct {
	PostalCode string `form:"postalCode"`
	City       string `form:"city"`
	Page       int    `form:"page"`
}

type BrimoExpeditionResponse struct {
	Id         int    `json:"id"`
	PostalCode string `json:"postalCode"`
	TimeZone   string `json:"timeZone"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Expedition string `json:"expedition"`
	Sla        int    `json:"sla"`
}
