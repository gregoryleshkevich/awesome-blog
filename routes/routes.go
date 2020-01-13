package routes

import (
	"awesome-blog/internal/controllers"
	"github.com/gorilla/mux"
)

var postController controllers.PostController = controllers.PostController{}

//Routes struct
type Routes struct {
	R *mux.Router
}
	
// InitRoutes routes
func (routes *Routes) InitRoutes() {
	routes.R.HandleFunc("/posts", postController.Create).Methods("POST")
	routes.R.HandleFunc("/posts", postController.Shows).Methods("GET")
	routes.R.HandleFunc("/posts/{id}", postController.Show).Methods("GET")
}
