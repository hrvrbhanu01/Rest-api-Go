package main

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

var profiles []Profile = []Profile{}

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

func additem(q http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	q.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)

	json.NewEncoder(q).Encode(profiles)
}

func getAllProfiles(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profiles)
}

func getProfile(q http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam) //converting var idParam to integer
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted into integer."))
		return
	}

	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("No profile found with the given ID."))
		return
	}
	profile := profiles[id]
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profile)

}

func updateProfile(q http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam) //converting var idParam to integer
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted into integer."))
		return
	}

	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("No profile found with the given ID."))
		return
	}
	var updatedProfile Profile
	json.NewDecoder(r.Body).Decode(&updatedProfile) //address where we want to add the updated data i.e. the var=updatedProfile

	profiles[id] = updatedProfile

	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(updatedProfile) //q is the response writer
}

func deleteProfile(q http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam) //converting var idParam to integer
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted into integer."))
		return
	}

	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("No profile found with the given ID."))
		return
	}

	profiles = append(profiles[:id], profiles[:id+1]...)

	q.WriteHeader(200)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profiles", additem).Methods("Post")

	router.HandleFunc("/profiles", getAllProfiles).Methods("Get")

	router.HandleFunc("/profiles/{id}", getProfile).Methods("Get")

	router.HandleFunc("/profiles/{id}", updateProfile).Methods("PUT")

	router.HandleFunc("/profiles/{id}", deleteProfile).Methods("DELETE")

	http.ListenAndServe(":5000", router)
}
