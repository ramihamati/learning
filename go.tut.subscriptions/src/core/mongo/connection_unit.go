package mongo

import "go.mongodb.org/mongo-driver/mongo"

type ConnectionUnit struct {
	_client mongo.Client
}

func New(
	connStringProvider ConnectionSettingsProvider) *ConnectionUnit {

	var connection_string = connStringProvider.Get()

}
