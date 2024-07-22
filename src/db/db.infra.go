package db

import (
	"fmt"

	"github.com/FlorVeneziano/gymbro-login-go/providers/envs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initializeDatabase() {
	var err error

	// * Get envs
	env := envs.GetEnvs()

	// * Setup mongo client
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(env.MONGO_HOST))
	if err != nil {
		fmt.Printf("Error connecting to database: %s", err)
	}

	// ? Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("Error connecting to database: %s", err)
	}

	// * Set database and collections
	db = client.Database(env.MONGO_DATABASE)
	usersCollection = db.Collection(usersCollectionName)

	fmt.Printf("Connected to database")
}

func terminateDatabase() {
	if client == nil {
		fmt.Println("Error disconnecting from database: client is nil")
	}

	err := client.Disconnect(ctx)
	if err != nil {
		fmt.Printf("Error disconnecting from database: %s", err)
	}

	fmt.Print("Disconnected from database")
}
