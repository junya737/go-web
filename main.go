package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type PageData struct {
	Title   string
	Message string
	Links   []Link
	Names   []string // 保存されたname
}

type Link struct {
	Text string
	URL  string
}

var db *sql.DB

var Links = []Link{
	{Text: "Home", URL: "/"},
	{Text: "About", URL: "/about"},
}

var savedNames []string // 保存されたname

// テンプレートを使ってHTMLを生成する関数
func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// データベースにnameを保存する関数
func saveName(name string) error {
	_, err := db.Exec("INSERT INTO names (name) VALUES (?)", name)
	return err
}

// ハンドラ関数を定義
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// w: レスポンスを管理する構造体
	// r: リクエストを保持する構造体
	// // レスポンスを書き込む
	// fmt.Fprintln(w, "Hello, World!")
	// テンプレートを使ってHTMLを生成

	var name string
	if r.Method == "POST" {
		name = r.FormValue("name")
		if name != "" {
			savedNames = append(savedNames, name)
			err := saveName(name)
			if err != nil {
				http.Error(w, "Failed to save name", http.StatusInternalServerError)
				return
			}
		}
	}

	pageData := PageData{
		Title:   "Home",
		Message: "Hello " + name + "!",
		Links:   Links,
		Names:   savedNames,
	}
	renderTemplate(w, "template.html", pageData)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		Title:   "About",
		Message: "This is the about page.",
		Links:   Links,
	}

	renderTemplate(w, "template.html", data)

}

func executeSQLFile(db *sql.DB, filePath string) error {
	//// ファイルを読み込む
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	// SQLファイルの内容を実行する
	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("failed to execute SQL: %w", err)
	}
	return nil
}

func main() {
	// データベースを開く
	var err error
	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	err = executeSQLFile(db, "./schema.sql")
	if err != nil {
		log.Fatal("Failed to execute schema:", err)
	}

	// "/" パスにリクエストが来たときにhelloHandlerを呼び出す
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about", aboutHandler)

	// サーバーを開始する
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
