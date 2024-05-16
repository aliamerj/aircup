package modules

type Disk struct {
	DiskPath string `json:"disk_path" gorm:"primaryKey;not null;"`
}
