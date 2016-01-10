package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/grunmax/git-go-websiteskeleton/app/common"
	"github.com/grunmax/git-go-websiteskeleton/app/home"
	"github.com/grunmax/git-go-websiteskeleton/app/user"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	router := httprouter.New()
	http.Handle("/", httpInterceptor(router))

	router.GET("/", home.GetHomePage)
	router.GET("/user", user.GetHomePage)

	router.GET("/user/:id", user.GetViewPage)

	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fileServer)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func httpInterceptor(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		router.ServeHTTP(w, req)

		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)

		switch req.Method {
		case "GET":
			// We may not always want to StatusOK, but for the sake of
			// this example we will
			common.LogAccess(w, req, elapsedTime)
		case "POST":
			// here we might use http.StatusCreated
		}

	})
}
