package mongo

import (
	"context"
	"errors"
	"noraclock/src/configs"
	"noraclock/src/logger"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientSingleton *mongo.Client
var clientOnce = &sync.Once{}

var conf = configs.Get()
var log = logger.General()

func getClient() *mongo.Client {
	clientOnce.Do(func() {
		client, err := connectClient()
		if err != nil {
			log.Sugar().Errorf("Failed to connect to database because: %s", err.Error())
			panic(err)
		}

		clientSingleton = client
	})

	return clientSingleton
}

func connectClient() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.Database.Address))
	if err != nil {
		return nil, err
	}

	{
		ctx, cancel := getTimeoutContext()
		defer cancel()

		if err := client.Connect(ctx); err != nil {
			return nil, err
		}
	}

	{
		ctx, cancel := getTimeoutContext()
		defer cancel()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func getTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(conf.Database.RequestTimeoutSeconds)*time.Second)
}

// IsDuplicateIDError verifies if the given error is a duplicate ID error.
func IsDuplicateIDError(err error) bool {
	var writeExc mongo.WriteException

	isWriteErr := errors.As(err, &writeExc)
	if !isWriteErr {
		return false
	}

	for _, err := range writeExc.WriteErrors {
		if err.Code == 11000 {
			return true
		}
	}
	return false
}
