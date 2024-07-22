package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDatabase() *mongo.Client {
	once.Do(initializeDatabase)
	return client
}

func GetUsersCollection() *mongo.Collection {
	once.Do(initializeDatabase)
	return usersCollection
}

func DisconnectDatabase() {
	terminateDatabase()
}
