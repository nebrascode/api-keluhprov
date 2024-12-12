package complaint_file

import (
	"e-complaint-api/entities"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type ComplaintFileUseCase struct {
	repository entities.ComplaintFileRepositoryInterface
	//gcs_api    entities.ComplaintFileGCSAPIInterface
}

func NewComplaintFileUseCase(repository entities.ComplaintFileRepositoryInterface, gcs_api entities.ComplaintFileGCSAPIInterface) *ComplaintFileUseCase {
	return &ComplaintFileUseCase{
		repository: repository,
		//gcs_api:    gcs_api,
	}
}

//func (u *ComplaintFileUseCase) Create(files []*multipart.FileHeader, complaintID string) ([]entities.ComplaintFile, error) {
//	filepaths, err_upload := u.gcs_api.Upload(files)
//	if err_upload != nil {
//		return []entities.ComplaintFile{}, err_upload
//	}
//
//	var complaintFiles []*entities.ComplaintFile
//	for _, filepath := range filepaths {
//		complaintFile := &entities.ComplaintFile{
//			ComplaintID: complaintID,
//			Path:        filepath,
//		}
//		complaintFiles = append(complaintFiles, complaintFile)
//	}
//
//	err_create := u.repository.Create(complaintFiles)
//	if err_create != nil {
//		return []entities.ComplaintFile{}, err_create
//	}
//
//	var convertedComplaintFiles []entities.ComplaintFile
//	for _, cf := range complaintFiles {
//		convertedComplaintFiles = append(convertedComplaintFiles, *cf)
//	}
//
//	return convertedComplaintFiles, nil
//}

func (u *ComplaintFileUseCase) Create(files []*multipart.FileHeader, complaintID string) ([]entities.ComplaintFile, error) {
	var complaintFiles []*entities.ComplaintFile

	// Tentukan path penyimpanan lokal
	destinationDir := "./uploads/complaint_files"
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("gagal membuat direktori: %v", err)
	}

	// Simpan file satu per satu ke direktori lokal
	for _, file := range files {
		// Tentukan nama file dan path lengkap
		fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		filePath := filepath.Join(destinationDir, fileName)

		// Buka file multipart untuk membaca isinya
		src, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("gagal membuka file: %v", err)
		}
		defer src.Close()

		// Buka file tujuan untuk menulis
		dst, err := os.Create(filePath)
		if err != nil {
			return nil, fmt.Errorf("gagal membuat file: %v", err)
		}
		defer dst.Close()

		// Salin isi file dari sumber ke tujuan
		_, err = io.Copy(dst, src)
		if err != nil {
			return nil, fmt.Errorf("gagal menyalin file: %v", err)
		}

		// Buat entri ComplaintFile untuk database
		complaintFile := &entities.ComplaintFile{
			ComplaintID: complaintID,
			Path:        filePath, // Menyimpan path lokal
		}
		complaintFiles = append(complaintFiles, complaintFile)
	}

	// Simpan ke database
	err_create := u.repository.Create(complaintFiles)
	if err_create != nil {
		return nil, err_create
	}

	var convertedComplaintFiles []entities.ComplaintFile
	for _, cf := range complaintFiles {
		convertedComplaintFiles = append(convertedComplaintFiles, *cf)
	}

	return convertedComplaintFiles, nil
}

//func (u *ComplaintFileUseCase) DeleteByComplaintID(complaintID string) error {
//	err_delete := u.repository.DeleteByComplaintID(complaintID)
//	if err_delete != nil {
//		return err_delete
//	}
//
//	return nil
//}

func (u *ComplaintFileUseCase) DeleteByComplaintID(complaintID string) error {
	// Ambil file path dari database
	complaintFiles, err := u.repository.FindByComplaintID(complaintID)
	if err != nil {
		return err
	}

	// Hapus file fisik dari sistem lokal
	for _, file := range complaintFiles {
		err := os.Remove(file.Path)
		if err != nil {
			return fmt.Errorf("gagal menghapus file: %v", err)
		}
	}

	// Hapus entri dari database
	err_delete := u.repository.DeleteByComplaintID(complaintID)
	if err_delete != nil {
		return err_delete
	}

	return nil
}
