package service

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"github.com/kataras/iris/hero"
	"github.com/zi-yang-zhang/audible-silence-backend/repository"
	"time"
)

type MarkdownService interface {
	GetMarkdownByID(id uint) (*Markdown, error)
	SaveMarkdown(title string, content string) (err error)
}

type markdownService struct {
	dao    repository.MarkdownDao
	logger *golog.Logger
}

/**
 * markdown object format
 * {
 *  url:
 *  title:
 *  id:
 *  likes:
 *  createDate:
 *  comments:
 * }
 *
 */
type Markdown struct {
	ID         uint      `json:"id"`
	Content    string    `json:"content"`
	Title      string    `json:"title"`
	Likes      uint      `json:"likes"`
	Comments   []string  `json:"comments"`
	CreateDate time.Time `json:"createDate"`
}

func InitMarkdownService(db *gorm.DB, logger *golog.Logger) {
	hero.Register(&markdownService{
		dao:    repository.NewMarkdownDao(db),
		logger: logger,
	})
}

func (s *markdownService) GetMarkdownByID(id uint) (*Markdown, error) {
	md, err := s.dao.GetMarkdownByID(id)
	if err != nil {
		return nil, err
	}
	return &Markdown{
		ID:         md.ID,
		Content:    md.Content,
		Title:      md.Title,
		Likes:      md.Likes,
		CreateDate: md.CreatedAt,
	}, err
}

func (s *markdownService) SaveMarkdown(title string, content string) (err error) {
	return s.dao.CreateMarkdown(title, content)
}
