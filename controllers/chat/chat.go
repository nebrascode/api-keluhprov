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
	// Data model untuk body input
	var request struct {
		Name string `json:"name" form:"name"`
	}

	// Parsing data dari body (mendukung JSON, form-data, x-www-form-urlencoded)
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if request.Name == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Room name is required"})
	}

	// Buat room menggunakan data dari body
	room, err := c.chatUsecase.CreateRoom(request.Name)
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

	// Konversi room ID ke integer
	roomIDInt, err := strconv.Atoi(roomID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Room ID"})
	}

	// Data model untuk body input
	var request struct {
		SenderID   int    `json:"senderID" form:"senderID"`
		SenderType string `json:"senderType" form:"senderType"`
		Message    string `json:"message" form:"message"`
	}

	// Parsing data dari body (mendukung JSON, form-data, x-www-form-urlencoded)
	if err := ctx.Bind(&request); err != nil {
		log.Println("Invalid input data:", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if request.SenderID <= 0 || request.SenderType == "" || request.Message == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
	}

	// Buat pesan
	msg := entities.Message{
		RoomID:     roomIDInt,
		SenderID:   request.SenderID,
		SenderType: request.SenderType,
		Message:    request.Message,
	}

	// Kirim pesan menggunakan usecase
	err = c.chatUsecase.SendMessage(&msg)
	if err != nil {
		log.Println("Failed to send message:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send message"})
	}

	return ctx.JSON(http.StatusOK, msg)
}

func (c *ChatController) GetMessagesByRoomID(ctx echo.Context) error {
	// Ambil parameter Room ID dari URL
	roomID := ctx.Param("room-id")
	if roomID == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Room ID is required"})
	}

	// Konversi Room ID ke integer
	roomIDInt, err := strconv.Atoi(roomID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Room ID"})
	}

	// Ambil pesan-pesan berdasarkan Room ID
	messages, err := c.chatUsecase.GetMessagesByRoomID(roomIDInt)
	if err != nil {
		log.Println("Failed to fetch messages:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch messages"})
	}

	// Kirimkan data pesan ke klien
	return ctx.JSON(http.StatusOK, messages)
}

func (c *ChatController) GetRoomByID(ctx echo.Context) error {
	// Ambil parameter room ID dari URL
	ID := ctx.Param("id")
	if ID == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Room ID is required"})
	}

	// Konversi room ID ke integer
	IDInt, err := strconv.Atoi(ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Room ID"})
	}

	// Ambil detail room dari usecase
	room, err := c.chatUsecase.GetRoomByID(IDInt)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Room not found"})
	}

	// Kirimkan data room ke klien
	return ctx.JSON(http.StatusOK, room)
}