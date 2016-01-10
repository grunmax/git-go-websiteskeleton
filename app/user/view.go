package user

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/grunmax/git-go-websiteskeleton/app/common"
	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("id"))
}

func GetViewPage(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	type Page struct {
		Title  string
		UserId string
	}

	userId := ps.ByName("id")

	p := Page{
		Title:  "user_view",
		UserId: userId,
	}

	common.Templates = template.Must(template.ParseFiles("templates/user/view.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
