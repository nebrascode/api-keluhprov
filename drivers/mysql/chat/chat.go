package chat

import (
	"e-complaint-api/entities"
	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

// CreateChat saves a new message to the database
func (r *chatRepository) CreateChat(chat *entities.Message) error {
	return r.db.Create(chat).Error
}

// GetChatsByUserID retrieves all messages sent by a specific user
func (r *chatRepository) GetChatsByUserID(userID int) ([]entities.Message, error) {
	var messages []entities.Message
	err := r.db.Where("sender_id = ? AND sender_type = ?", userID, "user").Order("created_at ASC").Find(&messages).Error
	return messages, err
}

// GetChatsByAdminID retrieves all messages sent by a specific admin
func (r *chatRepository) GetChatsByAdminID(adminID int) ([]entities.Message, error) {
	var messages []entities.Message
	err := r.db.Where("sender_id = ? AND sender_type = ?", adminID, "admin").Order("created_at ASC").Find(&messages).Error
	return messages, err
}

// GetChatsBetweenUserAndAdmin retrieves all messages exchanged between a user and an admin
func (r *chatRepository) GetChatsBetweenUserAndAdmin(userID, adminID int) ([]entities.Message, error) {
	var messages []entities.Message
	err := r.db.Where("(sender_id = ? AND sender_type = 'user') OR (sender_id = ? AND sender_type = 'admin')", userID, adminID).
		Order("created_at ASC").Find(&messages).Error
	return messages, err
}

// CreateRoom creates a new chat room
func (r *chatRepository) CreateRoom(name string) (*entities.Room, error) {
	room := &entities.Room{Name: name}
	if err := r.db.Create(room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

// CreateMessage saves a new message in a specific room
func (r *chatRepository) CreateMessage(message *entities.Message) error {
	return r.db.Create(message).Error
}

// GetAllRooms retrieves all chat rooms with their associated messages
func (r *chatRepository) GetAllRooms() ([]entities.Room, error) {
	var rooms []entities.Room
	err := r.db.Preload("Messages").Find(&rooms).Error
	return rooms, err
}

// GetMessagesByRoomID retrieves all messages for a specific room
func (r *chatRepository) GetMessagesByRoomID(roomID int) ([]entities.Message, error) {
	var messages []entities.Message
	err := r.db.Where("room_id = ?", roomID).Order("created_at ASC").Find(&messages).Error
	return messages, err
}

// NewChatRepository initializes a new chat repository
func NewChatRepository(db *gorm.DB) entities.ChatRepositoryInterface {
	return &chatRepository{db: db}
}

// GetRoomByID retrieves a room by its ID
func (r *chatRepository) GetRoomByID(ID int) (*entities.Room, error) {
	var room entities.Room
	err := r.db.Preload("Messages").First(&room, ID).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}