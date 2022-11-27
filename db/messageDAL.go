package db

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageLog struct {
	UserId  string              `bson:"userId"`
	Message string              `bson:"message"`
	Time    primitive.Timestamp `bson:"ts"`
}

func (m *MongoDB) InsertMessageLog(ctx context.Context, userId string, message string) {
	collection := m.Database.Collection(viper.GetString("mongo.MessageCollection"))
	collection.InsertOne(ctx, bson.D{{"ts", primitive.Timestamp{T: uint32(time.Now().Unix())}}, {"userId", userId}, {"message", message}})
}
