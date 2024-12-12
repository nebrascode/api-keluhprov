package news_file

import (
	"e-complaint-api/entities"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type NewsFileUseCase struct {
	repository entities.NewsFileRepositoryInterface
	//gcs_api    entities.NewsFileGCSAPIInterface
}

func NewNewsFileUseCase(repository entities.NewsFileRepositoryInterface, gcs_api entities.NewsFileGCSAPIInterface) *NewsFileUseCase {
	return &NewsFileUseCase{
		repository: repository,
		//gcs_api:    gcs_api,
	}
}

//func (u *NewsFileUseCase) Create(files []*multipart.FileHeader, newsID int) ([]entities.NewsFile, error) {
//	filepaths, err_upload := u.gcs_api.Upload(files)
//	if err_upload != nil {
//		return []entities.NewsFile{}, err_upload
//	}
//
//	var newsFiles []*entities.NewsFile
//	for _, filepath := range filepaths {
//		newsFile := &entities.NewsFile{
//			NewsID: newsID,
//			Path:   filepath,
//		}
//		newsFiles = append(newsFiles, newsFile)
//	}
//
//	err_create := u.repository.Create(newsFiles)
//	if err_create != nil {
//		return []entities.NewsFile{}, err_create
//	}
//
//	var convertedNewsFiles []entities.NewsFile
//	for _, nf := range newsFiles {
//		convertedNewsFiles = append(convertedNewsFiles, *nf)
//	}
//
//	return convertedNewsFiles, nil
//}

func (u *NewsFileUseCase) Create(files []*multipart.FileHeader, newsID int) ([]entities.NewsFile, error) {
	uploadDir := "./uploads/news_files" // Path direktori lokal

	// Buat direktori jika belum ada
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	var newsFiles []*entities.NewsFile

	for _, file := range files {
		// Gunakan path.Join untuk membuat path
		filePath := path.Join(uploadDir, file.Filename)

		// Pastikan path menggunakan forward slash
		filePath = strings.ReplaceAll(filePath, "\\", "/")

		// Buka file
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		// Buat file lokal
		dst, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		// Salin konten file
		if _, err := io.Copy(dst, src); err != nil {
			return nil, err
		}

		// Tambahkan file ke list untuk disimpan di database
		newsFile := &entities.NewsFile{
			NewsID: newsID,
			Path:   filePath, // Path dengan forward slash
		}
		newsFiles = append(newsFiles, newsFile)
	}

	// Simpan data file ke database
	err := u.repository.Create(newsFiles)
	if err != nil {
		return nil, err
	}

	var convertedNewsFiles []entities.NewsFile
	for _, nf := range newsFiles {
		convertedNewsFiles = append(convertedNewsFiles, *nf)
	}

	return convertedNewsFiles, nil
}

//func (u *NewsFileUseCase) DeleteByNewsID(newsID int) error {
//	err_delete := u.repository.DeleteByNewsID(newsID)
//	if err_delete != nil {
//		return err_delete
//	}
//
//	return nil
//}

func (u *NewsFileUseCase) DeleteByNewsID(newsID int) error {
	// Ambil semua file terkait dari repository
	files, err := u.repository.FindByNewsID(newsID) // Gunakan fungsi baru
	if err != nil {
		return err
	}

	// Hapus file dari direktori lokal
	for _, file := range files {
		if err := os.Remove(file.Path); err != nil && !os.IsNotExist(err) {
			return err
		}
	}

	// Hapus data dari database
	err = u.repository.DeleteByNewsID(newsID)
	if err != nil {
		return err
	}

	return nil
}
