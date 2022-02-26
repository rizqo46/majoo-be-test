package config

import (
	"majoo-be-test/config/database"
	"os"
	"strconv"

	"github.com/apex/log"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	PORT                   int
	DB                     *gorm.DB
	JwtTokenExpiryDuration int
	JwtSecretKey           string
}

func InitConfig() (*Config, error) {
	// Set Default
	c := new(Config)

	err := godotenv.Load(".env")
	if err != nil {
		log.Errorf("failed to load .env file, error :%v", err)
		return nil, err
	}

	c.PORT, err = strconv.Atoi(os.Getenv(`PORT`))
	if err != nil {
		log.Errorf("failed to get PORT, error :%v", err)
		return nil, err
	}

	c.JwtTokenExpiryDuration, err = strconv.Atoi(os.Getenv(`JWT_TOKEN_EXPIRY_DURATION`))
	if err != nil {
		log.Errorf("failed to get PORT, error :%v", err)
		return nil, err
	}

	c.JwtSecretKey = os.Getenv(`JWT_SECRET_KEY`)

	c.DB, err = database.InitDb()
	if err != nil {
		log.Errorf("failed to connect to database, error :%v", err)
		return nil, err
	}

	return c, nil
}
