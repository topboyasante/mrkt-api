package images

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/topboyasante/mrkt-api/api/v1/pkg/cloudinary"
	"github.com/topboyasante/mrkt-api/internal/config"
)

func UploadToCloudinary(file multipart.File, filePath string) (string, string, error) {
	ctx := context.Background()
	cld, err := cloudinary.SetupCloudinary()
	if err != nil {
		return "", "", err
	}
	uploadParams := uploader.UploadParams{
		PublicID:     filePath,
		Folder:       "mrkt",
		UploadPreset: config.ENV.CloudinaryUploadPreset,
	}

	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", "", err
	}

	imageUrl := result.SecureURL
	imagePublicId := result.PublicID

	return imageUrl, imagePublicId, nil
}

func DeleteFromCloudinary(publicId string) error {
	ctx := context.Background()
	cld, err := cloudinary.SetupCloudinary()
	if err != nil {
		return err
	}

	deletionParams := uploader.DestroyParams{
		PublicID: publicId,
	}

	_, err = cld.Upload.Destroy(ctx, deletionParams)
	if err != nil {
		return err
	}

	return nil
}

func UpdateImageOnCloudinary(newFile multipart.File, existingPublicId string, newFilePath string) (string, string, error) {
	err := DeleteFromCloudinary(existingPublicId)
	if err != nil {
		return "", "", fmt.Errorf("failed to delete existing image: %w", err)
	}

	newImageUrl, newImagePublicId, err := UploadToCloudinary(newFile, newFilePath)
	if err != nil {
		return "", "", fmt.Errorf("failed to upload new image: %w", err)
	}

	return newImageUrl, newImagePublicId, nil
}
