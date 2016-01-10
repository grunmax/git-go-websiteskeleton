package user

import (
	"html/template"
	"net/http"

	"github.com/grunmax/git-go-websiteskeleton/app/common"

	"github.com/gorilla/mux"
)

func GetViewPage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title  string
		UserId string
	}

	params := mux.Vars(req)
	userId := params["id"]

	p := Page{
		Title:  "user_view",
		UserId: userId,
	}

	common.Templates = template.Must(template.ParseFiles("templates/user/view.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
