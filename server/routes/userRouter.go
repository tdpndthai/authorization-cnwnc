package routers

import (
	"admin-go/controllers"

	"github.com/alexellis/faas/gateway/metrics"
	"github.com/gorilla/mux"
)

func SetUsersRouters(router *mux.Router) *mux.Router {
	metricsHandler := metrics.PrometheusHandler()
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/user/getall", controllers.GetAll).Methods("GET")
	router.HandleFunc("/user/getById/{id}", controllers.GetById).Methods("GET")
	router.HandleFunc("/user/getDetail/{id}", controllers.GetDetailUser).Methods("GET")
	router.HandleFunc("/user/update", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/add", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/delete/{id}", controllers.DeleteUser).Methods("DELETE")
	router.Handle("/metrics", metricsHandler)
	return router
}
