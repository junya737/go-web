package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type PageData struct {
	Title   string
	Message string
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
	}
	// テンプレートをパースしてレスポンスに書き込む
	tmpl, err := template.ParseFiles("template.html")

	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(w, pageData)

	if err != nil {
		fmt.Println(err)
	}

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "About",
		Message: "This is the about page.",
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	// "/" パスにリクエストが来たときにhelloHandlerを呼び出す
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about", aboutHandler)

	// サーバーを開始する
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
