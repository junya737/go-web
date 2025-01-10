package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"go-web/database"
	"go-web/handlers"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	// データベースを開く
	var err error
	db, err = sql.Open("sqlite3", "./database/data.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// スキーマを実行
	err = database.ExecuteSQLFile(db, "./database/schema.sql")
	if err != nil {
		log.Fatal("Failed to execute schema:", err)
	}

	// ハンドラを設定
	http.HandleFunc("/", handlers.HelloHandler(db))
	http.HandleFunc("/about", handlers.AboutHandler())

	// サーバーを開始する
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
