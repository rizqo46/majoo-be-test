package test

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"majoo-be-test/repository"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupUserRepo() repository.UserRepo {
	DB, err := InitDb()
	if err != nil {
		panic("cannot connect to database")
	}
	return repository.NewUserRepo(DB)
}

func HTTPRequest(method string, url string, body io.Reader) (res []byte, statusCode int, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	req.Header.Set("Content-Type", "application/json")

	timeout := 60
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	r, err := client.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	resp := buf.Bytes()

	defer func() {
		r.Body.Close()
	}()

	return resp, r.StatusCode, nil
}

func InitDb() (*gorm.DB, error) {
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	config := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 gormLogger,
	}

	err := godotenv.Load("../.env")
	if err != nil {
		panic(".env is not loaded properly")
	}

	var DB_USERNAME = os.Getenv(`DB_USERNAME`)
	var DB_PASSWORD = os.Getenv(`DB_PASSWORD`)
	var DB_NAME = os.Getenv(`DB_NAME`)
	var DB_HOST = os.Getenv(`DB_HOST`)
	var DB_PORT = os.Getenv(`DB_PORT`)

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local`,
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	openConnection, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: openConnection,
	}), config)

	if err != nil {
		return nil, err
	}

	return db, nil
}
