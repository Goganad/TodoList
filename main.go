package todo

import (
	"fmt"
	"github.com/Goganad/TodoList-REST-API/config"
	"github.com/Goganad/TodoList-REST-API/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func createRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/auth/signin", handlers.SignIn).Methods("POST")
	r.HandleFunc("/auth/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/api/lists/", handlers.GetAllLists).Methods("GET")
	r.HandleFunc("/api/lists/", handlers.CreateList).Methods("POST")
	r.HandleFunc("/api/lists/{id}", handlers.GetListById).Methods("GET")
	r.HandleFunc("/api/lists/{id}", handlers.UpdateList).Methods("PUT")
	r.HandleFunc("/api/lists/{id}", handlers.DeleteList).Methods("DELETE")
	r.HandleFunc("/api/lists/{id}/items", handlers.GetAllItems).Methods("GET")
	r.HandleFunc("/api/lists/{id}/items", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/api/lists/{id}/items/{item_id}", handlers.GetItemById).Methods("GET")
	r.HandleFunc("/api/lists/{id}/items/{item_id}", handlers.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/lists/{id}/items/{item_id}", handlers.DeleteItem).Methods("DELETE")

	return r
}

func main(){
	r := createRouter()

	PORT := config.ServerAddress
	fmt.Println("Serving on port: ", PORT)

	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}