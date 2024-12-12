package unggah_bukti

import (
	"e-complaint-api/entities"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type UnggahBuktiController struct {
	usecase entities.UnggahBuktiUseCaseInterface
}

func NewUnggahBuktiController(usecase entities.UnggahBuktiUseCaseInterface) *UnggahBuktiController {
	return &UnggahBuktiController{usecase: usecase}
}

func (c *UnggahBuktiController) Create(ctx echo.Context) error {
	// Parsing form-data
	complaintID := ctx.FormValue("complaint_id")
	penanggungJawab := ctx.FormValue("penanggung_jawab")
	finishedOn := ctx.FormValue("finished_on")

	// Ambil file dari form-data
	file, err := ctx.FormFile("path")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "File is required"})
	}

	// Buka file untuk disimpan
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open file"})
	}
	defer src.Close()

	// Buat direktori jika belum ada
	uploadDir := "./uploads/bukti_unggah"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create upload directory"})
		}
	}

	// Generate nama file unik
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename
	// Simpan file dengan path tanpa "./"
	filePath := "uploads/bukti_unggah/" + fileName

	// Simpan file ke direktori lokal
	dst, err := os.Create(uploadDir + "/" + fileName)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save file"})
	}
	defer dst.Close()

	// Copy isi file dari form-data ke file tujuan
	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to copy file"})
	}

	// Konversi finishedOn ke time.Time
	finishedOnTime, err := time.Parse("2006-01-02", finishedOn)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid finished_on format (use YYYY-MM-DD)"})
	}

	// Buat objek untuk disimpan ke database
	unggahBukti := &entities.UnggahBukti{
		ComplaintID:     complaintID,
		Path:            filePath, // Path yang disimpan di database
		PenanggungJawab: penanggungJawab,
		FinishedOn:      finishedOnTime,
	}

	// Simpan ke database menggunakan UseCase
	if err := c.usecase.Create(unggahBukti); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create record"})
	}

	// Berikan respon sukses
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message":        "Data uploaded successfully",
		"uploaded_bukti": unggahBukti,
	})
}


func (c *UnggahBuktiController) GetAll(ctx echo.Context) error {
	data, err := c.usecase.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch"})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *UnggahBuktiController) GetByComplaintID(ctx echo.Context) error {
	complaintID := ctx.Param("complaint-id")
	data, err := c.usecase.GetByComplaintID(complaintID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch"})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *UnggahBuktiController) Update(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	var data entities.UnggahBukti
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	if err := c.usecase.Update(id, &data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update"})
	}
	return ctx.JSON(http.StatusOK, data)
}

// Fungsi untuk menghapus unggah bukti berdasarkan ID
func (c *UnggahBuktiController) Delete(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	// Ambil data berdasarkan ID
	data, err := c.usecase.GetByID(id)
	if err != nil || data == nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Record not found"})
	}

	// Hapus file terkait
	if err := os.Remove(data.Path); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete file"})
	}

	// Hapus data dari database
	if err := c.usecase.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete record"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Deleted successfully"})
}