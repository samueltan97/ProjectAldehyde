package main

import (
	"fmt"
    "log"
    "net/http"
	//"net/url"
	"encoding/json"
	"crypto/rand"
	"github.com/gorilla/mux"
)

type Member struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Job string `json:"job"`
	Age int `json:"age"`
}


func CreateMemberHandler(w http.ResponseWriter, r *http.Request) {
		var member Member
		err := json.NewDecoder(r.Body).Decode(&member)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		member.Id = generateGUID()
		members = append(members, member)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(member)
}

func GetAllMembersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(members)	
}

func GetMemberHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
		for _, member := range members {
			if member.Id == params["id"] {
				json.NewEncoder(w).Encode(member)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
}

func PutMemberHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var member Member
	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
			w.WriteHeader(http.StatusBadRequest)
	}
	member.Id = params["id"]
	if indexOf(params["id"], members) >= 0 	{
		members[indexOf(params["id"], members)] = member
		json.NewEncoder(w).Encode(member)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func DeleteMemberHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, member := range members {
        if member.Id == params["id"] {
            members = append(members[:index], members[index+1:]...)
            break
        } else {
			w.WriteHeader(http.StatusNotFound)
		}
    }
	w.WriteHeader(http.StatusNoContent)
}

func generateGUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
    b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func indexOf(word string, members []Member) int {
    for index, member := range members {
        if word == member.Id {
            return index
        }
    }
    return -1
}

var members []Member

func main() {
	router := mux.NewRouter()
	members = append(members, Member{Id:"153a77a4-0f5f-5ae6-447c-548995440113", Name:"Samuel",Job:"Developer", Age:15})
	router.HandleFunc("/members", CreateMemberHandler).Methods("POST")
	router.HandleFunc("/members", GetAllMembersHandler).Methods("GET")
	router.HandleFunc("/members/{id}", GetMemberHandler).Methods("GET")
	router.HandleFunc("/members/{id}", PutMemberHandler).Methods("PUT")
	router.HandleFunc("/members/{id}", DeleteMemberHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
}


