package response

import "fresh-shop/server/model/file"

type ExaFileResponse struct {
	File file.ExaFileUploadAndDownload `json:"file"`
}
