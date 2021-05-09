package mongo

import (
	"context"
	"noraclock/src/exception"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var memSingleton *MemoryService
var memOnce = &sync.Once{}

// MemoryService implements the methods related to memory collection.
type MemoryService struct{}

// GetMemoryService returns the MemoryService singleton object.
func GetMemoryService() *MemoryService {
	memOnce.Do(func() {
		memSingleton = &MemoryService{}
	})

	return memSingleton
}

// Create creates a new Memory.
func (m MemoryService) Create(memory *Memory) error {
	memColl := getClient().Database(conf.Database.DatabaseName).Collection(conf.Database.MemoryCollectionName)

	ctx, cancel := getTimeoutContext()
	defer cancel()

	current := time.Now().Unix() * 1000
	memory.DocCreatedAt = current
	memory.DocUpdatedAt = current

	_, err := memColl.InsertOne(ctx, *memory)
	if err != nil {
		if IsDuplicateIDError(err) {
			return exception.MemoryAlreadyExists()
		}
		log.Sugar().Errorf("Unexpected error while creating memory: %s", err.Error())
		return err
	}

	return nil
}

// GetByID gets a Memory by ID.
func (m MemoryService) GetByID(memoryID string) (*Memory, error) {
	memColl := getClient().Database(conf.Database.DatabaseName).Collection(conf.Database.MemoryCollectionName)

	ctx, cancel := getTimeoutContext()
	defer cancel()

	memory := &Memory{}
	if err := memColl.FindOne(ctx, bson.M{"_id": memoryID}).Decode(memory); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.MemoryNotFound()
		}
		log.Sugar().Errorf("Unexpected error while getting memory: %s", err.Error())
		return nil, err
	}

	return memory, nil
}

// GetCount gets memory count.
func (m MemoryService) GetCount() (int, error) {
	memColl := getClient().Database(conf.Database.DatabaseName).Collection(conf.Database.MemoryCollectionName)

	ctx, cancel := getTimeoutContext()
	defer cancel()

	count, err := memColl.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Sugar().Errorf("Unexpected error while counting memories: %s", err.Error())
		return 0, err
	}

	return int(count), nil
}

// GetList gets memory list.
func (m MemoryService) GetList(limit int64, skip int64) ([]*Memory, error) {
	memColl := getClient().Database(conf.Database.DatabaseName).Collection(conf.Database.MemoryCollectionName)

	ctx, cancel := getTimeoutContext()
	defer cancel()

	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSkip(skip)

	cursor, err := memColl.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Sugar().Errorf("Unexpected error while getting memories: %s", err.Error())
		return nil, err
	}

	var memories []*Memory
	if err := cursor.All(context.Background(), &memories); err != nil {
		log.Sugar().Errorf("Unexpected error while decoding memories: %s", err.Error())
		return nil, err
	}

	return memories, nil
}

// Update updates a Memory.
func (m MemoryService) Update(update *Memory) error {
	memColl := getClient().Database(conf.Database.DatabaseName).Collection(conf.Database.MemoryCollectionName)

	ctx, cancel := getTimeoutContext()
	defer cancel()

	current := time.Now().Unix() * 1000
	update.DocUpdatedAt = current

	result, err := memColl.UpdateOne(ctx, bson.M{"_id": update.ID}, bson.M{"$set": *update})
	if err != nil {
		log.Sugar().Errorf("Unexpected error while updating memory: %s", err.Error())
		return err
	}

	if result.MatchedCount == 0 {
		return exception.MemoryNotFound()
	}

	return nil
}

// Delete deletes a memory.
func (m MemoryService) Delete(memoryID string) error {
	memColl := getClient().Database(conf.Database.DatabaseName).Collection(conf.Database.MemoryCollectionName)

	ctx, cancel := getTimeoutContext()
	defer cancel()

	result, err := memColl.DeleteOne(ctx, bson.M{"_id": memoryID})
	if err != nil {
		log.Sugar().Errorf("Unexpected error while deleting memory: %s", err.Error())
		return err
	}

	if result.DeletedCount == 0 {
		return exception.MemoryNotFound()
	}

	return nil
}
