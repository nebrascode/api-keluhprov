package entities

type Schedule struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Email     string `gorm:"size:255;not null" json:"email"`
	Job       string `gorm:"size:255;not null" json:"job"`
	StartDate string `gorm:"not null" json:"start_date"`
	EndDate   string `gorm:"not null" json:"end_date"`
	Status    string `gorm:"size:50;not null" json:"status"`
}

type ScheduleRepositoryInterface interface {
	Create(schedule *Schedule) error
	GetAll() ([]Schedule, error)
	GetByID(id int64) (*Schedule, error)
	Update(id int64, schedule *Schedule) error
	Delete(id int64) error
}

type ScheduleUseCaseInterface interface {
	Create(schedule *Schedule) error
	GetAll() ([]Schedule, error)
	GetByID(id int64) (*Schedule, error)
	Update(id int64, schedule *Schedule) error
	Delete(id int64) error
}
