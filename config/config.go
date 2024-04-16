package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		App
		HTTP
		Log
		DB
	}

	App struct {
		Name    string `env-required:"true" env:"APP_NAME"`
		Version string `env-required:"true" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" env:"LOG_LEVEL"`
	}

	DB struct {
		//Host     string `env-required:"true" env:"DB_HOST"`
		//Port     string `env-required:"true" env:"DB_PORT"`
		//User     string `env-required:"true" env:"DB_USER"`
		//Password string `env-required:"true" env:"DB_PASSWORD"`
		Name string `env-required:"true" env:"DB_NAME"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			return nil, err
		}
	}
	return &cfg, nil
}
