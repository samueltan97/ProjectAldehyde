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

func main() {
	database := make(map[string]Member, 100) 
	http.Handle("/register", NewRepoHTTPHandler(database, RegisterHandler))
	http.Handle("/members", NewRepoHTTPHandler(database, MembersHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type RequestMember struct {
	Name string `json:"name"`
	Job string `json:"job"`
	Age int `json:"age"`
}

type Member struct {
	Id string	`json:"id"`
	Name string `json:"name"`
	Job string `json:"job"`
	Age int `json:"age"`
}

type RepoHTTPHandler struct {
      repo map[string]Member
	  handler func(http.ResponseWriter, *http.Request, map[string]Member)
    }
    func (th *RepoHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		th.handler(w, r, th.repo)
    }

func NewRepoHTTPHandler(repo map[string]Member, handler func(http.ResponseWriter, *http.Request, map[string]Member)) *RepoHTTPHandler {
			fmt.Println("Built")
	return &RepoHTTPHandler{repo, handler}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request, repo map[string]Member) {
		fmt.Println("Serving")
		decoder := json.NewDecoder(r.Body)
		var data RequestMember
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Print(err, "help")
		}

		var newMember Member
		newMember.Id = generateGUID()
		newMember.Name = data.Name
		newMember.Job = data.Job
		newMember.Age = data.Age

		repo[newMember.Id] = newMember

		newMemberString, jsonMarshalerr := json.Marshal(newMember)
		if jsonMarshalerr != nil {
			fmt.Println(jsonMarshalerr)
			return
		}
		fmt.Fprintf(w, string(newMemberString));
}

func MembersHandler(w http.ResponseWriter, r *http.Request, repo map[string]Member) {
	if r.URL.RequestURI() == "/members" {
        arr := []Member{} 
		for _, value := range repo {  
			arr = append(arr, value)
		}
		memberString, jsonMarshalerr := json.Marshal(arr)
		if jsonMarshalerr != nil {
			fmt.Println(jsonMarshalerr)
			return
		}
		fmt.Fprintf(w, string(memberString));
    } else {
        fmt.Println(r.URL.RequestURI())
        fmt.Println("7 is odd")
    }
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


