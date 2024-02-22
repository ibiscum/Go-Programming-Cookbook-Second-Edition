package orchestrate

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Setup initializes a mongo client
func Setup(ctx context.Context, address string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	fmt.Println(address)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		return nil, err
	}

	return client, nil
}
