package db

import (
	"context"
	"fmt"
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

func (m *MongoDB) GetMessageLogByUserId(ctx context.Context, userId string) []MessageLog {
	collection := m.Database.Collection(viper.GetString("mongo.MessageCollection"))

	filter := bson.D{{"userId", userId}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("find fail")
	}

	var MessageLogs []MessageLog
	err = cursor.All(ctx, &MessageLogs)
	if err != nil {
		fmt.Println("cursor fail")
	}

	defer cursor.Close(ctx)

	return MessageLogs
}

func (m *MongoDB) GetAllMessageLog(ctx context.Context) []MessageLog {
	collection := m.Database.Collection(viper.GetString("mongo.MessageCollection"))

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("find fail")
	}

	var MessageLogs []MessageLog
	err = cursor.All(ctx, &MessageLogs)
	if err != nil {
		fmt.Println("cursor fail")
	}

	defer cursor.Close(ctx)

	return MessageLogs
}
