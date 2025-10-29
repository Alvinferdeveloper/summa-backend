package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/Alvinferdeveloper/summa-backend/dto" // Importacion del DTO
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 1024                // Maximum message size allowed from peer.
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// In production, should be a proper origin check.
		return true
	},
}

type IncomingMessage struct {
	Type           string    `json:"type"` // "chat" o cualquier otro tipo que venga del frontend
	ConversationID uint      `json:"conversation_id"`
	RecipientID    uuid.UUID `json:"recipient_id"`
	RecipientType  string    `json:"recipient_type"`
	Content        string    `json:"content"`
}

type Client struct {
	Hub      *Hub
	conn     *websocket.Conn
	send     chan []byte
	UserID   uuid.UUID
	UserType string
}

type PrivateMessage struct {
	RecipientID   uuid.UUID
	RecipientType string
	Message       []byte
}

// Hub maintains the set of active clients and broadcasts messages to them.
type Hub struct {
	clients    map[string]map[uuid.UUID]*Client // [userType][userID]
	private    chan *PrivateMessage
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		private:    make(chan *PrivateMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]map[uuid.UUID]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			if h.clients[client.UserType] == nil {
				h.clients[client.UserType] = make(map[uuid.UUID]*Client)
			}
			h.clients[client.UserType][client.UserID] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.UserType][client.UserID]; ok {
				delete(h.clients[client.UserType], client.UserID)
				close(client.send)
				if len(h.clients[client.UserType]) == 0 {
					delete(h.clients, client.UserType)
				}
			}
		case privateMessage := <-h.private:
			if recipientClient, ok := h.clients[privateMessage.RecipientType][privateMessage.RecipientID]; ok {
				select {
				case recipientClient.send <- privateMessage.Message:
				default:
					close(recipientClient.send)
					delete(h.clients[privateMessage.RecipientType], privateMessage.RecipientID)
				}
			} else {
				// Recipient is offline, publish to RabbitMQ
				log.Printf("Recipient %s %d is offline. Publishing to RabbitMQ...", privateMessage.RecipientType, privateMessage.RecipientID)
				if err := services.PublishChatMessageNotification(privateMessage.Message); err != nil {
					log.Printf("Failed to publish chat message to RabbitMQ: %v", err)
				}
			}
		}
	}
}

func (h *Hub) BroadcastToUser(userID string, payload []byte) {
	// Envolver el payload en un WebSocketMessage con tipo "notification"
	wsMessage := dto.WebSocketMessage{
		Type:    "notification",
		Payload: json.RawMessage(payload), // Usar RawMessage para el payload ya marshalizado
	}
	message, err := json.Marshal(wsMessage)
	if err != nil {
		log.Printf("Error marshalling WebSocket notification message: %v", err)
		return
	}

	if clients, ok := h.clients["job_seeker"]; ok {
		if client, ok := clients[uuid.MustParse(userID)]; ok {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients["job_seeker"], client.UserID)
			}
			return
		}
	}

	if clients, ok := h.clients["employer"]; ok {
		if client, ok := clients[uuid.MustParse(userID)]; ok {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients["employer"], client.UserID)
			}
			return
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.Hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, messageBytes, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var wsIncomingMessage dto.WebSocketMessage
		if err := json.Unmarshal(messageBytes, &wsIncomingMessage); err != nil {
			log.Printf("error unmarshalling incoming WebSocket message: %v", err)
			continue
		}

		payloadBytes, err := json.Marshal(wsIncomingMessage.Payload)
		if err != nil {
			log.Printf("error marshalling payload to bytes: %v", err)
			continue
		}

		if wsIncomingMessage.Type == "chat" {
			var msg IncomingMessage
			if err := json.Unmarshal(payloadBytes, &msg); err != nil {
				log.Printf("error unmarshalling chat payload: %v", err)
				continue
			}

			dbMessage, err := services.CreateMessage(msg.ConversationID, c.UserID, c.UserType, msg.RecipientID, msg.RecipientType, msg.Content)
			if err != nil {
				log.Printf("error saving message to db: %v", err)
				continue
			}

			outgoingChat := dto.WebSocketMessage{
				Type:    "chat",
				Payload: dbMessage,
			}
			outgoingBytes, err := json.Marshal(outgoingChat)
			if err != nil {
				log.Printf("error marshalling outgoing chat message: %v", err)
				continue
			}

			privateMessage := &PrivateMessage{
				RecipientID:   msg.RecipientID,
				RecipientType: msg.RecipientType,
				Message:       outgoingBytes,
			}
			c.Hub.private <- privateMessage
		} else {
			log.Printf("Received unhandled WebSocket message type: %s", wsIncomingMessage.Type)
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, userID uuid.UUID, userType string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{Hub: hub, conn: conn, send: make(chan []byte, 256), UserID: userID, UserType: userType}
	hub.register <- client

	go client.writePump()
	go client.readPump()
}
