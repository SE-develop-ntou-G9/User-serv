package usecase

import "mime/multipart"

type ImageUsecase struct {
	repo ImageRepository
}

func NewImageUsecase(imageRepo ImageRepository) *ImageUsecase {
	return &ImageUsecase{
		repo: imageRepo,
	}
}

func (uc *ImageUsecase) UploadAvatarToS3(file multipart.File, fileName string, contentType string) (string, error) {
	return uc.repo.UploadAvatarToS3(file, fileName, contentType)
}

func (uc *ImageUsecase) UploadLicenseToS3(file multipart.File, fileName string, contentType string) (string, error) {
	return uc.repo.UploadLicenseToS3(file, fileName, contentType)
}
