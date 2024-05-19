package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/aliamerj/aircup/server/apk/modules"
	resType "github.com/aliamerj/aircup/server/internal/constant"
	"github.com/aliamerj/aircup/server/internal/utility"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DirectoryInfo struct {
	Path          string     `json:"path"`
	TotalSize     string     `json:"totalSize"`
	AvailableSize string     `json:"availableSize"`
	Contents      []FileInfo `json:"contents"`
}

type FileInfo struct {
	Name         string     `json:"name"`
	Path         string     `json:"path"`
	Size         *string    `json:"size"`
	Extension    *string    `json:"extension"`
	ModifiedTime *time.Time `json:"modifiedTime"`
}

func CheckdirPath(c echo.Context, db *gorm.DB) error {
	path := c.QueryParam("path")
	if path == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Path is required")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return echo.NewHTTPError(http.StatusBadRequest, "Path not exist")
	}

	return c.JSON(http.StatusOK, utility.SuccesRespnse("path is available", resType.AvailablePath, true))
}
func SaveNewDirPath(c echo.Context, db *gorm.DB) error {
	path := c.QueryParam("path")
	if path == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Path is required")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return echo.NewHTTPError(http.StatusBadRequest, "Path not exist")
	}

	disk := modules.Disk{
		DiskPath: path,
	}
	if err := db.Create(&disk).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: disks.disk_path") {
			return echo.NewHTTPError(http.StatusBadRequest, "Path already exist")
		}

		return err
	}
	return c.JSON(http.StatusOK, utility.SuccesRespnse("saved path", resType.AvailablePath, true))
}

func GetSavedDisk(c echo.Context, db *gorm.DB) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot find the user's home directory: " + err.Error())
	}
	disk := modules.Disk{
		DiskPath: homeDir,
	}
	if err := db.Create(&disk).Error; err != nil {
		if !strings.Contains(err.Error(), "UNIQUE constraint failed: disks.disk_path") {
			return err
		}
	}
	var disks []modules.Disk
	if err := db.Find(&disks).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, utility.SuccesRespnse("Fetch Succesfully", resType.GetDisks, disks))
}

func GetDirInfo(c echo.Context, db *gorm.DB) error {
	path := c.QueryParam("path")
	if path == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Path is required")
	}
	// Get directory details
	dirInfo, err := getDirectoryInfo(path)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, utility.SuccesRespnse("Fetch Succesfully", resType.GetDirInfo, dirInfo))
}

func getDirectoryInfo(path string) (*DirectoryInfo, error) {
	// Read directory contents
	files, err := os.ReadDir(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "directory not found")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to read directory")
	}
	// Get disk usage info
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to get disk info")
	}

	total := 0
	available := stat.Bfree * uint64(stat.Bsize)

	var contents []FileInfo
	fileInfoChan := make(chan FileInfo, len(files))
	defer close(fileInfoChan)

	for _, file := range files {
		go func(file os.DirEntry) {
			isDir := file.IsDir()
			fileInfo, err := file.Info()
			if err != nil {
				fileInfoChan <- FileInfo{
					Name:         file.Name(),
					Path:         fmt.Sprintf("%s/%s", path, file.Name()),
					Size:         nil,
					Extension:    getFileExtension(file.Name(), isDir),
					ModifiedTime: nil,
				}
				return
			}

			sizeValue := fileInfo.Size()
			total += int(sizeValue)
			size := humanReadableSize(sizeValue)

			modifiedTimeValue := fileInfo.ModTime()
			modifiedTime := &modifiedTimeValue

			fileInfoChan <- FileInfo{
				Name:         file.Name(),
				Path:         fmt.Sprintf("%s/%s", path, file.Name()),
				Size:         &size,
				Extension:    getFileExtension(file.Name(), isDir),
				ModifiedTime: modifiedTime,
			}
		}(file)
	}

	for i := 0; i < len(files); i++ {
		contents = append(contents, <-fileInfoChan)
	}

	dirInfo := &DirectoryInfo{
		Path:          path,
		TotalSize:     humanReadableSize(int64(total)),
		AvailableSize: humanReadableSize(int64(available)),
		Contents:      contents,
	}

	return dirInfo, nil
}

func humanReadableSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}
	sizeInKB := float64(size) / 1024
	if sizeInKB < 1024 {
		return fmt.Sprintf("%.2f KB", sizeInKB)
	}
	sizeInMB := sizeInKB / 1024
	if sizeInMB < 1024 {
		return fmt.Sprintf("%.2f MB", sizeInMB)
	}
	sizeInGB := sizeInMB / 1024
	return fmt.Sprintf("%.2f GB", sizeInGB)
}

func getFileExtension(fileName string, isDirectory bool) *string {
	if isDirectory {
		folder := "folder"
		return &folder
	}
	parts := strings.Split(fileName, ".")
	if len(parts) > 1 {
		return &parts[len(parts)-1]
	}
	noExt := "noext"
	return &noExt
}

