package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API ApiConfig
	DB  DBConfig
}

type ApiConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func LoadConfigs() error {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		// * Caso o erro seja diferente de não encontrar o arquivo não parar a aplicação
		// ! Ja que existem valores defaults setados acima
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = ApiConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Port:     viper.GetString("database.port"),
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDb() DBConfig {
	return cfg.DB
}

func GetPort() string {
	return cfg.API.Port
}
