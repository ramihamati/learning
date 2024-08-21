package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"strings"
	"subscriptions/core/errors"
	"subscriptions/core/helpers"
	"time"
)

const (
	ConnectionFailed = "Db:Conn:Fail"
)

type ConnectionManager struct {
	Client  *mongo.Client
	Context context.Context
}

func NewManager(setting ConnectionSettings) ConnectionManager {
	var _context = context.TODO()
	context.WithTimeout(_context, time.Second*30)

	client, err := mongo.Connect(
		_context,
		options.Client().ApplyURI(getConnectionString(setting)))

	if err != nil {
		panic(errors.CodeError{
			Code:    ConnectionFailed,
			Message: "Failed To Create The Manager",
		})
	}

	return ConnectionManager{
		Client:  client,
		Context: _context,
	}
}

func (c ConnectionManager) GetCollection(collectionName string, databaseName string) {
	var collectionOptions = &options.CollectionOptions{
		ReadConcern: &readconcern.ReadConcern{
			Level: "majority",
		},
		WriteConcern: &writeconcern.WriteConcern{
			W: "majority",
		},
		ReadPreference: readpref.PrimaryPreferred(),
	}
	c.Client.
		Database(databaseName).
		Collection(collectionName, collectionOptions)
}

func (c ConnectionManager) GetDatabase(databaseName string) {
	c.Client.Database(databaseName)
}

func (c ConnectionManager) HasCollection(collectionName string, databaseName string) (bool, error) {
	var database = c.Client.Database(databaseName)
	var options = &options.ListCollectionsOptions{}

	collections, err := database.ListCollectionNames(c.Context, bson.D{}, options)

	if err != nil {
		return false, err
	}

	for _, a := range collections {
		if a == collectionName {
			return true, nil
		}
	}
	return false, nil
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
