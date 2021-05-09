package mongo

// Memory is the schema of a memory object in database.
type Memory struct {
	ID           string `json:"id" bson:"_id"`
	Title        string `json:"title" bson:"title"`
	Body         string `json:"body" bson:"body"`
	DocCreatedAt int64  `json:"doc_created_at" bson:"doc_created_at"`
	DocUpdatedAt int64  `json:"doc_updated_at" bson:"doc_updated_at"`
}
