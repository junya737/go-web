package handlers

import (
	"go-web/utils"
	"net/http"
)

// "/about" のハンドラ
func AboutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := utils.PageData{
			Title:   "About",
			Message: "This is the about page.",
			Links:   []utils.Link{{Text: "Home", URL: "/"}, {Text: "About", URL: "/about"}},
		}
		utils.RenderTemplate(w, "template.html", data)
	}
}
