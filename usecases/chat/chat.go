package chat

import (
	"e-complaint-api/entities"
)

type chatUseCase struct {
	chatRepo entities.ChatRepositoryInterface
}

// SendMessage adds a new message to a room
func (uc *chatUseCase) SendMessage(chat *entities.Message) error {
	return uc.chatRepo.CreateChat(chat)
}

// NewChatUseCase initializes a new chat use case
func NewChatUseCase(chatRepo entities.ChatRepositoryInterface) entities.ChatUseCaseInterface {
	return &chatUseCase{chatRepo: chatRepo}
}

// CreateRoom creates a new chat room
func (uc *chatUseCase) CreateRoom(name string) (*entities.Room, error) {
	return uc.chatRepo.CreateRoom(name)
}

// GetAllRooms retrieves all chat rooms with their associated messages
func (uc *chatUseCase) GetAllRooms() ([]entities.Room, error) {
	return uc.chatRepo.GetAllRooms()
}

// GetMessagesByRoomID retrieves all messages for a specific room
func (uc *chatUseCase) GetMessagesByRoomID(roomID int) ([]entities.Message, error) {
	return uc.chatRepo.GetMessagesByRoomID(roomID)
}

// GetUserChats retrieves all messages sent by a specific user
func (uc *chatUseCase) GetUserChats(userID int) ([]entities.Message, error) {
	return uc.chatRepo.GetChatsByUserID(userID)
}

// GetAdminChats retrieves all messages sent by a specific admin
func (uc *chatUseCase) GetAdminChats(adminID int) ([]entities.Message, error) {
	return uc.chatRepo.GetChatsByAdminID(adminID)
}

// GetConversation retrieves all messages exchanged between a user and an admin
func (uc *chatUseCase) GetConversation(userID, adminID int) ([]entities.Message, error) {
	return uc.chatRepo.GetChatsBetweenUserAndAdmin(userID, adminID)
}

// GetAllChatsByUser retrieves all chats where the user is involved
func (uc *chatUseCase) GetAllChatsByUser(userID int) ([]entities.Message, error) {
	messages, err := uc.chatRepo.GetChatsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
