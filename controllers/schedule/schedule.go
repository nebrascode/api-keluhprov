package schedule

import (
	"e-complaint-api/entities"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ScheduleController struct {
	usecase entities.ScheduleUseCaseInterface
}

func NewScheduleController(usecase entities.ScheduleUseCaseInterface) *ScheduleController {
	return &ScheduleController{usecase: usecase}
}

// Create Schedule
func (c *ScheduleController) Create(ctx echo.Context) error {
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	job := ctx.FormValue("job")
	status := ctx.FormValue("status")
	startDate := ctx.FormValue("start_date")
	endDate := ctx.FormValue("end_date")

	// Parsing tanggal
	startDateTime, err := time.Parse("02/01/2006", startDate)
	if err != nil {
		log.Printf("Error parsing start_date: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid start_date format (use DD/MM/YYYY)"})
	}

	endDateTime, err := time.Parse("02/01/2006", endDate)
	if err != nil {
		log.Printf("Error parsing end_date: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid end_date format (use DD/MM/YYYY)"})
	}

	// Simpan tanggal dalam format YYYY-MM-DD
	schedule := &entities.Schedule{
		Name:      name,
		Email:     email,
		Job:       job,
		Status:    status,
		StartDate: startDateTime.Format("2006-01-02"),
		EndDate:   endDateTime.Format("2006-01-02"),
	}

	// Simpan ke database
	if err := c.usecase.Create(schedule); err != nil {
		log.Printf("Error creating schedule: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create schedule"})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message":  "Schedule created successfully",
		"schedule": schedule,
	})
}

// Get All Schedules
func (c *ScheduleController) GetAll(ctx echo.Context) error {
	data, err := c.usecase.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch schedules"})
	}
	return ctx.JSON(http.StatusOK, data)
}

// Get Schedule by ID
func (c *ScheduleController) GetByID(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data, err := c.usecase.GetByID(id)
	if err != nil || data == nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Schedule not found"})
	}

	return ctx.JSON(http.StatusOK, data)
}

// Update Schedule
func (c *ScheduleController) Update(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	job := ctx.FormValue("job")
	status := ctx.FormValue("status")
	startDate := ctx.FormValue("start_date")
	endDate := ctx.FormValue("end_date")

	// Parsing tanggal
	startDateTime, err := time.Parse("02/01/2006", startDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid start_date format (use DD/MM/YYYY)"})
	}

	endDateTime, err := time.Parse("02/01/2006", endDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid end_date format (use DD/MM/YYYY)"})
	}

	// Update schedule dengan tanggal format YYYY-MM-DD
	updatedSchedule := &entities.Schedule{
		Name:      name,
		Email:     email,
		Job:       job,
		Status:    status,
		StartDate: startDateTime.Format("2006-01-02"),
		EndDate:   endDateTime.Format("2006-01-02"),
	}

	if err := c.usecase.Update(id, updatedSchedule); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update schedule"})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Schedule updated successfully",
		"schedule": updatedSchedule,
	})
}

// Delete Schedule
func (c *ScheduleController) Delete(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err := c.usecase.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete schedule"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Schedule deleted successfully"})
}
