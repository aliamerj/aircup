package controllers

import (
	"errors"
	"net/http"

	"github.com/aliamerj/aircup/server/apk/modules"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetConfig(c echo.Context, db *gorm.DB) error {
	var account modules.Account
	var disks []modules.Disk

	if err := db.First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusOK, map[string]string{"type": "NO_ACCOUNT"})
		}
		return err
	}
	if err := db.Find(&disks).Error; err != nil {
		return err
	}

	payload := struct {
		Account modules.Account `json:"account"`
		Disks   []modules.Disk  `json:"disks"`
		Type    string          `json:"type"`
	}{
		Account: account,
		Disks:   disks,
		Type:    "OK",
	}

	return c.JSON(http.StatusOK, payload)

}
