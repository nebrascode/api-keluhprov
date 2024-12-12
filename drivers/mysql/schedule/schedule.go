package schedule

import (
	"e-complaint-api/entities"
	"gorm.io/gorm"
)

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) entities.ScheduleRepositoryInterface {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) Create(schedule *entities.Schedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepository) GetAll() ([]entities.Schedule, error) {
	var results []entities.Schedule
	err := r.db.Find(&results).Error
	return results, err
}

func (r *scheduleRepository) GetByID(id int64) (*entities.Schedule, error) {
	var result entities.Schedule
	err := r.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *scheduleRepository) Update(id int64, schedule *entities.Schedule) error {
	return r.db.Model(&entities.Schedule{}).Where("id = ?", id).Updates(schedule).Error
}

func (r *scheduleRepository) Delete(id int64) error {
	return r.db.Delete(&entities.Schedule{}, id).Error
}
