package models

type CategoriesDTO struct {
	CategoryName string `orm:"size(40)"`
	ImagePath    string `orm:"size(250)"`
	// Image      *multipart.File
	// FileHeader *multipart.FileHeader
}

// type File struct {
// 	*file // os specific

// }

// type file struct {
// 	fd      int
// 	name    string
// 	dirinfo string // nil unless directory being read
// }
