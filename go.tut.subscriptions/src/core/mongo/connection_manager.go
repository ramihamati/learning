package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"subscriptions/core/errors"
	"subscriptions/core/helpers"
	"time"
)

const (
	ConnectionFailed = "Db:Conn:Fail"
)

type ConnectionManager struct {
	Client *mongo.Client
}

func NewManager(setting ConnectionSettings) ConnectionManager {
	con := context.TODO()
	context.WithTimeout(con, time.Second*30)

	client, err := mongo.Connect(
		con,
		options.Client().ApplyURI(getConnectionString(setting)))

	if err != nil {
		panic(errors.CodeError{
			Code:    ConnectionFailed,
			Message: "Failed To Create The Manager",
		})
	}

	return ConnectionManager{
		Client: client,
	}
}

func getConnectionString(settings ConnectionSettings) string {
	var computed = settings.connectionString

	if strings.Contains(computed, "?") {
		computed += fmt.Sprintf("&connectTimeoutMS=%d", settings.timeout.Milliseconds())
	} else {
		if !helpers.EndsWith(computed, "/") {
			computed += "/"
		}
		computed += fmt.Sprintf("?connectTimeoutMS=%d", settings.timeout.Milliseconds())
	}

	if !strings.Contains(computed, "keepAlive") {
		computed += "&keepAlive=true"
	}
	if !strings.Contains(computed, "autoReconnect") {
		computed += "&autoReconnect=true"
	}
	if !strings.Contains(computed, "socketTimeoutMS") {
		computed += fmt.Sprintf("&socketTimeoutMS=%d", settings.timeout.Milliseconds())
	}

	return computed
}
