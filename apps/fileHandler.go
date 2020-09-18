package apps

import (
	"CloudStorage/utils"
	"os"
	"path/filepath"
	"time"
)

type FileMetaData struct {
	FileName      string
	FilePath      string
	FileSize      int64
	UploadTime    string
	FileSha1Value string
}


func NewFileMeta(path string, f *os.File) *FileMetaData {
	return &FileMetaData{
		FileName:      filepath.Base(path),
		FilePath:      path,
		FileSize:      utils.GetFileSize(path),
		UploadTime:    time.Now().Format("2006-01-02 15:04:05"),
		FileSha1Value: utils.FileSha1(f),
	}
}
