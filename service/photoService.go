package service

import (
	"github.com/disintegration/imaging"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"github.com/kataras/iris/hero"
	"github.com/zi-yang-zhang/audible-silence-backend/repository"
	"io"
	"time"
)

type PhotoService interface {
	GetPhotoList() (*[]PhotoMeta, error)
	GetPhoto(uint) (*PhotoMeta, error)
	UploadPhoto(string, io.Reader) (err error)
}

type photoService struct {
	dao                 repository.PhotoMetaDao
	photoLocationConfig *PhotoLocationConfig
	logger              *golog.Logger
}

/**
 * gallery list object format
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
type PhotoMeta struct {
	ID           uint      `json:"id"`
	URL          string    `json:"url"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	Title        string    `json:"title"`
	Likes        uint      `json:"likes"`
	Comments     []string  `json:"comments"`
	CreateDate   time.Time `json:"createDate"`
}

type comment struct {
	UserName   string `json:"userName"`
	Comment    string `json:"comment"`
	CreateDate string `json:"createDate"`
	ModifyDate string `json:"modifyDate"`
}

type PhotoLocationConfig struct {
	Photo     string
	ThumbNail string
}

func InitPhotoService(db *gorm.DB, photoLocationConfig *PhotoLocationConfig, logger *golog.Logger) {
	hero.Register(&photoService{
		dao:                 repository.NewPhotoDao(db),
		photoLocationConfig: photoLocationConfig,
		logger:              logger,
	})
}

func (s *photoService) GetPhotoList() (*[]PhotoMeta, error) {
	result, err := s.dao.GetAllPhotoMeta()
	if err != nil {
		return nil, err
	}
	var response []PhotoMeta
	for _, photo := range *result {
		response = append(response, PhotoMeta{
			ID:         photo.ID,
			URL:        photo.URL,
			Title:      photo.Title,
			Likes:      photo.Likes,
			CreateDate: photo.CreatedAt,
		})
	}
	return &response, nil
}

func (s *photoService) GetPhoto(id uint) (*PhotoMeta, error) {
	photo, err := s.dao.GetPhotoMeta(id)
	if err != nil {
		return nil, err
	}
	return &PhotoMeta{
		ID:           photo.ID,
		URL:          photo.URL,
		ThumbnailURL: photo.ThumbnailURL,
		Title:        photo.Title,
		Likes:        photo.Likes,
		CreateDate:   photo.CreatedAt,
	}, nil
}
func (s *photoService) UploadPhoto(title string, reader io.Reader) (err error) {
	return nil
}
func (s *photoService) persistImage(fileName string, reader io.Reader) (err error) {
	srcImage, err := imaging.Decode(reader)
	if err != nil {
		s.logger.Fatalf("failed to open image: %v", err)
		return
	}
	thumbnail := imaging.Thumbnail(srcImage, 100, 100, imaging.Lanczos)
	// Save the resulting image as JPEG.
	err = imaging.Save(srcImage, s.photoLocationConfig.Photo+fileName+".jpg")
	if err != nil {
		s.logger.Fatalf("failed to save image: %v", err)
		return
	}
	err = imaging.Save(thumbnail, s.photoLocationConfig.ThumbNail+fileName+"-thumbnail.jpg")
	if err != nil {
		s.logger.Fatalf("failed to save image thumnail: %v", err)
		return
	}
	return
}
