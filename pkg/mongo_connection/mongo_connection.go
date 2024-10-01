package mongo_connection

import (
	"log"

	"github.com/LastDarkNes/go-dns/pkg/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(
		config.CommonConfig.MongoDBUrl,
	)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Panic(err)
	}

	return client
}
