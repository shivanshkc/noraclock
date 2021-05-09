package configs

// Values represents the configuration schema.
type Values struct {
	// Database holds the database related configs.
	Database struct {
		// Address is database of the database.
		Address string `json:"address"`
		// DatabaseName is the name of the database.
		DatabaseName string `json:"database_name" default:"noraclock"`
		// RequestTimeoutSeconds is the timeout in seconds for any database request.
		RequestTimeoutSeconds int `json:"request_timeout_seconds" default:"10"`
		// MemoryCollectionName is the name of the memory collection.
		MemoryCollectionName string `json:"memory_collection_name" default:"memories"`
	} `json:"database"`

	// Logger configs.
	Logger struct {
		// GeneralFilePath is the location of the general log file.
		GeneralFilePath string `json:"generalFilePath" default:"/var/log/general.log"`
		// AccessFilePath is the location of the access log file.
		AccessFilePath string `json:"accessFilePath" default:"/var/log/access.log"`
		// Level is the log level.
		Level string `json:"level" default:"info"`
	} `json:"logger"`

	// Server configs.
	Server struct {
		// Address is where the HTTP server will run.
		Address string `json:"address" default:"0.0.0.0:8080"`
	} `json:"server"`

	// Service configs.
	Service struct {
		// Name of the service.
		Name string `json:"name" default:"Noraclock"`
		// Version of the service.
		Version string `json:"version" default:"v1.0.0"`
		// Password of the service.
		Password string `json:"password" default:"secret"`
	} `json:"service"`
}
