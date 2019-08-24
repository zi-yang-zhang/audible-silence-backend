package repository

import (
	"github.com/jinzhu/gorm"
)

type PhotoMetaDao interface {
	GetAllPhotoMeta() (results *[]PhotoMeta, err error)
	GetPhotoMeta(id uint) (result *PhotoMeta, err error)
	InsertPhotoMeta(title string, url string, thumbnailURL string) (err error)
}

type photoMetaDB struct {
	*gorm.DB
}

// PhotoMeta is the db model of PhotoMeta
type PhotoMeta struct {
	BaseDatabaseModel
	URL          string `gorm:"type:varchar(1024);not null"`
	ThumbnailURL string `gorm:"type:varchar(1024);not null"`
	Title        string `gorm:"type:varchar(128);not null"`
	Likes        uint
}

func (db *photoMetaDB) InsertPhotoMeta(title string, url string, thumbnailURL string) (err error) {
	newPhoto := &PhotoMeta{
		URL:          url,
		Title:        title,
		ThumbnailURL: thumbnailURL,
	}
	err = db.Create(newPhoto).Error
	return
}

func (db *photoMetaDB) GetAllPhotoMeta() (results *[]PhotoMeta, err error) {
	results = &[]PhotoMeta{}
	err = db.Find(&results).Error
	return
}

func (db *photoMetaDB) GetPhotoMeta(id uint) (result *PhotoMeta, err error) {
	result = &PhotoMeta{}
	err = db.Where("id = ?", id).First(&result).Error
	return
}
