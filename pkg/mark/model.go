package mark

const (
	TopicTypeKV       = "kv"       // 无限制，根据key和value存取
	TopicTypeDaliy    = "daily"    // 每日一条
	TopicTypeMinutely = "minutely" // 每分钟一条
	TopicTypeSecondly = "secondly" // 每秒一条
)

// 话题
type Topic struct {
	Key        string `gorm:"column:key;primary_key"`
	Name       string `gorm:"column:name"`
	Type       string `gorm:"column:type"`
	CreateTime string `gorm:"column:create_time"`
	UpdateTime string `gorm:"column:update_time"`
}

func (Topic) TableName() string {
	return "topic"
}

type Content struct {
	ID         uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Topic      string `gorm:"column:topic;NOT NULL"`
	Key        string `gorm:"column:key;NOT NULL"`
	Data       string `gorm:"column:data"`
	CreateTime string `gorm:"column:create_time"`
	UpdateTime string `gorm:"column:update_time"`
}

func (Content) TableName() string {
	return "content"
}

// create table topic (key text primary_key,name text,type text,create_time text,update_time text);
// create table content (id integer primary_key AUTO_INCREMENT,topic text,data text,key text, create_time text,update_time text);
