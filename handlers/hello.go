package handlers

import (
	"database/sql"
	"go-web/database"
	"go-web/utils"
	"net/http"
)

// "/" のハンドラ
func HelloHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var name string
		if r.Method == "POST" {
			name = r.FormValue("name")
			if name != "" {
				if err := database.SaveName(db, name); err != nil {
					http.Error(w, "Failed to save name", http.StatusInternalServerError)
					return
				}
			}
		}

		names, err := database.GetNames(db)
		if err != nil {
			http.Error(w, "Failed to load names", http.StatusInternalServerError)
			return
		}

		data := utils.PageData{
			Title:   "Home",
			Message: "Hello " + name + "!",
			Links:   []utils.Link{{Text: "Home", URL: "/"}, {Text: "About", URL: "/about"}},
			Names:   names,
		}
		utils.RenderTemplate(w, "template.html", data)
	}
}
