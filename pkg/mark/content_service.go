package mark

import (
	"errors"
	"fmt"
	"time"

	"github.com/Cuuube/dit/pkg/utils/dateutil"
	"gorm.io/gorm"
)

func NewContentService() ContentService {
	s := SQLContentService{
		db: GetSQLiteGorm(),
	}
	return &s
}

type ContentService interface {
	GetContents(string) ([]*Content, error)
	GetContentByTopicAndKey(string, string) (*Content, error)
	CreateContent(string, string, string) error
	UpdateContent(string, string, string) error
	UpsertContent(string, string, string) error
	DeleteContent(string, string) error
}

type SQLContentService struct {
	db *gorm.DB
}

func (srv *SQLContentService) GetContents(topicKey string) ([]*Content, error) {
	c := make([]*Content, 0)
	err := srv.db.Where("topic=?", topicKey).Order("create_time desc").Find(&c).Error
	return c, err
}

func (srv *SQLContentService) GetContentByTopicAndKey(topicKey string, key string) (*Content, error) {
	c := Content{}
	err := srv.db.Table("content").Where("topic=? and key=?", topicKey, key).Find(&c).Error
	return &c, err
}

func (srv *SQLContentService) CreateContent(topicKey, key, data string) error {
	now := time.Now()
	t := Content{
		Topic:      topicKey,
		Data:       data,
		Key:        key,
		CreateTime: dateutil.FormatDatetime(now),
		UpdateTime: dateutil.FormatDatetime(now),
	}

	err := srv.db.Table("content").Create(&t).Error
	if errors.Is(err, gorm.ErrInvalidDB) {
		fmt.Println("invalid db!", err)
		return err
	}

	return err
}

func (srv *SQLContentService) UpdateContent(topicKey string, key string, data string) error {
	updateData := map[string]interface{}{
		"data":        data,
		"update_time": dateutil.FormatDatetime(time.Now()),
	}
	return srv.db.Table("content").Where("topic=? and key=?", topicKey, key).Updates(updateData).Error
}

func (srv *SQLContentService) UpsertContent(topicKey string, key string, data string) error {
	res, _ := srv.GetContentByTopicAndKey(topicKey, key)
	if res.Key == "" {
		return srv.CreateContent(topicKey, key, data)
	}
	return srv.UpdateContent(topicKey, key, data)
}

func (srv *SQLContentService) DeleteContent(topicKey string, key string) error {

	return srv.db.Table("content").Where("topic=? and key=?", topicKey, key).Delete(&Content{}).Error
}
