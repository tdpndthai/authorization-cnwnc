package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = SetUsersRouters(router)
	return router
}
