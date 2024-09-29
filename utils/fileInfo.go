package utils

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

type FileTransfer struct {
	TransferUUID uuid.UUID
	FilePath     string
	FileInfo     os.FileInfo
}

func ParseFiles(files []string) ([]FileTransfer, error) {
	ftList := make([]FileTransfer, 0, len(files))
	for _, relFP := range files {
		fp, err := filepath.Abs(relFP)
		if err != nil {
			return nil, fmt.Errorf("error with file '%s': %v", relFP, err)
		}
		fi, err := os.Stat(fp)
		if err != nil {
			return nil, fmt.Errorf("error with file '%s': %v", fp, err)
		}
		ftList = append(ftList, FileTransfer{
			TransferUUID: uuid.New(),
			FilePath:     fp,
			FileInfo:     fi,
		})
	}
	return ftList, nil
}
