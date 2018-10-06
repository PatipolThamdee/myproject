package routes

import (
	"fmt"
	ctrl "myproject/controller"
	"myproject/middleware"

	"github.com/gorilla/mux"
)

func InitialRoutes(r *mux.Router) {

	fmt.Println(`ssss`)
	s := r.PathPrefix("/api").Subrouter()
	apiRouter(s)
}

func apiRouter(s *mux.Router) {
	fmt.Println(`eeee`)
	m := s.PathPrefix("/ok").Subrouter()
	m.HandleFunc("/test", ctrl.TestCtrl).Methods("GET")
	m.HandleFunc("/testpost", ctrl.TestCtrlPost).Methods("POST")
	m.HandleFunc("/gettoken", ctrl.GetToken).Methods("POST")
	m.HandleFunc("/checktoken", (middleware.ValidateMiddleware(ctrl.CheckToken))).Methods("POST")

}
