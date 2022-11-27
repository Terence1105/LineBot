package db

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMessage interface {
	InsertMessageLog(context.Context, string, string)
}

type MongoDB struct {
	Database *mongo.Database
}

func NewDB() IMessage {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var messageDB IMessage = &MongoDB{Database: ConnectDB(ctx)}

	return messageDB
}

func ConnectDB(ctx context.Context) *mongo.Database {
	credential := options.Credential{
		AuthMechanism: viper.GetString("mongo.Credential.AuthMechanism"),
		AuthSource:    viper.GetString("mongo.Credential.AuthSource"),
		Username:      viper.GetString("mongo.Credential.Username"),
		Password:      viper.GetString("mongo.Credential.Password"),
	}
	options := options.Client().ApplyURI(viper.GetString("mongo.URI")).SetAuth(credential)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		panic(fmt.Errorf("fatal error mongodb connect fail: %w", err))
	}

	return client.Database(viper.GetString("mongo.DataBase"))
}
