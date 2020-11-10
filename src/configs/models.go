package configs

type logger struct {
	Level string `json:"level"`
	File  string `json:"file"`
}

type postgres struct {
	Host            string `json:"host"`
	Port            string `json:"port"`
	User            string `json:"user"`
	Password        string `json:"password"`
	DBName          string `json:"dbName"`
	RequestTimeout  int    `json:"requestTimeout"`
	MemoryTableName string `json:"memoryTableName"`
}

type server struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type service struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Password    string `json:"password"`
}

// Logger : Contains the logger configs.
var Logger = &logger{}

// Postgres : Contains the postgres configs.
var Postgres = &postgres{}

// Server : Contains the server configs.
var Server = &server{}

// Service : Contains the service configs.
var Service = &service{}
