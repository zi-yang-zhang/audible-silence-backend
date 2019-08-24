package repository

import (
	"github.com/jinzhu/gorm"
)

type MarkdownDao interface {
	CreateMarkdown(title string, url string) (err error)
	GetMarkdownByID(id uint) (result *Markdown, err error)
}
type Markdown struct {
	BaseDatabaseModel
	Content string `gorm:"type:text;not null"`
	Title   string `gorm:"type:varchar(128);not null"`
	Likes   uint
}

type markdownDB struct {
	*gorm.DB
}

func (db *markdownDB) CreateMarkdown(title string, content string) (err error) {
	newMarkdown := &Markdown{
		Content: content,
		Title:   title,
	}
	return db.Create(newMarkdown).Error
}
func (db *markdownDB) GetMarkdownByID(id uint) (result *Markdown, err error) {
	result = &Markdown{}
	err = db.Where("id = ?", id).First(&result).Error
	return
}
