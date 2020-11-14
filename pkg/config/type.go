package config

// Config struct
type Config struct {
	App        AppConfig        `json:"app"`
	DB         DatabasesConfig  `json:"db"`
	Server     ServerConfig     `json:"server"`
	GRPCServer GRPCServerConfig `json:"grpc_server"`
}

// Server struct
type ServerConfig struct {
	Port string `json:"port"`
}

// AppConfig struct
type AppConfig struct {
	Salt      string `json:"salt"`
	JWTSecret string `json:"jwt_secret"`
}

// DatabasesConfig struct
type DatabasesConfig struct {
	Postgres string `json:"postgres"`
}

type GRPCServerConfig struct {
	Account string `json:"account"`
	Seller  string `json:"seller"`
	Product string `json:"product"`
}
