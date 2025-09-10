package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Database, err error) {

	var mongoDBURL string
	var isAuth bool

	if username == "" && password == "" {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}

	//Connect

	clientOptions := options.Client().ApplyURI(mongoDBURL)

	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOptions.SetAuth(options.Credential{
			Username:   username,
			Password:   password,
			AuthSource: authDB, // Specifies the database where the user is authenticated
		})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Cannot to connect to mongoDB due to error: %w ", err)
	}

	//Ping
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Cannot to ping to mongoDB due to error: %w ", err)
	}

	return client.Database(database), nil
}
