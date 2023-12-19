package bootstrap

import (
	"context"
	"time"

	"srating/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
)

func NewMongoDatabase(env *Env) (client *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := env.DatabaseURL

	clientOptions := options.Client().ApplyURI(mongodbURI)

	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = pingDatabase(ctx, client)
	if err != nil {
		err := client.Disconnect(context.TODO()) // Disconnect the client on error
		if err != nil {
			utils.LogFatal(err, "Failed to disconnect MongoDB")
		}
		return nil, err
	}

	return client, nil
}

func pingDatabase(ctx context.Context, client *mongo.Client) error {
	pingReadPref := readpref.Primary()
	err := client.Ping(ctx, pingReadPref)
	if err != nil {
		return err
	}
	return nil
}

func CloseMongoDBConnection(client driver.Disconnector) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		utils.LogFatal(err, "Failed to disconnect MongoDB")
	}

	utils.LogInfo("MongoDB connection closed")
}
