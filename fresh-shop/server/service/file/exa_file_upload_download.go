package file

import (
	"errors"
	"mime/multipart"
	"strings"

	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/file"
	"fresh-shop/server/utils/upload"
)

//@author: [dalefeng](https://github.com/dalefeng)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) Upload(file file.ExaFileUploadAndDownload) error {
	return global.DB.Create(&file).Error
}

//@author: [dalefeng](https://github.com/dalefeng)
//@function: FindFile
//@description: 查询文件记录
//@param: id uint
//@return: model.ExaFileUploadAndDownload, error

func (e *FileUploadAndDownloadService) FindFile(id uint) (file.ExaFileUploadAndDownload, error) {
	var f file.ExaFileUploadAndDownload
	err := global.DB.Where("id = ?", id).First(&f).Error
	return f, err
}

//@author: [dalefeng](https://github.com/dalefeng)
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.ExaFileUploadAndDownload
//@return: err error

func (e *FileUploadAndDownloadService) DeleteFile(f file.ExaFileUploadAndDownload) (err error) {
	var fileFromDb file.ExaFileUploadAndDownload
	fileFromDb, err = e.FindFile(f.ID)
	if err != nil {
		return
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.DB.Where("id = ?", f.ID).Unscoped().Delete(&f).Error
	return err
}

// EditFileName 编辑文件名或者备注
func (e *FileUploadAndDownloadService) EditFileName(f file.ExaFileUploadAndDownload) (err error) {
	var fileFromDb file.ExaFileUploadAndDownload
	return global.DB.Where("id = ?", f.ID).First(&fileFromDb).Update("name", f.Name).Error
}

//@author: [dalefeng](https://github.com/dalefeng)
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (e *FileUploadAndDownloadService) GetFileRecordInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword
	db := global.DB.Model(&file.ExaFileUploadAndDownload{})
	var fileLists []file.ExaFileUploadAndDownload
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return fileLists, total, err
}

//@author: [dalefeng](https://github.com/dalefeng)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: file model.ExaFileUploadAndDownload, err error

func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (f file.ExaFileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := file.ExaFileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return f, e.Upload(f)
	}
	return
}
