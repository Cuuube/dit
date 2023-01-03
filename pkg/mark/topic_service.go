package mark

import (
	"errors"
	"fmt"
	"time"

	"github.com/Cuuube/dit/pkg/dateutil"
	"gorm.io/gorm"
)

type TopicService interface {
	GetTopics() ([]*Topic, error)
	GetTopicByTopicKey(string) (*Topic, error)
	CreateTopic(string, string, string) error
	DeleteTopic(string) error
}

func NewTopicService() TopicService {
	s := SQLTopicService{
		db: GetSQLiteGorm(),
	}
	return &s
}

type SQLTopicService struct {
	db *gorm.DB
}

func (srv *SQLTopicService) GetTopics() ([]*Topic, error) {
	t := make([]*Topic, 0)
	err := srv.db.Find(&t).Order("create_time desc").Error
	return t, err
}

func (srv *SQLTopicService) GetTopicByTopicKey(key string) (*Topic, error) {
	t := Topic{}
	err := srv.db.Where("key=?", key).First(&t).Error
	return &t, err
}

func (srv *SQLTopicService) CreateTopic(key string, name string, typ string) error {
	now := time.Now()
	t := Topic{
		Key:        key,
		Name:       name,
		Type:       typ,
		CreateTime: dateutil.FormatDatetime(now),
		UpdateTime: dateutil.FormatDatetime(now),
	}

	err := srv.db.Table("topic").Create(&t).Error
	if errors.Is(err, gorm.ErrInvalidDB) {
		fmt.Println("invalid db!", err)
		return err
	}

	return err
}

func (srv *SQLTopicService) DeleteTopic(key string) error {
	return srv.db.Table("topic").Where("key=?", key).Delete(&Topic{}).Error
}
