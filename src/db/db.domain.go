package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var once sync.Once
var ctx = context.Background()

var db *mongo.Database
var usersCollection *mongo.Collection

const usersCollectionName = "users"
