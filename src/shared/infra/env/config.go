package env

import (
	"htmx/src/shared/utils"
	"htmx/src/shared/utils/logger"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string `json:"PORT,omitempty"`
}

func NewConfig() (Config, error) {

	if err := godotenv.Load(); err != nil {
		logger.Default.Error("Error loading .env file: " + err.Error())
	}

	return Config{
		PORT: GetEnv("PORT", "3000"),
	}, nil
}

func (c Config) String() string {
	return utils.JsonifyPretty(c)
}
