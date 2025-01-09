package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type PageData struct {
	Title   string
	Message string
	Links   []Link
}

type Link struct {
	Text string
	URL  string
}

var Links = []Link{
	{Text: "Home", URL: "/"},
	{Text: "About", URL: "/about"},
}

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

// ハンドラ関数を定義
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// w: レスポンスを管理する構造体
	// r: リクエストを保持する構造体
	// // レスポンスを書き込む
	// fmt.Fprintln(w, "Hello, World!")
	// テンプレートを使ってHTMLを生成

	pageData := PageData{
		Title:   "Hello, World!",
		Message: "This is a message from the server.",
		Links:   Links,
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

func main() {
	// "/" パスにリクエストが来たときにhelloHandlerを呼び出す
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about", aboutHandler)

	// サーバーを開始する
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
