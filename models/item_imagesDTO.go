package models

type ItemImagesDTO struct {
	ImageName string `orm:"size(40)"`
	ImagePath    string `orm:"size(250)"`
	// Image      *multipart.File
	// FileHeader *multipart.FileHeader
}
