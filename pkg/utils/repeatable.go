package repeatable

import (
	"bd-backend/internal/config"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}

	return
}

func CrateDir() error {
	cfg := config.GetConfig()
	return os.MkdirAll(cfg.PublicFilePath, os.ModePerm)
}

func CreateFile(folderName string, file *multipart.FileHeader, r *gin.Context) (string, error) {
	cfg := config.GetConfig()
	appName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(filepath.Base(file.Filename)))
	appDirPath := cfg.PublicFilePath + "/" + folderName

	err := os.MkdirAll(appDirPath, os.ModePerm)
	if err != nil {
		fmt.Println("errr------------------", err)
		return "str", err
	}

	pathApp := fmt.Sprintf("%s/%s", appDirPath, appName)
	if err := r.SaveUploadedFile(file, pathApp); err != nil {
		fmt.Println("errr------------------", err)
		return "", err
	}

	pathApp = "/public/" + folderName + "/" + appName
	return pathApp, err
}

func CreateFileGin(folderName string, file *multipart.FileHeader, r *gin.Context) (string, error) {
	cfg := config.GetConfig()
	appName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(filepath.Base(file.Filename)))
	appDirPath := cfg.PublicFilePath + "/" + folderName

	err := os.MkdirAll(appDirPath, os.ModePerm)
	if err != nil {
		return "str", err
	}

	pathApp := fmt.Sprintf("%s/%s", appDirPath, appName)
	if err := r.SaveUploadedFile(file, pathApp); err != nil {
		return "", err
	}

	pathApp = "/public/" + folderName + "/" + appName
	return pathApp, err
}
