package schedule

import "e-complaint-api/entities"

type scheduleUseCase struct {
	repo entities.ScheduleRepositoryInterface
}

func NewScheduleUseCase(repo entities.ScheduleRepositoryInterface) entities.ScheduleUseCaseInterface {
	return &scheduleUseCase{repo: repo}
}

func (uc *scheduleUseCase) Create(schedule *entities.Schedule) error {
	return uc.repo.Create(schedule)
}

func (uc *scheduleUseCase) GetAll() ([]entities.Schedule, error) {
	return uc.repo.GetAll()
}

func (uc *scheduleUseCase) GetByID(id int64) (*entities.Schedule, error) {
	return uc.repo.GetByID(id)
}

func (uc *scheduleUseCase) Update(id int64, schedule *entities.Schedule) error {
	return uc.repo.Update(id, schedule)
}

func (uc *scheduleUseCase) Delete(id int64) error {
	return uc.repo.Delete(id)
}
