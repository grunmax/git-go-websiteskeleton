package home

import (
	"github.com/grunmax/git-go-websiteskeleton/app/common"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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
