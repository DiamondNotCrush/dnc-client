package router

import (
	"github.com/DiamondNotCrush/dnc-client/router/admin"
	"github.com/DiamondNotCrush/dnc-client/router/info"
	"github.com/DiamondNotCrush/dnc-client/router/share"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", admin.MainPage).Methods("GET")

	router.HandleFunc("/", info.ChangeConfig).Methods("POST")

	router.HandleFunc("/signup", admin.GetSignup).Methods("GET")
	//adds user to main server db + error handling
	router.HandleFunc("/signup", admin.PostSignup).Methods("POST")

	router.HandleFunc("/login", admin.GetLogin).Methods("GET")
	//login verification from main server db + error handling
	router.HandleFunc("/login", admin.PostLogin).Methods("POST")

	router.HandleFunc("/verify", admin.Verify).Methods("GET")

	router.HandleFunc("/library", share.Library).Methods("GET")
	//link is path name/file name and if non existent print blocking file on terminal
	router.HandleFunc("/shared/{path:.*}", share.Shared).Methods("GET")

	return router
}
