package unggah_bukti

import (
	"e-complaint-api/entities"
	"gorm.io/gorm"
)

type unggahBuktiRepository struct {
	db *gorm.DB
}

func NewUnggahBuktiRepository(db *gorm.DB) entities.UnggahBuktiRepositoryInterface {
	return &unggahBuktiRepository{db: db}
}

func (r *unggahBuktiRepository) Create(unggahBukti *entities.UnggahBukti) error {
	return r.db.Create(unggahBukti).Error
}

func (r *unggahBuktiRepository) GetAll() ([]entities.UnggahBukti, error) {
	var results []entities.UnggahBukti
	err := r.db.Find(&results).Error
	return results, err
}

func (r *unggahBuktiRepository) GetByComplaintID(complaintID string) ([]entities.UnggahBukti, error) {
	var results []entities.UnggahBukti
	err := r.db.Where("complaint_id = ?", complaintID).Find(&results).Error
	return results, err
}

// Implementasi GetByID untuk mengambil data berdasarkan ID
func (r *unggahBuktiRepository) GetByID(id int64) (*entities.UnggahBukti, error) {
	var result entities.UnggahBukti
	err := r.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *unggahBuktiRepository) Update(id int64, unggahBukti *entities.UnggahBukti) error {
	return r.db.Model(&entities.UnggahBukti{}).Where("id = ?", id).Updates(unggahBukti).Error
}

func (r *unggahBuktiRepository) Delete(id int64) error {
	return r.db.Delete(&entities.UnggahBukti{}, id).Error
}
