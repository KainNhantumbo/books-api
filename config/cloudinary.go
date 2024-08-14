package config

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadResult struct {
	SecureUrl string
	PublicId  string
	Url       string
}

func credentials() (*cloudinary.Cloudinary, context.Context) {
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	cld.Config.Cloud.APIKey = Config("CLOUDINARY_API_KEY")
	cld.Config.Cloud.APISecret = Config("CLOUDINARY_API_SECRET")
	cld.Config.Cloud.CloudName = Config("CLOUDINARY_NAME")

	ctx := context.Background()
	return cld, ctx
}

func UploadAsset(base64StringFile *string, publicId *string) (*uploader.UploadResult, error) {
	api, ctx := credentials()
	res, err := api.Upload.Upload(ctx, &base64StringFile, uploader.UploadParams{
		PublicID: *publicId,
		Folder:   "/golang-books-api/assets",
	})

	if err != nil {
		return &uploader.UploadResult{}, err
	}
	return res, nil
}

func DeleteAsset(publicId *string) (*uploader.DestroyResult, error) {
	var invalidate bool = true
	api, ctx := credentials()

	res, err := api.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:   *publicId,
		Invalidate: &invalidate,
	})

	if err != nil {
		return &uploader.DestroyResult{}, nil
	}
	return res, nil
}
