package config

type Config struct {
	ItemsPerPaginationPage int
}

func InitConfig() Config {
	return Config{
		ItemsPerPaginationPage: 10,
	}
}
