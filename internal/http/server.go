package http

import (
	"net/http"

	"github.com/redis/go-redis/v9"
)

// todolist CRUD

func createList(w http.ResponseWriter, req *http.Request) {
	// TODO: create new todo list

}

func getLists(w http.ResponseWriter, req *http.Request) {
	// TODO: add deletion of access token id from db

}

func deleteList(w http.ResponseWriter, req *http.Request) {
	// TODO: add deletion of access token id from db

}

func renameList(w http.ResponseWriter, req *http.Request) {
	// TODO: add deletion of access token id from db

}

// list item CRUD
func getListItems(w http.ResponseWriter, req *http.Request) {
	// TODO: add deletion of access token id from db

}

func addListItem(w http.ResponseWriter, req *http.Request) {
	// TODO: add deletion of access token id from db

}

func deleteListItem(w http.ResponseWriter, req *http.Request) {
	// TODO: add deletion of access token id from db

}

func updateListItem(w http.ResponseWriter, req *http.Request) {
	// TODO: add deletion of access token id from db

}

func getListItem(w http.ResponseWriter, req *http.Request) {
	// TODO: add deletion of access token id from db

}

// func headers(w http.ResponseWriter, req *http.Request) {
//
// 	for name, headers := range req.Header {
// 		for _, h := range headers {
// 			fmt.Fprintf(w, "%v: %v\n", name, h)
// 		}
// 	}
// }

func RootRouter(redis redis.Client) *http.ServeMux {
	root := http.NewServeMux()
	auth := newAuthHandler(redis)

	root.Handle("/api/v1/auth/", http.StripPrefix("/api/v1/auth", auth.authRouter()))

	return root
}
