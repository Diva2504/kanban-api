package repository

import (
	"github.com/takadev15/kanban-api/models"
	"gorm.io/gorm"
)

func GetAllCategories(db *gorm.DB) ([]models.Category, error) {
	var categories []models.Category

	res := db.Find(&categories)

	if res.Error != nil {
		return nil, res.Error
	} else {
		if res.RowsAffected <= 0 {
			return nil, res.Error
		} else {
			return categories, res.Error
		}
	}
}

func CreateCategory(data *models.Category, db *gorm.DB) error {

	err := db.Debug().Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(id int, db *gorm.DB) error {
	var categories models.Category

	err := db.Delete(&categories, id).Error

	if err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateCategory(id int, input *models.Category, db *gorm.DB) (models.Category, error) {
	var category models.Category
	err := db.Model(&category).Where("id = ?", id).Updates(&input).Error

	if err != nil {
		return models.Category{}, err
	} else {
		return category, err
	}
}
