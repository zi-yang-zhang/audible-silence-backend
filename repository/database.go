package repository

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BaseDatabaseModel struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string `gorm:"type:varchar(32)"`
	UpdatedBy string `gorm:"type:varchar(32)"`
}

func NewPhotoDao(db *gorm.DB) PhotoMetaDao {
	return &photoMetaDB{
		db,
	}
}

func NewMarkdownDao(db *gorm.DB) MarkdownDao {
	return &markdownDB{
		db,
	}
}
