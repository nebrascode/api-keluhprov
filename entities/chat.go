package entities

import (
	"time"
)

// Room represents a chat room where users and admins can communicate
type Room struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Messages  []Message `gorm:"foreignKey:RoomID"`
}

// Message represents a chat message within a room
type Message struct {
	ID         int       `gorm:"primaryKey"`
	RoomID     int       `gorm:"not null"`
	SenderID   int       `gorm:"not null"`
	SenderType string    `gorm:"type:ENUM('user', 'admin');not null"`
	Message    string    `gorm:"type:text;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

// ChatRepositoryInterface defines methods for interacting with chat data
type ChatRepositoryInterface interface {
	CreateChat(chat *Message) error
	GetChatsByUserID(userID int) ([]Message, error)
	GetChatsByAdminID(adminID int) ([]Message, error)
	GetChatsBetweenUserAndAdmin(userID, adminID int) ([]Message, error)
	GetAllRooms() ([]Room, error)
	CreateRoom(name string) (*Room, error)
	GetMessagesByRoomID(roomID int) ([]Message, error)
}

// ChatUseCaseInterface defines the business logic methods for chat interactions
type ChatUseCaseInterface interface {
	SendMessage(chat *Message) error
	GetUserChats(userID int) ([]Message, error)
	GetAdminChats(adminID int) ([]Message, error)
	GetConversation(userID, adminID int) ([]Message, error)
	GetAllChatsByUser(userID int) ([]Message, error)
	GetAllRooms() ([]Room, error)
	CreateRoom(name string) (*Room, error)
	GetMessagesByRoomID(roomID int) ([]Message, error)
}
