package file

import (
	"errors"

	"fresh-shop/server/global"
	"fresh-shop/server/model/file"
	"gorm.io/gorm"
)

type FileUploadAndDownloadService struct{}

//@author: [dalefeng](https://github.com/dalefeng)
//@function: FindOrCreateFile
//@description: 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
//@param: fileMd5 string, fileName string, chunkTotal int
//@return: file model.ExaFile, err error

func (e *FileUploadAndDownloadService) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (f file.ExaFile, err error) {
	var cfile file.ExaFile
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal

	if errors.Is(global.DB.Where("file_md5 = ? AND is_finish = ?", fileMd5, true).First(&f).Error, gorm.ErrRecordNotFound) {
		err = global.DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("ExaFileChunk").FirstOrCreate(&f, cfile).Error
		return f, err
	}
	cfile.IsFinish = true
	cfile.FilePath = f.FilePath
	err = global.DB.Create(&cfile).Error
	return cfile, err
}

//@author: [dalefeng](https://github.com/dalefeng)
//@function: CreateFileChunk
//@description: 创建文件切片记录
//@param: id uint, fileChunkPath string, fileChunkNumber int
//@return: error

func (e *FileUploadAndDownloadService) CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {
	var chunk file.ExaFileChunk
	chunk.FileChunkPath = fileChunkPath
	chunk.ExaFileID = id
	chunk.FileChunkNumber = fileChunkNumber
	err := global.DB.Create(&chunk).Error
	return err
}

//@author: [dalefeng](https://github.com/dalefeng)
//@function: DeleteFileChunk
//@description: 删除文件切片记录
//@param: fileMd5 string, fileName string, filePath string
//@return: error

func (e *FileUploadAndDownloadService) DeleteFileChunk(fileMd5 string, filePath string) error {
	var chunks []file.ExaFileChunk
	var file file.ExaFile
	err := global.DB.Where("file_md5 = ? ", fileMd5).First(&file).
		Updates(map[string]interface{}{
			"IsFinish":  true,
			"file_path": filePath,
		}).Error
	if err != nil {
		return err
	}
	err = global.DB.Where("exa_file_id = ?", file.ID).Delete(&chunks).Unscoped().Error
	return err
}
