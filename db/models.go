package db

import (
	"time"
)

// User represents the user table
type User struct {
	ID        int64     `xorm:"pk autoincr 'id'" json:"id"`
	Nickname  string    `xorm:"varchar(100) notnull 'nickname'" json:"nickname"`
	Email     string    `xorm:"varchar(255) notnull unique 'email'" json:"email"`
	CreatedAt time.Time `xorm:"created 'created_at'" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated 'updated_at'" json:"updatedAt"`
}

// TableName returns the table name for User
func (User) TableName() string {
	return "user"
}

// UserDevice represents the user_device table
type UserDevice struct {
	ID        int64     `xorm:"pk autoincr 'id'" json:"id"`
	UserID    int64     `xorm:"notnull 'user_id'" json:"userId"`
	DeviceID  string    `xorm:"varchar(255) notnull 'device_id'" json:"deviceId"`
	CreatedAt time.Time `xorm:"created 'created_at'" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated 'updated_at'" json:"updatedAt"`
}

// TableName returns the table name for UserDevice
func (UserDevice) TableName() string {
	return "user_device"
}

type Image struct {
	ID             int64     `xorm:"'id' pk autoincr" json:"id"`
	Filename       string    `xorm:"'filename' varchar(255) notnull" json:"filename"`
	OriginFilename string    `xorm:"'origin_filename' varchar(255) notnull" json:"originFilename"`
	FileExtension  string    `xorm:"'file_extension' varchar(10) notnull" json:"fileExtension"`
	Bucket         string    `xorm:"'bucket' varchar(255) notnull" json:"bucket"`
	ObjectKey      string    `xorm:"'object_key' varchar(255) notnull" json:"objectKey"`
	Uploaded       bool      `xorm:"'uploaded' tinyint(1) notnull default(0)" json:"uploaded"`
	LabelDetected  bool      `xorm:"'label_detected' tinyint(1) notnull default(0)" json:"labelDetected"`
	TextDetected   bool      `xorm:"'text_detected' tinyint(1) notnull default(0)" json:"textDetected"`
	CreatedAt      time.Time `xorm:"'created_at' created" json:"createdAt"`
	UpdatedAt      time.Time `xorm:"'updated_at' updated" json:"updatedAt"`
}

func (Image) TableName() string {
	return "image"
}

type Label struct {
	ID        int64     `xorm:"pk autoincr 'id'" json:"id"`
	Name      string    `xorm:"varchar(255) notnull unique 'name'" json:"name"`
	CreatedAt time.Time `xorm:"created 'created_at'" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated 'updated_at'" json:"updatedAt"`
}

func (Label) TableName() string {
	return "label"
}

type ImageLabel struct {
	ID        int64     `xorm:"pk autoincr 'id'" json:"id"`
	ImageID   int64     `xorm:"notnull 'image_id' unique(image_label)" json:"imageId"`
	LabelID   int64     `xorm:"notnull 'label_id' unique(image_label)" json:"labelId"`
	CreatedAt time.Time `xorm:"created 'created_at'" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated 'updated_at'" json:"updatedAt"`
}

func (ImageLabel) TableName() string {
	return "image_label"
}

type TextKeyword struct {
	ID        int64     `xorm:"pk autoincr 'id'" json:"id"`
	Keyword   string    `xorm:"varchar(255) notnull unique 'keyword'" json:"keyword"`
	CreatedAt time.Time `xorm:"created 'created_at'" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated 'updated_at'" json:"updatedAt"`
}

func (TextKeyword) TableName() string {
	return "text_keyword"
}

type ImageTextKeyword struct {
	ID            int64     `xorm:"pk autoincr 'id'" json:"id"`
	ImageID       int64     `xorm:"notnull 'image_id' unique(image_text_keyword)" json:"imageId"`
	TextKeywordId int64     `xorm:"notnull 'text_keyword_id' unique(image_text_keyword)" json:"text_keyword_id"`
	CreatedAt     time.Time `xorm:"created 'created_at'" json:"createdAt"`
	UpdatedAt     time.Time `xorm:"updated 'updated_at'" json:"updatedAt"`
}

func (ImageTextKeyword) TableName() string {
	return "image_text_keyword"
}

// Chat represents the chat table for chat sessions
type Chat struct {
	ID        int64     `xorm:"pk autoincr 'id'" json:"id"`
	UserID    int64     `xorm:"notnull 'user_id'" json:"userId"`
	Title     string    `xorm:"varchar(255) 'title'" json:"title"`
	BotName   string    `xorm:"varchar(100) 'bot_name'" json:"botName"`
	BotId     string    `xorm:"varchar(100) 'bot_id'" json:"botId"`
	BotAlias  string    `xorm:"varchar(100) 'bot_alias'" json:"botAlias"`
	LocaleId  string    `xorm:"varchar(20) 'locale_id'" json:"localeId"`
	SessionId string    `xorm:"varchar(255) 'session_id'" json:"sessionId"`
	CreatedAt time.Time `xorm:"created 'created_at'" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated 'updated_at'" json:"updatedAt"`
}

func (Chat) TableName() string {
	return "chat"
}

// ChatMessage represents the chat_message table for individual messages
type ChatMessage struct {
	ID        int64     `xorm:"pk autoincr 'id'" json:"id"`
	ChatID    int64     `xorm:"notnull 'chat_id'" json:"chatId"`
	Content   string    `xorm:"text 'content'" json:"content"`
	IsUser    bool      `xorm:"tinyint(1) notnull 'is_user'" json:"isUser"`
	Intent    string    `xorm:"varchar(100) 'intent'" json:"intent"`
	CreatedAt time.Time `xorm:"created 'created_at'" json:"createdAt"`
}

func (ChatMessage) TableName() string {
	return "chat_message"
}
