package home

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
		Title: "home",
	}

	common.Templates = template.Must(template.ParseFiles("templates/home/home.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
