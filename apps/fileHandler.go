package apps

import (
	"CloudStorage/utils"
	"os"
	"path/filepath"
	"time"
)

const (
	FILE_STATUS_AVALIABLE = iota
	FILE_STATUS_UNAVALIABLE
	FILE_STATUS_REMOVED
)

type FileMetaData struct {
	Id            int64
	FileName      string
	FilePath      string
	FileSize      int64
	UploadTime    string
	FileSha1Value string
	FileStatus    int
}

func NewFileMeta(path string, f *os.File) *FileMetaData {
	return &FileMetaData{
		FileName:      filepath.Base(path),
		FilePath:      path,
		FileSize:      utils.GetFileSize(path),
		UploadTime:    time.Now().Format("2006-01-02 15:04:05"),
		FileSha1Value: utils.FileSha1(f),
		FileStatus:    FILE_STATUS_AVALIABLE,
	}
}
