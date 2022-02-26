package model

type MerchantResponse struct {
	Name  string `json:"merchant_name" gorm:"column:merchant_name"`
	Omzet int    `json:"omzet" gorm:"column:omzet"`
	Date  string `json:"date" gorm:"column:date"`
}

type OutletResponse struct {
	MerchantName string `json:"merchant_name" gorm:"column:merchant_name"`
	OutletName   string `json:"outlet_name" gorm:"column:outlet_name"`
	Omzet        int    `json:"omzet" gorm:"column:omzet"`
	Date         string `json:"date" gorm:"column:date"`
}
