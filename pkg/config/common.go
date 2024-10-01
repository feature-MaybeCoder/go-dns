package config

type Config struct {
	ItemsPerPaginationPage int
	MongoDBUrl             string `env:"MONGO_URL"`
}

func InitConfig() Config {
	config := Config{
		ItemsPerPaginationPage: 10,
	}
	config = parseEnvConfig(config)

	return config
}

var CommonConfig Config = InitConfig()
