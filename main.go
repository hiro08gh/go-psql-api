package main

import (
	"github.com/go-psql-api/model"
	"github.com/go-psql-api/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	//.envファイルの読み込み
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//DB setup
	db := model.GetDBConn()

	defer db.Close()

	router.Init()
}
