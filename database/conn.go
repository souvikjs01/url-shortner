package database

import (
	"context"
	"fmt"
	"os"
	"time"
	"url-shortner/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type manager struct {
	connection *mongo.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

var Mgr Manager

type Manager interface {
	Insert(interface{}, string) (interface{}, error)
	GetUrlFromCode(string, string) (types.UrlDb, error)
}

func ConnectDB() {
	uri := os.Getenv("DB_URL")
	if uri == "" {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		cancel() // Ensure the context is canceled on error
		return
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		cancel()
		return
	}

	Mgr = &manager{
		connection: client,
		ctx:        ctx,
		cancel:     cancel,
	}
	fmt.Println("DB is now connected")
}
