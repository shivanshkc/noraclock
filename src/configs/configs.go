package configs

// Values represents the configuration schema.
type Values struct {
	// CouchDB configs.
	CouchDB struct {
		Address  string `json:"address" default:"http://127.0.0.1:5984"`
		Username string `json:"username" default:"dev"`
		Password string `json:"password" default:"dev"`
		Database string `json:"database" default:"nora-db"`
	} `json:"couchdb"`
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
		Version  string `json:"version" default:"v1.0.0"`
		Password string `json:"password" default:"secret"`
	} `json:"service"`
}
