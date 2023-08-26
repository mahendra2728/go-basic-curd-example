package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users = []User{
	{Id: 1, Firstname: "Dean", Lastname: "Ambrose"},
	{Id: 2, Firstname: "Seth", Lastname: "Rollins"},
	{Id: 3, Firstname: "Roman", Lastname: "Reign"},
}

func UserHandler() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", userHome)
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUserById)
	router.HandleFunc("/users", addUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":60000", router))
}

func userHome(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writter, "Welcome to Users API")
}

func getAllUsers(writter http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writter).Encode(users)
}

func getUserById(writter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId, _ := strconv.Atoi(params["id"])
	for _, user := range users {
		if user.Id == userId {
			json.NewEncoder(writter).Encode(user)
		}
	}
}

func addUser(writter http.ResponseWriter, request *http.Request) {
	requestBody, error := ioutil.ReadAll(request.Body)
	if error != nil {
		fmt.Println("Error : ", error)
	}
	var newUser User
	json.Unmarshal(requestBody, &newUser)
	users = append(users, newUser)
	json.NewEncoder(writter).Encode(newUser)
}
