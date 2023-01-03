package mark

type MarkService interface {
	TopicService
	ContentService
}

type BaseMarkService struct {
	TopicService
	ContentService
}

func NewMarkService() MarkService {
	s := BaseMarkService{
		TopicService:   NewTopicService(),
		ContentService: NewContentService(),
	}
	return &s
}
