package chat

import (
	"e-complaint-api/entities"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type ChatController struct {
	chatUsecase entities.ChatUseCaseInterface
}

func NewChatController(chatUsecase entities.ChatUseCaseInterface) *ChatController {
	return &ChatController{chatUsecase: chatUsecase}
}

func (c *ChatController) CreateRoom(ctx echo.Context) error {
	name := ctx.QueryParam("name")
	if name == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Room name is required"})
	}

	room, err := c.chatUsecase.CreateRoom(name)
	if err != nil {
		log.Println("Failed to create room:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create room"})
	}

	return ctx.JSON(http.StatusOK, room)
}

func (c *ChatController) GetAllRooms(ctx echo.Context) error {
	rooms, err := c.chatUsecase.GetAllRooms()
	if err != nil {
		log.Println("Failed to fetch rooms:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch rooms"})
	}

	return ctx.JSON(http.StatusOK, rooms)
}

func (c *ChatController) SendMessage(ctx echo.Context) error {
	roomID := ctx.Param("room-id")
	if roomID == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Room ID is required"})
	}

	// Ubah roomID ke integer
	roomIDInt, err := strconv.Atoi(roomID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Room ID"})
	}

	var request struct {
		SenderID   int    `json:"senderID" form:"senderID"`
		SenderType string `json:"senderType" form:"senderType"`
		Message    string `json:"message" form:"message"`
	}

	if err := ctx.Bind(&request); err != nil {
		log.Println("Invalid input data:", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if request.SenderID <= 0 || request.SenderType == "" || request.Message == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
	}

	msg := entities.Message{
		RoomID:     roomIDInt,
		SenderID:   request.SenderID,
		SenderType: request.SenderType,
		Message:    request.Message,
	}

	err = c.chatUsecase.SendMessage(&msg)
	if err != nil {
		log.Println("Failed to send message:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send message"})
	}

	return ctx.JSON(http.StatusOK, msg)
}
