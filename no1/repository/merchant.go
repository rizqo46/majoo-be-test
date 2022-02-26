package repository

import (
	"majoo-be-test/model"
	"majoo-be-test/util"

	"gorm.io/gorm"
)

type MerchantRepo interface {
	GetMerchantOmzet(merchantID, limit, page int) (*util.Pagination, error)
	GetOutletOmzet(outletID, limit, page int) (*util.Pagination, error)
	IsUserAuthToMerchant(userID, merchantID int) (bool, error)
	IsUserAuthToOutlet(userID, oultetID int) (bool, error)
}

type merchantRepoImpl struct {
	dB *gorm.DB
}

func NewMerchantRepo(dB *gorm.DB) MerchantRepo {
	return &merchantRepoImpl{dB}
}

func (r *merchantRepoImpl) IsUserAuthToMerchant(userID, merchantID int) (bool, error) {
	var count int64
	if err := r.dB.Raw(`SELECT COUNT(*) FROM Merchants m WHERE m.user_id = ? AND m.id = ?`, userID, merchantID).Count(&count).Error; err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}

func (r *merchantRepoImpl) IsUserAuthToOutlet(userID, oultetID int) (bool, error) {
	var count int64
	if err := r.dB.Raw(`SELECT COUNT(*) FROM Merchants m LEFT JOIN Outlets o ON o.merchant_id = m.id WHERE m.user_id = ? AND m.id = ?`, userID, oultetID).Count(&count).Error; err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}

func (r *merchantRepoImpl) GetMerchantOmzet(merchantID, limit, page int) (*util.Pagination, error) {
	var merchantName string
	if err := r.dB.Raw(`SELECT merchant_name FROM Merchants m WHERE m.id = ? LIMIT 1`, merchantID).
		Pluck("merchant_name", &merchantName).Error; err != nil {
		return nil, err
	}
	res := []model.MerchantResponse{}
	pagination := util.Pagination{
		Pagination: util.PaginationInfo{
			Limit:     limit,
			Page:      page,
			TotalRows: 30,
		},
	}

	if err := r.dB.Transaction(func(tx *gorm.DB) error {
		createDumbDateTable(tx)
		if err := tx.Raw(`SELECT ? as merchant_name, DATE_FORMAT(date, "%Y-%m-%d") as date, omzet FROM temp_date AS D LEFT JOIN
		(SELECT m.id, m.merchant_name, SUM(t.bill_total) as omzet, DATE(DATE_FORMAT(t.created_at, "%Y-%m-%d")) AS datef
		FROM Merchants m LEFT JOIN Transactions t ON t.merchant_id = m.id GROUP BY datef HAVING m.id=?) 
		as t ON t.datef = D.date LIMIT ? OFFSET ?;`, merchantName, merchantID, pagination.GetLimit(), pagination.GetOffset()).Find(&res).Error; err != nil {
			return nil
		}
		return nil
	}); err != nil {
		return nil, err
	}
	pagination.Pagination.TotalPages = 30 / pagination.GetLimit()
	if 30%pagination.GetLimit() != 0 {
		pagination.Pagination.TotalPages += 1
	}

	pagination.Data = res

	return &pagination, nil
}

func (r *merchantRepoImpl) GetOutletOmzet(outletID, limit, page int) (*util.Pagination, error) {
	var outlet map[string]interface{}
	if err := r.dB.Raw(`SELECT outlet_name, merchant_id FROM Outlets o WHERE o.id = ? LIMIT 1`, outletID).
		Find(&outlet).Error; err != nil {
		return nil, err
	}
	var merchantName string
	if err := r.dB.Raw(`SELECT merchant_name FROM Merchants m WHERE m.id = ? LIMIT 1`, outlet["merchant_id"]).
		Pluck("merchant_name", &merchantName).Error; err != nil {
		return nil, err
	}
	res := []model.OutletResponse{}
	pagination := util.Pagination{
		Pagination: util.PaginationInfo{
			Limit:     limit,
			Page:      page,
			TotalRows: 30,
		},
	}

	if err := r.dB.Transaction(func(tx *gorm.DB) error {
		createDumbDateTable(tx)
		if err := tx.Raw(`SELECT ? as merchant_name, ? as outlet_name, DATE_FORMAT(date, "%Y-%m-%d") as date, omzet FROM temp_date AS D LEFT JOIN
		(SELECT o.id, SUM(t.bill_total) as omzet, DATE(DATE_FORMAT(t.created_at, "%Y-%m-%d")) AS datef
		FROM Outlets o LEFT JOIN Transactions t ON t.outlet_id = o.id GROUP BY datef HAVING o.id=?) 
		as t ON t.datef = D.date LIMIT ? OFFSET ?;`, merchantName, outlet["outlet_name"], outletID, pagination.GetLimit(), pagination.GetOffset()).Find(&res).Error; err != nil {
			return nil
		}
		return nil
	}); err != nil {
		return nil, err
	}

	pagination.Pagination.TotalPages = 30 / pagination.GetLimit()
	if 30%pagination.GetLimit() != 0 {
		pagination.Pagination.TotalPages += 1
	}
	pagination.Data = res

	return &pagination, nil
}

func createDumbDateTable(tx *gorm.DB) {
	tx.Exec(`CREATE TEMPORARY TABLE IF NOT EXISTS temp_date(date DATE);
		`)
	tx.Exec(`
		INSERT INTO temp_date(date) VALUES ("2021-11-01"),
		("2021-11-02"),("2021-11-03"),("2021-11-04"),
		("2021-11-05"),("2021-11-06"),("2021-11-07"),
		("2021-11-08"),("2021-11-09"),
		("2021-11-10"),("2021-11-11"),("2021-11-12"),
		("2021-11-13"),("2021-11-14"),("2021-11-15"),
		("2021-11-16"),("2021-11-17"),("2021-11-18"),
		("2021-11-19"),("2021-11-20"),("2021-11-21"),
		("2021-11-22"),("2021-11-23"),("2021-11-24"),
		("2021-11-25"),("2021-11-26"),("2021-11-27"),
		("2021-11-28"),("2021-11-29"),("2021-11-30");
		`)
}
