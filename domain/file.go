package domain

import "mime/multipart"

type UploadFileInput struct {
	FileHeader *multipart.FileHeader `json:"file_header" form:"file_header"`
	FileName   string                `json:"file_name"`
}
