package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Terence1105/LineBot/util"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"
)

func NewLineBot() *linebot.Client {
	bot, err := linebot.New(viper.GetString("linebot.secret"), viper.GetString("linebot.token"))
	if err != nil {
		panic(fmt.Errorf("fatal error line bot connect fail: %w", err))
	}
	return bot
}

func (s *Server) CallbackHandler(c *gin.Context) {

	events, err := s.bot.ParseRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				s.MessageLogDB.InsertMessageLog(ctx, event.Source.UserID, message.Text)
			case *linebot.ImageMessage:
				s.MessageLogDB.InsertMessageLog(ctx, event.Source.UserID, "Image type data")
			case *linebot.VideoMessage:
				s.MessageLogDB.InsertMessageLog(ctx, event.Source.UserID, "Video type data")
			case *linebot.AudioMessage:
				s.MessageLogDB.InsertMessageLog(ctx, event.Source.UserID, "Audio type data")
			case *linebot.FileMessage:
				s.MessageLogDB.InsertMessageLog(ctx, event.Source.UserID, "File type data")
			case *linebot.LocationMessage:
				s.MessageLogDB.InsertMessageLog(ctx, event.Source.UserID, "Location type data")
			case *linebot.StickerMessage:
				s.MessageLogDB.InsertMessageLog(ctx, event.Source.UserID, "Sticker type data")
			default:
				s.MessageLogDB.InsertMessageLog(ctx, event.Source.UserID, "unknown type message")
			}

			randomNum := util.Get_Random_Int(0, 100)
			random_Meme := s.memePictures.Data.Memes[randomNum].Url
			originalContentUrl := random_Meme
			previewImageUrl := random_Meme

			if _, err = s.bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(originalContentUrl, previewImageUrl)).Do(); err != nil {
				c.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
	}
}

func (s *Server) PushMessage(c *gin.Context) {
	_, err := s.bot.PushMessage(viper.GetString("linebot.userId"), linebot.NewTextMessage("A meme a day keeps the doctor away")).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}
