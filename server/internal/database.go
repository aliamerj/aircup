package internal

import (
	"os"
	"path/filepath"

	"github.com/aliamerj/aircup/server/apk/modules"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConfigDB() *gorm.DB {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot find the user's home directory: " + err.Error())
	}
	folderConf := filepath.Join(homeDir, ".aircup")

	dirErr := os.MkdirAll(folderConf, os.ModePerm)
	if dirErr != nil {
		panic("Failed to create config files")
	}
	dbPath := filepath.Join(folderConf, "config.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database : " + err.Error())
	}
	db.AutoMigrate(&modules.Account{}, &modules.Disk{})
	return db
}
