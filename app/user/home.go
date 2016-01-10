package user

import (
	"html/template"
	"net/http"

	"github.com/grunmax/git-go-websiteskeleton/app/common"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title string
	}

	p := Page{
		Title: "user_home",
	}

	common.Templates = template.Must(template.ParseFiles("templates/user/home.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
