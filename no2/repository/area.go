package repository

import (
	Model "no2-majoo-be-test/model"

	"gorm.io/gorm"
)

type AreaRepo interface {
	InsertArea(param1 int64, param2 int64, areaType string, ar *Model.Area) (err error)
}

func NewAreaRepo(db *gorm.DB) AreaRepo {
	return &AreaRepository{
		DB: db,
	}
}

type AreaRepository struct {
	DB *gorm.DB
}

func (_r *AreaRepository) InsertArea(param1 int64, param2 int64, areaType string, ar *Model.Area) (err error) {
	ar.AreaType = areaType
	switch ar.AreaType {
	case "persegi panjang":
		ar.AreaValue = param1 * param2
	case "persegi":
		ar.AreaValue = param1 * param2
	case "segitiga":
		ar.AreaValue = param1 * param2 / 2
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data"
	}
	err = _r.DB.Create(&ar).Error
	if err != nil {
		return err
	}
	return nil
}
