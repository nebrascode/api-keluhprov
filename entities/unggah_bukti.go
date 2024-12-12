package entities

import "time"

type UnggahBukti struct {
	ID              int64     `gorm:"primaryKey" json:"id"`
	ComplaintID     string    `gorm:"size:15;not null" json:"complaint_id"`
	Path            string    `gorm:"size:255;not null" json:"path"`
	PenanggungJawab string    `gorm:"size:255;not null" json:"penanggung_jawab"`
	FinishedOn      time.Time `gorm:"not null" json:"finished_on"`
}

type UnggahBuktiRepositoryInterface interface {
	Create(unggahBukti *UnggahBukti) error
	GetAll() ([]UnggahBukti, error)
	GetByComplaintID(complaintID string) ([]UnggahBukti, error)
	GetByID(id int64) (*UnggahBukti, error) // Tambahkan GetByID
	Update(id int64, unggahBukti *UnggahBukti) error
	Delete(id int64) error
}

type UnggahBuktiUseCaseInterface interface {
	Create(unggahBukti *UnggahBukti) error
	GetAll() ([]UnggahBukti, error)
	GetByComplaintID(complaintID string) ([]UnggahBukti, error)
	GetByID(id int64) (*UnggahBukti, error) // Tambahkan GetByID
	Update(id int64, unggahBukti *UnggahBukti) error
	Delete(id int64) error
}
