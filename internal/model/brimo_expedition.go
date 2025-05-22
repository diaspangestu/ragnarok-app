package model

type BrimoExpedition struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	PostalCode string `gorm:"column:postal_code" json:"postalCode"`
	TimeZone   string `gorm:"column:timezone" json:"timeZone"`
	Province   string `gorm:"column:province" json:"province"`
	City       string `gorm:"column:city" json:"city"`
	Expedition string `gorm:"column:expedition" json:"expedition"`
	Sla        int    `gorm:"column:sla" json:"sla"`
}
