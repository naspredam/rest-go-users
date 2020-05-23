package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if payload == nil {
		return
	}
	response, _ := json.Marshal(payload)
	w.Write(response)
}

func fetchAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := FindAll()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorMessage{Message: "Could not retrieve the users..."})
		return
	}
	jsonResponse(w, http.StatusOK, users)
}

func fetchUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := FindByID(id)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorMessage{Message: "Could not retrieve the user..."})
		return
	}
	jsonResponse(w, http.StatusOK, user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	errDecode := decoder.Decode(&user)
	if errDecode != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorMessage{Message: "Some problem occurred."})
		return
	}

	user, err := Save(user)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorMessage{Message: "Could not save the user..."})
		return
	}
	jsonResponse(w, http.StatusCreated, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := Delete(id)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorMessage{Message: "Could not connect to the database"})
		return
	}
	jsonResponse(w, http.StatusNoContent, nil)
}

// Router blah
func Router() *mux.Router {

	var router *mux.Router
	
	router = mux.NewRouter().StrictSlash(true)

	router.Path("/users").HandlerFunc(fetchAllUsers).Methods("GET")
	router.Path("/users/{id}").HandlerFunc(fetchUserByID).Methods("GET")
	router.Path("/users").HandlerFunc(createUser).Methods("POST")
	router.Path("/users/{id}").HandlerFunc(deleteUser).Methods("DELETE")

	return router
}
