package utils

import (
	"context"
	"mime/multipart"
	"slot/models"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func CreateValidator(event *models.Event) string {
	if event.Duration != 30 {
		return "Duration must be 30 min."
	}

	return ""
}

func RemoveExt(filename string) string {
    newName := strings.Split(filename, ".")
    return newName[0]
}

var UrlChan = make(chan string)
var ErrChan = make(chan string)

func UploadToCloud(file multipart.File, filename string) {
    var ctx = context.Background()
	cld, _ := cloudinary.NewFromParams("roshandev2000", "259612962344283", "FLEZzfwrnl2D8X6xbr8MiVDGV90")
    resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: filename})

    if err != nil {
        ErrChan <- "Failed to upload file."
    }

    UrlChan <- resp.SecureURL
}