package database

import (
	"fmt"
	"noraclock/src/constants"
)

func init() {
	data := []byte(fmt.Sprintf(`
{
	"views": {
		"%s": {
			"map": "function (doc) { emit(doc.createdAt, null); }"
		}
	}
}
	`, constants.CouchListMemoriesView))

	err := CouchDB.CreateDesign(conf.CouchDB.Database, constants.CouchDesign, data)
	if err == nil {
		log.Sugar().Infof("CouchDB views created successfully.")
		return
	}
	if err.Error() == constants.CouchUpdateConflictReason {
		log.Sugar().Infof("CouchDB %s design doc found already created.", constants.CouchDesign)
		return
	}
	log.Sugar().Errorf("Failed to create CouchDB %s design doc.", constants.CouchDesign)
	panic(err)
}
