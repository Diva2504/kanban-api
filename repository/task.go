package repository

import (
	"github.com/takadev15/kanban-api/models"
	"gorm.io/gorm"
)

func GetAllTask(db *gorm.DB) ([]models.Task, error) {
	var task []models.Task

	res := db.Find(&task)

	if res.Error != nil {
		return nil, res.Error
	} else {
		if res.RowsAffected <= 0 {
			return nil, res.Error
		} else {
			return task, res.Error
		}
	}
}

func CreateTask(data models.Task, db *gorm.DB) error {
	err := db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil

}

func UpdateTask(id int, data models.Task, db *gorm.DB) (models.Task, error) {
	var task models.Task
	err := db.Model(&task).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return models.Task{}, err
	}
	return task, err
}

func DeleteTask(id int, db *gorm.DB) error {
	var task models.Task

	del := db.Delete(&task, id)

	if del.Error != nil {
		return del.Error
	} else {
		return nil
	}
}

func UpdateCategoryTask(db *gorm.DB, id int, categoryId uint) error {
	var task models.Task
	err := db.Model(&task).Where("id = ?", id).Updates(models.Task{CategoryId: categoryId}).Error
	return err
}
