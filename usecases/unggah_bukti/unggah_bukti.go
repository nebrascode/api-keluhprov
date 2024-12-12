package unggah_bukti

import "e-complaint-api/entities"

type unggahBuktiUseCase struct {
	repo entities.UnggahBuktiRepositoryInterface
}

func NewUnggahBuktiUseCase(repo entities.UnggahBuktiRepositoryInterface) entities.UnggahBuktiUseCaseInterface {
	return &unggahBuktiUseCase{repo: repo}
}

func (uc *unggahBuktiUseCase) Create(unggahBukti *entities.UnggahBukti) error {
	return uc.repo.Create(unggahBukti)
}

func (uc *unggahBuktiUseCase) GetAll() ([]entities.UnggahBukti, error) {
	return uc.repo.GetAll()
}

func (uc *unggahBuktiUseCase) GetByComplaintID(complaintID string) ([]entities.UnggahBukti, error) {
	return uc.repo.GetByComplaintID(complaintID)
}

// Menambahkan method GetByID
func (uc *unggahBuktiUseCase) GetByID(id int64) (*entities.UnggahBukti, error) {
	return uc.repo.GetByID(id)
}

func (uc *unggahBuktiUseCase) Update(id int64, unggahBukti *entities.UnggahBukti) error {
	return uc.repo.Update(id, unggahBukti)
}

func (uc *unggahBuktiUseCase) Delete(id int64) error {
	return uc.repo.Delete(id)
}
