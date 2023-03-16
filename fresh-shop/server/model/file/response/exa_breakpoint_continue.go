package response

import "fresh-shop/server/model/file"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File file.ExaFile `json:"file"`
}
