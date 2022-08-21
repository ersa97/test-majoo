package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	majoo "github.com/ersa97/test-majoo"
	md "github.com/ersa97/test-majoo/middleware"
	"github.com/ersa97/test-majoo/models"
	"github.com/gorilla/mux"
)

func Mux(majooService majoo.MajooService) {
	r := mux.NewRouter()
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(rw, r)
		})
	})

	r.HandleFunc("/", majooService.TestAPI).Methods("GET")
	//login
	r.HandleFunc("/login", majooService.Login).Methods("POST")

	//report
	r.Handle("/merchant/omzet", md.Auth(majooService.GetMerchantOmzet)).Methods("GET")
	r.Handle("/merchant/outlet/omzet", md.Auth(majooService.GetMerchantOutletOmzet)).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))

	r.NotFoundHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		json.NewEncoder(response).Encode(models.Response{
			Message: "route not found",
			Data:    nil,
		})
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		json.NewEncoder(response).Encode(models.Response{
			Message: "method not allowed",
			Data:    nil,
		})

	})

	appPort := os.Getenv("APPLICATION_PORT")

	log.Println("Running at " + os.Getenv("APP_URL") + ":" + appPort + "/")

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", appPort), r)

}
