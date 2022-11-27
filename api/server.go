package api

import (
	"github.com/Terence1105/LineBot/db"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Server struct {
	router       *gin.Engine
	bot          *linebot.Client
	MessageLogDB db.IMessage
}

func NewServer(mp MemePictures, bot *linebot.Client, messageLogDB db.IMessage) *Server {

	server := &Server{
		bot:          bot,
		MessageLogDB: messageLogDB,
	}
	server.setUpRouter()
	return server
}

func (s *Server) setUpRouter() {
	router := gin.Default()
	router.POST("/callback", s.CallbackHandler)
	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
