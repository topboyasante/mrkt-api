package cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/topboyasante/mrkt-api/internal/config"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cldSecret := config.ENV.CloudinaryAPISecret
	cldName := config.ENV.CloudinaryCloudName
	cldKey := config.ENV.CloudinaryAPIKey

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
